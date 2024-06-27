package trenovaorm

import "testing"

func TestNumericField_Definition(t *testing.T) {
	tests := []struct {
		name     string
		field    NumericField
		expected string
	}{
		{
			name: "Basic NumericField",
			field: NumericField{
				ColumnName: "value",
				Precision:  10,
				Scale:      2,
				Nullable:   false,
				Unique:     true,
				Default:    123.45,
			},
			expected: `"value" NUMERIC(10, 2) NOT NULL UNIQUE DEFAULT 123.45`,
		},
		{
			name: "Nullable NumericField",
			field: NumericField{
				ColumnName: "value",
				Precision:  10,
				Scale:      2,
				Nullable:   true,
			},
			expected: `"value" NUMERIC(10, 2)`,
		},
		{
			name: "NumericField with Comment",
			field: NumericField{
				ColumnName: "value",
				Precision:  10,
				Scale:      2,
				Nullable:   false,
				Unique:     true,
				Default:    123.45,
				Comment:    "Numeric value",
			},
			expected: `"value" NUMERIC(10, 2) NOT NULL UNIQUE DEFAULT 123.45`,
		},
		{
			name: "NumericField with Custom Type",
			field: NumericField{
				ColumnName: "value",
				Precision:  10,
				Scale:      2,
				Nullable:   false,
				Unique:     true,
				CustomType: "DECIMAL(10, 2)",
				Default:    123.45,
			},
			expected: `"value" DECIMAL(10, 2) NOT NULL UNIQUE DEFAULT 123.45`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.Definition()
			if got != tt.expected {
				t.Errorf("NumericField.Definition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNumericField_Name(t *testing.T) {
	field := NumericField{
		ColumnName: "value",
	}
	expected := "value"
	if got := field.Name(); got != expected {
		t.Errorf("NumericField.Name() = %v, want %v", got, expected)
	}
}

func TestNumericField_CommentSQL(t *testing.T) {
	field := NumericField{
		ColumnName: "value",
		Comment:    "Numeric value",
	}
	expected := `COMMENT ON COLUMN "users"."value" IS 'Numeric value';`
	if got := field.CommentSQL("users"); got != expected {
		t.Errorf("NumericField.CommentSQL() = %v, want %v", got, expected)
	}
}

func TestNumericField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		field   NumericField
		wantErr bool
	}{
		{
			name: "Valid NumericField",
			field: NumericField{
				ColumnName: "value",
				Precision:  10,
				Scale:      2,
				Default:    123.45,
			},
			wantErr: false,
		},
		{
			name: "Invalid NumericField with empty ColumnName",
			field: NumericField{
				ColumnName: "",
				Precision:  10,
				Scale:      2,
				Default:    123.45,
			},
			wantErr: true,
		},
		{
			name: "Invalid NumericField with Negative Scale",
			field: NumericField{
				ColumnName: "value",
				Precision:  10,
				Scale:      -1,
			},
			wantErr: true,
		},
		{
			name: "Invalid NumericField with Default Exceeding Precision and Scale",
			field: NumericField{
				ColumnName: "value",
				Precision:  5,
				Scale:      2,
				Default:    123456.78,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("NumericField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNumericField_GoType(t *testing.T) {
	tests := []struct {
		name     string
		field    NumericField
		expected string
	}{
		{
			name: "Non-Nullable NumericField",
			field: NumericField{
				ColumnName: "value",
				Nullable:   false,
			},
			expected: "float64",
		},
		{
			name: "Nullable NumericField",
			field: NumericField{
				ColumnName: "value",
				Nullable:   true,
			},
			expected: "*float64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.GoType(); got != tt.expected {
				t.Errorf("NumericField.GoType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNumericField_IndexSQL(t *testing.T) {
	tests := []struct {
		name     string
		field    NumericField
		table    string
		expected string
	}{
		{
			name: "NumericField with Index",
			field: NumericField{
				ColumnName: "value",
				Index:      true,
			},
			table:    "users",
			expected: `CREATE INDEX "idx_users_value" ON "users" ("value");`,
		},
		{
			name: "NumericField without Index",
			field: NumericField{
				ColumnName: "value",
			},
			table:    "users",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.IndexSQL(tt.table); got != tt.expected {
				t.Errorf("NumericField.IndexSQL() = %v, want %v", got, tt.expected)
			}
		})
	}
}
