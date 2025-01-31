package tmpl

import (
	"bytes"
	"fmt"
	"text/template"
)

type Embedder struct {
	variablesMap map[string]string
}

func NewEmbedder(variablesMap map[string]string) *Embedder {
	return &Embedder{
		variablesMap: variablesMap,
	}
}

func (e *Embedder) GenerateByStrTmpl(textTmpl string) (*string, error) {
	// テンプレートを解析
	t, err := template.New("text").Parse(textTmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	// bytes.Bufferを使用してテンプレート出力をバッファに書き込み
	var buf bytes.Buffer
	if err := t.Execute(&buf, e.variablesMap); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	// バッファの内容を文字列として取得
	res := buf.String()
	return &res, nil
}
