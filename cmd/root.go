package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type CommandConfig struct {
	TemplateDir string            `mapstructure:"template-dir"`
	Output      map[string]string `mapstructure:"output"`
}

// flags
var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-scaffolding",
	Short: "'go-scaffolding' is a CLI tool to generate golang project scaffolding",
	Long:  `'go-scaffolding' is a CLI tool to generate golang project scaffolding.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

func init() {
	// global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .go-scaffolding.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "config file (default is .go-scaffolding.yaml)")

	// local flags
}
