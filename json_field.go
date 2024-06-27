package trenovaorm

import (
	"fmt"
	"strings"
)

// JSONField represents a JSON field in the database.
type JSONField struct {
	ColumnName  string
	Nullable    bool
	Unique      bool
	Default     string
	Index       bool
	Comment     string
	CustomType  string
	Constraints []string
	StructTag   string
}

// Definition generates the SQL definition for the JSONField.
func (f *JSONField) Definition() string {
	typ := "JSONB"
	if f.CustomType != "" {
		typ = f.CustomType
	}
	def := fmt.Sprintf(`"%s" %s`, f.ColumnName, typ)

	if !f.Nullable {
		def += " NOT NULL"
	}

	if f.Unique {
		def += " UNIQUE"
	}

	if f.Default != "" {
		def += fmt.Sprintf(" DEFAULT '%s'", f.Default)
	}

	if len(f.Constraints) > 0 {
		def += " " + strings.Join(f.Constraints, " ")
	}

	return def
}

// Name returns the column name for the JSONField.
func (f *JSONField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the JSONField.
func (f *JSONField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *JSONField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}
	return nil
}

// GoType returns the Go type for the JSONField.
func (f *JSONField) GoType() string {
	if f.Nullable {
		return "*map[string]any"
	}
	return "map[string]any"
}

// IndexSQL generates the SQL statement for creating an index if Index or Unique is true.
func (f *JSONField) IndexSQL(tableName string) string {
	if !f.Index && !f.Unique {
		return ""
	}
	indexType := "INDEX"
	if f.Unique {
		indexType = "UNIQUE INDEX"
	}
	indexName := fmt.Sprintf("%s_%s_idx", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE %s "%s" ON "%s" ("%s");`, indexType, indexName, tableName, f.ColumnName)
}
