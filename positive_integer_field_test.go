package trenovaorm

import "testing"

func TestPositiveIntegerField_Definition(t *testing.T) {
	tests := []struct {
		name     string
		field    PositiveIntegerField
		expected string
	}{
		{
			name: "Basic PositiveIntegerField",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Nullable:   false,
				Unique:     true,
				Default:    1,
			},
			expected: `"positive_value" INTEGER NOT NULL UNIQUE DEFAULT 1 CHECK (positive_value > 0)`,
		},
		{
			name: "Nullable PositiveIntegerField",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Nullable:   true,
			},
			expected: `"positive_value" INTEGER CHECK (positive_value > 0)`,
		},
		{
			name: "PositiveIntegerField with Comment",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Nullable:   false,
				Unique:     true,
				Default:    1,
				Comment:    "Positive integer value",
			},
			expected: `"positive_value" INTEGER NOT NULL UNIQUE DEFAULT 1 CHECK (positive_value > 0)`,
		},
		{
			name: "PositiveIntegerField with Custom Type",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Nullable:   false,
				Unique:     true,
				CustomType: "BIGINT",
				Default:    1,
			},
			expected: `"positive_value" BIGINT NOT NULL UNIQUE DEFAULT 1 CHECK (positive_value > 0)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.Definition()
			if got != tt.expected {
				t.Errorf("PositiveIntegerField.Definition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPositiveIntegerField_Name(t *testing.T) {
	field := PositiveIntegerField{
		ColumnName: "positive_value",
	}
	expected := "positive_value"
	if got := field.Name(); got != expected {
		t.Errorf("PositiveIntegerField.Name() = %v, want %v", got, expected)
	}
}

func TestPositiveIntegerField_CommentSQL(t *testing.T) {
	field := PositiveIntegerField{
		ColumnName: "positive_value",
		Comment:    "Positive integer value",
	}
	expected := `COMMENT ON COLUMN "users"."positive_value" IS 'Positive integer value';`
	if got := field.CommentSQL("users"); got != expected {
		t.Errorf("PositiveIntegerField.CommentSQL() = %v, want %v", got, expected)
	}
}

func TestPositiveIntegerField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		field   PositiveIntegerField
		wantErr bool
	}{
		{
			name: "Valid PositiveIntegerField",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Default:    1,
			},
			wantErr: false,
		},
		{
			name: "Invalid PositiveIntegerField with empty ColumnName",
			field: PositiveIntegerField{
				ColumnName: "",
				Default:    1,
			},
			wantErr: true,
		},
		{
			name: "Invalid PositiveIntegerField with Negative Default",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Default:    -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("PositiveIntegerField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPositiveIntegerField_GoType(t *testing.T) {
	tests := []struct {
		name     string
		field    PositiveIntegerField
		expected string
	}{
		{
			name: "Non-Nullable PositiveIntegerField",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Nullable:   false,
			},
			expected: "int",
		},
		{
			name: "Nullable PositiveIntegerField",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Nullable:   true,
			},
			expected: "*int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.GoType(); got != tt.expected {
				t.Errorf("PositiveIntegerField.GoType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPositiveIntegerField_IndexSQL(t *testing.T) {
	tests := []struct {
		name     string
		field    PositiveIntegerField
		table    string
		expected string
	}{
		{
			name: "PositiveIntegerField with Index",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Index:      true,
			},
			table:    "users",
			expected: `CREATE INDEX "users_positive_value_idx" ON "users" ("positive_value");`,
		},
		{
			name: "PositiveIntegerField with Unique Index",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
				Unique:     true,
			},
			table:    "users",
			expected: `CREATE UNIQUE INDEX "users_positive_value_idx" ON "users" ("positive_value");`,
		},
		{
			name: "PositiveIntegerField without Index",
			field: PositiveIntegerField{
				ColumnName: "positive_value",
			},
			table:    "users",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.IndexSQL(tt.table); got != tt.expected {
				t.Errorf("PositiveIntegerField.IndexSQL() = %v, want %v", got, tt.expected)
			}
		})
	}
}
