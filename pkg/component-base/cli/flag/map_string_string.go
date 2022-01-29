package flag

import (
	"fmt"
	"sort"
	"strings"
)

// MapStringString can be set from the command line with the format `--flag "string=string"`.
// Multiple flag invocations are supported. For example: `--flag "a=foo" --flag "b=bar"`.
// If this is desired to be the only type invocation `NoSplit` should be set to true.
// Multiple comma-separated key-value pairs in a single invocation are supported if `NoSplit`
// is set to false. For example: `--flag "a=foo,b=bar"`.
type MapStringString struct {
	Map         *map[string]string
	initialized bool
	NoSplit     bool
}

func NewMapStringString(m *map[string]string) *MapStringString {
	return &MapStringString{Map: m}
}
func NewMapStringStringNoSplit(m *map[string]string) *MapStringString {
	return &MapStringString{
		Map:     m,
		NoSplit: true,
	}
}

func (m *MapStringString) Set(value string) error {
	if m.Map == nil {
		return fmt.Errorf("no target (nil point to map[string]string)")
	}
	if !m.initialized || *m.Map == nil {
		// clear default values, or allocate if no existing map
		*m.Map = make(map[string]string)
		m.initialized = true
	}

	if !m.NoSplit {
		for _, s := range strings.Split(value, ",") {
			if len(s) == 0 {
				continue
			}
			arr := strings.SplitN(s, "=", 2)
			if len(arr) != 2 {
				return fmt.Errorf("malformed pair, expect string=string")
			}
			k := strings.TrimSpace(arr[0])
			v := strings.TrimSpace(arr[1])

			(*m.Map)[k] = v
		}
	}

	// account for only one key-value pair in a single invocation
	arr := strings.SplitN(value, "=", 2)
	if len(arr) != 2 {
		return fmt.Errorf("malformed pair, expect string=string")
	}
	k := strings.TrimSpace(arr[0])
	v := strings.TrimSpace(arr[1])
	(*m.Map)[k] = v
	return nil
}

func (m *MapStringString) String() string {
	if m == nil || m.Map == nil {
		return ""
	}
	var pairs []string
	for k, v := range *m.Map {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(pairs)
	return strings.Join(pairs, ",")
}

func (m *MapStringString) Type() string {
	return "mapStringString"
}

func (m *MapStringString) Empty() bool {
	return len(*m.Map) == 0
}
