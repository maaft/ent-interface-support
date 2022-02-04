// Code generated by entc, DO NOT EDIT.

package ent

import (
	"backend/ent/car"
	"backend/ent/schema/pulid"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Car is the model entity for the Car schema.
type Car struct {
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
	// WheelPressure holds the value of the "wheelPressure" field.
	WheelPressure float64 `json:"wheelPressure,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			values[i] = new(pulid.ID)
		case car.FieldWheelPressure:
			values[i] = new(sql.NullFloat64)
		case car.FieldName, car.FieldDescription:
			values[i] = new(sql.NullString)
		case car.FieldCreatedAt, car.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Car", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (c *Car) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case car.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case car.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case car.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case car.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case car.FieldWheelPressure:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field wheelPressure", values[i])
			} else if value.Valid {
				c.WheelPressure = value.Float64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Car.
// Note that you need to call Car.Unwrap() before calling this method if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return (&CarClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Car entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Car is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", description=")
	builder.WriteString(c.Description)
	builder.WriteString(", wheelPressure=")
	builder.WriteString(fmt.Sprintf("%v", c.WheelPressure))
	builder.WriteByte(')')
	return builder.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car

func (c Cars) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
