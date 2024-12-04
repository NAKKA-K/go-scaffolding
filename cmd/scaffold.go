package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	logging.Verbose(verbose, "Abs template directory: %s\n", absTemplateDir)

	for templateFileName, outputPathTmpl := range config.Run.Output {
		// ディレクトリ名やファイル名に動的な名前が含まれることがあるので置換する
		outputPath, err := generateStrByTemplate(outputPathTmpl, caseNames)
		if err != nil || outputPath == nil {
			log.Printf("Failed to generate output path by output path template %s: %v", outputPathTmpl, err)
			continue
		}
		logging.Verbose(verbose, "Output path: %s", *outputPath)

		templateFilePath := filepath.Join(absTemplateDir, templateFileName)
		if err := writeFileByTemplate(templateFilePath, *outputPath); err != nil {
			log.Printf("Failed to write file by template %s -> %s: %v", templateFilePath, *outputPath, err)
			continue
		}

		fmt.Printf("Generated: \"%s\" -> \"%s\"\n", templateFilePath, *outputPath)
	}

	logging.Verbose(verbose, "Complete.")

	return nil
}

func generateStrByTemplate(textTmpl string, data any) (*string, error) {
	// テンプレートを解析
	t, err := template.New("text").Parse(textTmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %v", err)
	}

	// bytes.Bufferを使用してテンプレート出力をバッファに書き込み
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %v", err)
	}

	// バッファの内容を文字列として取得
	res := buf.String()
	return &res, nil
}

func writeFileByTemplate(templateFilePath string, outputPath string) error {
	logging.Verbose(verbose, "Template file path: %s", templateFilePath)

	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %v", templateFilePath, err)
	}

	// 出力先のディレクトリを生成する
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", outputDir, err)
	}

	// 出力ファイルを作成する
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("could not create output file %s: %v", outputPath, err)
	}
	defer outputFile.Close()

	// テンプレートを元にデータを埋め込み、ファイルを生成する
	err = tmpl.Execute(outputFile, caseNames)
	if err != nil {
		return fmt.Errorf("failed to execute template %s -> %s: %v", templateFilePath, outputPath, err)
	}

	return nil
}
