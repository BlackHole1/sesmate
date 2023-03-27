package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/BlackHole1/sesmate/internal/gen"
	"github.com/BlackHole1/sesmate/pkg/char"
)

var genCmd = &cobra.Command{
	Use: "gen",
}

func init() {
	const (
		flagDir         = "dir"
		flagOutput      = "output"
		flagFilename    = "filename"
		flagPackageName = "package-name"
		flagPrefix      = "prefix"
		flagCharCase    = "case"
	)

	var (
		dir         string
		output      string
		filename    string
		packageName string
		prefix      string
		charCase    = char.CasePascal
	)

	genCmd.RunE = func(cmd *cobra.Command, args []string) error {
		g := gen.New(dir, output, filename, packageName, prefix, charCase)

		return g.Execute()
	}

	genCmd.Flags().StringVar(&dir, flagDir, "", "SES template directory")
	if err := genCmd.MarkFlagRequired(flagDir); err != nil {
		log.Fatalln(err.Error())
	}
	if err := genCmd.MarkFlagDirname(flagDir); err != nil {
		log.Fatalln(err.Error())
	}

	genCmd.Flags().StringVar(&output, flagOutput, "./sestemplate", "Output directory")
	genCmd.Flags().StringVar(&filename, flagFilename, "name", "Output go filename")
	genCmd.Flags().StringVar(&packageName, flagPackageName, "sestemplate", "GO package name")
	genCmd.Flags().StringVar(&prefix, flagPrefix, "", "Prefix of generated const")
	genCmd.Flags().VarP(&charCase, flagCharCase, "", "Case of generated const, allowed values: lower, upper, camel, pascal, snake, screaming_snake, capitalized_snake")
}
