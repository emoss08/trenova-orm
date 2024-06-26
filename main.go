package main

import (
	"fmt"
	"log"

	"github.com/emoss08/trenova-orm/orm"
)

func exampleIndexDefinition() orm.Index {
	return orm.Index{
		Name:        "equipment_types_code_organization_id_unq",
		Columns:     []string{"organization_id"},                 // Direct column name
		Expressions: []orm.Expression{orm.Lower{Column: "code"}}, // Using the Lower expression
		Unique:      true,
	}
}

func main() {
	// Define a simple index
	index := orm.Index{
		Columns: []string{"first_name", "last_name"},
		Unique:  false,
	}

	sql, err := index.SQL("users")
	if err != nil {
		log.Fatalf("Error generating SQL: %v", err)
	}
	fmt.Println(sql)

	// Define an index with custom expressions
	exprIndex := orm.Index{
		Expressions: []orm.Expression{
			orm.Lower{Column: "first_name"},
			orm.Concat{Columns: []string{"first_name", "last_name"}},
			orm.Gist{Column: "location"},
			orm.Gin{Column: "tags"},
		},
		Unique: true,
	}

	sql, err = exprIndex.SQL("users")
	if err != nil {
		log.Fatalf("Error generating SQL: %v", err)
	}
	fmt.Println(sql)

	orgIndex := orm.Index{
		Name:        "equipment_types_code_organization_id_unq",
		Columns:     []string{"organization_id"},                 // Direct column name
		Expressions: []orm.Expression{orm.Lower{Column: "code"}}, // Using the Lower expression
		Unique:      true,
	}

	sql, err = orgIndex.SQL("equipment_types")
	if err != nil {
		log.Fatalf("Error generating SQL: %v", err)
	}
	fmt.Println(sql)
}
