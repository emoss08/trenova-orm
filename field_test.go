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
				ColumnName: "amount",
				Precision:  10,
				Scale:      2,
				Nullable:   false,
				Unique:     true,
			},
			expected: `"amount" NUMERIC(10, 2) NOT NULL UNIQUE`,
		},
		{
			name: "NumericField with Default",
			field: NumericField{
				ColumnName: "amount",
				Precision:  10,
				Scale:      2,
				Nullable:   false,
				Default:    "0.00",
			},
			expected: `"amount" NUMERIC(10, 2) NOT NULL DEFAULT 0.00`,
		},
		{
			name: "NumericField with Constraints",
			field: NumericField{
				ColumnName:  "amount",
				Precision:   10,
				Scale:       2,
				Nullable:    false,
				Constraints: []string{"CHECK (amount >= 0)"},
			},
			expected: `"amount" NUMERIC(10, 2) NOT NULL CHECK (amount >= 0)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.Definition(); got != tt.expected {
				t.Errorf("NumericField.Definition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNumericField_CommentSQL(t *testing.T) {
	tests := []struct {
		name     string
		field    NumericField
		table    string
		expected string
	}{
		{
			name: "NumericField with Comment",
			field: NumericField{
				ColumnName: "amount",
				Precision:  10,
				Scale:      2,
				Nullable:   true,
				Comment:    "Amount of money",
			},
			table:    "transactions",
			expected: `COMMENT ON COLUMN "transactions"."amount" IS 'Amount of money';`,
		},
		{
			name: "NumericField without Comment",
			field: NumericField{
				ColumnName: "amount",
				Precision:  10,
				Scale:      2,
				Nullable:   false,
			},
			table:    "transactions",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.field.CommentSQL(tt.table); got != tt.expected {
				t.Errorf("NumericField.CommentSQL() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNumericField_Name(t *testing.T) {
	field := NumericField{
		ColumnName: "amount",
	}
	expected := "amount"
	if got := field.Name(); got != expected {
		t.Errorf("NumericField.Name() = %v, want %v", got, expected)
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
				ColumnName: "amount",
				Precision:  10,
				Scale:      2,
			},
			wantErr: false,
		},
		{
			name: "Invalid NumericField with Precision <= 0",
			field: NumericField{
				ColumnName: "amount",
				Precision:  -1,
				Scale:      2,
			},
			wantErr: true,
		},
		{
			name: "Invalid NumericField with Scale < 0",
			field: NumericField{
				ColumnName: "amount",
				Precision:  10,
				Scale:      -1,
			},
			wantErr: true,
		},
		{
			name: "Invalid NumericField with Precision < Scale",
			field: NumericField{
				ColumnName: "amount",
				Precision:  2,
				Scale:      10,
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
