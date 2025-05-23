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
	"entgo.io/ent/entc/integration/gremlin/ent/builder"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// BuilderQuery is the builder for querying Builder entities.
type BuilderQuery struct {
	config
	ctx        *QueryContext
	order      []builder.OrderOption
	inters     []Interceptor
	predicates []predicate.Builder
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the BuilderQuery builder.
func (_q *BuilderQuery) Where(ps ...predicate.Builder) *BuilderQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *BuilderQuery) Limit(limit int) *BuilderQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *BuilderQuery) Offset(offset int) *BuilderQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *BuilderQuery) Unique(unique bool) *BuilderQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *BuilderQuery) Order(o ...builder.OrderOption) *BuilderQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// First returns the first Builder entity from the query.
// Returns a *NotFoundError when no Builder was found.
func (_q *BuilderQuery) First(ctx context.Context) (*Builder, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{builder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *BuilderQuery) FirstX(ctx context.Context) *Builder {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Builder ID from the query.
// Returns a *NotFoundError when no Builder ID was found.
func (_q *BuilderQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{builder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *BuilderQuery) FirstIDX(ctx context.Context) string {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Builder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Builder entity is found.
// Returns a *NotFoundError when no Builder entities are found.
func (_q *BuilderQuery) Only(ctx context.Context) (*Builder, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{builder.Label}
	default:
		return nil, &NotSingularError{builder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *BuilderQuery) OnlyX(ctx context.Context) *Builder {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Builder ID in the query.
// Returns a *NotSingularError when more than one Builder ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *BuilderQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{builder.Label}
	default:
		err = &NotSingularError{builder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *BuilderQuery) OnlyIDX(ctx context.Context) string {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Builders.
func (_q *BuilderQuery) All(ctx context.Context) ([]*Builder, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Builder, *BuilderQuery]()
	return withInterceptors[[]*Builder](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *BuilderQuery) AllX(ctx context.Context) []*Builder {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Builder IDs.
func (_q *BuilderQuery) IDs(ctx context.Context) (ids []string, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(builder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *BuilderQuery) IDsX(ctx context.Context) []string {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *BuilderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*BuilderQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *BuilderQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *BuilderQuery) Exist(ctx context.Context) (bool, error) {
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
func (_q *BuilderQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BuilderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *BuilderQuery) Clone() *BuilderQuery {
	if _q == nil {
		return nil
	}
	return &BuilderQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]builder.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Builder{}, _q.predicates...),
		// clone intermediate query.
		gremlin: _q.gremlin.Clone(),
		path:    _q.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (_q *BuilderQuery) GroupBy(field string, fields ...string) *BuilderGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BuilderGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = builder.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (_q *BuilderQuery) Select(fields ...string) *BuilderSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &BuilderSelect{BuilderQuery: _q}
	sbuild.label = builder.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BuilderSelect configured with the given aggregations.
func (_q *BuilderQuery) Aggregate(fns ...AggregateFunc) *BuilderSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *BuilderQuery) prepareQuery(ctx context.Context) error {
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

func (_q *BuilderQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Builder, error) {
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
	var _ms Builders
	if err := _ms.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range _ms {
		_ms[i].config = _q.config
	}
	return _ms, nil
}

func (_q *BuilderQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _q.gremlinQuery(ctx).Count().Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (_q *BuilderQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(builder.Label)
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

// BuilderGroupBy is the group-by builder for Builder entities.
type BuilderGroupBy struct {
	selector
	build *BuilderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (_g *BuilderGroupBy) Aggregate(fns ...AggregateFunc) *BuilderGroupBy {
	_g.fns = append(_g.fns, fns...)
	return _g
}

// Scan applies the selector query and scans the result into the given value.
func (_g *BuilderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _g.build.ctx, ent.OpQueryGroupBy)
	if err := _g.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BuilderQuery, *BuilderGroupBy](ctx, _g.build, _g, _g.build.inters, v)
}

func (_g *BuilderGroupBy) gremlinScan(ctx context.Context, root *BuilderQuery, v any) error {
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

// BuilderSelect is the builder for selecting fields of Builder entities.
type BuilderSelect struct {
	*BuilderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (_s *BuilderSelect) Aggregate(fns ...AggregateFunc) *BuilderSelect {
	_s.fns = append(_s.fns, fns...)
	return _s
}

// Scan applies the selector query and scans the result into the given value.
func (_s *BuilderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _s.ctx, ent.OpQuerySelect)
	if err := _s.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BuilderQuery, *BuilderSelect](ctx, _s.BuilderQuery, _s, _s.inters, v)
}

func (_s *BuilderSelect) gremlinScan(ctx context.Context, root *BuilderQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := _s.ctx.Fields; len(fields) == 1 {
		if fields[0] != builder.FieldID {
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
