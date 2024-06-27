package trenovaorm

import (
	"testing"
)

func TestForeignKeyField_Definition(t *testing.T) {
	tests := []struct {
		name     string
		field    ForeignKeyField
		expected string
	}{
		{
			name: "Basic ForeignKeyField",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				ReferenceTable: "users",
				ReferenceField: "id",
				Nullable:       false,
				Unique:         true,
				Default:        "1",
				Annotations: Annotation{
					OnDelete: OnDeleteCascade,
					OnUpdate: OnUpdateCascade,
				},
			},
			expected: `"user_id" INTEGER NOT NULL UNIQUE DEFAULT '1'`,
		},
		{
			name: "Nullable ForeignKeyField",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				ReferenceTable: "users",
				ReferenceField: "id",
				Nullable:       true,
			},
			expected: `"user_id" INTEGER`,
		},
		{
			name: "ForeignKeyField with Comment",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				ReferenceTable: "users",
				ReferenceField: "id",
				Nullable:       false,
				Unique:         true,
				Default:        "1",
				Comment:        "Foreign key to users table",
				Annotations: Annotation{
					OnDelete: OnDeleteCascade,
					OnUpdate: OnUpdateCascade,
				},
			},
			expected: `"user_id" INTEGER NOT NULL UNIQUE DEFAULT '1'`,
		},
		{
			name: "ForeignKeyField with Custom Type",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				ReferenceTable: "users",
				ReferenceField: "id",
				Nullable:       false,
				Unique:         true,
				CustomType:     "BIGINT",
				Default:        "1",
				Annotations: Annotation{
					OnDelete: OnDeleteCascade,
					OnUpdate: OnUpdateCascade,
				},
			},
			expected: `"user_id" BIGINT NOT NULL UNIQUE DEFAULT '1'`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.Definition()
			if got != tt.expected {
				t.Errorf("ForeignKeyField.Definition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestForeignKeyField_Name(t *testing.T) {
	field := ForeignKeyField{
		ColumnName: "user_id",
	}
	expected := "user_id"
	if got := field.Name(); got != expected {
		t.Errorf("ForeignKeyField.Name() = %v, want %v", got, expected)
	}
}

func TestForeignKeyField_CommentSQL(t *testing.T) {
	field := ForeignKeyField{
		ColumnName: "user_id",
		Comment:    "Foreign key to users table",
	}
	expected := `COMMENT ON COLUMN "orders"."user_id" IS 'Foreign key to users table';`
	if got := field.CommentSQL("orders"); got != expected {
		t.Errorf("ForeignKeyField.CommentSQL() = %v, want %v", got, expected)
	}
}

func TestForeignKeyField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		field   ForeignKeyField
		wantErr bool
	}{
		{
			name: "Valid ForeignKeyField",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				ReferenceTable: "users",
				ReferenceField: "id",
				Default:        "1",
			},
			wantErr: false,
		},
		{
			name: "Invalid ForeignKeyField with empty ColumnName",
			field: ForeignKeyField{
				ColumnName:     "",
				ReferenceTable: "users",
				ReferenceField: "id",
				Default:        "1",
			},
			wantErr: true,
		},
		{
			name: "Invalid ForeignKeyField with empty References",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				ReferenceTable: "",
				ReferenceField: "",
				Default:        "1",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ForeignKeyField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestForeignKeyField_GoType(t *testing.T) {
	tests := []struct {
		name     string
		field    ForeignKeyField
		expected string
	}{
		{
			name: "Non-Nullable ForeignKeyField",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				Nullable:       false,
				ReferencedType: "User",
			},
			expected: "User",
		},
		{
			name: "Nullable ForeignKeyField",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				Nullable:       true,
				ReferencedType: "User",
			},
			expected: "*User",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.GoType(); got != tt.expected {
				t.Errorf("ForeignKeyField.GoType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestForeignKeyField_IndexSQL(t *testing.T) {
	tests := []struct {
		name     string
		field    ForeignKeyField
		table    string
		expected string
	}{
		{
			name: "ForeignKeyField with Index",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				Index:          true,
				ReferenceTable: "users",
				ReferenceField: "id",
			},
			table:    "orders",
			expected: `CREATE INDEX "orders_user_id_idx" ON "orders" ("user_id");`,
		},
		{
			name: "ForeignKeyField with Unique Index",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				Unique:         true,
				ReferenceTable: "users",
				ReferenceField: "id",
			},
			table:    "orders",
			expected: `CREATE UNIQUE INDEX "orders_user_id_idx" ON "orders" ("user_id");`,
		},
		{
			name: "ForeignKeyField without Index",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				ReferenceTable: "users",
				ReferenceField: "id",
			},
			table:    "orders",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.IndexSQL(tt.table); got != tt.expected {
				t.Errorf("ForeignKeyField.IndexSQL() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestForeignKeyField_ForeignKeyConstraint(t *testing.T) {
	tests := []struct {
		name     string
		field    ForeignKeyField
		table    string
		expected string
	}{
		{
			name: "Basic ForeignKeyField constraint",
			field: ForeignKeyField{
				ColumnName:     "user_id",
				ReferenceTable: "users",
				ReferenceField: "id",
				Annotations: Annotation{
					OnDelete: OnDeleteCascade,
					OnUpdate: OnUpdateCascade,
				},
			},
			table:    "orders",
			expected: `FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.ForeignKeyConstraint(tt.table); got != tt.expected {
				t.Errorf("ForeignKeyField.ForeignKeyConstraint() = %v, want %v", got, tt.expected)
			}
		})
	}
}
