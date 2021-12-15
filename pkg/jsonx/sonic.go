package jsonx

import (
	"encoding/json"
)

func Marshal(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}

func Unmarshal(buf []byte, val interface{}) error {
	return json.Unmarshal(buf, val)
}
