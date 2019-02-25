package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "kmode",
}

var filename string
var varFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&filename, "filename", "filename.lua", "kmode filename")
	rootCmd.PersistentFlags().StringVar(&varFile, "var-file", "", "variable file")
	rootCmd.AddCommand(outputCmd)
	rootCmd.AddCommand(applyCmd)
}

// Execute is the entry point for a kmode
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
