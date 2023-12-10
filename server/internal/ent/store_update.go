// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hktrib/RetailGo/internal/ent/category"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/ent/predicate"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/hktrib/RetailGo/internal/ent/user"
)

// StoreUpdate is the builder for updating Store entities.
type StoreUpdate struct {
	config
	hooks    []Hook
	mutation *StoreMutation
}

// Where appends a list predicates to the StoreUpdate builder.
func (su *StoreUpdate) Where(ps ...predicate.Store) *StoreUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUUID sets the "uuid" field.
func (su *StoreUpdate) SetUUID(s string) *StoreUpdate {
	su.mutation.SetUUID(s)
	return su
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (su *StoreUpdate) SetNillableUUID(s *string) *StoreUpdate {
	if s != nil {
		su.SetUUID(*s)
	}
	return su
}

// SetStoreName sets the "store_name" field.
func (su *StoreUpdate) SetStoreName(s string) *StoreUpdate {
	su.mutation.SetStoreName(s)
	return su
}

// SetNillableStoreName sets the "store_name" field if the given value is not nil.
func (su *StoreUpdate) SetNillableStoreName(s *string) *StoreUpdate {
	if s != nil {
		su.SetStoreName(*s)
	}
	return su
}

// SetCreatedBy sets the "created_by" field.
func (su *StoreUpdate) SetCreatedBy(s string) *StoreUpdate {
	su.mutation.SetCreatedBy(s)
	return su
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (su *StoreUpdate) SetNillableCreatedBy(s *string) *StoreUpdate {
	if s != nil {
		su.SetCreatedBy(*s)
	}
	return su
}

// SetOwnerEmail sets the "owner_email" field.
func (su *StoreUpdate) SetOwnerEmail(s string) *StoreUpdate {
	su.mutation.SetOwnerEmail(s)
	return su
}

// SetNillableOwnerEmail sets the "owner_email" field if the given value is not nil.
func (su *StoreUpdate) SetNillableOwnerEmail(s *string) *StoreUpdate {
	if s != nil {
		su.SetOwnerEmail(*s)
	}
	return su
}

// ClearOwnerEmail clears the value of the "owner_email" field.
func (su *StoreUpdate) ClearOwnerEmail() *StoreUpdate {
	su.mutation.ClearOwnerEmail()
	return su
}

// SetStoreAddress sets the "store_address" field.
func (su *StoreUpdate) SetStoreAddress(s string) *StoreUpdate {
	su.mutation.SetStoreAddress(s)
	return su
}

// SetNillableStoreAddress sets the "store_address" field if the given value is not nil.
func (su *StoreUpdate) SetNillableStoreAddress(s *string) *StoreUpdate {
	if s != nil {
		su.SetStoreAddress(*s)
	}
	return su
}

// ClearStoreAddress clears the value of the "store_address" field.
func (su *StoreUpdate) ClearStoreAddress() *StoreUpdate {
	su.mutation.ClearStoreAddress()
	return su
}

// SetStorePhone sets the "store_phone" field.
func (su *StoreUpdate) SetStorePhone(s string) *StoreUpdate {
	su.mutation.SetStorePhone(s)
	return su
}

// SetNillableStorePhone sets the "store_phone" field if the given value is not nil.
func (su *StoreUpdate) SetNillableStorePhone(s *string) *StoreUpdate {
	if s != nil {
		su.SetStorePhone(*s)
	}
	return su
}

// ClearStorePhone clears the value of the "store_phone" field.
func (su *StoreUpdate) ClearStorePhone() *StoreUpdate {
	su.mutation.ClearStorePhone()
	return su
}

// SetStripeAccountID sets the "stripe_account_id" field.
func (su *StoreUpdate) SetStripeAccountID(s string) *StoreUpdate {
	su.mutation.SetStripeAccountID(s)
	return su
}

// SetNillableStripeAccountID sets the "stripe_account_id" field if the given value is not nil.
func (su *StoreUpdate) SetNillableStripeAccountID(s *string) *StoreUpdate {
	if s != nil {
		su.SetStripeAccountID(*s)
	}
	return su
}

// ClearStripeAccountID clears the value of the "stripe_account_id" field.
func (su *StoreUpdate) ClearStripeAccountID() *StoreUpdate {
	su.mutation.ClearStripeAccountID()
	return su
}

// SetStoreType sets the "store_type" field.
func (su *StoreUpdate) SetStoreType(s string) *StoreUpdate {
	su.mutation.SetStoreType(s)
	return su
}

// SetNillableStoreType sets the "store_type" field if the given value is not nil.
func (su *StoreUpdate) SetNillableStoreType(s *string) *StoreUpdate {
	if s != nil {
		su.SetStoreType(*s)
	}
	return su
}

// ClearStoreType clears the value of the "store_type" field.
func (su *StoreUpdate) ClearStoreType() *StoreUpdate {
	su.mutation.ClearStoreType()
	return su
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (su *StoreUpdate) AddItemIDs(ids ...int) *StoreUpdate {
	su.mutation.AddItemIDs(ids...)
	return su
}

// AddItems adds the "items" edges to the Item entity.
func (su *StoreUpdate) AddItems(i ...*Item) *StoreUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return su.AddItemIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (su *StoreUpdate) AddCategoryIDs(ids ...int) *StoreUpdate {
	su.mutation.AddCategoryIDs(ids...)
	return su
}

// AddCategories adds the "categories" edges to the Category entity.
func (su *StoreUpdate) AddCategories(c ...*Category) *StoreUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCategoryIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (su *StoreUpdate) AddUserIDs(ids ...int) *StoreUpdate {
	su.mutation.AddUserIDs(ids...)
	return su
}

// AddUser adds the "user" edges to the User entity.
func (su *StoreUpdate) AddUser(u ...*User) *StoreUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.AddUserIDs(ids...)
}

// Mutation returns the StoreMutation object of the builder.
func (su *StoreUpdate) Mutation() *StoreMutation {
	return su.mutation
}

// ClearItems clears all "items" edges to the Item entity.
func (su *StoreUpdate) ClearItems() *StoreUpdate {
	su.mutation.ClearItems()
	return su
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (su *StoreUpdate) RemoveItemIDs(ids ...int) *StoreUpdate {
	su.mutation.RemoveItemIDs(ids...)
	return su
}

// RemoveItems removes "items" edges to Item entities.
func (su *StoreUpdate) RemoveItems(i ...*Item) *StoreUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return su.RemoveItemIDs(ids...)
}

// ClearCategories clears all "categories" edges to the Category entity.
func (su *StoreUpdate) ClearCategories() *StoreUpdate {
	su.mutation.ClearCategories()
	return su
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (su *StoreUpdate) RemoveCategoryIDs(ids ...int) *StoreUpdate {
	su.mutation.RemoveCategoryIDs(ids...)
	return su
}

// RemoveCategories removes "categories" edges to Category entities.
func (su *StoreUpdate) RemoveCategories(c ...*Category) *StoreUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCategoryIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (su *StoreUpdate) ClearUser() *StoreUpdate {
	su.mutation.ClearUser()
	return su
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (su *StoreUpdate) RemoveUserIDs(ids ...int) *StoreUpdate {
	su.mutation.RemoveUserIDs(ids...)
	return su
}

// RemoveUser removes "user" edges to User entities.
func (su *StoreUpdate) RemoveUser(u ...*User) *StoreUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return su.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StoreUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StoreUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StoreUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StoreUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(store.Table, store.Columns, sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UUID(); ok {
		_spec.SetField(store.FieldUUID, field.TypeString, value)
	}
	if value, ok := su.mutation.StoreName(); ok {
		_spec.SetField(store.FieldStoreName, field.TypeString, value)
	}
	if value, ok := su.mutation.CreatedBy(); ok {
		_spec.SetField(store.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := su.mutation.OwnerEmail(); ok {
		_spec.SetField(store.FieldOwnerEmail, field.TypeString, value)
	}
	if su.mutation.OwnerEmailCleared() {
		_spec.ClearField(store.FieldOwnerEmail, field.TypeString)
	}
	if value, ok := su.mutation.StoreAddress(); ok {
		_spec.SetField(store.FieldStoreAddress, field.TypeString, value)
	}
	if su.mutation.StoreAddressCleared() {
		_spec.ClearField(store.FieldStoreAddress, field.TypeString)
	}
	if value, ok := su.mutation.StorePhone(); ok {
		_spec.SetField(store.FieldStorePhone, field.TypeString, value)
	}
	if su.mutation.StorePhoneCleared() {
		_spec.ClearField(store.FieldStorePhone, field.TypeString)
	}
	if value, ok := su.mutation.StripeAccountID(); ok {
		_spec.SetField(store.FieldStripeAccountID, field.TypeString, value)
	}
	if su.mutation.StripeAccountIDCleared() {
		_spec.ClearField(store.FieldStripeAccountID, field.TypeString)
	}
	if value, ok := su.mutation.StoreType(); ok {
		_spec.SetField(store.FieldStoreType, field.TypeString, value)
	}
	if su.mutation.StoreTypeCleared() {
		_spec.ClearField(store.FieldStoreType, field.TypeString)
	}
	if su.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ItemsTable,
			Columns: []string{store.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedItemsIDs(); len(nodes) > 0 && !su.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ItemsTable,
			Columns: []string{store.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ItemsTable,
			Columns: []string{store.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.CategoriesTable,
			Columns: []string{store.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !su.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.CategoriesTable,
			Columns: []string{store.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.CategoriesTable,
			Columns: []string{store.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   store.UserTable,
			Columns: store.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedUserIDs(); len(nodes) > 0 && !su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   store.UserTable,
			Columns: store.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   store.UserTable,
			Columns: store.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{store.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StoreUpdateOne is the builder for updating a single Store entity.
type StoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StoreMutation
}

// SetUUID sets the "uuid" field.
func (suo *StoreUpdateOne) SetUUID(s string) *StoreUpdateOne {
	suo.mutation.SetUUID(s)
	return suo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableUUID(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetUUID(*s)
	}
	return suo
}

// SetStoreName sets the "store_name" field.
func (suo *StoreUpdateOne) SetStoreName(s string) *StoreUpdateOne {
	suo.mutation.SetStoreName(s)
	return suo
}

// SetNillableStoreName sets the "store_name" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableStoreName(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetStoreName(*s)
	}
	return suo
}

// SetCreatedBy sets the "created_by" field.
func (suo *StoreUpdateOne) SetCreatedBy(s string) *StoreUpdateOne {
	suo.mutation.SetCreatedBy(s)
	return suo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableCreatedBy(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetCreatedBy(*s)
	}
	return suo
}

// SetOwnerEmail sets the "owner_email" field.
func (suo *StoreUpdateOne) SetOwnerEmail(s string) *StoreUpdateOne {
	suo.mutation.SetOwnerEmail(s)
	return suo
}

// SetNillableOwnerEmail sets the "owner_email" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableOwnerEmail(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetOwnerEmail(*s)
	}
	return suo
}

// ClearOwnerEmail clears the value of the "owner_email" field.
func (suo *StoreUpdateOne) ClearOwnerEmail() *StoreUpdateOne {
	suo.mutation.ClearOwnerEmail()
	return suo
}

// SetStoreAddress sets the "store_address" field.
func (suo *StoreUpdateOne) SetStoreAddress(s string) *StoreUpdateOne {
	suo.mutation.SetStoreAddress(s)
	return suo
}

// SetNillableStoreAddress sets the "store_address" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableStoreAddress(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetStoreAddress(*s)
	}
	return suo
}

// ClearStoreAddress clears the value of the "store_address" field.
func (suo *StoreUpdateOne) ClearStoreAddress() *StoreUpdateOne {
	suo.mutation.ClearStoreAddress()
	return suo
}

// SetStorePhone sets the "store_phone" field.
func (suo *StoreUpdateOne) SetStorePhone(s string) *StoreUpdateOne {
	suo.mutation.SetStorePhone(s)
	return suo
}

// SetNillableStorePhone sets the "store_phone" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableStorePhone(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetStorePhone(*s)
	}
	return suo
}

// ClearStorePhone clears the value of the "store_phone" field.
func (suo *StoreUpdateOne) ClearStorePhone() *StoreUpdateOne {
	suo.mutation.ClearStorePhone()
	return suo
}

// SetStripeAccountID sets the "stripe_account_id" field.
func (suo *StoreUpdateOne) SetStripeAccountID(s string) *StoreUpdateOne {
	suo.mutation.SetStripeAccountID(s)
	return suo
}

// SetNillableStripeAccountID sets the "stripe_account_id" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableStripeAccountID(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetStripeAccountID(*s)
	}
	return suo
}

// ClearStripeAccountID clears the value of the "stripe_account_id" field.
func (suo *StoreUpdateOne) ClearStripeAccountID() *StoreUpdateOne {
	suo.mutation.ClearStripeAccountID()
	return suo
}

// SetStoreType sets the "store_type" field.
func (suo *StoreUpdateOne) SetStoreType(s string) *StoreUpdateOne {
	suo.mutation.SetStoreType(s)
	return suo
}

// SetNillableStoreType sets the "store_type" field if the given value is not nil.
func (suo *StoreUpdateOne) SetNillableStoreType(s *string) *StoreUpdateOne {
	if s != nil {
		suo.SetStoreType(*s)
	}
	return suo
}

// ClearStoreType clears the value of the "store_type" field.
func (suo *StoreUpdateOne) ClearStoreType() *StoreUpdateOne {
	suo.mutation.ClearStoreType()
	return suo
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (suo *StoreUpdateOne) AddItemIDs(ids ...int) *StoreUpdateOne {
	suo.mutation.AddItemIDs(ids...)
	return suo
}

// AddItems adds the "items" edges to the Item entity.
func (suo *StoreUpdateOne) AddItems(i ...*Item) *StoreUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return suo.AddItemIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (suo *StoreUpdateOne) AddCategoryIDs(ids ...int) *StoreUpdateOne {
	suo.mutation.AddCategoryIDs(ids...)
	return suo
}

// AddCategories adds the "categories" edges to the Category entity.
func (suo *StoreUpdateOne) AddCategories(c ...*Category) *StoreUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCategoryIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (suo *StoreUpdateOne) AddUserIDs(ids ...int) *StoreUpdateOne {
	suo.mutation.AddUserIDs(ids...)
	return suo
}

// AddUser adds the "user" edges to the User entity.
func (suo *StoreUpdateOne) AddUser(u ...*User) *StoreUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.AddUserIDs(ids...)
}

// Mutation returns the StoreMutation object of the builder.
func (suo *StoreUpdateOne) Mutation() *StoreMutation {
	return suo.mutation
}

// ClearItems clears all "items" edges to the Item entity.
func (suo *StoreUpdateOne) ClearItems() *StoreUpdateOne {
	suo.mutation.ClearItems()
	return suo
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (suo *StoreUpdateOne) RemoveItemIDs(ids ...int) *StoreUpdateOne {
	suo.mutation.RemoveItemIDs(ids...)
	return suo
}

// RemoveItems removes "items" edges to Item entities.
func (suo *StoreUpdateOne) RemoveItems(i ...*Item) *StoreUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return suo.RemoveItemIDs(ids...)
}

// ClearCategories clears all "categories" edges to the Category entity.
func (suo *StoreUpdateOne) ClearCategories() *StoreUpdateOne {
	suo.mutation.ClearCategories()
	return suo
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (suo *StoreUpdateOne) RemoveCategoryIDs(ids ...int) *StoreUpdateOne {
	suo.mutation.RemoveCategoryIDs(ids...)
	return suo
}

// RemoveCategories removes "categories" edges to Category entities.
func (suo *StoreUpdateOne) RemoveCategories(c ...*Category) *StoreUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCategoryIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (suo *StoreUpdateOne) ClearUser() *StoreUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (suo *StoreUpdateOne) RemoveUserIDs(ids ...int) *StoreUpdateOne {
	suo.mutation.RemoveUserIDs(ids...)
	return suo
}

// RemoveUser removes "user" edges to User entities.
func (suo *StoreUpdateOne) RemoveUser(u ...*User) *StoreUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return suo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the StoreUpdate builder.
func (suo *StoreUpdateOne) Where(ps ...predicate.Store) *StoreUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StoreUpdateOne) Select(field string, fields ...string) *StoreUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Store entity.
func (suo *StoreUpdateOne) Save(ctx context.Context) (*Store, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StoreUpdateOne) SaveX(ctx context.Context) *Store {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StoreUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StoreUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StoreUpdateOne) sqlSave(ctx context.Context) (_node *Store, err error) {
	_spec := sqlgraph.NewUpdateSpec(store.Table, store.Columns, sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Store.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, store.FieldID)
		for _, f := range fields {
			if !store.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != store.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UUID(); ok {
		_spec.SetField(store.FieldUUID, field.TypeString, value)
	}
	if value, ok := suo.mutation.StoreName(); ok {
		_spec.SetField(store.FieldStoreName, field.TypeString, value)
	}
	if value, ok := suo.mutation.CreatedBy(); ok {
		_spec.SetField(store.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := suo.mutation.OwnerEmail(); ok {
		_spec.SetField(store.FieldOwnerEmail, field.TypeString, value)
	}
	if suo.mutation.OwnerEmailCleared() {
		_spec.ClearField(store.FieldOwnerEmail, field.TypeString)
	}
	if value, ok := suo.mutation.StoreAddress(); ok {
		_spec.SetField(store.FieldStoreAddress, field.TypeString, value)
	}
	if suo.mutation.StoreAddressCleared() {
		_spec.ClearField(store.FieldStoreAddress, field.TypeString)
	}
	if value, ok := suo.mutation.StorePhone(); ok {
		_spec.SetField(store.FieldStorePhone, field.TypeString, value)
	}
	if suo.mutation.StorePhoneCleared() {
		_spec.ClearField(store.FieldStorePhone, field.TypeString)
	}
	if value, ok := suo.mutation.StripeAccountID(); ok {
		_spec.SetField(store.FieldStripeAccountID, field.TypeString, value)
	}
	if suo.mutation.StripeAccountIDCleared() {
		_spec.ClearField(store.FieldStripeAccountID, field.TypeString)
	}
	if value, ok := suo.mutation.StoreType(); ok {
		_spec.SetField(store.FieldStoreType, field.TypeString, value)
	}
	if suo.mutation.StoreTypeCleared() {
		_spec.ClearField(store.FieldStoreType, field.TypeString)
	}
	if suo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ItemsTable,
			Columns: []string{store.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !suo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ItemsTable,
			Columns: []string{store.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.ItemsTable,
			Columns: []string{store.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.CategoriesTable,
			Columns: []string{store.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !suo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.CategoriesTable,
			Columns: []string{store.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   store.CategoriesTable,
			Columns: []string{store.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   store.UserTable,
			Columns: store.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedUserIDs(); len(nodes) > 0 && !suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   store.UserTable,
			Columns: store.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   store.UserTable,
			Columns: store.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Store{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{store.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
