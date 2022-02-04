// Code generated by entc, DO NOT EDIT.

package ent

import (
	"backend/ent/car"
	"backend/ent/plane"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: car.FieldID,
			},
		},
		Type: "Car",
		Fields: map[string]*sqlgraph.FieldSpec{
			car.FieldCreatedAt:     {Type: field.TypeTime, Column: car.FieldCreatedAt},
			car.FieldUpdatedAt:     {Type: field.TypeTime, Column: car.FieldUpdatedAt},
			car.FieldName:          {Type: field.TypeString, Column: car.FieldName},
			car.FieldDescription:   {Type: field.TypeString, Column: car.FieldDescription},
			car.FieldWheelPressure: {Type: field.TypeFloat64, Column: car.FieldWheelPressure},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   plane.Table,
			Columns: plane.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: plane.FieldID,
			},
		},
		Type: "Plane",
		Fields: map[string]*sqlgraph.FieldSpec{
			plane.FieldCreatedAt:   {Type: field.TypeTime, Column: plane.FieldCreatedAt},
			plane.FieldUpdatedAt:   {Type: field.TypeTime, Column: plane.FieldUpdatedAt},
			plane.FieldName:        {Type: field.TypeString, Column: plane.FieldName},
			plane.FieldDescription: {Type: field.TypeString, Column: plane.FieldDescription},
			plane.FieldAltitude:    {Type: field.TypeFloat64, Column: plane.FieldAltitude},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *CarQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CarQuery builder.
func (cq *CarQuery) Filter() *CarFilter {
	return &CarFilter{cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CarMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CarMutation builder.
func (m *CarMutation) Filter() *CarFilter {
	return &CarFilter{m}
}

// CarFilter provides a generic filtering capability at runtime for CarQuery.
type CarFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *CarFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *CarFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(car.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the createdAt field.
func (f *CarFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(car.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updatedAt field.
func (f *CarFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(car.FieldUpdatedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *CarFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(car.FieldName))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *CarFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(car.FieldDescription))
}

// WhereWheelPressure applies the entql float64 predicate on the wheelPressure field.
func (f *CarFilter) WhereWheelPressure(p entql.Float64P) {
	f.Where(p.Field(car.FieldWheelPressure))
}

// addPredicate implements the predicateAdder interface.
func (pq *PlaneQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PlaneQuery builder.
func (pq *PlaneQuery) Filter() *PlaneFilter {
	return &PlaneFilter{pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PlaneMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PlaneMutation builder.
func (m *PlaneMutation) Filter() *PlaneFilter {
	return &PlaneFilter{m}
}

// PlaneFilter provides a generic filtering capability at runtime for PlaneQuery.
type PlaneFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *PlaneFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql string predicate on the id field.
func (f *PlaneFilter) WhereID(p entql.StringP) {
	f.Where(p.Field(plane.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the createdAt field.
func (f *PlaneFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(plane.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updatedAt field.
func (f *PlaneFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(plane.FieldUpdatedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *PlaneFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(plane.FieldName))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *PlaneFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(plane.FieldDescription))
}

// WhereAltitude applies the entql float64 predicate on the altitude field.
func (f *PlaneFilter) WhereAltitude(p entql.Float64P) {
	f.Where(p.Field(plane.FieldAltitude))
}