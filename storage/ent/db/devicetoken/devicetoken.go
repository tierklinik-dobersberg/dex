// Code generated by entc, DO NOT EDIT.

package devicetoken

const (
	// Label holds the string label denoting the devicetoken type in the database.
	Label = "device_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeviceCode holds the string denoting the device_code field in the database.
	FieldDeviceCode = "device_code"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// FieldLastRequest holds the string denoting the last_request field in the database.
	FieldLastRequest = "last_request"
	// FieldPollInterval holds the string denoting the poll_interval field in the database.
	FieldPollInterval = "poll_interval"
	// Table holds the table name of the devicetoken in the database.
	Table = "device_tokens"
)

// Columns holds all SQL columns for devicetoken fields.
var Columns = []string{
	FieldID,
	FieldDeviceCode,
	FieldStatus,
	FieldToken,
	FieldExpiry,
	FieldLastRequest,
	FieldPollInterval,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DeviceCodeValidator is a validator for the "device_code" field. It is called by the builders before save.
	DeviceCodeValidator func(string) error
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(string) error
)