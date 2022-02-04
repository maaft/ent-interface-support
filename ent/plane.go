// Code generated by entc, DO NOT EDIT.

package ent

import (
	"backend/ent/plane"
	"backend/ent/schema/pulid"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Plane is the model entity for the Plane schema.
type Plane struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Altitude holds the value of the "altitude" field.
	Altitude float64 `json:"altitude,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Plane) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case plane.FieldID:
			values[i] = new(pulid.ID)
		case plane.FieldAltitude:
			values[i] = new(sql.NullFloat64)
		case plane.FieldName, plane.FieldDescription:
			values[i] = new(sql.NullString)
		case plane.FieldCreatedAt, plane.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Plane", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Plane fields.
func (pl *Plane) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case plane.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pl.ID = *value
			}
		case plane.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				pl.CreatedAt = value.Time
			}
		case plane.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				pl.UpdatedAt = value.Time
			}
		case plane.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case plane.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pl.Description = value.String
			}
		case plane.FieldAltitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field altitude", values[i])
			} else if value.Valid {
				pl.Altitude = value.Float64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Plane.
// Note that you need to call Plane.Unwrap() before calling this method if this Plane
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Plane) Update() *PlaneUpdateOne {
	return (&PlaneClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the Plane entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Plane) Unwrap() *Plane {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Plane is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Plane) String() string {
	var builder strings.Builder
	builder.WriteString("Plane(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(pl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(pl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", description=")
	builder.WriteString(pl.Description)
	builder.WriteString(", altitude=")
	builder.WriteString(fmt.Sprintf("%v", pl.Altitude))
	builder.WriteByte(')')
	return builder.String()
}

// Planes is a parsable slice of Plane.
type Planes []*Plane

func (pl Planes) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}
