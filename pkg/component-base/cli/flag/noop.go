package flag

import (
	goflag "flag"
	"github.com/spf13/pflag"
)

// NoOp implements goflag.Value and plfag.Value,
// but has a noop Set implementation.
type NoOp struct {
}

var (
	_ goflag.Value = NoOp{}
	_ pflag.Value  = NoOp{}
)

func (n NoOp) Type() string {
	return "noop"
}

func (n NoOp) String() string {
	return ""
}

func (n NoOp) Set(s string) error {
	return nil
}
