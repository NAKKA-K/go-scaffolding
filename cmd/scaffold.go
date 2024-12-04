package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/NAKKA-K/go-scaffolding/internal/logging"
	"github.com/NAKKA-K/go-scaffolding/internal/naming"
	"github.com/spf13/cobra"
)

var (
	resource string
)

var caseNames *naming.CaseNames

// scaffoldCmd represents the scaffold command
var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "一番最初の雛形を生成する.",
	Long:  `新しいリソースのAPIを作り始めたい時に実行するコマンドです.`,
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

	// TODO: 外部依存に影響されて生成されるコードをテンプレートに反映させるためのフラグを追加する
	// 例えばGraphQLのスキーマから生成されたmodelや、DBのスキーマから生成されたentを利用する場合など
}

func executeScaffold(cmd *cobra.Command, args []string) error {
	// :=をするとcaseNamesがローカル変数扱いされてしまうので、再代入で書くために事前定義
	var err error
	caseNames, err = naming.NewCaseNames(resource)
	if err != nil {
		log.Fatalf("Failed to create naming: %v", err)
	}

	// テンプレートディレクトリの絶対パスを取得します
	absTemplateDir, err := filepath.Abs(config.Run.TemplateDir)
	if err != nil {
		log.Fatalf("Failed to determine absolute path: %v", err)
	}
	logging.Verbose(verbose, "Abs template directory: %s", absTemplateDir)

	for templateFileName, outputPath := range config.Run.Output {
		templateFilePath := filepath.Join(absTemplateDir, templateFileName)
		logging.Verbose(verbose, "Template file path: %s", templateFilePath)

		tmpl, err := template.ParseFiles(templateFilePath)
		if err != nil {
			log.Printf("Failed to parse template %s: %v", templateFilePath, err)
			continue
		}

		// ディレクトリ名やファイル名にリソース名が含まれることがあるので、`{resource}`をリソース名に置換する
		outputPath = strings.Replace(outputPath, "{resource}", caseNames.SnakeCase, -1)
		logging.Verbose(verbose, "Output path: %s", outputPath)

		// 出力先のディレクトリを生成する
		outputDir := filepath.Dir(outputPath)
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			log.Printf("Failed to create directory %s: %v", outputDir, err)
			continue
		}

		// 出力ファイルを作成する
		outputFile, err := os.Create(outputPath)
		if err != nil {
			log.Printf("Could not create output file %s: %v", outputPath, err)
			continue
		}
		defer outputFile.Close() // リソースがリークしている可能性があります。'defer' が 'for' ループで呼び出されています

		err = tmpl.Execute(outputFile, caseNames)
		if err != nil {
			log.Printf("Failed to execute template %s -> %s: %v", templateFilePath, outputPath, err)
		}

		fmt.Printf("Generated: \"%s\" -> \"%s\"\n", templateFilePath, outputPath)
	}

	logging.Verbose(verbose, "Complete.")

	return nil
}
