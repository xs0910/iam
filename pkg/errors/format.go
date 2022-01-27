package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// formatInfo contains all the error information.
type formatInfo struct {
	code    int
	message string
	err     string
	stack   *stack
}

// Format implements fmt.Formatter. https://golang.org/pkg/fmt/#hdr-Printing
//
// Verbs:
//     %s  - Returns the user-safe error string mapped to the error code or
//       ┊   the error message if none is specified.
//     %v      Alias for %s
//
// Flags:
//      #      JSON formatted output, useful for logging
//      -      Output caller details, useful for troubleshooting
//      +      Output full error stack details, useful for debugging
//
// Examples:
//      %s:    error for internal read B
//      %v:    error for internal read B
//      %-v:   error for internal read B - #0 [/**/main.go:12 (main.main)] (#100102) Internal Server Error
//      %+v:   error for internal read B - #0 [/**/main.go:12 (main.main)] (#100102) Internal Server Error; error for internal read A - #1 [/**/main.go:35 (main.newErrorB)] (#100104) Validation failed
//      %#v:   [{"error":"error for internal read B"}]
//      %#-v:  [{"caller":"#0 /**/main.go:12 (main.main)","error":"error for internal read B","message":"(#100102) Internal Server Error"}]
//      %#+v:  [{"caller":"#0 /**/main.go:12 (main.main)","error":"error for internal read B","message":"(#100102) Internal Server Error"},{"caller":"#1 /**/main.go:35 (main.newErrorB)","error":"error for internal read A","message":"(#100104) Validation failed"}]
func (w *withCode) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		str := bytes.NewBuffer([]byte{})
		var jsonData []map[string]interface{}

		var (
			modeJSON   bool
			flagDetail bool
			flagTrace  bool
		)

		if state.Flag('#') {
			modeJSON = true
		}

		if state.Flag('-') {
			flagDetail = true
		}
		if state.Flag('+') {
			flagTrace = true
		}

		sep := ""
		errs := list(w)
		length := len(errs)
		for k, e := range errs {
			info := buildFormatInfo(e)
			jsonData, str = format(length-k-1, jsonData, str, info, sep, flagDetail, flagTrace, modeJSON)
			sep = "; "

			if !flagTrace {
				break
			}

			if !flagDetail && !flagTrace && !modeJSON {
				break
			}
		}
		if modeJSON {
			var bytes []byte
			bytes, _ = json.Marshal(jsonData)

			str.Write(bytes)
		}

		fmt.Fprintf(state, "%s", strings.Trim(str.String(), "\r\n\t"))
	default:
		info := buildFormatInfo(w)
		// Externally-safe error message
		fmt.Fprintf(state, info.message)
	}
}

func format(k int, jsonData []map[string]interface{}, str *bytes.Buffer, info *formatInfo,
	sep string, flagDetail, flagTrace, modeJSON bool) ([]map[string]interface{}, *bytes.Buffer) {
	if modeJSON {
		data := map[string]interface{}{}
		if flagDetail || flagTrace {
			data = map[string]interface{}{
				"message": info.message,
				"code":    info.code,
				"error":   info.err,
			}

			caller := fmt.Sprintf("#%d", k)
			if info.stack != nil {
				f := Frame((*info.stack)[0])
				caller = fmt.Sprintf("%s %s:%d (%s)",
					caller,
					f.file(),
					f.line(),
					f.name(),
				)
			}
			data["caller"] = caller
		} else {
			data["error"] = info.message
		}
		jsonData = append(jsonData, data)
	} else {
		if flagDetail || flagTrace {
			if info.stack != nil {
				f := Frame((*info.stack)[0])
				fmt.Fprintf(str, "%s%s - #%d [%s:%d (%s)] (%d) %s",
					sep,
					info.err,
					k,
					f.file(),
					f.line(),
					f.name(),
					info.code,
					info.message,
				)
			} else {
				fmt.Fprintf(str, "%s%s - #%d %s", sep, info.err, k, info.message)
			}

		} else {
			fmt.Fprintf(str, info.message)
		}
	}

	return jsonData, str
}

// list will convert the error stack into a simple array.
func list(e error) []error {
	var ret []error

	if e != nil {
		if w, ok := e.(interface{ Unwrap() error }); ok {
			ret = append(ret, e)
			ret = append(ret, list(w.Unwrap())...)
		} else {
			ret = append(ret, e)
		}
	}

	return ret
}

func buildFormatInfo(e error) *formatInfo {
	var info *formatInfo

	switch err := e.(type) {
	case *fundamental:
		info = &formatInfo{
			code:    unknownCoder.Code(),
			message: err.msg,
			err:     err.msg,
			stack:   err.stack,
		}
	case *withStack:
		info = &formatInfo{
			code:    unknownCoder.Code(),
			message: err.Error(),
			err:     err.Error(),
			stack:   err.stack,
		}
	case *withCode:
		coder, ok := codes[err.code]
		if !ok {
			coder = unknownCoder
		}

		extMsg := coder.String()
		if extMsg == "" {
			extMsg = err.err.Error()
		}

		info = &formatInfo{
			code:    coder.Code(),
			message: extMsg,
			err:     err.err.Error(),
			stack:   err.stack,
		}
	default:
		info = &formatInfo{
			code:    unknownCoder.Code(),
			message: err.Error(),
			err:     err.Error(),
		}
	}

	return info
}
