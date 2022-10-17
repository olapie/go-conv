package conv

import (
	"encoding/json"
	"encoding/xml"
)

func init() {
	RegisterUnmarshalFunc(mimetypeJSON, json.Unmarshal)
	RegisterUnmarshalFunc(jsonUTF8, json.Unmarshal)
	RegisterUnmarshalFunc(mimetypeXML, xml.Unmarshal)
	RegisterUnmarshalFunc(mimetypeXML2, xml.Unmarshal)
	RegisterUnmarshalFunc(xmlUTF8, xml.Unmarshal)
}
