// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/migration/ent/card"
	"entgo.io/ent/examples/migration/ent/payment"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CardID holds the value of the "card_id" field.
	CardID int `json:"card_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency payment.Currency `json:"currency,omitempty"`
	// Time holds the value of the "time" field.
	Time time.Time `json:"time,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Status holds the value of the "status" field.
	Status payment.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentQuery when eager-loading is set.
	Edges        PaymentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PaymentEdges holds the relations/edges for other nodes in the graph.
type PaymentEdges struct {
	// Card holds the value of the card edge.
	Card *Card `json:"card,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) CardOrErr() (*Card, error) {
	if e.Card != nil {
		return e.Card, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: card.Label}
	}
	return nil, &NotLoadedError{edge: "card"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case payment.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case payment.FieldID, payment.FieldCardID:
			values[i] = new(sql.NullInt64)
		case payment.FieldCurrency, payment.FieldDescription, payment.FieldStatus:
			values[i] = new(sql.NullString)
		case payment.FieldTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (_m *Payment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case payment.FieldCardID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field card_id", values[i])
			} else if value.Valid {
				_m.CardID = int(value.Int64)
			}
		case payment.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				_m.Amount = value.Float64
			}
		case payment.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				_m.Currency = payment.Currency(value.String)
			}
		case payment.FieldTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				_m.Time = value.Time
			}
		case payment.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				_m.Description = value.String
			}
		case payment.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				_m.Status = payment.Status(value.String)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Payment.
// This includes values selected through modifiers, order, etc.
func (_m *Payment) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryCard queries the "card" edge of the Payment entity.
func (_m *Payment) QueryCard() *CardQuery {
	return NewPaymentClient(_m.config).QueryCard(_m)
}

// Update returns a builder for updating this Payment.
// Note that you need to call Payment.Unwrap() before calling this method if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Payment) Update() *PaymentUpdateOne {
	return NewPaymentClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Payment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Payment) Unwrap() *Payment {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("card_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.CardID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", _m.Amount))
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(fmt.Sprintf("%v", _m.Currency))
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(_m.Time.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(_m.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", _m.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment
