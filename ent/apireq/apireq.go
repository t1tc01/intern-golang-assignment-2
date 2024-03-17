// Code generated by ent, DO NOT EDIT.

package apireq

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the apireq type in the database.
	Label = "apireq"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReqTime holds the string denoting the req_time field in the database.
	FieldReqTime = "req_time"
	// FieldReqParam holds the string denoting the req_param field in the database.
	FieldReqParam = "req_param"
	// FieldReqBody holds the string denoting the req_body field in the database.
	FieldReqBody = "req_body"
	// FieldReqHeaders holds the string denoting the req_headers field in the database.
	FieldReqHeaders = "req_headers"
	// FieldReqMetadata holds the string denoting the req_metadata field in the database.
	FieldReqMetadata = "req_metadata"
	// FieldRespTime holds the string denoting the resp_time field in the database.
	FieldRespTime = "resp_time"
	// FieldRespStatus holds the string denoting the resp_status field in the database.
	FieldRespStatus = "resp_status"
	// FieldRespBody holds the string denoting the resp_body field in the database.
	FieldRespBody = "resp_body"
	// FieldRespHeaders holds the string denoting the resp_headers field in the database.
	FieldRespHeaders = "resp_headers"
	// FieldRespMetadata holds the string denoting the resp_metadata field in the database.
	FieldRespMetadata = "resp_metadata"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the apireq in the database.
	Table = "apireq"
)

// Columns holds all SQL columns for apireq fields.
var Columns = []string{
	FieldID,
	FieldReqTime,
	FieldReqParam,
	FieldReqBody,
	FieldReqHeaders,
	FieldReqMetadata,
	FieldRespTime,
	FieldRespStatus,
	FieldRespBody,
	FieldRespHeaders,
	FieldRespMetadata,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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

// OrderOption defines the ordering options for the Apireq queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByReqTime orders the results by the req_time field.
func ByReqTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReqTime, opts...).ToFunc()
}

// ByReqParam orders the results by the req_param field.
func ByReqParam(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReqParam, opts...).ToFunc()
}

// ByReqBody orders the results by the req_body field.
func ByReqBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReqBody, opts...).ToFunc()
}

// ByReqHeaders orders the results by the req_headers field.
func ByReqHeaders(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReqHeaders, opts...).ToFunc()
}

// ByReqMetadata orders the results by the req_metadata field.
func ByReqMetadata(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReqMetadata, opts...).ToFunc()
}

// ByRespTime orders the results by the resp_time field.
func ByRespTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRespTime, opts...).ToFunc()
}

// ByRespStatus orders the results by the resp_status field.
func ByRespStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRespStatus, opts...).ToFunc()
}

// ByRespBody orders the results by the resp_body field.
func ByRespBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRespBody, opts...).ToFunc()
}

// ByRespHeaders orders the results by the resp_headers field.
func ByRespHeaders(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRespHeaders, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}
