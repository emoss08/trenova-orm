package trenovaorm

import (
	"fmt"
	"strings"
)

// NumericField represents a numeric field in the database.
type NumericField struct {
	ColumnName  string
	Precision   int
	Scale       int
	Nullable    bool
	Unique      bool
	Default     float64
	Index       bool
	Comment     string
	CustomType  string
	Constraints []string
	StructTag   string
}

// Definition generates the SQL definition for the NumericField.
func (f *NumericField) Definition() string {
	typ := fmt.Sprintf(`NUMERIC(%d, %d)`, f.Precision, f.Scale)
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
		def += fmt.Sprintf(" DEFAULT %.*f", f.Scale, f.Default)
	}

	if len(f.Constraints) > 0 {
		def += " " + strings.Join(f.Constraints, " ")
	}

	return def
}

// Name returns the column name for the NumericField.
func (f *NumericField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the NumericField.
func (f *NumericField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *NumericField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	// Ensure precision and scale are positive and precision is greater than or equal to scale.
	if f.Precision <= 0 || f.Scale < 0 || f.Precision < f.Scale {
		return fmt.Errorf("invalid precision or scale for NumericField: precision %d, scale %d", f.Precision, f.Scale)
	}

	// Check the precision and scale of the default value
	defaultStr := fmt.Sprintf("%.*f", f.Scale, f.Default)
	if len(defaultStr)-1 > f.Precision+1 {
		return fmt.Errorf("default value %f exceeds defined precision %d and scale %d", f.Default, f.Precision, f.Scale)
	}

	return nil
}

// GoType returns the Go type for the NumericField.
func (f *NumericField) GoType() string {
	if f.Nullable {
		return "*float64"
	}
	return "float64"
}

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *NumericField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
}
