package trenovaorm

// Mixin interface for defining mixin fields.
type Mixin interface {
	Fields() []Field
}

// BaseMixin struct to be embedded in mixins.
type BaseMixin struct{}

// Fields returns an empty slice for base mixin.
func (BaseMixin) Fields() []Field {
	return []Field{}
}

// TimestampedMixin provides common fields for tracking creation and update times.
type TimestampedMixin struct {
	BaseMixin
}

// Fields returns the common timestamp fields.
func (t TimestampedMixin) Fields() []Field {
	return []Field{
		&DateField{
			ColumnName: "created_at",
			Nullable:   false,
			Default:    "CURRENT_TIMESTAMP",
			Comment:    "Creation timestamp",
			StructTag:  `json:"created_at" validate:"required"`,
		},
		&DateField{
			ColumnName: "updated_at",
			Nullable:   false,
			Default:    "CURRENT_TIMESTAMP",
			Comment:    "Update timestamp",
			StructTag:  `json:"updated_at" validate:"required"`,
		},
	}
}
