package jsonutil

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JSONRawMessage []byte

func (m JSONRawMessage) Find(key string) JSONRawMessage {
	var objMap map[string]json.RawMessage
	err := json.Unmarshal(m, &objMap)
	if err != nil {
		fmt.Printf("Resolve JSON Key failed, find key = %s, err= %s", key, err)
		return nil
	}
	return JSONRawMessage(objMap[key])
}

func (m JSONRawMessage) ToList() []JSONRawMessage {
	var lists []json.RawMessage
	err := json.Unmarshal(m, &lists)
	if err != nil {
		fmt.Printf("Resolve JSON List failed, err=%s", err)
		return nil
	}
	var res []JSONRawMessage
	for _, v := range lists {
		res = append(res, JSONRawMessage(v))
	}
	return res
}

func (m JSONRawMessage) ToString() string {
	res := strings.ReplaceAll(string(m[:]), "\"", "")
	return res
}
