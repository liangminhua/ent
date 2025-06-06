// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/car"
	"entgo.io/ent/entc/integration/customid/ent/pet"
)

// Car is the model entity for the Car schema.
type Car struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BeforeID holds the value of the "before_id" field.
	BeforeID float64 `json:"before_id,omitempty"`
	// AfterID holds the value of the "after_id" field.
	AfterID float64 `json:"after_id,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarQuery when eager-loading is set.
	Edges        CarEdges `json:"edges"`
	pet_cars     *string
	selectValues sql.SelectValues
}

// CarEdges holds the relations/edges for other nodes in the graph.
type CarEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Pet `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarEdges) OwnerOrErr() (*Pet, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: pet.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case car.FieldBeforeID, car.FieldAfterID:
			values[i] = new(sql.NullFloat64)
		case car.FieldID:
			values[i] = new(sql.NullInt64)
		case car.FieldModel:
			values[i] = new(sql.NullString)
		case car.ForeignKeys[0]: // pet_cars
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (_m *Car) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case car.FieldBeforeID:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field before_id", values[i])
			} else if value.Valid {
				_m.BeforeID = value.Float64
			}
		case car.FieldAfterID:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field after_id", values[i])
			} else if value.Valid {
				_m.AfterID = value.Float64
			}
		case car.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				_m.Model = value.String
			}
		case car.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pet_cars", values[i])
			} else if value.Valid {
				_m.pet_cars = new(string)
				*_m.pet_cars = value.String
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Car.
// This includes values selected through modifiers, order, etc.
func (_m *Car) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Car entity.
func (_m *Car) QueryOwner() *PetQuery {
	return NewCarClient(_m.config).QueryOwner(_m)
}

// Update returns a builder for updating this Car.
// Note that you need to call Car.Unwrap() before calling this method if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Car) Update() *CarUpdateOne {
	return NewCarClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Car entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Car) Unwrap() *Car {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Car is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("before_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.BeforeID))
	builder.WriteString(", ")
	builder.WriteString("after_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.AfterID))
	builder.WriteString(", ")
	builder.WriteString("model=")
	builder.WriteString(_m.Model)
	builder.WriteByte(')')
	return builder.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car
