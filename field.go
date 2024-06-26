package trenovaorm

import (
	"errors"
	"fmt"
	"strings"
)

type Constraint string

const (
	ConstraintUnqiue     = Constraint("UNIQUE")
	ConstraintNotNull    = Constraint("NOT NULL")
	ConstraintCheck      = Constraint("CHECK")
	ConstraintPrimaryKey = Constraint("PRIMARY KEY")
	ConstraintDefault    = Constraint("DEFAULT")
)

func (c Constraint) String() string {
	return string(c)
}

// Field defines a general database field interface.
type Field interface {
	Definition() string
	Name() string
	Validate() error
	CommentSQL(tableName string) string
	GoType() string
}

// BaseField provides common properties for all fields.
type BaseField struct {
	ColumnName  string
	Nullable    bool
	Unique      bool
	Index       bool
	Comment     string
	CustomType  string
	Constraints []string
	StructTag   string
}

func (f *BaseField) Name() string {
	return f.ColumnName
}

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
	return nil
}

// GoType returns the Go type for the IntegerField.
func (f *IntegerField) GoType() string {
	if f.Nullable {
		return "*int"
	}
	return "int"
}

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
	return nil
}

// GoType returns the Go type for the BooleanField.
func (f *BooleanField) GoType() string {
	if f.Nullable {
		return "*bool"
	}
	return "bool"
}

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
		def += fmt.Sprintf(" DEFAULT '%s'", f.Default.String())
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
	return nil
}

// GoType returns the Go type for the DateField.
func (f *DateField) GoType() string {
	if f.Nullable {
		return "*time.Time"
	}
	return "time.Time"
}

// NumericField represents a numeric field in the database.
type NumericField struct {
	ColumnName  string
	Precision   int
	Scale       int
	Nullable    bool
	Unique      bool
	Default     string
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
	if f.Default != "" {
		def += fmt.Sprintf(" DEFAULT %s", f.Default)
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
	// Example validation: Ensure precision and scale are positive and precision is greater than or equal to scale.
	if f.Precision <= 0 || f.Scale < 0 || f.Precision < f.Scale {
		return fmt.Errorf("Invalid precision or scale for NumericField: precision %d, scale %d", f.Precision, f.Scale)
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
	return nil
}

// GoType returns the Go type for the TextField.
func (f *TextField) GoType() string {
	if f.Nullable {
		return "*string"
	}
	return "string"
}

// UUIDField represents a UUID field in the database.
type UUIDField struct {
	ColumnName  string
	Nullable    bool
	Blank       bool
	Unique      bool
	Default     PSQLFunction
	Index       bool
	Comment     string
	CustomType  string
	Constraints []string
	PrimaryKey  bool
	StructTag   string
}

func (f *UUIDField) Definition() string {
	def := fmt.Sprintf(`"%s" UUID`, f.ColumnName)

	if !f.Nullable {
		def += fmt.Sprintf(" %s", ConstraintNotNull.String())
	}

	if f.Unique {
		def += fmt.Sprintf(" %s", ConstraintUnqiue.String())
	}

	if f.Default != "" {
		def += fmt.Sprintf(" DEFAULT %s", f.Default.String())
	}

	if f.PrimaryKey {
		def += fmt.Sprintf(" %s", ConstraintPrimaryKey.String())
	}

	if len(f.Constraints) > 0 {
		def += " " + strings.Join(f.Constraints, " ")
	}

	return def
}

func (f *UUIDField) Name() string {
	return f.ColumnName
}

func (f *UUIDField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

func (f *UUIDField) Validate() error {
	return nil
}

func (f *UUIDField) GoType() string {
	if f.Nullable {
		return "*uuid.UUID"
	}
	return "uuid.UUID"
}
