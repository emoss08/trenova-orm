package orm

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
