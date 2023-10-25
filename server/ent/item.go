// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hktrib/RetailGo/ent/item"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Photo holds the value of the "photo" field.
	Photo []byte `json:"photo,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// StoreID holds the value of the "store_id" field.
	StoreID int `json:"store_id,omitempty"`
	// Category holds the value of the "category" field.
	Category     string `json:"category,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldPhoto:
			values[i] = new([]byte)
		case item.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case item.FieldID, item.FieldQuantity, item.FieldStoreID:
			values[i] = new(sql.NullInt64)
		case item.FieldName, item.FieldCategory:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case item.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case item.FieldPhoto:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field photo", values[j])
			} else if value != nil {
				i.Photo = *value
			}
		case item.FieldQuantity:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[j])
			} else if value.Valid {
				i.Quantity = int(value.Int64)
			}
		case item.FieldPrice:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[j])
			} else if value.Valid {
				i.Price = value.Float64
			}
		case item.FieldStoreID:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field store_id", values[j])
			} else if value.Valid {
				i.StoreID = int(value.Int64)
			}
		case item.FieldCategory:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[j])
			} else if value.Valid {
				i.Category = value.String
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Item.
// This includes values selected through modifiers, order, etc.
func (i *Item) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return NewItemClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("photo=")
	builder.WriteString(fmt.Sprintf("%v", i.Photo))
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", i.Quantity))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", i.Price))
	builder.WriteString(", ")
	builder.WriteString("store_id=")
	builder.WriteString(fmt.Sprintf("%v", i.StoreID))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(i.Category)
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item
