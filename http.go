package conv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"

	"code.olapie.com/errors"
)

const (
	mimetypeTextPlain = "text/plain"
	mimetypeHTML      = "text/html"
	mimetypeXML2      = "text/xml"
	mimetypeCSS       = "text/css"
	mimetypeXML       = "application/xml"
	mimetypeXHTML     = "application/xhtml+xml"
	mimetypeProtobuf  = "application/x-protobuf"

	mimetypeFormData = "multipart/form-data"
	mimetypeGIF      = "image/gif"
	mimetypeJPEG     = "image/jpeg"
	mimetypePNG      = "image/png"
	mimetypeWEBP     = "image/webp"
	mimetypeICON     = "image/x-icon"

	mimetypeMPEG = "video/mpeg"

	mimetypeFormURLEncoded = "application/x-www-form-urlencoded"
	mimetypeOctetStream    = "application/octet-stream"
	mimetypeJSON           = "application/json"
	mimetypePDF            = "application/pdf"
	mimetypeMSWord         = "application/msword"
	mimetypeGZIP           = "application/x-gzip"
	mimetypeWASM           = "application/wasm"
	mimetypeContentType    = "Content-Type"
	mimetypeAcceptEncoding = "Accept-Encoding"
)

const (
	charsetUTF8       = "charset=utf-8"
	charsetUTF8Suffix = "; " + charsetUTF8
	textPlainUTF8     = mimetypeTextPlain + charsetUTF8Suffix
	htmlUTF8          = mimetypeHTML + charsetUTF8Suffix
	jsonUTF8          = mimetypeJSON + charsetUTF8Suffix
	xmlUTF8           = mimetypeXML + charsetUTF8Suffix
)

// ToHTTPAttachment returns value for Content-Disposition
// e.g. Content-Disposition: attachment; filename=test.txt
func ToHTTPAttachment(filename string) string {
	return fmt.Sprintf(`attachment; filename="%s"`, filename)
}

func GetHTTPContentType(h http.Header) string {
	t := h.Get(mimetypeContentType)
	for i, ch := range t {
		if ch == ' ' || ch == ';' {
			t = t[:i]
			break
		}
	}
	return t
}

func IsMIMETextType(typ string) bool {
	switch typ {
	case mimetypeTextPlain, mimetypeHTML, mimetypeCSS, mimetypeXML, mimetypeXML2, mimetypeXHTML, mimetypeJSON, textPlainUTF8, htmlUTF8, jsonUTF8, xmlUTF8:
		return true
	default:
		return false
	}
}

func GetHTTPAcceptEncodings(h http.Header) []string {
	a := strings.Split(h.Get(mimetypeAcceptEncoding), ",")
	for i, s := range a {
		a[i] = strings.TrimSpace(s)
	}

	// Remove empty strings
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == "" {
			a = append(a[:i], a[i+1:]...)
		}
	}
	return a
}

func HTTPCookiesToMap(cookies []*http.Cookie) map[string]any {
	params := map[string]any{}
	for _, c := range cookies {
		params[c.Name] = c.Value
	}
	return params
}

func HTTPHeaderToMap(h http.Header) map[string]any {
	params := map[string]any{}
	for k, v := range h {
		k = strings.ToLower(k)
		if strings.HasPrefix(k, "x-") {
			k = k[2:]
			k = strings.Replace(k, "-", "_", -1)
			params[k] = v
		}
	}
	return params
}

func URLValuesToMap(values url.Values) map[string]any {
	m := map[string]any{}
	for k, va := range values {
		isArray := strings.HasSuffix(k, "[]")
		if isArray {
			k = k[0 : len(k)-2]
			if k == "" {
				continue
			}

			if len(va) == 1 {
				va = strings.Split(va[0], ",")
			}
		}

		if len(va) == 0 {
			continue
		}

		k = strings.ToLower(k)
		if isArray || len(va) > 1 {
			// value is an array or expected to be an array
			m[k] = va
		} else {
			m[k] = va[0]
		}
	}

	if jsonStr, _ := m["json"].(string); jsonStr != "" {
		var j map[string]any
		err := json.Unmarshal([]byte(jsonStr), &j)
		if err == nil {
			for k, v := range m {
				j[k] = v
			}
			m = j
		}
	}
	return m
}

