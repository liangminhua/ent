// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationship"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
)

// RelationshipQuery is the builder for querying Relationship entities.
type RelationshipQuery struct {
	config
	ctx          *QueryContext
	order        []relationship.OrderOption
	inters       []Interceptor
	predicates   []predicate.Relationship
	withUser     *UserQuery
	withRelative *UserQuery
	withInfo     *RelationshipInfoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RelationshipQuery builder.
func (_q *RelationshipQuery) Where(ps ...predicate.Relationship) *RelationshipQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *RelationshipQuery) Limit(limit int) *RelationshipQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *RelationshipQuery) Offset(offset int) *RelationshipQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *RelationshipQuery) Unique(unique bool) *RelationshipQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *RelationshipQuery) Order(o ...relationship.OrderOption) *RelationshipQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryUser chains the current query on the "user" edge.
func (_q *RelationshipQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relationship.Table, relationship.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, relationship.UserTable, relationship.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRelative chains the current query on the "relative" edge.
func (_q *RelationshipQuery) QueryRelative() *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relationship.Table, relationship.RelativeColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, relationship.RelativeTable, relationship.RelativeColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInfo chains the current query on the "info" edge.
func (_q *RelationshipQuery) QueryInfo() *RelationshipInfoQuery {
	query := (&RelationshipInfoClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relationship.Table, relationship.InfoColumn, selector),
			sqlgraph.To(relationshipinfo.Table, relationshipinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, relationship.InfoTable, relationship.InfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Relationship entity from the query.
// Returns a *NotFoundError when no Relationship was found.
func (_q *RelationshipQuery) First(ctx context.Context) (*Relationship, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{relationship.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *RelationshipQuery) FirstX(ctx context.Context) *Relationship {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single Relationship entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Relationship entity is found.
// Returns a *NotFoundError when no Relationship entities are found.
func (_q *RelationshipQuery) Only(ctx context.Context) (*Relationship, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{relationship.Label}
	default:
		return nil, &NotSingularError{relationship.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *RelationshipQuery) OnlyX(ctx context.Context) *Relationship {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of Relationships.
func (_q *RelationshipQuery) All(ctx context.Context) ([]*Relationship, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Relationship, *RelationshipQuery]()
	return withInterceptors[[]*Relationship](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *RelationshipQuery) AllX(ctx context.Context) []*Relationship {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (_q *RelationshipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*RelationshipQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *RelationshipQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *RelationshipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *RelationshipQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RelationshipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *RelationshipQuery) Clone() *RelationshipQuery {
	if _q == nil {
		return nil
	}
	return &RelationshipQuery{
		config:       _q.config,
		ctx:          _q.ctx.Clone(),
		order:        append([]relationship.OrderOption{}, _q.order...),
		inters:       append([]Interceptor{}, _q.inters...),
		predicates:   append([]predicate.Relationship{}, _q.predicates...),
		withUser:     _q.withUser.Clone(),
		withRelative: _q.withRelative.Clone(),
		withInfo:     _q.withInfo.Clone(),
		// clone intermediate query.
		sql:  _q.sql.Clone(),
		path: _q.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *RelationshipQuery) WithUser(opts ...func(*UserQuery)) *RelationshipQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withUser = query
	return _q
}

// WithRelative tells the query-builder to eager-load the nodes that are connected to
// the "relative" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *RelationshipQuery) WithRelative(opts ...func(*UserQuery)) *RelationshipQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withRelative = query
	return _q
}

// WithInfo tells the query-builder to eager-load the nodes that are connected to
// the "info" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *RelationshipQuery) WithInfo(opts ...func(*RelationshipInfoQuery)) *RelationshipQuery {
	query := (&RelationshipInfoClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withInfo = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Weight int `json:"weight,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Relationship.Query().
//		GroupBy(relationship.FieldWeight).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *RelationshipQuery) GroupBy(field string, fields ...string) *RelationshipGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RelationshipGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = relationship.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Weight int `json:"weight,omitempty"`
//	}
//
//	client.Relationship.Query().
//		Select(relationship.FieldWeight).
//		Scan(ctx, &v)
func (_q *RelationshipQuery) Select(fields ...string) *RelationshipSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &RelationshipSelect{RelationshipQuery: _q}
	sbuild.label = relationship.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RelationshipSelect configured with the given aggregations.
func (_q *RelationshipQuery) Aggregate(fns ...AggregateFunc) *RelationshipSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *RelationshipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	for _, f := range _q.ctx.Fields {
		if !relationship.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.sql = prev
	}
	if relationship.Policy == nil {
		return errors.New("ent: uninitialized relationship.Policy (forgotten import ent/runtime?)")
	}
	if err := relationship.Policy.EvalQuery(ctx, _q); err != nil {
		return err
	}
	return nil
}

func (_q *RelationshipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Relationship, error) {
	var (
		nodes       = []*Relationship{}
		_spec       = _q.querySpec()
		loadedTypes = [3]bool{
			_q.withUser != nil,
			_q.withRelative != nil,
			_q.withInfo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Relationship).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Relationship{config: _q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, _q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := _q.withUser; query != nil {
		if err := _q.loadUser(ctx, query, nodes, nil,
			func(n *Relationship, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := _q.withRelative; query != nil {
		if err := _q.loadRelative(ctx, query, nodes, nil,
			func(n *Relationship, e *User) { n.Edges.Relative = e }); err != nil {
			return nil, err
		}
	}
	if query := _q.withInfo; query != nil {
		if err := _q.loadInfo(ctx, query, nodes, nil,
			func(n *Relationship, e *RelationshipInfo) { n.Edges.Info = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *RelationshipQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Relationship, init func(*Relationship), assign func(*Relationship, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Relationship)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (_q *RelationshipQuery) loadRelative(ctx context.Context, query *UserQuery, nodes []*Relationship, init func(*Relationship), assign func(*Relationship, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Relationship)
	for i := range nodes {
		fk := nodes[i].RelativeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "relative_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (_q *RelationshipQuery) loadInfo(ctx context.Context, query *RelationshipInfoQuery, nodes []*Relationship, init func(*Relationship), assign func(*Relationship, *RelationshipInfo)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Relationship)
	for i := range nodes {
		fk := nodes[i].InfoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(relationshipinfo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "info_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (_q *RelationshipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *RelationshipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(relationship.Table, relationship.Columns, nil)
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if _q.withUser != nil {
			_spec.Node.AddColumnOnce(relationship.FieldUserID)
		}
		if _q.withRelative != nil {
			_spec.Node.AddColumnOnce(relationship.FieldRelativeID)
		}
		if _q.withInfo != nil {
			_spec.Node.AddColumnOnce(relationship.FieldInfoID)
		}
	}
	if ps := _q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := _q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := _q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := _q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (_q *RelationshipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(relationship.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = relationship.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range _q.predicates {
		p(selector)
	}
	for _, p := range _q.order {
		p(selector)
	}
	if offset := _q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := _q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RelationshipGroupBy is the group-by builder for Relationship entities.
type RelationshipGroupBy struct {
	selector
	build *RelationshipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (_g *RelationshipGroupBy) Aggregate(fns ...AggregateFunc) *RelationshipGroupBy {
	_g.fns = append(_g.fns, fns...)
	return _g
}

// Scan applies the selector query and scans the result into the given value.
func (_g *RelationshipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _g.build.ctx, ent.OpQueryGroupBy)
	if err := _g.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RelationshipQuery, *RelationshipGroupBy](ctx, _g.build, _g, _g.build.inters, v)
}

func (_g *RelationshipGroupBy) sqlScan(ctx context.Context, root *RelationshipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(_g.fns))
	for _, fn := range _g.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*_g.flds)+len(_g.fns))
		for _, f := range *_g.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*_g.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := _g.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RelationshipSelect is the builder for selecting fields of Relationship entities.
type RelationshipSelect struct {
	*RelationshipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (_s *RelationshipSelect) Aggregate(fns ...AggregateFunc) *RelationshipSelect {
	_s.fns = append(_s.fns, fns...)
	return _s
}

// Scan applies the selector query and scans the result into the given value.
func (_s *RelationshipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _s.ctx, ent.OpQuerySelect)
	if err := _s.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RelationshipQuery, *RelationshipSelect](ctx, _s.RelationshipQuery, _s, _s.inters, v)
}

func (_s *RelationshipSelect) sqlScan(ctx context.Context, root *RelationshipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(_s.fns))
	for _, fn := range _s.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*_s.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := _s.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
