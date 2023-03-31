package main

import (
	"log"

	"github.com/BlackHole1/sesmate/internal/version"
)

func main() {
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.Version = version.Version

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
