// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/attachedfile"
	"entgo.io/ent/entc/integration/edgeschema/ent/file"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/process"
	"entgo.io/ent/schema/field"
)

// ProcessQuery is the builder for querying Process entities.
type ProcessQuery struct {
	config
	ctx               *QueryContext
	order             []process.OrderOption
	inters            []Interceptor
	predicates        []predicate.Process
	withFiles         *FileQuery
	withAttachedFiles *AttachedFileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProcessQuery builder.
func (_q *ProcessQuery) Where(ps ...predicate.Process) *ProcessQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *ProcessQuery) Limit(limit int) *ProcessQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *ProcessQuery) Offset(offset int) *ProcessQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *ProcessQuery) Unique(unique bool) *ProcessQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *ProcessQuery) Order(o ...process.OrderOption) *ProcessQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryFiles chains the current query on the "files" edge.
func (_q *ProcessQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(process.Table, process.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, process.FilesTable, process.FilesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttachedFiles chains the current query on the "attached_files" edge.
func (_q *ProcessQuery) QueryAttachedFiles() *AttachedFileQuery {
	query := (&AttachedFileClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(process.Table, process.FieldID, selector),
			sqlgraph.To(attachedfile.Table, attachedfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, process.AttachedFilesTable, process.AttachedFilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Process entity from the query.
// Returns a *NotFoundError when no Process was found.
func (_q *ProcessQuery) First(ctx context.Context) (*Process, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{process.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *ProcessQuery) FirstX(ctx context.Context) *Process {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Process ID from the query.
// Returns a *NotFoundError when no Process ID was found.
func (_q *ProcessQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{process.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *ProcessQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Process entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Process entity is found.
// Returns a *NotFoundError when no Process entities are found.
func (_q *ProcessQuery) Only(ctx context.Context) (*Process, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{process.Label}
	default:
		return nil, &NotSingularError{process.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *ProcessQuery) OnlyX(ctx context.Context) *Process {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Process ID in the query.
// Returns a *NotSingularError when more than one Process ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *ProcessQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{process.Label}
	default:
		err = &NotSingularError{process.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *ProcessQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Processes.
func (_q *ProcessQuery) All(ctx context.Context) ([]*Process, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Process, *ProcessQuery]()
	return withInterceptors[[]*Process](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *ProcessQuery) AllX(ctx context.Context) []*Process {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Process IDs.
func (_q *ProcessQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(process.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *ProcessQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *ProcessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*ProcessQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *ProcessQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *ProcessQuery) Exist(ctx context.Context) (bool, error) {
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
func (_q *ProcessQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProcessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *ProcessQuery) Clone() *ProcessQuery {
	if _q == nil {
		return nil
	}
	return &ProcessQuery{
		config:            _q.config,
		ctx:               _q.ctx.Clone(),
		order:             append([]process.OrderOption{}, _q.order...),
		inters:            append([]Interceptor{}, _q.inters...),
		predicates:        append([]predicate.Process{}, _q.predicates...),
		withFiles:         _q.withFiles.Clone(),
		withAttachedFiles: _q.withAttachedFiles.Clone(),
		// clone intermediate query.
		sql:  _q.sql.Clone(),
		path: _q.path,
	}
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *ProcessQuery) WithFiles(opts ...func(*FileQuery)) *ProcessQuery {
	query := (&FileClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withFiles = query
	return _q
}

// WithAttachedFiles tells the query-builder to eager-load the nodes that are connected to
// the "attached_files" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *ProcessQuery) WithAttachedFiles(opts ...func(*AttachedFileQuery)) *ProcessQuery {
	query := (&AttachedFileClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withAttachedFiles = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (_q *ProcessQuery) GroupBy(field string, fields ...string) *ProcessGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProcessGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = process.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (_q *ProcessQuery) Select(fields ...string) *ProcessSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &ProcessSelect{ProcessQuery: _q}
	sbuild.label = process.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProcessSelect configured with the given aggregations.
func (_q *ProcessQuery) Aggregate(fns ...AggregateFunc) *ProcessSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *ProcessQuery) prepareQuery(ctx context.Context) error {
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
		if !process.ValidColumn(f) {
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
	return nil
}

func (_q *ProcessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Process, error) {
	var (
		nodes       = []*Process{}
		_spec       = _q.querySpec()
		loadedTypes = [2]bool{
			_q.withFiles != nil,
			_q.withAttachedFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Process).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Process{config: _q.config}
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
	if query := _q.withFiles; query != nil {
		if err := _q.loadFiles(ctx, query, nodes,
			func(n *Process) { n.Edges.Files = []*File{} },
			func(n *Process, e *File) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withAttachedFiles; query != nil {
		if err := _q.loadAttachedFiles(ctx, query, nodes,
			func(n *Process) { n.Edges.AttachedFiles = []*AttachedFile{} },
			func(n *Process, e *AttachedFile) { n.Edges.AttachedFiles = append(n.Edges.AttachedFiles, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *ProcessQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*Process, init func(*Process), assign func(*Process, *File)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Process)
	nids := make(map[int]map[*Process]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(process.FilesTable)
		s.Join(joinT).On(s.C(file.FieldID), joinT.C(process.FilesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(process.FilesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(process.FilesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Process]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*File](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "files" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (_q *ProcessQuery) loadAttachedFiles(ctx context.Context, query *AttachedFileQuery, nodes []*Process, init func(*Process), assign func(*Process, *AttachedFile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Process)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(attachedfile.FieldProcID)
	}
	query.Where(predicate.AttachedFile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(process.AttachedFilesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProcID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "proc_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (_q *ProcessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *ProcessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(process.Table, process.Columns, sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, process.FieldID)
		for i := range fields {
			if fields[i] != process.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
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

func (_q *ProcessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(process.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = process.Columns
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

// ProcessGroupBy is the group-by builder for Process entities.
type ProcessGroupBy struct {
	selector
	build *ProcessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (_g *ProcessGroupBy) Aggregate(fns ...AggregateFunc) *ProcessGroupBy {
	_g.fns = append(_g.fns, fns...)
	return _g
}

// Scan applies the selector query and scans the result into the given value.
func (_g *ProcessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _g.build.ctx, ent.OpQueryGroupBy)
	if err := _g.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProcessQuery, *ProcessGroupBy](ctx, _g.build, _g, _g.build.inters, v)
}

func (_g *ProcessGroupBy) sqlScan(ctx context.Context, root *ProcessQuery, v any) error {
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

// ProcessSelect is the builder for selecting fields of Process entities.
type ProcessSelect struct {
	*ProcessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (_s *ProcessSelect) Aggregate(fns ...AggregateFunc) *ProcessSelect {
	_s.fns = append(_s.fns, fns...)
	return _s
}

// Scan applies the selector query and scans the result into the given value.
func (_s *ProcessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _s.ctx, ent.OpQuerySelect)
	if err := _s.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProcessQuery, *ProcessSelect](ctx, _s.ProcessQuery, _s, _s.inters, v)
}

func (_s *ProcessSelect) sqlScan(ctx context.Context, root *ProcessQuery, v any) error {
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
