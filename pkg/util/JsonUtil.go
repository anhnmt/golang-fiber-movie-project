package util

import "encoding/json"

func JSONMarshal(data interface{}) ([]byte, error) {
	raw, err := json.Marshal(data)

	return raw, err
}

func JSONUnmarshal(data []byte, object interface{}) error {
	err := json.Unmarshal(data, &object)

	return err
}
