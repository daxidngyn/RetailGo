// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hktrib/RetailGo/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// IsOwner applies equality check predicate on the "is_owner" field. It's identical to IsOwnerEQ.
func IsOwner(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsOwner, v))
}

// StoreID applies equality check predicate on the "store_id" field. It's identical to StoreIDEQ.
func StoreID(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStoreID, v))
}

// ClerkUserID applies equality check predicate on the "clerk_user_id" field. It's identical to ClerkUserIDEQ.
func ClerkUserID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldClerkUserID, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastName, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// IsOwnerEQ applies the EQ predicate on the "is_owner" field.
func IsOwnerEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsOwner, v))
}

// IsOwnerNEQ applies the NEQ predicate on the "is_owner" field.
func IsOwnerNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsOwner, v))
}

// StoreIDEQ applies the EQ predicate on the "store_id" field.
func StoreIDEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStoreID, v))
}

// StoreIDNEQ applies the NEQ predicate on the "store_id" field.
func StoreIDNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStoreID, v))
}

// StoreIDIn applies the In predicate on the "store_id" field.
func StoreIDIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldStoreID, vs...))
}

// StoreIDNotIn applies the NotIn predicate on the "store_id" field.
func StoreIDNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStoreID, vs...))
}

// StoreIDGT applies the GT predicate on the "store_id" field.
func StoreIDGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldStoreID, v))
}

// StoreIDGTE applies the GTE predicate on the "store_id" field.
func StoreIDGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldStoreID, v))
}

// StoreIDLT applies the LT predicate on the "store_id" field.
func StoreIDLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldStoreID, v))
}

// StoreIDLTE applies the LTE predicate on the "store_id" field.
func StoreIDLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldStoreID, v))
}

// ClerkUserIDEQ applies the EQ predicate on the "clerk_user_id" field.
func ClerkUserIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldClerkUserID, v))
}

// ClerkUserIDNEQ applies the NEQ predicate on the "clerk_user_id" field.
func ClerkUserIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldClerkUserID, v))
}

// ClerkUserIDIn applies the In predicate on the "clerk_user_id" field.
func ClerkUserIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldClerkUserID, vs...))
}

// ClerkUserIDNotIn applies the NotIn predicate on the "clerk_user_id" field.
func ClerkUserIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldClerkUserID, vs...))
}

// ClerkUserIDGT applies the GT predicate on the "clerk_user_id" field.
func ClerkUserIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldClerkUserID, v))
}

// ClerkUserIDGTE applies the GTE predicate on the "clerk_user_id" field.
func ClerkUserIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldClerkUserID, v))
}

// ClerkUserIDLT applies the LT predicate on the "clerk_user_id" field.
func ClerkUserIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldClerkUserID, v))
}

// ClerkUserIDLTE applies the LTE predicate on the "clerk_user_id" field.
func ClerkUserIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldClerkUserID, v))
}

// ClerkUserIDContains applies the Contains predicate on the "clerk_user_id" field.
func ClerkUserIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldClerkUserID, v))
}

// ClerkUserIDHasPrefix applies the HasPrefix predicate on the "clerk_user_id" field.
func ClerkUserIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldClerkUserID, v))
}

// ClerkUserIDHasSuffix applies the HasSuffix predicate on the "clerk_user_id" field.
func ClerkUserIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldClerkUserID, v))
}

// ClerkUserIDIsNil applies the IsNil predicate on the "clerk_user_id" field.
func ClerkUserIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldClerkUserID))
}

// ClerkUserIDNotNil applies the NotNil predicate on the "clerk_user_id" field.
func ClerkUserIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldClerkUserID))
}

// ClerkUserIDEqualFold applies the EqualFold predicate on the "clerk_user_id" field.
func ClerkUserIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldClerkUserID, v))
}

// ClerkUserIDContainsFold applies the ContainsFold predicate on the "clerk_user_id" field.
func ClerkUserIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldClerkUserID, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameIsNil applies the IsNil predicate on the "first_name" field.
func FirstNameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldFirstName))
}

// FirstNameNotNil applies the NotNil predicate on the "first_name" field.
func FirstNameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldFirstName))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameIsNil applies the IsNil predicate on the "last_name" field.
func LastNameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldLastName))
}

// LastNameNotNil applies the NotNil predicate on the "last_name" field.
func LastNameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldLastName))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldLastName, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameIsNil applies the IsNil predicate on the "username" field.
func UsernameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldUsername))
}

// UsernameNotNil applies the NotNil predicate on the "username" field.
func UsernameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldUsername))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// HasStore applies the HasEdge predicate on the "store" edge.
func HasStore() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, StoreTable, StorePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStoreWith applies the HasEdge predicate on the "store" edge with a given conditions (other predicates).
func HasStoreWith(preds ...predicate.Store) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newStoreStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserToStore applies the HasEdge predicate on the "UserToStore" edge.
func HasUserToStore() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserToStoreTable, UserToStoreColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserToStoreWith applies the HasEdge predicate on the "UserToStore" edge with a given conditions (other predicates).
func HasUserToStoreWith(preds ...predicate.UserToStore) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserToStoreStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
