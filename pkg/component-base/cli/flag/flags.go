package flag

import (
	goflag "flag"
	"github.com/spf13/pflag"
	"log"
	"strings"
)

func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(goflag.CommandLine)
}

// WordSepNormalizeFunc changes all flags that contain "_ separators
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and warns for flags that contain "_" separators.
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		result := strings.ReplaceAll(name, "_", "-")
		log.Printf("%s is DEPECATED and will be removed in a future version. use %s instead.", name, result)
		return pflag.NormalizedName(result)
	}
	return pflag.NormalizedName(name)
}

// PrintFlags logs the flags in the FlagSet.
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}
