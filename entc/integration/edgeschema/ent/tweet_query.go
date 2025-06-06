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
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweetlike"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
)

// TweetQuery is the builder for querying Tweet entities.
type TweetQuery struct {
	config
	ctx            *QueryContext
	order          []tweet.OrderOption
	inters         []Interceptor
	predicates     []predicate.Tweet
	withLikedUsers *UserQuery
	withUser       *UserQuery
	withTags       *TagQuery
	withLikes      *TweetLikeQuery
	withTweetUser  *UserTweetQuery
	withTweetTags  *TweetTagQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TweetQuery builder.
func (_q *TweetQuery) Where(ps ...predicate.Tweet) *TweetQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *TweetQuery) Limit(limit int) *TweetQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *TweetQuery) Offset(offset int) *TweetQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *TweetQuery) Unique(unique bool) *TweetQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *TweetQuery) Order(o ...tweet.OrderOption) *TweetQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryLikedUsers chains the current query on the "liked_users" edge.
func (_q *TweetQuery) QueryLikedUsers() *UserQuery {
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
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tweet.LikedUsersTable, tweet.LikedUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (_q *TweetQuery) QueryUser() *UserQuery {
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
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tweet.UserTable, tweet.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (_q *TweetQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tweet.TagsTable, tweet.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLikes chains the current query on the "likes" edge.
func (_q *TweetQuery) QueryLikes() *TweetLikeQuery {
	query := (&TweetLikeClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(tweetlike.Table, tweetlike.TweetColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, tweet.LikesTable, tweet.LikesColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweetUser chains the current query on the "tweet_user" edge.
func (_q *TweetQuery) QueryTweetUser() *UserTweetQuery {
	query := (&UserTweetClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(usertweet.Table, usertweet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, tweet.TweetUserTable, tweet.TweetUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweetTags chains the current query on the "tweet_tags" edge.
func (_q *TweetQuery) QueryTweetTags() *TweetTagQuery {
	query := (&TweetTagClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, selector),
			sqlgraph.To(tweettag.Table, tweettag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, tweet.TweetTagsTable, tweet.TweetTagsColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tweet entity from the query.
// Returns a *NotFoundError when no Tweet was found.
func (_q *TweetQuery) First(ctx context.Context) (*Tweet, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tweet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *TweetQuery) FirstX(ctx context.Context) *Tweet {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tweet ID from the query.
// Returns a *NotFoundError when no Tweet ID was found.
func (_q *TweetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tweet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *TweetQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tweet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tweet entity is found.
// Returns a *NotFoundError when no Tweet entities are found.
func (_q *TweetQuery) Only(ctx context.Context) (*Tweet, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tweet.Label}
	default:
		return nil, &NotSingularError{tweet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *TweetQuery) OnlyX(ctx context.Context) *Tweet {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tweet ID in the query.
// Returns a *NotSingularError when more than one Tweet ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *TweetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tweet.Label}
	default:
		err = &NotSingularError{tweet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *TweetQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tweets.
func (_q *TweetQuery) All(ctx context.Context) ([]*Tweet, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Tweet, *TweetQuery]()
	return withInterceptors[[]*Tweet](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *TweetQuery) AllX(ctx context.Context) []*Tweet {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tweet IDs.
func (_q *TweetQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(tweet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *TweetQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *TweetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*TweetQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *TweetQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *TweetQuery) Exist(ctx context.Context) (bool, error) {
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
func (_q *TweetQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TweetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *TweetQuery) Clone() *TweetQuery {
	if _q == nil {
		return nil
	}
	return &TweetQuery{
		config:         _q.config,
		ctx:            _q.ctx.Clone(),
		order:          append([]tweet.OrderOption{}, _q.order...),
		inters:         append([]Interceptor{}, _q.inters...),
		predicates:     append([]predicate.Tweet{}, _q.predicates...),
		withLikedUsers: _q.withLikedUsers.Clone(),
		withUser:       _q.withUser.Clone(),
		withTags:       _q.withTags.Clone(),
		withLikes:      _q.withLikes.Clone(),
		withTweetUser:  _q.withTweetUser.Clone(),
		withTweetTags:  _q.withTweetTags.Clone(),
		// clone intermediate query.
		sql:  _q.sql.Clone(),
		path: _q.path,
	}
}

// WithLikedUsers tells the query-builder to eager-load the nodes that are connected to
// the "liked_users" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *TweetQuery) WithLikedUsers(opts ...func(*UserQuery)) *TweetQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withLikedUsers = query
	return _q
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *TweetQuery) WithUser(opts ...func(*UserQuery)) *TweetQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withUser = query
	return _q
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *TweetQuery) WithTags(opts ...func(*TagQuery)) *TweetQuery {
	query := (&TagClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withTags = query
	return _q
}

// WithLikes tells the query-builder to eager-load the nodes that are connected to
// the "likes" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *TweetQuery) WithLikes(opts ...func(*TweetLikeQuery)) *TweetQuery {
	query := (&TweetLikeClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withLikes = query
	return _q
}

// WithTweetUser tells the query-builder to eager-load the nodes that are connected to
// the "tweet_user" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *TweetQuery) WithTweetUser(opts ...func(*UserTweetQuery)) *TweetQuery {
	query := (&UserTweetClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withTweetUser = query
	return _q
}

// WithTweetTags tells the query-builder to eager-load the nodes that are connected to
// the "tweet_tags" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *TweetQuery) WithTweetTags(opts ...func(*TweetTagQuery)) *TweetQuery {
	query := (&TweetTagClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withTweetTags = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Tweet.Query().
//		GroupBy(tweet.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *TweetQuery) GroupBy(field string, fields ...string) *TweetGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TweetGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = tweet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Tweet.Query().
//		Select(tweet.FieldText).
//		Scan(ctx, &v)
func (_q *TweetQuery) Select(fields ...string) *TweetSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &TweetSelect{TweetQuery: _q}
	sbuild.label = tweet.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TweetSelect configured with the given aggregations.
func (_q *TweetQuery) Aggregate(fns ...AggregateFunc) *TweetSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *TweetQuery) prepareQuery(ctx context.Context) error {
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
		if !tweet.ValidColumn(f) {
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

func (_q *TweetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tweet, error) {
	var (
		nodes       = []*Tweet{}
		_spec       = _q.querySpec()
		loadedTypes = [6]bool{
			_q.withLikedUsers != nil,
			_q.withUser != nil,
			_q.withTags != nil,
			_q.withLikes != nil,
			_q.withTweetUser != nil,
			_q.withTweetTags != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Tweet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Tweet{config: _q.config}
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
	if query := _q.withLikedUsers; query != nil {
		if err := _q.loadLikedUsers(ctx, query, nodes,
			func(n *Tweet) { n.Edges.LikedUsers = []*User{} },
			func(n *Tweet, e *User) { n.Edges.LikedUsers = append(n.Edges.LikedUsers, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withUser; query != nil {
		if err := _q.loadUser(ctx, query, nodes,
			func(n *Tweet) { n.Edges.User = []*User{} },
			func(n *Tweet, e *User) { n.Edges.User = append(n.Edges.User, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withTags; query != nil {
		if err := _q.loadTags(ctx, query, nodes,
			func(n *Tweet) { n.Edges.Tags = []*Tag{} },
			func(n *Tweet, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withLikes; query != nil {
		if err := _q.loadLikes(ctx, query, nodes,
			func(n *Tweet) { n.Edges.Likes = []*TweetLike{} },
			func(n *Tweet, e *TweetLike) { n.Edges.Likes = append(n.Edges.Likes, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withTweetUser; query != nil {
		if err := _q.loadTweetUser(ctx, query, nodes,
			func(n *Tweet) { n.Edges.TweetUser = []*UserTweet{} },
			func(n *Tweet, e *UserTweet) { n.Edges.TweetUser = append(n.Edges.TweetUser, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withTweetTags; query != nil {
		if err := _q.loadTweetTags(ctx, query, nodes,
			func(n *Tweet) { n.Edges.TweetTags = []*TweetTag{} },
			func(n *Tweet, e *TweetTag) { n.Edges.TweetTags = append(n.Edges.TweetTags, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *TweetQuery) loadLikedUsers(ctx context.Context, query *UserQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Tweet)
	nids := make(map[int]map[*Tweet]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tweet.LikedUsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(tweet.LikedUsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tweet.LikedUsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tweet.LikedUsersPrimaryKey[1]))
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
					nids[inValue] = map[*Tweet]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "liked_users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (_q *TweetQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Tweet)
	nids := make(map[int]map[*Tweet]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tweet.UserTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(tweet.UserPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tweet.UserPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tweet.UserPrimaryKey[1]))
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
					nids[inValue] = map[*Tweet]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (_q *TweetQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Tweet)
	nids := make(map[int]map[*Tweet]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tweet.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(tweet.TagsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tweet.TagsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tweet.TagsPrimaryKey[1]))
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
					nids[inValue] = map[*Tweet]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (_q *TweetQuery) loadLikes(ctx context.Context, query *TweetLikeQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *TweetLike)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Tweet)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(tweetlike.FieldTweetID)
	}
	query.Where(predicate.TweetLike(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(tweet.LikesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TweetID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "tweet_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}
func (_q *TweetQuery) loadTweetUser(ctx context.Context, query *UserTweetQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *UserTweet)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Tweet)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(usertweet.FieldTweetID)
	}
	query.Where(predicate.UserTweet(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(tweet.TweetUserColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TweetID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "tweet_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (_q *TweetQuery) loadTweetTags(ctx context.Context, query *TweetTagQuery, nodes []*Tweet, init func(*Tweet), assign func(*Tweet, *TweetTag)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Tweet)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(tweettag.FieldTweetID)
	}
	query.Where(predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(tweet.TweetTagsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TweetID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "tweet_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (_q *TweetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *TweetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for i := range fields {
			if fields[i] != tweet.FieldID {
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

func (_q *TweetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(tweet.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = tweet.Columns
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

// TweetGroupBy is the group-by builder for Tweet entities.
type TweetGroupBy struct {
	selector
	build *TweetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (_g *TweetGroupBy) Aggregate(fns ...AggregateFunc) *TweetGroupBy {
	_g.fns = append(_g.fns, fns...)
	return _g
}

// Scan applies the selector query and scans the result into the given value.
func (_g *TweetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _g.build.ctx, ent.OpQueryGroupBy)
	if err := _g.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TweetQuery, *TweetGroupBy](ctx, _g.build, _g, _g.build.inters, v)
}

func (_g *TweetGroupBy) sqlScan(ctx context.Context, root *TweetQuery, v any) error {
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

// TweetSelect is the builder for selecting fields of Tweet entities.
type TweetSelect struct {
	*TweetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (_s *TweetSelect) Aggregate(fns ...AggregateFunc) *TweetSelect {
	_s.fns = append(_s.fns, fns...)
	return _s
}

// Scan applies the selector query and scans the result into the given value.
func (_s *TweetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _s.ctx, ent.OpQuerySelect)
	if err := _s.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TweetQuery, *TweetSelect](ctx, _s.TweetQuery, _s, _s.inters, v)
}

func (_s *TweetSelect) sqlScan(ctx context.Context, root *TweetQuery, v any) error {
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
