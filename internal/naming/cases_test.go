package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCaseNames(t *testing.T) {
	tests := []struct {
		s    string
		want *CaseNames
		err  error
	}{
		{
			s: "snake_case",
			want: &CaseNames{
				SnakeCase:      "snake_case",
				CamelCase:      "snakeCase",
				PascalCase:     "SnakeCase",
				ConnectionCase: "snakecase",
				ConstantCase:   "SNAKE_CASE",
				KebabCase:      "snake-case",
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			n, err := NewCaseNames(tt.s)
			assert.Equal(t, n, tt.want)
			assert.Equal(t, err, tt.err)
		})
	}
}

func TestCaseNames_ToMap(t *testing.T) {
	tests := []struct {
		name  string
		input CaseNames
		want  map[string]string
	}{
		{
			name: "正常にmapが生成される",
			input: CaseNames{
				SnakeCase:      "snake_case",
				CamelCase:      "snakeCase",
				PascalCase:     "SnakeCase",
				ConnectionCase: "snakecase",
				ConstantCase:   "SNAKE_CASE",
				KebabCase:      "snake-case",
			},
			want: map[string]string{
				"SnakeCase":      "snake_case",
				"CamelCase":      "snakeCase",
				"PascalCase":     "SnakeCase",
				"ConnectionCase": "snakecase",
				"ConstantCase":   "SNAKE_CASE",
				"KebabCase":      "snake-case",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ToMap(); !assert.Equal(t, got, tt.want) {
				t.Errorf("Name.ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidSnakeCase_NoErr(t *testing.T) {
	tests := []struct {
		s string
	}{
		{
			s: "valid_snake_case", // Normal SnakeCase
		},
		{
			s: "validsnakecase", // No underscores
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			ok := IsValidSnakeCase(tt.s)
			assert.True(t, ok)
		})
	}
}

func TestIsValidSnakeCase_Err(t *testing.T) {
	tests := []struct {
		s string
	}{
		{
			s: "invalid__case", // Consecutive underscores
		},
		{
			s: "_invalid_start", // Leading underscore
		},
		{
			s: "invalid_end_", // Trailing underscore
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			ok := IsValidSnakeCase(tt.s)
			assert.False(t, ok)
		})
	}
}

func TestSnakeCaseName_CamelCase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "camel_case",
			want: "camelCase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := SnakeCaseName(tt.name)
			if got := n.CamelCase(); got != tt.want {
				t.Errorf("Name.CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCaseName_PascalCase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "pascal_case",
			want: "PascalCase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := SnakeCaseName(tt.name)
			if got := n.PascalCase(); got != tt.want {
				t.Errorf("Name.PascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCaseName_ConnectionCase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "connection_case",
			want: "connectioncase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := SnakeCaseName(tt.name)
			if got := n.ConnectionCase(); got != tt.want {
				t.Errorf("Name.ConnectionCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCaseName_ConstantCase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "constant_case",
			want: "CONSTANT_CASE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := SnakeCaseName(tt.name)
			if got := n.ConstantCase(); got != tt.want {
				t.Errorf("Name.ConstantCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCaseName_KebabCase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "kebab_case",
			want: "kebab-case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := SnakeCaseName(tt.name)
			if got := n.KebabCase(); got != tt.want {
				t.Errorf("Name.KebabCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
