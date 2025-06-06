// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/license"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// LicenseQuery is the builder for querying License entities.
type LicenseQuery struct {
	config
	ctx        *QueryContext
	order      []license.OrderOption
	inters     []Interceptor
	predicates []predicate.License
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the LicenseQuery builder.
func (_q *LicenseQuery) Where(ps ...predicate.License) *LicenseQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *LicenseQuery) Limit(limit int) *LicenseQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *LicenseQuery) Offset(offset int) *LicenseQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *LicenseQuery) Unique(unique bool) *LicenseQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *LicenseQuery) Order(o ...license.OrderOption) *LicenseQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// First returns the first License entity from the query.
// Returns a *NotFoundError when no License was found.
func (_q *LicenseQuery) First(ctx context.Context) (*License, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{license.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *LicenseQuery) FirstX(ctx context.Context) *License {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first License ID from the query.
// Returns a *NotFoundError when no License ID was found.
func (_q *LicenseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{license.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *LicenseQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single License entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one License entity is found.
// Returns a *NotFoundError when no License entities are found.
func (_q *LicenseQuery) Only(ctx context.Context) (*License, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{license.Label}
	default:
		return nil, &NotSingularError{license.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *LicenseQuery) OnlyX(ctx context.Context) *License {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only License ID in the query.
// Returns a *NotSingularError when more than one License ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *LicenseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{license.Label}
	default:
		err = &NotSingularError{license.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *LicenseQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Licenses.
func (_q *LicenseQuery) All(ctx context.Context) ([]*License, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*License, *LicenseQuery]()
	return withInterceptors[[]*License](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *LicenseQuery) AllX(ctx context.Context) []*License {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of License IDs.
func (_q *LicenseQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(license.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *LicenseQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *LicenseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*LicenseQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *LicenseQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *LicenseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *LicenseQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LicenseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *LicenseQuery) Clone() *LicenseQuery {
	if _q == nil {
		return nil
	}
	return &LicenseQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]license.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.License{}, _q.predicates...),
		// clone intermediate query.
		gremlin: _q.gremlin.Clone(),
		path:    _q.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.License.Query().
//		GroupBy(license.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *LicenseQuery) GroupBy(field string, fields ...string) *LicenseGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LicenseGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = license.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.License.Query().
//		Select(license.FieldCreateTime).
//		Scan(ctx, &v)
func (_q *LicenseQuery) Select(fields ...string) *LicenseSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &LicenseSelect{LicenseQuery: _q}
	sbuild.label = license.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LicenseSelect configured with the given aggregations.
func (_q *LicenseQuery) Aggregate(fns ...AggregateFunc) *LicenseSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *LicenseQuery) prepareQuery(ctx context.Context) error {
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
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.gremlin = prev
	}
	return nil
}

func (_q *LicenseQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*License, error) {
	res := &gremlin.Response{}
	traversal := _q.gremlinQuery(ctx)
	if len(_q.ctx.Fields) > 0 {
		fields := make([]any, len(_q.ctx.Fields))
		for i, f := range _q.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var _ms Licenses
	if err := _ms.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range _ms {
		_ms[i].config = _q.config
	}
	return _ms, nil
}

func (_q *LicenseQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _q.gremlinQuery(ctx).Count().Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (_q *LicenseQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(license.Label)
	if _q.gremlin != nil {
		v = _q.gremlin.Clone()
	}
	for _, p := range _q.predicates {
		p(v)
	}
	if len(_q.order) > 0 {
		v.Order()
		for _, p := range _q.order {
			p(v)
		}
	}
	switch limit, offset := _q.ctx.Limit, _q.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := _q.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// LicenseGroupBy is the group-by builder for License entities.
type LicenseGroupBy struct {
	selector
	build *LicenseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (_g *LicenseGroupBy) Aggregate(fns ...AggregateFunc) *LicenseGroupBy {
	_g.fns = append(_g.fns, fns...)
	return _g
}

// Scan applies the selector query and scans the result into the given value.
func (_g *LicenseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _g.build.ctx, ent.OpQueryGroupBy)
	if err := _g.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LicenseQuery, *LicenseGroupBy](ctx, _g.build, _g, _g.build.inters, v)
}

func (_g *LicenseGroupBy) gremlinScan(ctx context.Context, root *LicenseQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range _g.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *_g.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*_g.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := _g.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*_g.flds)+len(_g.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// LicenseSelect is the builder for selecting fields of License entities.
type LicenseSelect struct {
	*LicenseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (_s *LicenseSelect) Aggregate(fns ...AggregateFunc) *LicenseSelect {
	_s.fns = append(_s.fns, fns...)
	return _s
}

// Scan applies the selector query and scans the result into the given value.
func (_s *LicenseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _s.ctx, ent.OpQuerySelect)
	if err := _s.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LicenseQuery, *LicenseSelect](ctx, _s.LicenseQuery, _s, _s.inters, v)
}

func (_s *LicenseSelect) gremlinScan(ctx context.Context, root *LicenseQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := _s.ctx.Fields; len(fields) == 1 {
		if fields[0] != license.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(_s.ctx.Fields))
		for i, f := range _s.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := _s.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
