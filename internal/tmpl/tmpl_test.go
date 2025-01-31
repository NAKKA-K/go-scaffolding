package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GenerateByStrTmplのテストを作る
// テストの内容は、GenerateByStrTmplの引数に渡した文字列が、埋め込みデータを使って正しく埋め込まれるかどうかを確認する
// 例えば、"Hello, {{.Name}}"という文字列に、{"Name": "World"}という埋め込みデータを使って、"Hello, World"という文字列が生成されるかどうかを確認する
func TestEmbedder_GenerateByStrTmpl(t *testing.T) {
	// テストケースの定義
	testCases := []struct {
		name         string
		textTmpl     string
		variablesMap map[string]string
		want         string
	}{
		{
			name:         "正常に生成される",
			textTmpl:     "usecase/{{.SnakeCase}}.go",
			variablesMap: map[string]string{"SnakeCase": "resource"},
			want:         "usecase/resource.go",
		},
	}

	// テストケースの数だけ繰り返す
	for _, tt := range testCases {
		// テストケースの名前を表示
		t.Run(tt.name, func(t *testing.T) {
			// テストケースの期待値と実際の値が一致しているかどうかを確認する
			// 一致していない場合は、テストを失敗させる
			embedder := NewEmbedder(tt.variablesMap, false)
			got, err := embedder.GenerateByStrTmpl(tt.textTmpl)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, *got)
		})
	}
}
