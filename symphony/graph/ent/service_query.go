// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/customer"
	"github.com/facebookincubator/symphony/graph/ent/link"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/property"
	"github.com/facebookincubator/symphony/graph/ent/service"
	"github.com/facebookincubator/symphony/graph/ent/serviceendpoint"
	"github.com/facebookincubator/symphony/graph/ent/servicetype"
)

// ServiceQuery is the builder for querying Service entities.
type ServiceQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Service
	// eager-loading edges.
	withType       *ServiceTypeQuery
	withDownstream *ServiceQuery
	withUpstream   *ServiceQuery
	withProperties *PropertyQuery
	withLinks      *LinkQuery
	withCustomer   *CustomerQuery
	withEndpoints  *ServiceEndpointQuery
	withFKs        bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (sq *ServiceQuery) Where(ps ...predicate.Service) *ServiceQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *ServiceQuery) Limit(limit int) *ServiceQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *ServiceQuery) Offset(offset int) *ServiceQuery {
	sq.offset = &offset
	return sq
}

// Order adds an order step to the query.
func (sq *ServiceQuery) Order(o ...Order) *ServiceQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryType chains the current query on the type edge.
func (sq *ServiceQuery) QueryType() *ServiceTypeQuery {
	query := &ServiceTypeQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(service.Table, service.FieldID, sq.sqlQuery()),
		sqlgraph.To(servicetype.Table, servicetype.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, service.TypeTable, service.TypeColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryDownstream chains the current query on the downstream edge.
func (sq *ServiceQuery) QueryDownstream() *ServiceQuery {
	query := &ServiceQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(service.Table, service.FieldID, sq.sqlQuery()),
		sqlgraph.To(service.Table, service.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, service.DownstreamTable, service.DownstreamPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryUpstream chains the current query on the upstream edge.
func (sq *ServiceQuery) QueryUpstream() *ServiceQuery {
	query := &ServiceQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(service.Table, service.FieldID, sq.sqlQuery()),
		sqlgraph.To(service.Table, service.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, service.UpstreamTable, service.UpstreamPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryProperties chains the current query on the properties edge.
func (sq *ServiceQuery) QueryProperties() *PropertyQuery {
	query := &PropertyQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(service.Table, service.FieldID, sq.sqlQuery()),
		sqlgraph.To(property.Table, property.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, service.PropertiesTable, service.PropertiesColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryLinks chains the current query on the links edge.
func (sq *ServiceQuery) QueryLinks() *LinkQuery {
	query := &LinkQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(service.Table, service.FieldID, sq.sqlQuery()),
		sqlgraph.To(link.Table, link.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, service.LinksTable, service.LinksPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryCustomer chains the current query on the customer edge.
func (sq *ServiceQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(service.Table, service.FieldID, sq.sqlQuery()),
		sqlgraph.To(customer.Table, customer.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, service.CustomerTable, service.CustomerPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryEndpoints chains the current query on the endpoints edge.
func (sq *ServiceQuery) QueryEndpoints() *ServiceEndpointQuery {
	query := &ServiceEndpointQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(service.Table, service.FieldID, sq.sqlQuery()),
		sqlgraph.To(serviceendpoint.Table, serviceendpoint.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, service.EndpointsTable, service.EndpointsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// First returns the first Service entity in the query. Returns *NotFoundError when no service was found.
func (sq *ServiceQuery) First(ctx context.Context) (*Service, error) {
	sSlice, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sSlice) == 0 {
		return nil, &NotFoundError{service.Label}
	}
	return sSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ServiceQuery) FirstX(ctx context.Context) *Service {
	s, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return s
}

// FirstID returns the first Service id in the query. Returns *NotFoundError when no id was found.
func (sq *ServiceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{service.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (sq *ServiceQuery) FirstXID(ctx context.Context) string {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Service entity in the query, returns an error if not exactly one entity was returned.
func (sq *ServiceQuery) Only(ctx context.Context) (*Service, error) {
	sSlice, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sSlice) {
	case 1:
		return sSlice[0], nil
	case 0:
		return nil, &NotFoundError{service.Label}
	default:
		return nil, &NotSingularError{service.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ServiceQuery) OnlyX(ctx context.Context) *Service {
	s, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// OnlyID returns the only Service id in the query, returns an error if not exactly one id was returned.
func (sq *ServiceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{service.Label}
	default:
		err = &NotSingularError{service.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (sq *ServiceQuery) OnlyXID(ctx context.Context) string {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Services.
func (sq *ServiceQuery) All(ctx context.Context) ([]*Service, error) {
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *ServiceQuery) AllX(ctx context.Context) []*Service {
	sSlice, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sSlice
}

// IDs executes the query and returns a list of Service ids.
func (sq *ServiceQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := sq.Select(service.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ServiceQuery) IDsX(ctx context.Context) []string {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ServiceQuery) Count(ctx context.Context) (int, error) {
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ServiceQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ServiceQuery) Exist(ctx context.Context) (bool, error) {
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ServiceQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ServiceQuery) Clone() *ServiceQuery {
	return &ServiceQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]Order{}, sq.order...),
		unique:     append([]string{}, sq.unique...),
		predicates: append([]predicate.Service{}, sq.predicates...),
		// clone intermediate query.
		sql: sq.sql.Clone(),
	}
}

//  WithType tells the query-builder to eager-loads the nodes that are connected to
// the "type" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ServiceQuery) WithType(opts ...func(*ServiceTypeQuery)) *ServiceQuery {
	query := &ServiceTypeQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withType = query
	return sq
}

//  WithDownstream tells the query-builder to eager-loads the nodes that are connected to
// the "downstream" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ServiceQuery) WithDownstream(opts ...func(*ServiceQuery)) *ServiceQuery {
	query := &ServiceQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withDownstream = query
	return sq
}

//  WithUpstream tells the query-builder to eager-loads the nodes that are connected to
// the "upstream" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ServiceQuery) WithUpstream(opts ...func(*ServiceQuery)) *ServiceQuery {
	query := &ServiceQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withUpstream = query
	return sq
}

//  WithProperties tells the query-builder to eager-loads the nodes that are connected to
// the "properties" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ServiceQuery) WithProperties(opts ...func(*PropertyQuery)) *ServiceQuery {
	query := &PropertyQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withProperties = query
	return sq
}

//  WithLinks tells the query-builder to eager-loads the nodes that are connected to
// the "links" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ServiceQuery) WithLinks(opts ...func(*LinkQuery)) *ServiceQuery {
	query := &LinkQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withLinks = query
	return sq
}

//  WithCustomer tells the query-builder to eager-loads the nodes that are connected to
// the "customer" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ServiceQuery) WithCustomer(opts ...func(*CustomerQuery)) *ServiceQuery {
	query := &CustomerQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withCustomer = query
	return sq
}

//  WithEndpoints tells the query-builder to eager-loads the nodes that are connected to
// the "endpoints" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ServiceQuery) WithEndpoints(opts ...func(*ServiceEndpointQuery)) *ServiceQuery {
	query := &ServiceEndpointQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withEndpoints = query
	return sq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Service.Query().
//		GroupBy(service.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *ServiceQuery) GroupBy(field string, fields ...string) *ServiceGroupBy {
	group := &ServiceGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = sq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Service.Query().
//		Select(service.FieldCreateTime).
//		Scan(ctx, &v)
//
func (sq *ServiceQuery) Select(field string, fields ...string) *ServiceSelect {
	selector := &ServiceSelect{config: sq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = sq.sqlQuery()
	return selector
}

func (sq *ServiceQuery) sqlAll(ctx context.Context) ([]*Service, error) {
	var (
		nodes   []*Service = []*Service{}
		withFKs            = sq.withFKs
		_spec              = sq.querySpec()
	)
	if sq.withType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, service.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Service{config: sq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withType; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Service)
		for i := range nodes {
			if fk := nodes[i].type_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(servicetype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "type_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Type = n
			}
		}
	}

	if query := sq.withDownstream; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*Service, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []string
			edges   = make(map[string][]*Service)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   service.DownstreamTable,
				Columns: service.DownstreamPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(service.DownstreamPrimaryKey[1], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := strconv.FormatInt(eout.Int64, 10)
				inValue := strconv.FormatInt(ein.Int64, 10)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, sq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "downstream": %v`, err)
		}
		query.Where(service.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "downstream" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Downstream = append(nodes[i].Edges.Downstream, n)
			}
		}
	}

	if query := sq.withUpstream; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*Service, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []string
			edges   = make(map[string][]*Service)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   service.UpstreamTable,
				Columns: service.UpstreamPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(service.UpstreamPrimaryKey[0], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := strconv.FormatInt(eout.Int64, 10)
				inValue := strconv.FormatInt(ein.Int64, 10)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, sq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "upstream": %v`, err)
		}
		query.Where(service.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "upstream" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Upstream = append(nodes[i].Edges.Upstream, n)
			}
		}
	}

	if query := sq.withProperties; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*Service)
		for i := range nodes {
			id, err := strconv.Atoi(nodes[i].ID)
			if err != nil {
				return nil, err
			}
			fks = append(fks, id)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Property(func(s *sql.Selector) {
			s.Where(sql.InValues(service.PropertiesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.service_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "service_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "service_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Properties = append(node.Edges.Properties, n)
		}
	}

	if query := sq.withLinks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*Service, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []string
			edges   = make(map[string][]*Service)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   service.LinksTable,
				Columns: service.LinksPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(service.LinksPrimaryKey[0], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := strconv.FormatInt(eout.Int64, 10)
				inValue := strconv.FormatInt(ein.Int64, 10)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, sq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "links": %v`, err)
		}
		query.Where(link.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "links" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Links = append(nodes[i].Edges.Links, n)
			}
		}
	}

	if query := sq.withCustomer; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*Service, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []string
			edges   = make(map[string][]*Service)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   service.CustomerTable,
				Columns: service.CustomerPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(service.CustomerPrimaryKey[0], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := strconv.FormatInt(eout.Int64, 10)
				inValue := strconv.FormatInt(ein.Int64, 10)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, sq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "customer": %v`, err)
		}
		query.Where(customer.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "customer" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Customer = append(nodes[i].Edges.Customer, n)
			}
		}
	}

	if query := sq.withEndpoints; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*Service)
		for i := range nodes {
			id, err := strconv.Atoi(nodes[i].ID)
			if err != nil {
				return nil, err
			}
			fks = append(fks, id)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ServiceEndpoint(func(s *sql.Selector) {
			s.Where(sql.InValues(service.EndpointsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.service_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "service_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "service_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Endpoints = append(node.Edges.Endpoints, n)
		}
	}

	return nodes, nil
}

func (sq *ServiceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ServiceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sq *ServiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   service.Table,
			Columns: service.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: service.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ServiceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(service.Table)
	selector := builder.Select(t1.Columns(service.Columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(service.Columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServiceGroupBy is the builder for group-by Service entities.
type ServiceGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ServiceGroupBy) Aggregate(fns ...Aggregate) *ServiceGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scan the result into the given value.
func (sgb *ServiceGroupBy) Scan(ctx context.Context, v interface{}) error {
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *ServiceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sgb *ServiceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ServiceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *ServiceGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sgb *ServiceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ServiceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *ServiceGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sgb *ServiceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ServiceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *ServiceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sgb *ServiceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ServiceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *ServiceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *ServiceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgb.sqlQuery().Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *ServiceGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql
	columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
	columns = append(columns, sgb.fields...)
	for _, fn := range sgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sgb.fields...)
}

// ServiceSelect is the builder for select fields of Service entities.
type ServiceSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ss *ServiceSelect) Scan(ctx context.Context, v interface{}) error {
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *ServiceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ss *ServiceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ServiceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *ServiceSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ss *ServiceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ServiceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *ServiceSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ss *ServiceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ServiceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *ServiceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ss *ServiceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ServiceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *ServiceSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *ServiceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sqlQuery().Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ss *ServiceSelect) sqlQuery() sql.Querier {
	selector := ss.sql
	selector.Select(selector.Columns(ss.fields...)...)
	return selector
}
