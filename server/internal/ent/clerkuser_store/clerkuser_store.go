// Code generated by ent, DO NOT EDIT.

package clerkuser_store

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the clerkuser_store type in the database.
	Label = "clerk_user_store"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldClerkID holds the string denoting the clerk_id field in the database.
	FieldClerkID = "clerk_id"
	// FieldStoreID holds the string denoting the store_id field in the database.
	FieldStoreID = "store_id"
	// Table holds the table name of the clerkuser_store in the database.
	Table = "clerk_user_stores"
)

// Columns holds all SQL columns for clerkuser_store fields.
var Columns = []string{
	FieldID,
	FieldClerkID,
	FieldStoreID,
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

// OrderOption defines the ordering options for the ClerkUser_Store queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByClerkID orders the results by the clerk_id field.
func ByClerkID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClerkID, opts...).ToFunc()
}

// ByStoreID orders the results by the store_id field.
func ByStoreID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoreID, opts...).ToFunc()
}