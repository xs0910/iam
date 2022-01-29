package flag

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// MapStringBool can be set from the command line with the format `--flag "string=bool"`.
// Multiple comma-separated key-value pairs in a single invocation are supported.
// For example: `--flag "a=true,b=false"`.
// Multiple flag invocations are supported. For example: `--flag "a=true" --flag "b=false"`.
type MapStringBool struct {
	Map         *map[string]bool
	initialized bool
}

func NewMapStringBool(m *map[string]bool) *MapStringBool {
	return &MapStringBool{Map: m}
}

func (m *MapStringBool) Set(value string) error {
	if m.Map == nil {
		return fmt.Errorf("no target (nil point to map[string]bool)")
	}
	if !m.initialized || *m.Map == nil {
		// clear default values, or allocate if no existing map
		*m.Map = make(map[string]bool)
		m.initialized = true
	}

	for _, s := range strings.Split(value, ",") {
		if len(s) == 0 {
			continue
		}
		arr := strings.SplitN(s, "=", 2)
		if len(arr) != 2 {
			return fmt.Errorf("malformed pair, expect string=bool")
		}
		k := strings.TrimSpace(arr[0])
		v := strings.TrimSpace(arr[1])
		boolValue, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("invalid value of %s: %s, err: %v", k, v, err)
		}
		(*m.Map)[k] = boolValue
	}
	return nil
}

func (m *MapStringBool) String() string {
	if m == nil || m.Map == nil {
		return ""
	}
	var pairs []string
	for k, v := range *m.Map {
		pairs = append(pairs, fmt.Sprintf("%s=%t", k, v))
	}
	sort.Strings(pairs)
	return strings.Join(pairs, ",")
}

func (m *MapStringBool) Type() string {
	return "mapStringBool"
}

func (m *MapStringBool) Empty() bool {
	return len(*m.Map) == 0
}
