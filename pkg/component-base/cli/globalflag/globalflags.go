package globalflag

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"strings"
)

// AddGlobalFlags explicitly registers flags that libraries (log, ver flag, etc.) register
// against the global flagSets from "flag".
// We do this in order to prevent unwanted flags from leaking into the component's flagSets.
func AddGlobalFlags(fs *pflag.FlagSet, name string) {
	fs.BoolP("help", "h", false, fmt.Sprintf("help for %s", name))
}

// Register adds a flag to local that targets the value associated with the flag named globalName in flag.CommandLine.
func Register(local *pflag.FlagSet, globalName string) {
	if f := flag.CommandLine.Lookup(globalName); f != nil {
		pflagFlag := pflag.PFlagFromGoFlag(f)
		pflagFlag.Name = normalize(pflagFlag.Name)
		local.AddFlag(pflagFlag)
	} else {
		panic(fmt.Sprintf("failed to find flag in global flagset (flag): %s", globalName))
	}
}

func normalize(name string) string {
	return strings.ReplaceAll(name, "_", "-")
}
