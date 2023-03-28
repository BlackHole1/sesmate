package main

import (
	"log"
)

func main() {
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.Version = Version

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
