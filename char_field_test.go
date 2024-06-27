package trenovaorm

import "testing"

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
