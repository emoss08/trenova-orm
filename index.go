package trenovaorm

import (
	"errors"
	"fmt"
	"strings"
)

// Index defines the structure for database indices, supporting both simple and complex cases.
type Index struct {
	Name        string       // Index name
	Columns     []string     // Simple column names
	Expressions []Expression // Custom SQL expressions as Expression interface
	Unique      bool         // Whether the index is unique
}

// generateName generates an index name based on the table and column names.
func (idx *Index) generateName(tableName string) string {
	if idx.Name != "" {
		return idx.Name
	}

	var colNames []string

	colNames = append(colNames, idx.Columns...)

	for _, exp := range idx.Expressions {
		colNames = append(colNames, exp.ColumnName())
	}

	return fmt.Sprintf("%s_%s_idx", tableName, strings.Join(colNames, "_"))
}

// Validate checks the integrity of the Index struct.
func (idx *Index) Validate() error {
	if len(idx.Columns) == 0 && len(idx.Expressions) == 0 {
		return errors.New("at least one column or expression must be specified")
	}
	return nil
}

// SQL generates the SQL statement for creating the index on a given table.
func (idx *Index) SQL(tableName string) (string, error) {
	if err := idx.Validate(); err != nil {
		return "", err
	}

	// Generate index name if not provided
	idx.Name = idx.generateName(tableName)

	uniqueness := ""
	if idx.Unique {
		uniqueness = "UNIQUE "
	}

	var parts []string
	for _, col := range idx.Columns {
		parts = append(parts, quoteIdentifier(col))
	}
	for _, exp := range idx.Expressions {
		parts = append(parts, exp.Expression())
	}

	expressions := strings.Join(parts, ", ")
	return fmt.Sprintf(`CREATE %sINDEX IF NOT EXISTS "%s" ON "%s" (%s);`, uniqueness, idx.Name, tableName, expressions), nil
}
