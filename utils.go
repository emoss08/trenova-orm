package trenovaorm

import (
	"fmt"
	"reflect"
	"strings"
)

// Helper function to join and quote columns
func joinColumns(columns []string) string {
	quoted := make([]string, len(columns))
	for i, col := range columns {
		quoted[i] = quoteIdentifier(col)
	}
	return strings.Join(quoted, ", ")
}

// Helper function to quote PostgreSQL identifiers
func quoteIdentifier(identifier string) string {
	return fmt.Sprintf(`"%s"`, identifier)
}

// toSnakeCase converts a string to snake_case.
func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

// getTypeName uses reflection to derive the struct name that embeds BaseSchema.
func getTypeName(b *BaseSchema) string {
	// Iterate through the pointer chain to find the non-pointer type
	for typ := reflect.TypeOf(b); typ.Kind() == reflect.Ptr; typ = typ.Elem() {
		if typ.Kind() == reflect.Struct {
			// Check if this type is the base type or derived type
			if typ.Name() == "BaseSchema" && typ.PkgPath() == reflect.TypeOf(BaseSchema{}).PkgPath() {
				continue
			}
			return typ.Name()
		}
	}
	return "base_schema"
}
