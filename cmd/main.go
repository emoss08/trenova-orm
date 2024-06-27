package main

import (
	"fmt"
	"strings"

	trenovaorm "github.com/emoss08/trenova-orm"
)

// TimestampedModel provides common fields for tracking creation and update times.
type TimestampedModel struct {
	trenovaorm.Model
}

// Fields returns the common timestamp fields.
func (t TimestampedModel) Fields() []trenovaorm.Field {
	return []trenovaorm.Field{
		&trenovaorm.DateField{
			ColumnName: "created_at",
			Nullable:   false,
			Default:    trenovaorm.CurrentTimestamp,
			Comment:    "Creation timestamp",
			StructTag:  `json:"created_at" validate:"required"`,
		},
		&trenovaorm.DateField{
			ColumnName: "updated_at",
			Nullable:   false,
			Default:    trenovaorm.CurrentTimestamp,
			Comment:    "Update timestamp",
			StructTag:  `json:"updated_at" validate:"required"`,
		},
	}
}

// User holds the schema definition for the User entity.
type User struct {
	TimestampedModel
}

// Mixins of the User.
func (User) Mixins() []trenovaorm.Mixin {
	return []trenovaorm.Mixin{
		trenovaorm.TimestampedMixin{},
	}
}

func (User) TableName() string {
	return "users"
}

// Fields of the User.
func (User) Fields() []trenovaorm.Field {
	return []trenovaorm.Field{
		&trenovaorm.UUIDField{
			ColumnName: "id",
			Nullable:   false,
			Blank:      false,
			Unique:     true,
			Default:    trenovaorm.UUIDGenerateV4,
			PrimaryKey: true,
			Comment:    "Unique identifier of the user",
			CustomType: "uuid",
			StructTag:  `json:"id" validate:"required"`,
		},
		&trenovaorm.CharField{
			ColumnName: "username",
			MaxLength:  255,
			Nullable:   false,
			Blank:      false,
			Unique:     true,
			Comment:    "Username of the user",
			StructTag:  `json:"username" validate:"required"`,
		},
		&trenovaorm.CharField{
			ColumnName: "email",
			MaxLength:  255,
			Nullable:   false,
			Blank:      false,
			Unique:     true,
			Comment:    "Email address of the user",
			StructTag:  `json:"email" validate:"required,email"`,
		},
		&trenovaorm.TextField{
			ColumnName: "bio",
			Nullable:   true,
			Blank:      true,
			Comment:    "Biography of the user",
			StructTag:  `json:"bio" validate:"omitempty"`,
		},
		&trenovaorm.BooleanField{
			ColumnName: "is_active",
			Nullable:   false,
			Default:    true,
			Comment:    "Is the user active",
			StructTag:  `json:"is_active" validate:"required"`,
		},
		&trenovaorm.JSONField{
			ColumnName: "metadata",
			Nullable:   true,
			Comment:    "Metadata of the user",
			StructTag:  `json:"metadata" validate:"omitempty"`,
		},
		&trenovaorm.PositiveIntegerField{
			ColumnName: "age",
			Nullable:   true,
			Unique:     false,
			Default:    0,
			Index:      false,
			Comment:    "Age of the user",
			StructTag:  `json:"age" validate:"omitempty"`,
		},
		&trenovaorm.NumericField{
			ColumnName: "rating",
			Precision:  19,
			Scale:      2,
			Nullable:   true,
			Unique:     false,
			Default:    19.00,
			Index:      false,
			Comment:    "Rating of the user",
			StructTag:  `json:"rating" validate:"omitempty"`,
		},
		&trenovaorm.ForeignKeyField{
			ColumnName:     "role_id",
			ReferenceTable: "roles",
			ReferenceField: "id",
			Nullable:       false,
			Unique:         false,
			Default:        "1",
			Comment:        "Role of the user",
			Annotations: trenovaorm.Annotation{
				OnDelete: trenovaorm.OnDeleteCascade,
				OnUpdate: trenovaorm.OnUpdateCascade,
			},
			StructTag: `json:"role_id" validate:"required"`,
		},
	}
}

