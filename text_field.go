package trenovaorm

import (
	"fmt"
	"strings"
)

// TextField represents a text field in the database.
type TextField struct {
	ColumnName  string
	Nullable    bool
	Blank       bool
	Unique      bool
	Default     string
	Index       bool
	Comment     string
	CustomType  string
	Constraints []string
	StructTag   string
}

// Definition generates the SQL definition for the TextField.
func (f *TextField) Definition() string {
	typ := "TEXT"
	if f.CustomType != "" {
		typ = f.CustomType
	}
	def := fmt.Sprintf(`"%s" %s`, f.ColumnName, typ)

	if !f.Blank && !f.Nullable {
		def += fmt.Sprintf(" %s", ConstraintNotNull.String())
	}

	if f.Unique {
		def += fmt.Sprintf(" %s", ConstraintUnqiue.String())
	}

	if f.Default != "" {
		def += fmt.Sprintf(" DEFAULT '%s'", f.Default)
	}

	if len(f.Constraints) > 0 {
		def += " " + strings.Join(f.Constraints, " ")
	}

	return def
}

// Name returns the column name for the TextField.
func (f *TextField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the TextField.
func (f *TextField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *TextField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	return nil
}

// GoType returns the Go type for the TextField.
func (f *TextField) GoType() string {
	if f.Nullable {
		return "*string"
	}
	return "string"
}

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *TextField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
}
