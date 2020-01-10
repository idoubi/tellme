package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cliVersion = "0.3.1"
)

var rootCmd = &cobra.Command{
	Use:     "tellme",
	Version: cliVersion,
	Run: func(cmd *cobra.Command, args []string) {
		showVersion()
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "version of tellme")
}

// Execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func showVersion() {
	fmt.Printf("tellme version %s\n", cliVersion)
}
