package trenovaorm

import (
	"testing"
)

func TestCharField_Definition(t *testing.T) {
	tests := []struct {
		name     string
		field    CharField
		expected string
	}{
		{
			name: "Basic CharField",
			field: CharField{
				ColumnName: "name",
				MaxLength:  255,
				Nullable:   false,
				Blank:      false,
				Unique:     true,
			},
			expected: `"name" VARCHAR(255) NOT NULL UNIQUE`,
		},
		{
			name: "CharField with Default",
			field: CharField{
				ColumnName: "name",
				MaxLength:  255,
				Nullable:   false,
				Blank:      false,
				Default:    "unknown",
			},
			expected: `"name" VARCHAR(255) NOT NULL DEFAULT 'unknown'`,
		},
		{
			name: "CharField with Constraints",
			field: CharField{
				ColumnName:  "name",
				MaxLength:   255,
				Nullable:    false,
				Blank:       false,
				Constraints: []string{"CHECK (name <> '')"},
			},
			expected: `"name" VARCHAR(255) NOT NULL CHECK (name <> '')`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.Definition(); got != tt.expected {
				t.Errorf("CharField.Definition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCharField_CommentSQL(t *testing.T) {
	tests := []struct {
		name     string
		field    CharField
		table    string
		expected string
	}{
		{
			name: "CharField with Comment",
			field: CharField{
				ColumnName: "name",
				MaxLength:  255,
				Nullable:   false,
				Blank:      false,
				Comment:    "Name of the report",
			},
			table:    "custom_reports",
			expected: `COMMENT ON COLUMN "custom_reports"."name" IS 'Name of the report';`,
		},
		{
			name: "CharField without Comment",
			field: CharField{
				ColumnName: "name",
				MaxLength:  255,
				Nullable:   false,
				Blank:      false,
			},
			table:    "custom_reports",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.CommentSQL(tt.table); got != tt.expected {
				t.Errorf("CharField.CommentSQL() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCharField_Name(t *testing.T) {
	field := CharField{
		ColumnName: "username",
	}
	expected := "username"
	if got := field.Name(); got != expected {
		t.Errorf("CharField.Name() = %v, want %v", got, expected)
	}
}

func TestCharField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		field   CharField
		wantErr bool
	}{
		{
			name: "Valid CharField",
			field: CharField{
				ColumnName: "name",
				MaxLength:  255,
			},
			wantErr: false,
		},
		{
			name: "Invalid CharField with MaxLength <= 0",
			field: CharField{
				ColumnName: "name",
				MaxLength:  -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("CharField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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

func TestTextField_Definition(t *testing.T) {
	tests := []struct {
		name     string
		field    TextField
		expected string
	}{
		{
			name: "Basic TextField",
			field: TextField{
				ColumnName: "content",
				Nullable:   false,
				Blank:      false,
				Unique:     true,
			},
			expected: `"content" TEXT NOT NULL UNIQUE`,
		},
		{
			name: "TextField with Default",
			field: TextField{
				ColumnName: "content",
				Nullable:   false,
				Blank:      false,
				Default:    "default content",
			},
			expected: `"content" TEXT NOT NULL DEFAULT 'default content'`,
		},
		{
			name: "TextField with Constraints",
			field: TextField{
				ColumnName:  "content",
				Nullable:    false,
				Blank:       false,
				Constraints: []string{"CHECK (LENGTH(content) > 0)"},
			},
			expected: `"content" TEXT NOT NULL CHECK (LENGTH(content) > 0)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.Definition(); got != tt.expected {
				t.Errorf("TextField.Definition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTextField_CommentSQL(t *testing.T) {
	tests := []struct {
		name     string
		field    TextField
		table    string
		expected string
	}{
		{
			name: "TextField with Comment",
			field: TextField{
				ColumnName: "content",
				Nullable:   true,
				Blank:      true,
				Comment:    "Content of the text field",
			},
			table:    "articles",
			expected: `COMMENT ON COLUMN "articles"."content" IS 'Content of the text field';`,
		},
		{
			name: "TextField without Comment",
			field: TextField{
				ColumnName: "content",
				Nullable:   false,
				Blank:      false,
			},
			table:    "articles",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.CommentSQL(tt.table); got != tt.expected {
				t.Errorf("TextField.CommentSQL() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTextField_Name(t *testing.T) {
	field := TextField{
		ColumnName: "content",
	}
	expected := "content"
	if got := field.Name(); got != expected {
		t.Errorf("TextField.Name() = %v, want %v", got, expected)
	}
}

func TestTextField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		field   TextField
		wantErr bool
	}{
		{
			name: "Valid TextField",
			field: TextField{
				ColumnName: "content",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("TextField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
