package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/BlackHole1/sesmate/internal/server"
)

var serverCmd = &cobra.Command{
	Use: "server",
}

func init() {
	const (
		flagDir  = "dir"
		flagHost = "host"
		flagPort = "port"
	)

	var (
		dir, host string
		port      int
	)

	serverCmd.RunE = func(cmd *cobra.Command, args []string) error {
		s := server.New(dir, host, port)
		return s.Execute()
	}

	serverCmd.Flags().StringVar(&dir, flagDir, "", "SES template directory, automatically created if it does not exist")
	if err := serverCmd.MarkFlagDirname(flagDir); err != nil {
		log.Fatalln(err.Error())
	}

	serverCmd.Flags().StringVar(&host, flagHost, "127.0.0.1", "Server host")
	serverCmd.Flags().IntVar(&port, flagPort, 8091, "Server port")
}
