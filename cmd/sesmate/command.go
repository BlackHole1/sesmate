package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/BlackHole1/sesmate/internal/sync"
)

var ak, sk, endpoint, region, directory string
var remove bool

var rootCmd = &cobra.Command{
	Use: "sesmate",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		optionsValueWithEnv(&ak, "AWS_AK", flagAK, cmd)
		optionsValueWithEnv(&sk, "AWS_SK", flagSK, cmd)
		optionsValueWithEnv(&endpoint, "AWS_ENDPOINT", flagEndpoint, cmd)
		optionsValueWithEnv(&region, "AWS_REGION", flagRegion, cmd)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
	ValidArgs: []string{"sync"},
}

var syncCmd = &cobra.Command{
	Use: "sync",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := requireValue(&directory, flagDir, cmd); err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		d := sync.New(ak, sk, endpoint, region, directory, remove)
		if err := d.Execute(); err != nil {
			log.Fatalln(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().StringVar(&ak, flagAK, "", "AWS access key OR use AWS_AK env")
	syncCmd.Flags().StringVar(&sk, flagSK, "", "AWS secret key OR use AWS_SK env")
	syncCmd.Flags().StringVar(&endpoint, flagEndpoint, "", "AWS endpoint OR use AWS_ENDPOINT env")
	syncCmd.Flags().StringVar(&region, flagRegion, "", "AWS region OR use AWS_REGION env")
	syncCmd.Flags().StringVar(&directory, flagDir, "", "SES template directory")
	if err := syncCmd.MarkFlagRequired(flagDir); err != nil {
		log.Fatalln(err.Error())
	}
	if err := syncCmd.MarkFlagDirname(flagDir); err != nil {
		log.Fatalln(err.Error())
	}

	syncCmd.Flags().BoolVar(&remove, flagRemove, false, "Delete remote template when it is not found locally.")

}

func requireValue(p *string, flagName string, cmd *cobra.Command) error {
	var val string

	if flag := cmd.Flag(flagName); flag != nil && flag.Value.String() != "" {
		val = flag.Value.String()
	}

	if val == "" {
		return fmt.Errorf("missing flag %s", flagName)
	}

	*p = val

	return nil
}

func requireValueWithEnv(p *string, envName, flagName string, cmd *cobra.Command) error {
	var val string

	if v, ok := os.LookupEnv(envName); ok && v != "" {
		val = v
	}

	if flag := cmd.Flag(flagName); flag != nil && flag.Value.String() != "" {
		val = flag.Value.String()
	}

	if val == "" {
		return fmt.Errorf("missing env %s or flag %s", envName, flagName)
	}

	*p = val

	return nil
}

func optionsValueWithEnv(p *string, envName, flagName string, cmd *cobra.Command) {
	if err := requireValueWithEnv(p, envName, flagName, cmd); err != nil {
		*p = ""
	}
}
