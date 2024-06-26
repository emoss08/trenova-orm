package trenovaorm

// Schema interface for defining entity schemas.
type Schema interface {
	Fields() []Field
	Indexes() []Index
	TableName() string
	Mixins() []Mixin
}

// BaseSchema provides common fields and methods for all schemas.
type BaseSchema struct {
	tableName string
}

// Mixins returns the mixins for the schema. This method should be overridden by concrete schemas if they have mixins.
func (b *BaseSchema) Mixins() []Mixin {
	return []Mixin{}
}

// Fields returns the fields defined by the schema, including mixin fields.
func (b *BaseSchema) Fields() []Field {
	return []Field{}
}

// Indexes returns an empty slice for base schema.
func (b *BaseSchema) Indexes() []Index {
	return []Index{}
}

// TableName returns the table name for the schema.
func (b *BaseSchema) TableName() string {
	if b.tableName == "" {
		panic("Table name not set for schema")
	}
	return b.tableName
}

// SetTableName allows overriding the default table name.
func (b *BaseSchema) SetTableName(name string) {
	b.tableName = name
}
