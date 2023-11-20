package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name"),
		field.Bytes("photo"),
		field.Int("quantity"),
		field.Float("price").
			SchemaType(map[string]string{
				dialect.Postgres: "decimal(10,2)",
			}),
		field.Int("store_id"),
		field.String("stripe_price_id"),
	}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).Ref("items").Through("category_item", CategoryItem.Type),
		edge.From("store", Store.Type).Ref("items").Required().Field("store_id").Unique(),
	}
}
