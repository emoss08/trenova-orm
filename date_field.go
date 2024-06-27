package trenovaorm

import (
	"fmt"
	"strings"
)

// DateField represents a date field in the database.
type DateField struct {
	ColumnName  string
	Nullable    bool
	Unique      bool
	Default     PSQLFunction
	Index       bool
	Comment     string
	CustomType  string
	Constraints []string
	StructTag   string
}

// Definition generates the SQL definition for the DateField.
func (f *DateField) Definition() string {
	typ := "DATE"
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

	if f.Default != "" {
		def += fmt.Sprintf(" DEFAULT %s", f.Default.String())
	}

	if len(f.Constraints) > 0 {
		def += " " + strings.Join(f.Constraints, " ")
	}
	return def
}

// Name returns the column name for the DateField.
func (f *DateField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the DateField.
func (f *DateField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *DateField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	return nil
}

// GoType returns the Go type for the DateField.
func (f *DateField) GoType() string {
	if f.Nullable {
		return "*time.Time"
	}
	return "time.Time"
}

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *DateField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
}
