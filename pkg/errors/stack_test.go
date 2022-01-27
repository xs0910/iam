package errors

import (
	"runtime"
	"testing"
)

var initPc = caller()

type X struct {
}

// val returns a Frame pointing to itself.
func (x X) val() Frame {
	return caller()
}

// ptr returns a Frame pointing to itself.
func (x X) ptr() Frame {
	return caller()
}

func TestFrameFormat(t *testing.T) {
	var tests = []struct {
		Frame
		format string
		want   string
	}{{
		initPc,
		"%s",
		"stack_test.go",
	}, {
		initPc,
		"%+s",
		"github.com/xs0910/iam/pkg/errors.init\n" +
			"\t.+/github.com/xs0910/iam/pkg/errors/stack_test.go",
	}, {
		0,
		"%s",
		"unknown",
	}, {
		0,
		"%+s",
		"unknown",
	}, {
		initPc,
		"%d",
		"8",
	}, {
		0,
		"%d",
		"0",
	}, {
		initPc,
		"%n",
		"init",
	}, {
		func() Frame {
			var x X
			return x.ptr()
		}(),
		"%n",
		`X.ptr`,
	}, {
		func() Frame {
			var x X
			return x.val()
		}(),
		"%n",
		"X.val",
	}, {
		0,
		"%n",
		"",
	}, {
		initPc,
		"%v",
		"stack_test.go:8",
	}, {
		initPc,
		"%+v",
		"github.com/xs0910/iam/pkg/errors.init\n" +
			"\t.+/github.com/xs0910/iam/pkg/errors/stack_test.go:8",
	}, {
		0,
		"%v",
		"unknown:0",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.Frame, tt.format, tt.want)
	}
}

func TestFuncName(t *testing.T) {
	tests := []struct {
		name, want string
	}{
		{"", ""},
		{"runtime.main", "main"},
		{"github.com/xs0910/iam/pkg/errors.funcName", "funcName"},
		{"funcName", "funcName"},
		{"io.copyBuffer", "copyBuffer"},
		{"main.(*R).Write", "(*R).Write"},
	}
	for _, tt := range tests {
		got := funcName(tt.name)
		want := tt.want
		if got != want {
			t.Errorf("funcName(%q): want: %q, got: %q", tt.name, want, got)
		}
	}
}

// a version of runtime.Caller that returns a Frame, not an uintptr
func caller() Frame {
	var pcs [3]uintptr
	n := runtime.Callers(2, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	frame, _ := frames.Next()
	return Frame(frame.PC)
}
