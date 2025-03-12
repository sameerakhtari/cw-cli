package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cw-cli",
	Short: "cw-cli is a CLI tool for managing Cloudways services",
	Long:  `cw-cli provides a command-line interface to manage servers, applications, DNS, and more on Cloudways.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
