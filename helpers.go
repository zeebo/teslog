package main

import (
	"bytes"
	"encoding/json"
)

func jsonEncode(x interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	err := enc.Encode(x)
	return buf.Bytes(), err
}
