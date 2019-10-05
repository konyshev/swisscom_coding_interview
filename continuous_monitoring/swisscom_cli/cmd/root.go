package cmd

import (
	"fmt"
	"os"

	. "github.com/konyshev/swisscom_cli/config"
	"github.com/spf13/cobra"
)

var config = Config{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "swisscom_cli",
	Short: "Simple cli client for swisscom comodir service",
	Long:  `Simple cli client for swisscom comodir service`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	config.Read()
}
