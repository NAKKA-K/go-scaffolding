package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/NAKKA-K/go-scaffolding/internal/logging"
)

type Config struct {
	Run struct {
		TemplateDir string            `mapstructure:"template-dir"`
		Output      map[string]string `mapstructure:"output"`
	}
}

var config Config

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
	cobra.OnInitialize(initConfig)

	// global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .go-scaffolding.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "config file (default is .go-scaffolding.yaml)")

	// local flags
}

func initConfig() {
	if cfgFile != "" {
		// オプションで渡された設定ファイルを探す
		logging.Verbose(verbose, "set config file: %s", cfgFile)
		workDir, err := os.Getwd()
		cobra.CheckErr(err)

		f := filepath.Join(workDir, cfgFile)

		viper.SetConfigFile(f)
	} else {
		// デフォルト挙動として作業ディレクトリから設定ファイルを探す
		logging.Verbose(verbose, "default config file: %s", cfgFile)
		workDir, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(workDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-scaffolding")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// 設定ファイルを読み込む
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error: ", err)
	}
	logging.Verbose(verbose, "Using config file: %s", viper.ConfigFileUsed())

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln("Error: ", err)
	}
	logging.Verbose(verbose, "config: %+v", config)
}
