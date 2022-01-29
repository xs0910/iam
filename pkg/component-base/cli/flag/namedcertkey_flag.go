package flag

import (
	"errors"
	"flag"
	"strings"
)

// NamedCertKey is a flag value parsing "certFile,keyFile" and "certFile,keyFile:name,name,name".
type NamedCertKey struct {
	Names             []string
	CertFile, KeyFile string
}

var _ flag.Value = &NamedCertKey{}

func (nkc *NamedCertKey) String() string {
	s := nkc.CertFile + "," + nkc.KeyFile
	if len(nkc.Names) > 0 {
		s = s + ":" + strings.Join(nkc.Names, ",")
	}
	return s
}

func (nkc *NamedCertKey) Set(value string) error {
	cs := strings.SplitN(value, ":", 2)
	var keyCert string
	if len(cs) == 2 {
		var names string
		keyCert, names = strings.TrimSpace(cs[0]), strings.TrimSpace(cs[1])
		if names == "" {
			return errors.New("empty names list is not allowed")
		}
		nkc.Names = nil
		for _, name := range strings.Split(names, ",") {
			nkc.Names = append(nkc.Names, strings.TrimSpace(name))
		}
	} else {
		nkc.Names = nil
		keyCert = strings.TrimSpace(cs[0])
	}
	cs = strings.Split(keyCert, ",")
	if len(cs) != 2 {
		return errors.New("expected comma separated certificate and key file paths")
	}
	nkc.CertFile = strings.TrimSpace(cs[0])
	nkc.KeyFile = strings.TrimSpace(cs[1])
	return nil
}

func (*NamedCertKey) Type() string {
	return "namedCertKey"
}

// NamedCertKeyArray is a flag value parsing NamedCertKeys, each passed with its own
// flag instance (in contrast to comma separated slices).
type NamedCertKeyArray struct {
	value   *[]NamedCertKey
	changed bool
}

var _ flag.Value = &NamedCertKeyArray{}

// NewNamedCertKeyArray creates a new NamedCertKeyArray with the internal value
// pointing to p.
func NewNamedCertKeyArray(p *[]NamedCertKey) *NamedCertKeyArray {
	return &NamedCertKeyArray{
		value: p,
	}
}

func (a *NamedCertKeyArray) Set(val string) error {
	nkc := NamedCertKey{}
	err := nkc.Set(val)
	if err != nil {
		return err
	}
	if !a.changed {
		*a.value = []NamedCertKey{nkc}
		a.changed = true
	} else {
		*a.value = append(*a.value, nkc)
	}
	return nil
}

func (a *NamedCertKeyArray) Type() string {
	return "namedCertKey"
}

func (a *NamedCertKeyArray) String() string {
	m := make([]string, 0, len(*a.value))
	for i := range *a.value {
		m = append(m, (*a.value)[i].String())
	}
	return "[" + strings.Join(m, ";") + "]"
}
