package cmd

import (
	"github.com/lbernardo/generator/pkg/entity"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var rootPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "generator",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootPath, "root", ".", "rootPath")
}

func AddCommand(module *entity.Module) {
	rootCmd.AddCommand(module.Command())
}
