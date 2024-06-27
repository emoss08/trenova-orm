package trenovaorm

import (
	"errors"
	"fmt"
	"strings"
)

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
