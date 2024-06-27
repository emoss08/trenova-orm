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

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *CharField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("%s_%s_idx", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
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

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *IntegerField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
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

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *BooleanField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
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

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *NumericField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
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

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *TextField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
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
	typ := "uuid"
	if f.CustomType != "" {
		typ = f.CustomType
	}
	def := fmt.Sprintf(`"%s" %s`, f.ColumnName, typ)

	if !f.Nullable {
		def += fmt.Sprintf(" %s", ConstraintNotNull.String())
	}

	if f.PrimaryKey {
		def += fmt.Sprintf(" %s", ConstraintPrimaryKey.String())
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
	if f.ColumnName == "" {
		return errors.New("column name cannot be empty")
	}
	if f.PrimaryKey && f.Nullable {
		return errors.New("primary key field cannot be nullable")
	}
	return nil
}

func (f *UUIDField) GoType() string {
	if f.Nullable {
		return "*uuid.UUID"
	}
	return "uuid.UUID"
}

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *UUIDField) IndexSQL(tableName string) string {
	if !f.Index {
		return ""
	}
	indexName := fmt.Sprintf("idx_%s_%s", tableName, f.ColumnName)
	return fmt.Sprintf(`CREATE INDEX "%s" ON "%s" ("%s");`, indexName, tableName, f.ColumnName)
}

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
	def += " CHECK (" + f.ColumnName + " > 0)"

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
