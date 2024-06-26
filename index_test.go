package trenovaorm

import "testing"

func TestIndex_Validate(t *testing.T) {
	tests := []struct {
		name    string
		index   Index
		wantErr bool
	}{
		{"Valid Index with Columns", Index{Columns: []string{"col1"}}, false},
		{"Valid Index with Expressions", Index{Expressions: []Expression{Lower{Column: "col1"}}}, false},
		{"Invalid Index with No Columns or Expressions", Index{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.index.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Index.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIndex_generateName(t *testing.T) {
	tests := []struct {
		name      string
		index     Index
		tableName string
		want      string
	}{
		{"Generate Name with Columns", Index{Columns: []string{"col1", "col2"}}, "table", "table_col1_col2_idx"},
		{"Generate Name with Expressions", Index{Expressions: []Expression{Lower{Column: "col1"}}}, "table", "table_col1_idx"},
		{"Use Provided Name", Index{Name: "custom_name"}, "table", "custom_name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.index.generateName(tt.tableName)
			if got != tt.want {
				t.Errorf("Index.generateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_SQL(t *testing.T) {
	tests := []struct {
		name      string
		index     Index
		tableName string
		want      string
		wantErr   bool
	}{
		{
			"SQL for Valid Index with Columns",
			Index{Columns: []string{"col1", "col2"}},
			"table",
			`CREATE INDEX IF NOT EXISTS "table_col1_col2_idx" ON "table" ("col1", "col2");`,
			false,
		},
		{
			"SQL for Valid Unique Index with Expressions",
			Index{Expressions: []Expression{Lower{Column: "col1"}}, Unique: true},
			"table",
			`CREATE UNIQUE INDEX IF NOT EXISTS "table_col1_idx" ON "table" (LOWER("col1"));`,
			false,
		},
		{
			"SQL for Invalid Index with No Columns or Expressions",
			Index{},
			"table",
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.index.SQL(tt.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index.SQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Index.SQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
