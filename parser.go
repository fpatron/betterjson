package betterjson

import (
	"encoding/json"
)

// Unmarshal parses the JSON-encoded data with comment support and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	processedData := decomment(data)
	return json.Unmarshal(processedData, v)
}