func ParseHTTPRequest(req *http.Request, memInBytes int64) (map[string]any, []byte, error) {
	typ := GetHTTPContentType(req.Header)
	params := map[string]any{}
	switch typ {
	case mimetypeHTML, mimetypeTextPlain:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return params, nil, fmt.Errorf("read html or plain body: %w", err)
		}
		return params, body, nil
	case mimetypeJSON:
		body, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return params, nil, fmt.Errorf("read json body: %w", err)
		}
		if len(body) == 0 {
			return params, nil, nil
		}
		decoder := json.NewDecoder(bytes.NewBuffer(body))
		decoder.UseNumber()
		err = decoder.Decode(&params)
		if err != nil {
			var obj any
			err = json.Unmarshal(body, &obj)
			if err != nil {
				return params, body, fmt.Errorf("unmarshal json %s: %w", string(body), err)
			}
		}
		return params, body, nil
	case mimetypeFormURLEncoded:
		// TODO: will crash
		//body, err := req.GetBody()
		//if err != nil {
		//	return params, nil, fmt.Errorf("get body: %w", err)
		//}
		//bodyData, err := ioutil.Read(body)
		//body.Close()
		//if err != nil {
		//	return params, nil, fmt.Errorf("read form body: %w", err)
		//}
		if err := req.ParseForm(); err != nil {
			return params, nil, fmt.Errorf("parse form: %w", err)
		}
		return URLValuesToMap(req.Form), nil, nil
	case mimetypeFormData:
		err := req.ParseMultipartForm(memInBytes)
		if err != nil {
			return nil, nil, fmt.Errorf("parse multipart form: %w", err)
		}

		if req.MultipartForm != nil && req.MultipartForm.File != nil {
			return URLValuesToMap(req.MultipartForm.Value), nil, nil
		}
		return params, nil, nil
	default:
		body, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return params, nil, fmt.Errorf("read json body: %w", err)
		}
		return params, body, nil
	}
}

func URLJoin(a ...string) string {
	if len(a) == 0 {
		return ""
	}
	// path.Join will convert // to be /
	p := path.Join(a...)
	p = strings.Replace(p, ":/", "://", 1)
	i := strings.Index(p, "://")
	s := p
	if i >= 0 {
		i += 3
		s = p[i:]
		l := strings.Split(s, "/")
		for i, v := range l {
			l[i] = url.PathEscape(v)
		}
		p = p[:i] + path.Join(l...)
	}
	return p
}

func ToURLValues(i any) (url.Values, error) {
	i = IndirectToStringerOrError(i)
	if i == nil {
		return nil, errors.New("nil values")
	}
	switch v := i.(type) {
	case url.Values:
		return v, nil
	}

	b, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("cannot convert %#v of type %T to url.Values", i, i)
	}
	var m map[string]any
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, fmt.Errorf("cannot convert %#v of type %T to url.Values", i, i)
	}
	uv := url.Values{}
	for k, v := range m {
		uv.Set(k, fmt.Sprint(v))
	}
	return uv, nil
}

func MustURLValues(i any) url.Values {
	v, err := ToURLValues(i)
	if err != nil {
		panic(err)
	}
	return v
}

func ToURLString(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("parse: %w", err)
	}
	if u.Scheme == "" {
		return "", errors.New("missing schema")
	}
	if u.Host == "" {
		return "", errors.New("missing host")
	}
	return u.String(), nil
}

func IsURL(s string) bool {
	u, _ := ToURLString(s)
	return u != ""
}

func VarargsToURLValues(keyAndValues ...any) (url.Values, error) {
	uv := url.Values{}
	keys, vals, err := VarargsToSlice(keyAndValues...)
	if err != nil {
		return nil, err
	}
	for i, k := range keys {
		a, err := ToStringSlice(vals[i])
		if err != nil {
			return nil, err
		}
		for _, v := range a {
			if v != "" {
				uv.Add(k, v)
			}
		}
	}
	return uv, nil
}

func MustVarargsToURLValues(keyAndValues ...any) url.Values {
	v, err := VarargsToURLValues(keyAndValues...)
	if err != nil {
		panic(err)
	}
	return v
}

type UnmarshalFunc func([]byte, any) error

var contentTypeToUnmarshalFunc sync.Map

func RegisterUnmarshalFunc(contentType string, f UnmarshalFunc) {
	contentTypeToUnmarshalFunc.Store(contentType, f)
}

func GetUnmarshalFunc(contentType string) UnmarshalFunc {
	v, ok := contentTypeToUnmarshalFunc.Load(contentType)
	if ok {
		u, _ := v.(UnmarshalFunc)
		return u
	}
	return nil
}

func GetHTTPResult[T any](resp *http.Response) (T, error) {
	var res T
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return res, fmt.Errorf("read resp body: %v", err)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return res, errors.Format(resp.StatusCode, string(body))
	}

	if any(res) == nil {
		return res, nil
	}

	ct := GetHTTPContentType(resp.Header)
	if f := GetUnmarshalFunc(ct); f != nil {
		err = f(body, &res)
		return res, errors.Wrapf(err, "unmarshal")
	}

	if len(body) == 0 {
		err = errors.New("no data")
	} else if _, ok := any(res).([]byte); ok {
		res = any(body).(T)
	} else {
		if err = SetBytes(&res, body); err != nil {
			err = fmt.Errorf("cannot handle %s: %w ", ct, err)
		}
	}
	return res, err
}
