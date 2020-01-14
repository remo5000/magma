// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/servicetype"
)

// ServiceType is the model entity for the ServiceType schema.
type ServiceType struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// HasCustomer holds the value of the "has_customer" field.
	HasCustomer bool `json:"has_customer,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServiceType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullTime{},
		&sql.NullTime{},
		&sql.NullString{},
		&sql.NullBool{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServiceType fields.
func (st *ServiceType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(servicetype.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	st.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		st.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		st.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		st.Name = value.String
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field has_customer", values[3])
	} else if value.Valid {
		st.HasCustomer = value.Bool
	}
	return nil
}

// QueryServices queries the services edge of the ServiceType.
func (st *ServiceType) QueryServices() *ServiceQuery {
	return (&ServiceTypeClient{st.config}).QueryServices(st)
}

// QueryPropertyTypes queries the property_types edge of the ServiceType.
func (st *ServiceType) QueryPropertyTypes() *PropertyTypeQuery {
	return (&ServiceTypeClient{st.config}).QueryPropertyTypes(st)
}

// Update returns a builder for updating this ServiceType.
// Note that, you need to call ServiceType.Unwrap() before calling this method, if this ServiceType
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *ServiceType) Update() *ServiceTypeUpdateOne {
	return (&ServiceTypeClient{st.config}).UpdateOne(st)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (st *ServiceType) Unwrap() *ServiceType {
	tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("ent: ServiceType is not a transactional entity")
	}
	st.config.driver = tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *ServiceType) String() string {
	var builder strings.Builder
	builder.WriteString("ServiceType(")
	builder.WriteString(fmt.Sprintf("id=%v", st.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(st.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(st.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(st.Name)
	builder.WriteString(", has_customer=")
	builder.WriteString(fmt.Sprintf("%v", st.HasCustomer))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (st *ServiceType) id() int {
	id, _ := strconv.Atoi(st.ID)
	return id
}

// ServiceTypes is a parsable slice of ServiceType.
type ServiceTypes []*ServiceType

func (st ServiceTypes) config(cfg config) {
	for _i := range st {
		st[_i].config = cfg
	}
}