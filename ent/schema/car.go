package schema

import (
	"backend/ent/schema/pulid"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

func (c Car) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinFromType(c),
		TimestampedMixin{},
		VehicleMixin{},
	}
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.Float("wheelPressure"),
	}
}
