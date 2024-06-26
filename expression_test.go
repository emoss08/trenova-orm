package trenovaorm

import "testing"

func TestLower_Expression(t *testing.T) {
	expr := Lower{Column: "col1"}
	expected := `LOWER("col1")`
	if expr.Expression() != expected {
		t.Errorf("Lower.Expression() = %v, want %v", expr.Expression(), expected)
	}
}

func TestLower_ColumnName(t *testing.T) {
	expr := Lower{Column: "col1"}
	expected := "col1"
	if expr.ColumnName() != expected {
		t.Errorf("Lower.ColumnName() = %v, want %v", expr.ColumnName(), expected)
	}
}

func TestUpper_Expression(t *testing.T) {
	expr := Upper{Column: "col1"}
	expected := `UPPER("col1")`
	if expr.Expression() != expected {
		t.Errorf("Upper.Expression() = %v, want %v", expr.Expression(), expected)
	}
}

func TestUpper_ColumnName(t *testing.T) {
	expr := Upper{Column: "col1"}
	expected := "col1"
	if expr.ColumnName() != expected {
		t.Errorf("Upper.ColumnName() = %v, want %v", expr.ColumnName(), expected)
	}
}

func TestConcat_Expression(t *testing.T) {
	expr := Concat{Columns: []string{"col1", "col2"}}
	expected := `CONCAT("col1", "col2")`
	if expr.Expression() != expected {
		t.Errorf("Concat.Expression() = %v, want %v", expr.Expression(), expected)
	}
}

func TestConcat_ColumnName(t *testing.T) {
	expr := Concat{Columns: []string{"col1", "col2"}}
	expected := "col1_col2"
	if expr.ColumnName() != expected {
		t.Errorf("Concat.ColumnName() = %v, want %v", expr.ColumnName(), expected)
	}
}

func TestGist_Expression(t *testing.T) {
	expr := Gist{Column: "col1"}
	expected := `USING GIST ("col1")`
	if expr.Expression() != expected {
		t.Errorf("Gist.Expression() = %v, want %v", expr.Expression(), expected)
	}
}

func TestGist_ColumnName(t *testing.T) {
	expr := Gist{Column: "col1"}
	expected := "col1"
	if expr.ColumnName() != expected {
		t.Errorf("Gist.ColumnName() = %v, want %v", expr.ColumnName(), expected)
	}
}

func TestGin_Expression(t *testing.T) {
	expr := Gin{Column: "col1"}
	expected := `USING GIN ("col1")`
	if expr.Expression() != expected {
		t.Errorf("Gin.Expression() = %v, want %v", expr.Expression(), expected)
	}
}

func TestGin_ColumnName(t *testing.T) {
	expr := Gin{Column: "col1"}
	expected := "col1"
	if expr.ColumnName() != expected {
		t.Errorf("Gin.ColumnName() = %v, want %v", expr.ColumnName(), expected)
	}
}

func TestBtree_Expression(t *testing.T) {
	expr := Btree{Column: "col1"}
	expected := `USING BTREE ("col1")`
	if expr.Expression() != expected {
		t.Errorf("Btree.Expression() = %v, want %v", expr.Expression(), expected)
	}
}

func TestBtree_ColumnName(t *testing.T) {
	expr := Btree{Column: "col1"}
	expected := "col1"
	if expr.ColumnName() != expected {
		t.Errorf("Btree.ColumnName() = %v, want %v", expr.ColumnName(), expected)
	}
}

func TestHash_Expression(t *testing.T) {
	expr := Hash{Column: "col1"}
	expected := `USING HASH ("col1")`
	if expr.Expression() != expected {
		t.Errorf("Hash.Expression() = %v, want %v", expr.Expression(), expected)
	}
}

func TestHash_ColumnName(t *testing.T) {
	expr := Hash{Column: "col1"}
	expected := "col1"
	if expr.ColumnName() != expected {
		t.Errorf("Hash.ColumnName() = %v, want %v", expr.ColumnName(), expected)
	}
}

func TestToTsVector_Expression(t *testing.T) {
	expr := ToTsVector{Config: "english", Column: "col1"}
	expected := `to_tsvector("english", "col1")`
	if expr.Expression() != expected {
		t.Errorf("ToTsVector.Expression() = %v, want %v", expr.Expression(), expected)
	}
}

func TestToTsVector_ColumnName(t *testing.T) {
	expr := ToTsVector{Config: "english", Column: "col1"}
	expected := "col1"
	if expr.ColumnName() != expected {
		t.Errorf("ToTsVector.ColumnName() = %v, want %v", expr.ColumnName(), expected)
	}
}
