package naming

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// SnakeCaseName 命名規則を変換するために、基準となる規則としてスネークケースを利用する
type SnakeCaseName string

func NewSnakeCaseName(s string) (*SnakeCaseName, error) {
	if ok := IsValidSnakeCase(s); !ok {
		return nil, fmt.Errorf("invalid name, must snake case")
	}

	n := SnakeCaseName(s)
	return &n, nil
}

var snakeCasePattern = regexp.MustCompile(`^[a-z]+(_[a-z]+)*$`)

func IsValidSnakeCase(s string) bool {
	if ok := snakeCasePattern.MatchString(s); ok {
		return true
	}
	return false
}

func (n SnakeCaseName) SnakeCase() string {
	return string(n)
}

func (n SnakeCaseName) CamelCase() string {
	parts := strings.Split(string(n), "_")
	titleCaser := cases.Title(language.Und, cases.NoLower)
	for i := range parts {
		if i > 0 {
			parts[i] = titleCaser.String(parts[i])
		}
	}
	return strings.Join(parts, "")
}

func (n SnakeCaseName) PascalCase() string {
	parts := strings.Split(string(n), "_")
	titleCaser := cases.Title(language.Und)
	for i := range parts {
		parts[i] = titleCaser.String(parts[i])
	}
	return strings.Join(parts, "")
}

func (n SnakeCaseName) ConnectionCase() string {
	return strings.ReplaceAll(string(n), "_", "")
}

func (n SnakeCaseName) ConstantCase() string {
	upperCaser := cases.Upper(language.Und)
	return upperCaser.String(string(n))
}

func (n SnakeCaseName) KebabCase() string {
	return strings.ReplaceAll(string(n), "_", "-")
}
