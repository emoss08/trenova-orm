package trenovaorm

import (
	"fmt"
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
