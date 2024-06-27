package trenovaorm

import (
	"fmt"
	"strings"
)

// PositiveIntegerField represents a positive integer field in the database.
type PositiveIntegerField struct {
	ColumnName  string
	Nullable    bool
	Unique      bool
	Default     int
	Index       bool
	Comment     string
	CustomType  string
	Constraints []string
	StructTag   string
}

// Definition generates the SQL definition for the PositiveIntegerField.
func (f *PositiveIntegerField) Definition() string {
	typ := "INTEGER"
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

	if f.Default != 0 {
		def += fmt.Sprintf(" DEFAULT %d", f.Default)
	}

	// Adding check constraint for positive integers
	def += fmt.Sprintf(" CHECK (%s > 0)", f.ColumnName)

	if len(f.Constraints) > 0 {
		def += " " + strings.Join(f.Constraints, " ")
	}

	return def
}

// Name returns the column name for the PositiveIntegerField.
func (f *PositiveIntegerField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the PositiveIntegerField.
func (f *PositiveIntegerField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *PositiveIntegerField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	if f.Default < 0 {
		return fmt.Errorf("default value for positive integer field must be positive")
	}

	return nil
}

// GoType returns the Go type for the PositiveIntegerField.
func (f *PositiveIntegerField) GoType() string {
	if f.Nullable {
		return "*int"
	}
	return "int"
}

// IndexSQL generates the SQL statement for creating an index if Index or Unique is true.
func (f *PositiveIntegerField) IndexSQL(tableName string) string {
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
