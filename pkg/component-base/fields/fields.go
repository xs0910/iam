package fields

import (
	"sort"
	"strings"
)

// Fields allows you to present fields independently of their storage.
type Fields interface {
	Has(filed string) (exists bool)
	Get(field string) (value string)
}

// Set is a map of field:value. It implements Fields.
type Set map[string]string

func (s Set) Has(filed string) (exists bool) {
	_, exists = s[filed]
	return
}

func (s Set) Get(field string) (value string) {
	return s[field]
}

func (s Set) String() string {
	selector := make([]string, 0, len(s))
	for key, value := range s {
		selector = append(selector, key+"="+value)
	}

	// Sort for determinism.
	sort.StringSlice(selector).Sort()
	return strings.Join(selector, ",")
}

// AsSelector converts fields into a selectors.
func (s Set) AsSelector() Selector {
	return SelectorFromSet(s)
}
