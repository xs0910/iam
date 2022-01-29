package term

import (
	"fmt"
	"io"

	"github.com/moby/term"
)

// TerminalSize returns the current width and height of the user's terminal.
// If it isn't a terminal, nil is returned.
func TerminalSize(w io.Writer) (int, int, error) {
	outFd, isTerminal := term.GetFdInfo(w)
	if !isTerminal {
		return 0, 0, fmt.Errorf("given writer is no terminal")
	}

	windowSize, err := term.GetWinsize(outFd)
	if err != nil {
		return 0, 0, err
	}
	return int(windowSize.Width), int(windowSize.Height), nil
}
