# go-scaffolding
※このライブラリは開発中によりアグレッシブに変更される可能性があります。

Golangプロジェクトで使用するscaffoldingツール

使用ライブラリ
* text/template
* spf13/cobra
    * spf13/cobra-cli で雛形を生成


## Usage

### config
`.go-scaffolding.yaml`
```
api:
  template-dir: 'template/api/'
  output:
    - 'model_to_entity_mapper.go.tmpl': "../presentation/{{.SnakeCase}}/mapper.go"
    - 'entity.go.tmpl': "../entity/{{.SnakeCase}}.go"
    - 'entity_test.go.tmpl': "../entity/{{.SnakeCase}}.go"
    - 'usecase.go.tmpl': "../usecase/{{.SnakeCase}}/{{.SnakeCase}}.go"
    - 'api_update_test.go.tmpl': "../test/{{.SnakeCase}}/create_{{.SnakeCase}}_test.go"
    - 'api_create_test.go.tmpl': "../test/{{.SnakeCase}}/update_{{.SnakeCase}}_test.go"

api_test:
  template-dir: 'template/api_test/'
  output:
    - 'api_update_test.go.tmpl': "../test/{{.SnakeCase}}/create_{{.SnakeCase}}_test.go"
    - 'api_create_test.go.tmpl': "../test/{{.SnakeCase}}/update_{{.SnakeCase}}_test.go"
```

### run
```
go-scaffolding scaffold api -r resource_snake_case
go-scaffolding scaffold api_test -r resource_snake_case
```

## Settings

### Variables in template
コマンドの`-r`オプションで渡したリソース名がそれぞれの記法で展開されます。

* `{{.SnakeCase}}`: resource `snake_case`
* `{{.CamelCase}}`: resource `camelCase`
* `{{.PascalCase}}`: resource `PascalCase`
* `{{.ConnectionCase}}`: resource `connectioncase`
* `{{.ConstantCase}}`: resource `CONSTANT_CASE`
* `{{.KebabCase}}`: resource `kebab-case`

### config
scaffoldingするためのテンプレートファイルの置き場所と出力先を指定します。
ファイルの出力先には、`{{.SnakeCase}}`などの変数を使用することができます。

サブコマンドより下位の階層にキーを受け取り、そのキーに対応する設定を読み込みます。
これによりユーザーは任意のキーごと設定でscaffoldingが可能です。

## Features

### 外から変数を渡せるようにする

現状、コマンドのオプションでリソース名を受け取り、リソース名の形式違いのみテンプレートに展開しています。
これを拡張し、外部から変数を渡すことで、任意の変数をテンプレートに展開できるようにします。
外部から変数を渡す方法は、オプションや設定ファイルから読み込むなどが考えられます。

## Contributing

### development

```shell
make test
make fix
make o # テスト実行
```

### release
gitでタグを打ち、GitHubにpushすれば自動でgoreleaserのGitHubActionsが発火しリリースされます。

```
git tag v0.3.0
git push --tags
```
