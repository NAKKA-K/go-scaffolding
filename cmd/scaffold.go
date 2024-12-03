package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	resource string
)

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
	fmt.Println("scaffold called")
	// テンプレートディレクトリの絶対パスを取得します
	absPath, err := filepath.Abs(config.Run.TemplateDir)
	if err != nil {
		log.Fatalf("Failed to determine absolute path: %v", err)
	}

	// Directory内のテンプレートファイル全てをロード
	tmpl, err := template.ParseGlob(filepath.Join(absPath, "*.tmpl"))
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	// テンプレートに埋め込むためのデータサンプル
	data := map[string]interface{}{
		"Title":   "Example Title",
		"Content": "This is an example content.",
	}

	// テンプレートを実行して標準出力に書き出す
	err = tmpl.ExecuteTemplate(os.Stdout, "your_template.tmpl", data)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	return nil
}
