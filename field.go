package trenovaorm

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
