package trenovaorm

import "testing"

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
