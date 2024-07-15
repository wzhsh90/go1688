package utils

import "encoding/json"

func Unmarshal(data []byte, pt interface{}) error {
	return json.Unmarshal(data, pt)
}
func Marshal(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}
func MarshalStr(val interface{}) string {
	byteData, merr := json.Marshal(val)
	if merr != nil {
		return ""
	}
	return string(byteData)
}
