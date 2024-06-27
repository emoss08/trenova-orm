package trenovaorm

import (
	"fmt"
	"strings"
)

// OnDeleteOption defines the possible options for the ON DELETE behavior.
type OnDeleteOption string

const (
	OnDeleteCascade  OnDeleteOption = "CASCADE"
	OnDeleteSetNull  OnDeleteOption = "SET NULL"
	OnDeleteRestrict OnDeleteOption = "RESTRICT"
	OnDeleteNoAction OnDeleteOption = "NO ACTION"
)

// OnUpdateOption defines the possible options for the ON UPDATE behavior.
type OnUpdateOption string

const (
	OnUpdateCascade  OnUpdateOption = "CASCADE"
	OnUpdateSetNull  OnUpdateOption = "SET NULL"
	OnUpdateRestrict OnUpdateOption = "RESTRICT"
	OnUpdateNoAction OnUpdateOption = "NO ACTION"
)

// Annotation represents the annotations for foreign key relationships.
type Annotation struct {
	OnDelete OnDeleteOption
	OnUpdate OnUpdateOption
}

func (a Annotation) String() string {
	return fmt.Sprintf("ON DELETE %s ON UPDATE %s", a.OnDelete, a.OnUpdate)
}

// ForeignKeyField represents a foreign key field in the database.
type ForeignKeyField struct {
	ColumnName     string
	ReferenceTable string
	ReferenceField string
	Annotations    Annotation
	Nullable       bool
	Unique         bool
	Default        string
	Index          bool
	Comment        string
	CustomType     string
	Constraints    []string
	StructTag      string
	ReferencedType string // The Go type of the referenced field
}

// Definition generates the SQL definition for the ForeignKeyField.
func (f *ForeignKeyField) Definition() string {
	typ := "INTEGER" // Default type for foreign keys
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

// ForeignKeyConstraint generates the SQL for the foreign key constraint.
func (f *ForeignKeyField) ForeignKeyConstraint(tableName string) string {
	constraint := fmt.Sprintf(`FOREIGN KEY ("%s") REFERENCES "%s"("%s")`, f.ColumnName, f.ReferenceTable, f.ReferenceField)
	if f.Annotations.OnDelete != "" {
		constraint += fmt.Sprintf(" ON DELETE %s", f.Annotations.OnDelete)
	}
	if f.Annotations.OnUpdate != "" {
		constraint += fmt.Sprintf(" ON UPDATE %s", f.Annotations.OnUpdate)
	}
	return constraint
}

// Name returns the column name for the ForeignKeyField.
func (f *ForeignKeyField) Name() string {
	return f.ColumnName
}

// CommentSQL generates the SQL statement for adding a comment to the ForeignKeyField.
func (f *ForeignKeyField) CommentSQL(tableName string) string {
	if f.Comment == "" {
		return ""
	}
	return fmt.Sprintf(`COMMENT ON COLUMN "%s"."%s" IS '%s';`, tableName, f.ColumnName, f.Comment)
}

// Validate checks if the field's configuration is valid.
func (f *ForeignKeyField) Validate() error {
	if f.ColumnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}
	if f.ReferenceTable == "" || f.ReferenceField == "" {
		return fmt.Errorf("references cannot be empty")
	}
	return nil
}

// GoType returns the Go type for the ForeignKeyField.
func (f *ForeignKeyField) GoType() string {
	if f.Nullable {
		return fmt.Sprintf("*%s", f.ReferencedType)
	}
	return f.ReferencedType
}

// IndexSQL generates the SQL statement for creating an index if Index is true.
func (f *ForeignKeyField) IndexSQL(tableName string) string {
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
