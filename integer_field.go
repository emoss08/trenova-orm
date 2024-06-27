package trenovaorm

import (
	"fmt"
	"strings"
)

// IntegerField represents an integer field in the database.
type IntegerField struct {
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

// Definition generates the SQL definition for the IntegerField.
func (f *IntegerField) Definition() string {
	typ := "INTEGER"
	if f.CustomType != "" {
		typ = f.CustomType
	}

	def := fmt.Sprintf(`"%s" %s`, f.ColumnName, typ)

	if !f.Nullable {
		def += fmt.Sprintf(" %s", ConstraintNotNull.String())
	}

	if f.Unique {
		def += fmt.Sprintf(" %s", ConstraintUnqiue.String())
	}

	if f.Default != 0 {
		def += fmt.Sprintf(" DEFAULT %d", f.Default)
	}

	if len(f.Constraints) > 0 {
		def += " " + strings.Join(f.Constraints, " ")
	}

	return def
}

// Name returns the column name for the IntegerField.
func (f *IntegerField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the IntegerField.
func (f *IntegerField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *IntegerField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	return nil
}

// GoType returns the Go type for the IntegerField.
func (f *IntegerField) GoType() string {
	if f.Nullable {
		return "*int"
	}
	return "int"
}

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *IntegerField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
}
