package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/BlackHole1/sesmate/internal/sync"
	"github.com/BlackHole1/sesmate/internal/utils"
)

var syncCmd = &cobra.Command{
	Use: "sync",
}

func init() {
	const (
		flagAK       = "ak"
		flagSK       = "sk"
		flagEndpoint = "endpoint"
		flagRegion   = "region"
		flagDir      = "dir"
		flagRemove   = "remove"
	)

	var (
		syncAK, syncSK, syncEndpoint, syncRegion, syncDirectory string
		syncRemove                                              bool
	)

	syncCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		utils.OptionsValueWithEnv(&syncAK, "AWS_AK", flagAK, cmd)
		utils.OptionsValueWithEnv(&syncSK, "AWS_SK", flagSK, cmd)
		utils.OptionsValueWithEnv(&syncEndpoint, "AWS_ENDPOINT", flagEndpoint, cmd)
		utils.OptionsValueWithEnv(&syncRegion, "AWS_REGION", flagRegion, cmd)

		return nil
	}
	syncCmd.RunE = func(cmd *cobra.Command, args []string) error {
		d := sync.New(syncAK, syncSK, syncEndpoint, syncRegion, syncDirectory, syncRemove)
		if err := d.Execute(); err != nil {
			log.Fatalln(err.Error())
		}

		return nil
	}

	syncCmd.Flags().StringVar(&syncAK, flagAK, "", "AWS access key OR use AWS_AK env")
	syncCmd.Flags().StringVar(&syncSK, flagSK, "", "AWS secret key OR use AWS_SK env")
	syncCmd.Flags().StringVar(&syncEndpoint, flagEndpoint, "", "AWS endpoint OR use AWS_ENDPOINT env")
	syncCmd.Flags().StringVar(&syncRegion, flagRegion, "", "AWS region OR use AWS_REGION env")
	syncCmd.Flags().StringVar(&syncDirectory, flagDir, "", "SES template directory")
	if err := syncCmd.MarkFlagRequired(flagDir); err != nil {
		log.Fatalln(err.Error())
	}
	if err := syncCmd.MarkFlagDirname(flagDir); err != nil {
		log.Fatalln(err.Error())
	}

	syncCmd.Flags().BoolVar(&syncRemove, flagRemove, false, "Delete remote template when it is not found locally.")

}
