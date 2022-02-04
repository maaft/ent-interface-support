package schema

import (
	"backend/ent/schema/pulid"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Plane holds the schema definition for the Plane entity.
type Plane struct {
	ent.Schema
}

func (c Plane) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinFromType(c),
		TimestampedMixin{},
		VehicleMixin{},
	}
}

// Fields of the Plane.
func (Plane) Fields() []ent.Field {
	return []ent.Field{
		field.Float("altitude"),
	}
}
