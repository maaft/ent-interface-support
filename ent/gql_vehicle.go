package ent

import (
	"backend/ent/car"
	"backend/ent/plane"
	"backend/ent/schema/pulid"
	"context"
	"encoding/json"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/hashicorp/go-multierror"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrVehicleNotFound creates a node not found graphql error.
func ErrVehicleNotFound(id interface{}) *gqlerror.Error {
	err := gqlerror.Errorf("Could not resolve to a vehicle with the global id of '%v'", id)
	errcode.Set(err, "NOT_FOUND")
	return err
}

// Vehicler wraps the basic Vehicle method.
type Vehicler interface {
	Vehicle(context.Context) (*Vehicle, error)
}

// Vehicle in the graph.
type Vehicle struct {
	ID     pulid.ID `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// // Field of a node.
// type Field struct {
// 	Type  string `json:"type,omitempty"`  // field type.
// 	Name  string `json:"name,omitempty"`  // field name (as in struct).
// 	Value string `json:"value,omitempty"` // stringified value.
// }

// // Edges between two nodes.
// type Edge struct {
// 	Type string         `json:"type,omitempty"` // edge type.
// 	Name string         `json:"name,omitempty"` // edge name.
// 	IDs  []pulid.ID `json:"ids,omitempty"`  // node ids (where this edge point to).
// }

func (c *Car) Vehicle(ctx context.Context) (node *Vehicle, err error) {
	node = &Vehicle{
		ID:     c.ID,
		Type:   "Car",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(c.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "createdAt",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updatedAt",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.WheelPressure); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "float64",
		Name:  "wheelPressure",
		Value: string(buf),
	}
	return node, nil
}

func (pl *Plane) Vehicle(ctx context.Context) (node *Vehicle, err error) {
	node = &Vehicle{
		ID:     pl.ID,
		Type:   "Plane",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(pl.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "createdAt",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pl.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updatedAt",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pl.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pl.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pl.Altitude); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "float64",
		Name:  "altitude",
		Value: string(buf),
	}
	return node, nil
}

func (c *Client) Vehicle(ctx context.Context, id pulid.ID) (*Vehicle, error) {
	n, err := c.Vehicler(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Vehicle(ctx)
}

var errVehicleInvalidID = &NotFoundError{"node"}

// VehicleOption allows configuring the Vehicler execution using functional options.
type VehicleOption func(*vehicleOptions)

// WithVehicleType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithVehicleType(f func(context.Context, pulid.ID) (string, error)) VehicleOption {
	return func(o *vehicleOptions) {
		o.nodeType = f
	}
}

// WithFixedVehicleType sets the Type of the node to a fixed value.
func WithFixedVehicleType(t string) VehicleOption {
	return WithVehicleType(func(context.Context, pulid.ID) (string, error) {
		return t, nil
	})
}

type vehicleOptions struct {
	nodeType func(context.Context, pulid.ID) (string, error)
}

func (c *Client) newVehicleOpts(opts []VehicleOption) *vehicleOptions {
	nopts := &vehicleOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id pulid.ID) (string, error) {
			return "", fmt.Errorf("cannot resolve vehicler (%v) without its type", id)
		}
	}
	return nopts
}

// Vehicler returns a Vehicle by its id. If the VehicleType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Vehicler(ctx, id)
//		c.Vehicler(ctx, id, ent.WithVehicleType(pet.Table))
//
func (c *Client) Vehicler(ctx context.Context, id pulid.ID, opts ...VehicleOption) (_ Vehicler, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, ErrVehicleNotFound(id))
		}
	}()
	table, err := c.newVehicleOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}

	return c.vehicler(ctx, table, id)
}

func (c *Client) vehicler(ctx context.Context, table string, id pulid.ID) (Vehicler, error) {
	fmt.Println(table, id)
	switch table {
	case car.Table:
		n, err := c.Car.Query().
			Where(car.ID(id)).
			CollectFields(ctx, "Car").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case plane.Table:
		n, err := c.Plane.Query().
			Where(plane.ID(id)).
			CollectFields(ctx, "Plane").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve vehicler from table %q: %w", table, errVehicleInvalidID)
	}
}

func (c *Client) Vehiclers(ctx context.Context, ids []pulid.ID, opts ...VehicleOption) ([]Vehicler, error) {
	switch len(ids) {
	case 1:
		vehicler, err := c.Vehicler(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Vehicler{vehicler}, nil
	case 0:
		return []Vehicler{}, nil
	}

	vehiclers := make([]Vehicler, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]pulid.ID)
	id2idx := make(map[pulid.ID][]int, len(ids))
	nopts := c.newVehicleOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.vehiclers(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					vehiclers[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if vehiclers[i] != nil {
				continue
			}
			errors[i] = ErrVehicleNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], ErrVehicleNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return vehiclers, nil
}

func (c *Client) vehiclers(ctx context.Context, table string, ids []pulid.ID) ([]Vehicler, error) {
	vehiclers := make([]Vehicler, len(ids))
	idmap := make(map[pulid.ID][]*Vehicler, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &vehiclers[i])
	}
	switch table {
	case car.Table:
		nodes, err := c.Car.Query().
			Where(car.IDIn(ids...)).
			CollectFields(ctx, "Car").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, vehicler := range idmap[node.ID] {
				*vehicler = node
			}
		}
	case plane.Table:
		nodes, err := c.Plane.Query().
			Where(plane.IDIn(ids...)).
			CollectFields(ctx, "Plane").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, vehicler := range idmap[node.ID] {
				*vehicler = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve vehiclers from table %q: %w", table, errVehicleInvalidID)
	}
	return vehiclers, nil
}
