package trenovaorm

import (
	"errors"
	"fmt"
	"strings"
)

// CharField represents a string field in the database.
type CharField struct {
	ColumnName  string
	MaxLength   int
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

// Definition generates the SQL definition for the CharField.
func (f *CharField) Definition() string {
	typ := fmt.Sprintf(`VARCHAR(%d)`, f.MaxLength)
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

// Name returns the column name for the CharField.
func (f *CharField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the CharField.
func (f *CharField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *CharField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	if f.Nullable && f.Default != "" {
		return fmt.Errorf("CharField %s is nullable and has a default value", f.ColumnName)
	}

	// Ensure maxLength is positive.
	if f.MaxLength <= 0 {
		return errors.New("MaxLength must be positive")
	}
	return nil
}

// GoType returns the Go type for the CharField.
func (f *CharField) GoType() string {
	if f.Nullable || f.Blank {
		return "*string"
	}
	return "string"
}

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *CharField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("%s_%s_idx", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
}
