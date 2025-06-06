// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv2/pet"
	"entgo.io/ent/entc/integration/migrate/entv2/user"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges        PetEdges `json:"edges"`
	owner_id     *int
	selectValues sql.SelectValues
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			values[i] = new(sql.NullInt64)
		case pet.FieldName:
			values[i] = new(sql.NullString)
		case pet.ForeignKeys[0]: // owner_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (_m *Pet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case pet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				_m.Name = value.String
			}
		case pet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
			} else if value.Valid {
				_m.owner_id = new(int)
				*_m.owner_id = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pet.
// This includes values selected through modifiers, order, etc.
func (_m *Pet) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Pet entity.
func (_m *Pet) QueryOwner() *UserQuery {
	return NewPetClient(_m.config).QueryOwner(_m)
}

// Update returns a builder for updating this Pet.
// Note that you need to call Pet.Unwrap() before calling this method if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Pet) Update() *PetUpdateOne {
	return NewPetClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Pet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Pet) Unwrap() *Pet {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("entv2: Pet is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet
