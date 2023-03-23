package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/BlackHole1/sesmate/internal/sync"
	"github.com/BlackHole1/sesmate/pkg/utils"
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
		aK, sK, endpoint, region, dir string
		remove                        bool
	)

	syncCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		utils.OptionsValueWithEnv(&aK, "AWS_AK", flagAK, cmd)
		utils.OptionsValueWithEnv(&sK, "AWS_SK", flagSK, cmd)
		utils.OptionsValueWithEnv(&endpoint, "AWS_ENDPOINT", flagEndpoint, cmd)
		utils.OptionsValueWithEnv(&region, "AWS_REGION", flagRegion, cmd)

		return nil
	}
	syncCmd.RunE = func(cmd *cobra.Command, args []string) error {
		d := sync.New(aK, sK, endpoint, region, dir, remove)
		return d.Execute()
	}

	syncCmd.Flags().StringVar(&aK, flagAK, "", "AWS access key OR use AWS_AK env")
	syncCmd.Flags().StringVar(&sK, flagSK, "", "AWS secret key OR use AWS_SK env")
	syncCmd.Flags().StringVar(&endpoint, flagEndpoint, "", "AWS endpoint OR use AWS_ENDPOINT env")
	syncCmd.Flags().StringVar(&region, flagRegion, "", "AWS region OR use AWS_REGION env")
	syncCmd.Flags().StringVar(&dir, flagDir, "", "SES template directory")
	if err := syncCmd.MarkFlagRequired(flagDir); err != nil {
		log.Fatalln(err.Error())
	}
	if err := syncCmd.MarkFlagDirname(flagDir); err != nil {
		log.Fatalln(err.Error())
	}

	syncCmd.Flags().BoolVar(&remove, flagRemove, false, "Delete remote template when it is not found locally.")

}
