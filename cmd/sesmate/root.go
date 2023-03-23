package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "sesmate",
}

func init() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	rootCmd.ValidArgs = []string{"sync"}
}
