//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

type RawMessage = json.RawMessage

var (
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewEncoder    = json.NewEncoder
	NewDecoder    = json.NewDecoder
)
