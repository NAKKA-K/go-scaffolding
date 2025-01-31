package tmpl

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/NAKKA-K/go-scaffolding/internal/logging"
)

type Embedder struct {
	variablesMap map[string]string
	verbose      bool
}

func NewEmbedder(variablesMap map[string]string, verbose bool) *Embedder {
	return &Embedder{
		variablesMap: variablesMap,
		verbose:      verbose,
	}
}

// GenerateByStrTmpl は、指定されたテンプレート文字列を元に、埋め込みデータを使って文字列を生成します.
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

// WriteFileByTemplate は、指定されたテンプレートファイルを元に、指定された出力先にファイルを生成します.
func (e *Embedder) WriteFileByTemplate(templateFilePath string, outputPath string) error {
	logging.Verbose(e.verbose, "Template file path: %s", templateFilePath)

	tpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templateFilePath, err)
	}

	// 出力先のディレクトリを生成する
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", outputDir, err)
	}

	// 出力ファイルを作成する
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("could not create output file %s: %w", outputPath, err)
	}
	defer outputFile.Close()

	// TODO: caseNamesだけではなく、埋め込み用のデータを増やす. ex) {{.GqlModel1}}.{{.Resource.PascalCase}} -> "resourcetable1.ResourcePascalCase"
	// テンプレートを元にデータを埋め込み、ファイルを生成する
	err = tpl.Execute(outputFile, e.variablesMap)
	if err != nil {
		return fmt.Errorf("failed to execute template %s -> %s: %w", templateFilePath, outputPath, err)
	}

	return nil
}
