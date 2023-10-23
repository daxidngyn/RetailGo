// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hktrib/RetailGo/ent/item"
	"github.com/hktrib/RetailGo/ent/predicate"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *ItemUpdate) SetName(s string) *ItemUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetPhoto sets the "photo" field.
func (iu *ItemUpdate) SetPhoto(b []byte) *ItemUpdate {
	iu.mutation.SetPhoto(b)
	return iu
}

// SetQuantity sets the "quantity" field.
func (iu *ItemUpdate) SetQuantity(i int) *ItemUpdate {
	iu.mutation.ResetQuantity()
	iu.mutation.SetQuantity(i)
	return iu
}

// AddQuantity adds i to the "quantity" field.
func (iu *ItemUpdate) AddQuantity(i int) *ItemUpdate {
	iu.mutation.AddQuantity(i)
	return iu
}

// SetStoreID sets the "store_id" field.
func (iu *ItemUpdate) SetStoreID(i int) *ItemUpdate {
	iu.mutation.ResetStoreID()
	iu.mutation.SetStoreID(i)
	return iu
}

// AddStoreID adds i to the "store_id" field.
func (iu *ItemUpdate) AddStoreID(i int) *ItemUpdate {
	iu.mutation.AddStoreID(i)
	return iu
}

// SetCategory sets the "category" field.
func (iu *ItemUpdate) SetCategory(s string) *ItemUpdate {
	iu.mutation.SetCategory(s)
	return iu
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Photo(); ok {
		_spec.SetField(item.FieldPhoto, field.TypeBytes, value)
	}
	if value, ok := iu.mutation.Quantity(); ok {
		_spec.SetField(item.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedQuantity(); ok {
		_spec.AddField(item.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := iu.mutation.StoreID(); ok {
		_spec.SetField(item.FieldStoreID, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedStoreID(); ok {
		_spec.AddField(item.FieldStoreID, field.TypeInt, value)
	}
	if value, ok := iu.mutation.Category(); ok {
		_spec.SetField(item.FieldCategory, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetName sets the "name" field.
func (iuo *ItemUpdateOne) SetName(s string) *ItemUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetPhoto sets the "photo" field.
func (iuo *ItemUpdateOne) SetPhoto(b []byte) *ItemUpdateOne {
	iuo.mutation.SetPhoto(b)
	return iuo
}

// SetQuantity sets the "quantity" field.
func (iuo *ItemUpdateOne) SetQuantity(i int) *ItemUpdateOne {
	iuo.mutation.ResetQuantity()
	iuo.mutation.SetQuantity(i)
	return iuo
}

// AddQuantity adds i to the "quantity" field.
func (iuo *ItemUpdateOne) AddQuantity(i int) *ItemUpdateOne {
	iuo.mutation.AddQuantity(i)
	return iuo
}

// SetStoreID sets the "store_id" field.
func (iuo *ItemUpdateOne) SetStoreID(i int) *ItemUpdateOne {
	iuo.mutation.ResetStoreID()
	iuo.mutation.SetStoreID(i)
	return iuo
}

// AddStoreID adds i to the "store_id" field.
func (iuo *ItemUpdateOne) AddStoreID(i int) *ItemUpdateOne {
	iuo.mutation.AddStoreID(i)
	return iuo
}

// SetCategory sets the "category" field.
func (iuo *ItemUpdateOne) SetCategory(s string) *ItemUpdateOne {
	iuo.mutation.SetCategory(s)
	return iuo
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iuo *ItemUpdateOne) Where(ps ...predicate.Item) *ItemUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Photo(); ok {
		_spec.SetField(item.FieldPhoto, field.TypeBytes, value)
	}
	if value, ok := iuo.mutation.Quantity(); ok {
		_spec.SetField(item.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedQuantity(); ok {
		_spec.AddField(item.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.StoreID(); ok {
		_spec.SetField(item.FieldStoreID, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedStoreID(); ok {
		_spec.AddField(item.FieldStoreID, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.Category(); ok {
		_spec.SetField(item.FieldCategory, field.TypeString, value)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
