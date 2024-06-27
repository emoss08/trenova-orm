package trenovaorm

import "testing"

func TestTimeField_Definition(t *testing.T) {
	tests := []struct {
		name     string
		field    TimeField
		expected string
	}{
		{
			name: "Basic TimeField",
			field: TimeField{
				ColumnName: "time_value",
				Nullable:   false,
				Unique:     true,
				Default:    CurrentTimestamp,
			},
			expected: `"time_value" TIME NOT NULL UNIQUE DEFAULT current_timestamp`,
		},
		{
			name: "Nullable TimeField",
			field: TimeField{
				ColumnName: "time_value",
				Nullable:   true,
			},
			expected: `"time_value" TIME`,
		},
		{
			name: "TimeField with Comment",
			field: TimeField{
				ColumnName: "time_value",
				Nullable:   false,
				Unique:     true,
				Default:    CurrentTimestamp,
				Comment:    "Time value",
			},
			expected: `"time_value" TIME NOT NULL UNIQUE DEFAULT current_timestamp`,
		},
		{
			name: "TimeField with Custom Type",
			field: TimeField{
				ColumnName: "time_value",
				Nullable:   false,
				Unique:     true,
				CustomType: "TIMESTAMP",
				Default:    CurrentTimestamp,
			},
			expected: `"time_value" TIMESTAMP NOT NULL UNIQUE DEFAULT current_timestamp`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.Definition()
			if got != tt.expected {
				t.Errorf("TimeField.Definition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTimeField_Name(t *testing.T) {
	field := TimeField{
		ColumnName: "time_value",
	}
	expected := "time_value"
	if got := field.Name(); got != expected {
		t.Errorf("TimeField.Name() = %v, want %v", got, expected)
	}
}

func TestTimeField_CommentSQL(t *testing.T) {
	field := TimeField{
		ColumnName: "time_value",
		Comment:    "Time value",
	}
	expected := `COMMENT ON COLUMN "users"."time_value" IS 'Time value';`
	if got := field.CommentSQL("users"); got != expected {
		t.Errorf("TimeField.CommentSQL() = %v, want %v", got, expected)
	}
}

func TestTimeField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		field   TimeField
		wantErr bool
	}{
		{
			name: "Valid TimeField",
			field: TimeField{
				ColumnName: "time_value",
				Default:    CurrentTimestamp,
			},
			wantErr: false,
		},
		{
			name: "Invalid TimeField with empty ColumnName",
			field: TimeField{
				ColumnName: "",
				Default:    CurrentTimestamp,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("TimeField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimeField_GoType(t *testing.T) {
	tests := []struct {
		name     string
		field    TimeField
		expected string
	}{
		{
			name: "Non-Nullable TimeField",
			field: TimeField{
				ColumnName: "time_value",
				Nullable:   false,
			},
			expected: "TimeOnly",
		},
		{
			name: "Nullable TimeField",
			field: TimeField{
				ColumnName: "time_value",
				Nullable:   true,
			},
			expected: "*TimeOnly",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.GoType(); got != tt.expected {
				t.Errorf("TimeField.GoType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTimeField_IndexSQL(t *testing.T) {
	tests := []struct {
		name     string
		field    TimeField
		table    string
		expected string
	}{
		{
			name: "TimeField with Index",
			field: TimeField{
				ColumnName: "time_value",
				Index:      true,
			},
			table:    "users",
			expected: `CREATE INDEX "users_time_value_idx" ON "users" ("time_value");`,
		},
		{
			name: "TimeField with Unique Index",
			field: TimeField{
				ColumnName: "time_value",
				Unique:     true,
			},
			table:    "users",
			expected: `CREATE UNIQUE INDEX "users_time_value_idx" ON "users" ("time_value");`,
		},
		{
			name: "TimeField without Index",
			field: TimeField{
				ColumnName: "time_value",
			},
			table:    "users",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.IndexSQL(tt.table); got != tt.expected {
				t.Errorf("TimeField.IndexSQL() = %v, want %v", got, tt.expected)
			}
		})
	}
}
