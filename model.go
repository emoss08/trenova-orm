package trenovaorm

// Model interface for defining entity models.
type Model interface {
	Fields() []Field
	Indexes() []Index
	TableName() string
	Mixins() []Mixin
}

// BaseModel provides common fields and methods for all models.
type BaseModel struct {
	tableName string
}

// Mixins returns the mixins for the model. This method should be overridden by concrete models if they have mixins.
func (b *BaseModel) Mixins() []Mixin {
	return []Mixin{}
}

// Fields returns the fields defined by the model, including mixin fields.
func (b *BaseModel) Fields() []Field {
	return []Field{}
}

// Indexes returns an empty slice for base model.
func (b *BaseModel) Indexes() []Index {
	return []Index{}
}

// TableName returns the table name for the model.
func (b *BaseModel) TableName() string {
	if b.tableName == "" {
		panic("Table name not set for schema")
	}
	return b.tableName
}

// SetTableName allows overriding the default table name.
func (b *BaseModel) SetTableName(name string) {
	b.tableName = name
}
