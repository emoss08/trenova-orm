package trenovaorm

import "testing"

func TestUUIDField_Definition(t *testing.T) {
	tests := []struct {
		name     string
		field    UUIDField
		expected string
	}{
		{
			name: "Basic UUIDField",
			field: UUIDField{
				ColumnName: "id",
				Nullable:   false,
				Unique:     true,
				Default:    UUIDGenerateV4,
			},
			expected: `"id" uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4()`,
		},
		{
			name: "Nullable UUIDField",
			field: UUIDField{
				ColumnName: "id",
				Nullable:   true,
			},
			expected: `"id" uuid`,
		},
		{
			name: "UUIDField with Comment",
			field: UUIDField{
				ColumnName: "id",
				Nullable:   false,
				Unique:     true,
				Default:    UUIDGenerateV4,
				Comment:    "Primary key",
			},
			expected: `"id" uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4()`,
		},
		{
			name: "UUIDField with Custom Type",
			field: UUIDField{
				ColumnName: "id",
				Nullable:   false,
				Unique:     true,
				CustomType: "CHAR(36)",
				Default:    UUIDGenerateV4,
			},
			expected: `"id" CHAR(36) NOT NULL UNIQUE DEFAULT uuid_generate_v4()`,
		},
		{
			name: "UUIDField as Primary Key",
			field: UUIDField{
				ColumnName: "id",
				Nullable:   false,
				PrimaryKey: true,
				Default:    UUIDGenerateV4,
			},
			expected: `"id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4()`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.Definition()
			if got != tt.expected {
				t.Errorf("UUIDField.Definition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestUUIDField_Name(t *testing.T) {
	field := UUIDField{
		ColumnName: "id",
	}
	expected := "id"
	if got := field.Name(); got != expected {
		t.Errorf("UUIDField.Name() = %v, want %v", got, expected)
	}
}

func TestUUIDField_CommentSQL(t *testing.T) {
	field := UUIDField{
		ColumnName: "id",
		Comment:    "Primary key",
	}
	expected := `COMMENT ON COLUMN "users"."id" IS 'Primary key';`
	if got := field.CommentSQL("users"); got != expected {
		t.Errorf("UUIDField.CommentSQL() = %v, want %v", got, expected)
	}
}

func TestUUIDField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		field   UUIDField
		wantErr bool
	}{
		{
			name: "Valid UUIDField",
			field: UUIDField{
				ColumnName: "id",
			},
			wantErr: false,
		},
		{
			name: "Invalid UUIDField with empty ColumnName",
			field: UUIDField{
				ColumnName: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid UUIDField with PrimaryKey and Nullable",
			field: UUIDField{
				ColumnName: "id",
				Nullable:   true,
				PrimaryKey: true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("UUIDField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUUIDField_GoType(t *testing.T) {
	tests := []struct {
		name     string
		field    UUIDField
		expected string
	}{
		{
			name: "Non-Nullable UUIDField",
			field: UUIDField{
				ColumnName: "id",
				Nullable:   false,
			},
			expected: "uuid.UUID",
		},
		{
			name: "Nullable UUIDField",
			field: UUIDField{
				ColumnName: "id",
				Nullable:   true,
			},
			expected: "*uuid.UUID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.GoType(); got != tt.expected {
				t.Errorf("UUIDField.GoType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestUUIDField_IndexSQL(t *testing.T) {
	field := UUIDField{
		ColumnName: "id",
		Index:      true,
	}
	expected := `CREATE INDEX "idx_users_id" ON "users" ("id");`
	if got := field.IndexSQL("users"); got != expected {
		t.Errorf("UUIDField.IndexSQL() = %v, want %v", got, expected)
	}
}
