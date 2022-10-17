package conv

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

func MustJSONString(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return string(b)
}

func MustJSONBytes(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Panic(err)
	}
	return b
}

func JSONExample(i any) string {
	v := DeepNew(reflect.TypeOf(i))
	return MustJSONString(v.Interface())
}

func FileToJSON(filename string, v any) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func JSONToFile(v any, filename string) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
