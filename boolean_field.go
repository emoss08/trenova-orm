package trenovaorm

import (
	"fmt"
	"strings"
)

// BooleanField represents a boolean field in the database.
type BooleanField struct {
	ColumnName  string
	Nullable    bool
	Unique      bool
	Default     bool
	Index       bool
	Comment     string
	CustomType  string
	Constraints []string
	StructTag   string
}

// Definition generates the SQL definition for the BooleanField.
func (f *BooleanField) Definition() string {
	typ := "BOOLEAN"
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
	if f.Default {
		def += " DEFAULT TRUE"
	} else {
		def += " DEFAULT FALSE"
	}
	if len(f.Constraints) > 0 {
		def += " " + strings.Join(f.Constraints, " ")
	}
	return def
}

// Name returns the column name for the BooleanField.
func (f *BooleanField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the BooleanField.
func (f *BooleanField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *BooleanField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	return nil
}

// GoType returns the Go type for the BooleanField.
func (f *BooleanField) GoType() string {
	if f.Nullable {
		return "*bool"
	}
	return "bool"
}

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *BooleanField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
}
