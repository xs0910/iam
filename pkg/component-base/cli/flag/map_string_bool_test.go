package flag

import (
	"reflect"
	"testing"
)

func TestMapStringBool_Set(t *testing.T) {
	var nilMap map[string]bool
	cases := []struct {
		desc   string
		vals   []string
		start  *MapStringBool
		expect *MapStringBool
		err    string
	}{
		{
			desc:  "clear defaults",
			vals:  []string{""},
			start: NewMapStringBool(&map[string]bool{"default": true}),
			expect: &MapStringBool{
				Map:         &map[string]bool{},
				initialized: true,
			},
			err: "",
		},
		{
			desc:  "allocates map if currently nil",
			vals:  []string{""},
			start: &MapStringBool{initialized: true, Map: &nilMap},
			expect: &MapStringBool{
				Map:         &map[string]bool{},
				initialized: true,
			},
			err: "",
		},
		{
			desc:  "one key",
			vals:  []string{"one=true"},
			start: NewMapStringBool(&nilMap),
			expect: &MapStringBool{
				Map:         &map[string]bool{"one": true},
				initialized: true,
			},
			err: "",
		},
		{
			desc:  "two keys",
			vals:  []string{"one=true,two=false"},
			start: NewMapStringBool(&nilMap),
			expect: &MapStringBool{
				Map:         &map[string]bool{"one": true, "two": false},
				initialized: true,
			},
			err: "",
		},
		{
			desc:  "two keys, multiple Set invocations",
			vals:  []string{"one=true", "two=false"},
			start: NewMapStringBool(&nilMap),
			expect: &MapStringBool{
				Map:         &map[string]bool{"one": true, "two": false},
				initialized: true,
			},
			err: "",
		},
	}

	for _, c := range cases {
		nilMap = nil
		t.Run(c.desc, func(t *testing.T) {
			var err error
			for _, val := range c.vals {
				err = c.start.Set(val)
				if err != nil {
					break
				}
			}

			if c.err != "" {
				if err == nil || err.Error() != c.err {
					t.Fatalf("expect error %s but got %v", c.err, err)
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(c.expect, c.start) {
				t.Fatalf("expect %#v but got %#v", c.expect, c.start)
			}
		})
	}
}
