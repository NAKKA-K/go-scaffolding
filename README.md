# go-scaffolding
Golangプロジェクトで使用するscaffoldingツール

使用ライブラリ
* text/template
* spf13/cobra
    * spf13/cobra-cli で雛形を生成


## use

.go-scaffolding.yaml
```
run:
  template-dir: 'template/'
  output:
    - 'translator.go.tmpl': "presentation/{{.SnakeCase}}/{{.SnakeCase}}.go"
    - 'model_to_entity_mapper.go.tmpl': "../presentation/{{.SnakeCase}}/mapper.go"
    - 'entity.go.tmpl': "../entity/{{.SnakeCase}}.go"
    - 'usecase.go.tmpl': "../usecase/{{.SnakeCase}}.go"
    - 'affiliation.go.tmpl': "../usecase/affiliation/{{.SnakeCase}}.go"
    - 'affiliation_test.go.tmpl': "../usecase/affiliation/{{.SnakeCase}}_test.go"
    - 'interactor_repository.go.tmpl': "../interactor/repository/{{.SnakeCase}}.go"
    - 'infrastructure_repository.go.tmpl': "../infrastructure/repository/{{.SnakeCase}}.go"
    - 'ent_to_entity_mapper.go.tmpl': "../infrastructure/mapper/{{.SnakeCase}}.go"
    - 'entity_to_model_mapper.go.tmpl': "../presentation/{{.SnakeCase}}.go"
```

```
make run RESOURCE=resource_snake_case
```

## docs

* SnakeCase: snake_case
* CamelCase: camelCase
* PascalCase: PascalCase
* ConnectionCase: 
* ConstantCase: CONSTANT_CASE
* KebabCase: kebab-case

## 展望
* 外から変数を渡せるようにする
