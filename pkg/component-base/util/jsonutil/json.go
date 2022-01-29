package jsonutil

import (
	"encoding/json"
	"log"

	"github.com/bitly/go-simplejson"
)

type fakeJson struct {
	*simplejson.Json
}

func NewJson(data []byte) (Json, error) {
	j, err := simplejson.NewJson(data)
	return &fakeJson{j}, err
}

func (j *fakeJson) Get(key string) Json {
	return &fakeJson{j.Json.Get(key)}
}

func (j *fakeJson) GetPath(branch ...string) Json {
	return &fakeJson{j.Json.GetPath(branch...)}
}
func (j *fakeJson) CheckGet(key string) (Json, bool) {
	result, ok := j.Json.CheckGet(key)
	return &fakeJson{result}, ok
}

func Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func ToString(v interface{}) string {
	b, err := Encode(v)
	if err != nil {
		return ""
	}

	return string(b)
}

func ToJson(data interface{}) Json {
	var j Json
	j = &fakeJson{simplejson.New()}
	b, err := Encode(data)
	if err != nil {
		log.Printf("Failed to encode [%+v] to []byte, error: %+v", data, err)
		return j
	}

	j, err = NewJson(b)
	if err != nil {
		log.Printf("Failed to decode [%+v] to Json, error: %+v", data, err)
	}
	return j
}
