package orm

import (
	"fmt"
	"strings"
)

// Expression defines a PostgreSQL expression.
type Expression interface {
	Expression() string
	ColumnName() string // To be used for naming the index
}

// Lower defines the LOWER expression in PostgreSQL.
type Lower struct {
	Column string
}

func (l Lower) Expression() string {
	return fmt.Sprintf("LOWER(%s)", quoteIdentifier(l.Column))
}

func (l Lower) ColumnName() string {
	return l.Column
}

// Upper defines the UPPER expression in PostgreSQL.
type Upper struct {
	Column string
}

func (u Upper) Expression() string {
	return fmt.Sprintf("UPPER(%s)", quoteIdentifier(u.Column))
}

func (u Upper) ColumnName() string {
	return u.Column
}

// Concat defines a CONCAT expression in PostgreSQL for string concatenation.
type Concat struct {
	Columns []string
}

func (c Concat) Expression() string {
	return fmt.Sprintf("CONCAT(%s)", joinColumns(c.Columns))
}

func (c Concat) ColumnName() string {
	return strings.Join(c.Columns, "_")
}

// Gist defines a GIST index in PostgreSQL.
type Gist struct {
	Column string
}

func (g Gist) Expression() string {
	return fmt.Sprintf("USING GIST (%s)", quoteIdentifier(g.Column))
}

func (g Gist) ColumnName() string {
	return g.Column
}

// Gin defines a GIN index in PostgreSQL.
type Gin struct {
	Column string
}

func (g Gin) Expression() string {
	return fmt.Sprintf("USING GIN (%s)", quoteIdentifier(g.Column))
}

func (g Gin) ColumnName() string {
	return g.Column
}
