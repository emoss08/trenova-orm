package trenovaorm

// Function represents a PostgreSQL function.
type PSQLFunction string

// Predefined PostgreSQL functions.
const (
	CurrentTimestamp = PSQLFunction("current_timestamp")
	UUIDGenerateV4   = PSQLFunction("uuid_generate_v4()")
)

// Default returns the function as a default value for a field.
func (f PSQLFunction) String() string {
	return string(f)
}
