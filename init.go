package conv

import (
	"encoding/json"
	"encoding/xml"
)

func init() {
	RegisterUnmarshalFunc(_JSON, json.Unmarshal)
	RegisterUnmarshalFunc(_JsonUTF8, json.Unmarshal)
	RegisterUnmarshalFunc(_XML, xml.Unmarshal)
	RegisterUnmarshalFunc(_XML2, xml.Unmarshal)
	RegisterUnmarshalFunc(_XmlUTF8, xml.Unmarshal)
}
