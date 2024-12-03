package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	resource   string
	configFile string
)

var (
	dir string
)

// scaffoldCmd represents the scaffold command
var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "一番最初の雛形を生成する.",
	Long:  `新しいリソースのAPIを作り始めたい時に実行するコマンドです.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("scaffold called")

		// TODO: リソース名を元にファイルを生成する処理を書く

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)

	// go-scaffolding scaffold -f configFile.yaml -r companion_ad
	scaffoldCmd.Flags().StringVarP(&configFile, "configFile-file", "f", "", "設定ファイルを指定します.")
	if err := scaffoldCmd.MarkFlagRequired("configFile"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scaffoldCmd.Flags().StringVarP(&resource, "resource", "r", "", "リソース名を指定します.")
	if err := scaffoldCmd.MarkFlagRequired("resource"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// TODO: 外部依存に影響されて生成されるコードをテンプレートに反映させるためのフラグを追加する
	// 例えばGraphQLのスキーマから生成されたmodelや、DBのスキーマから生成されたentを利用する場合など

}
