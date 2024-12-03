package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		os.Exit(1)
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
		verboseLog("set config file", cfgFile)
		workDir, err := os.Getwd()
		cobra.CheckErr(err)

		f := filepath.Join(workDir, cfgFile)

		viper.SetConfigFile(f)
	} else {
		// デフォルト挙動として作業ディレクトリから設定ファイルを探す
		verboseLog("default config file", cfgFile)
		workDir, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(workDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-scaffolding")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// 設定ファイルを読み込む
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
	verboseLog("config", config)
}

func verboseLog(a ...any) {
	if !verbose {
		return
	}

	fmt.Print("[LOG]: ")
	fmt.Println(a...)
}
