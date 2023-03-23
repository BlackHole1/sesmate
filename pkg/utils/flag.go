package utils

import (
	"os"

	"github.com/spf13/cobra"
)

func OptionsValueWithEnv(p *string, envName, flagName string, cmd *cobra.Command) {
	var val string

	if v, ok := os.LookupEnv(envName); ok && v != "" {
		val = v
	}

	if flag := cmd.Flag(flagName); flag != nil && flag.Value.String() != "" {
		val = flag.Value.String()
	}

	if val != "" {
		*p = val
	}
}
