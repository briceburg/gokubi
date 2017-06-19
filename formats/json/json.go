package json

import "encoding/json"

func Marshal(in map[string]interface{}) ([]byte, error) {
	return json.Marshal(in)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