// Indexes of the User.
func (User) Indexes() []trenovaorm.Index {
	return []trenovaorm.Index{
		{
			Columns: []string{"email"},
			Unique:  true,
		},
		{
			Columns: []string{"is_active"},
			Unique:  false,
		},
		{
			Name: "idx_unique_username_email",
			Expressions: []trenovaorm.Expression{
				trenovaorm.Lower{Column: "username"},
			},
			Columns: []string{"email"},
			Unique:  true,
		},
	}
}

// Example usage of the schema
func main() {
	user := &User{}

	// Generate SQL for creating the table
	createTableSQL := generateCreateTableSQL(user)
	fmt.Println(createTableSQL)

	// Generate SQL for adding comments
	commentSQLs := generateAddCommentsSQL(user)
	for _, sql := range commentSQLs {
		fmt.Println(sql)
	}

	// Generate SQL for creating indexes
	indexSQLs := generateCreateIndexesSQL(user)
	for _, sql := range indexSQLs {
		fmt.Println(sql)
	}

	// // Generate Go struct definition
	// goStruct := generateGoStruct(user)
	// fmt.Println(goStruct)
}

// Helper function to generate create table SQL
// Helper function to generate create table SQL
func generateCreateTableSQL(model trenovaorm.Model) string {
	var definitions []string
	var foreignKeys []string

	for _, field := range model.Fields() {
		definitions = append(definitions, field.Definition())
		if fkField, ok := field.(*trenovaorm.ForeignKeyField); ok {
			foreignKeys = append(foreignKeys, fkField.ForeignKeyConstraint(model.TableName()))
		}
	}

	// Include the mixin fields
	for _, mixin := range model.Mixins() {
		for _, field := range mixin.Fields() {
			definitions = append(definitions, field.Definition())
			if fkField, ok := field.(*trenovaorm.ForeignKeyField); ok {
				foreignKeys = append(foreignKeys, fkField.ForeignKeyConstraint(model.TableName()))
			}
		}
	}

	// Add foreign key constraints to definitions
	definitions = append(definitions, foreignKeys...)

	tableName := model.TableName()
	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (%s);`, tableName, strings.Join(definitions, ", "))
}

// Helper function to generate add comments SQL
func generateAddCommentsSQL(model trenovaorm.Model) []string {
	var comments []string
	tableName := model.TableName()
	for _, field := range model.Fields() {
		commentSQL := field.CommentSQL(tableName)
		if commentSQL != "" {
			comments = append(comments, commentSQL)
		}
	}

	// Include the mixin fields
	for _, mixin := range model.Mixins() {
		for _, field := range mixin.Fields() {
			commentSQL := field.CommentSQL(tableName)
			if commentSQL != "" {
				comments = append(comments, commentSQL)
			}
		}
	}

	return comments
}

// Helper function to generate create indexes SQL
func generateCreateIndexesSQL(model trenovaorm.Model) []string {
	var indexes []string
	tableName := model.TableName()
	for _, index := range model.Indexes() {
		indexSQL, _ := index.SQL(tableName)
		indexes = append(indexes, indexSQL)
	}
	return indexes
}

// Helper function to generate Go struct definition
func generateGoStruct(model trenovaorm.Model) string {
	var fields []string
	for _, field := range model.Fields() {
		fieldDef := fmt.Sprintf("%s %s `json:\"%s\"`", toCamelCase(field.Name()), field.GoType(), field.Name())
		fields = append(fields, fieldDef)
	}

	// Add the mixin fields
	for _, mixin := range model.Mixins() {
		for _, field := range mixin.Fields() {
			fieldDef := fmt.Sprintf("%s %s `json:\"%s\"`", toCamelCase(field.Name()), field.GoType(), field.Name())
			fields = append(fields, fieldDef)
		}
	}

	return fmt.Sprintf("type %s struct {\n\t%s\n}", toCamelCase(model.TableName()), strings.Join(fields, "\n\t"))
}

// toCamelCase converts snake_case to CamelCase
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
