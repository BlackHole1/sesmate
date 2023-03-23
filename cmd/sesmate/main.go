package main

import (
	"log"
)

func main() {
	rootCmd.AddCommand(syncCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
