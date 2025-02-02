package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/NAKKA-K/go-scaffolding/internal/logging"
	"github.com/NAKKA-K/go-scaffolding/internal/naming"
	"github.com/NAKKA-K/go-scaffolding/internal/tmpl"
)

var (
	resource string
)

var commandConfig CommandConfig
var caseNames *naming.CaseNames

// scaffoldCmd represents the scaffold command
var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "一番最初の雛形を生成する.",
	Long:  `新しいリソースのAPIを作り始めたい時に実行するコマンドです.`,
	Args:  cobra.ExactArgs(1), // 1つの引数を必須にする
	RunE:  executeScaffold,
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)

	// go-scaffolding scaffold -f configFile.yaml -r companion_ad
	scaffoldCmd.Flags().StringVarP(&resource, "resource", "r", "", "リソース名を指定します.")
	if err := scaffoldCmd.MarkFlagRequired("resource"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// local flags

	// TODO: 外部依存に影響されて生成されるコードをテンプレートに反映させるためのフラグを追加する
	// 例えばGraphQLのスキーマから生成されたmodelや、DBのスキーマから生成されたentを利用する場合など
}

func executeScaffold(cmd *cobra.Command, args []string) error {
	initConfig(args)

	// :=をするとcaseNamesがローカル変数扱いされてしまうので、再代入で書くために事前定義
	var err error
	caseNames, err = naming.NewCaseNames(resource)
	if err != nil {
		log.Fatalf("Failed to create naming: %v", err)
	}

	// テンプレートディレクトリの絶対パスを取得します
	absTemplateDir, err := filepath.Abs(commandConfig.TemplateDir)
	if err != nil {
		log.Fatalf("Failed to determine absolute path: %v", err)
	}
	logging.Verbose(verbose, "Abs template directory: %s\n", absTemplateDir)

	embedder := tmpl.NewEmbedder(caseNames.ToMap(), verbose)
	for templateFileName, outputPathTmpl := range commandConfig.Output {
		// ディレクトリ名やファイル名に動的な名前が含まれることがあるので置換する
		outputPath, err := embedder.GenerateByStrTmpl(outputPathTmpl)
		if err != nil || outputPath == nil {
			log.Printf("Failed to generate output path by output path template %s: %v", outputPathTmpl, err)
			continue
		}
		logging.Verbose(verbose, "Output path: %s", *outputPath)

		templateFilePath := filepath.Join(absTemplateDir, templateFileName)
		if err := embedder.WriteFileByTemplate(templateFilePath, *outputPath); err != nil {
			log.Printf("Failed to write file by template %s -> %s: %v", templateFilePath, *outputPath, err)
			continue
		}

		fmt.Printf("Generated: \"%s\" -> \"%s\"\n", templateFilePath, *outputPath)
	}

	logging.Verbose(verbose, "Complete.")

	return nil
}

func validArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, but got %d", len(args))
	}

	return nil
}

func initConfig(args []string) {
	if err := validArgs(args); err != nil {
		log.Fatalf("Invalid arguments: %v", err)
	}
	sectionName := args[0]

	setConfig()
	viper.AutomaticEnv() // read in environment variables that match

	// 設定ファイルを読み込む
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error: ", err)
	}
	logging.Verbose(verbose, "Using config file: %s", viper.ConfigFileUsed())

	// 指定されたセクションを取得
	if viper.IsSet(sectionName) {
		err := viper.UnmarshalKey(sectionName, &commandConfig)
		if err != nil {
			log.Fatalf("failed to unmarshal key %s: %v", sectionName, err)
		}
		fmt.Printf("Section %s found: %+v\n", sectionName, commandConfig)
	} else {
		log.Fatalf("Section %s not found in config yaml\n", sectionName)
	}
	logging.Verbose(verbose, "config: %+v", commandConfig)
}

func setConfig() {
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
}
