// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/a8m/entspatial/ent/migrate"

	"github.com/a8m/entspatial/ent/location"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Location is the client for interacting with the Location builders.
	Location *LocationClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Location = NewLocationClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Location: NewLocationClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		Location: NewLocationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Location.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Location.Use(hooks...)
}

// LocationClient is a client for the Location schema.
type LocationClient struct {
	config
}

// NewLocationClient returns a client for the Location from the given config.
func NewLocationClient(c config) *LocationClient {
	return &LocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `location.Hooks(f(g(h())))`.
func (c *LocationClient) Use(hooks ...Hook) {
	c.hooks.Location = append(c.hooks.Location, hooks...)
}

// Create returns a create builder for Location.
func (c *LocationClient) Create() *LocationCreate {
	mutation := newLocationMutation(c.config, OpCreate)
	return &LocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Location entities.
func (c *LocationClient) CreateBulk(builders ...*LocationCreate) *LocationCreateBulk {
	return &LocationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Location.
func (c *LocationClient) Update() *LocationUpdate {
	mutation := newLocationMutation(c.config, OpUpdate)
	return &LocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LocationClient) UpdateOne(l *Location) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocation(l))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LocationClient) UpdateOneID(id int) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocationID(id))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Location.
func (c *LocationClient) Delete() *LocationDelete {
	mutation := newLocationMutation(c.config, OpDelete)
	return &LocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LocationClient) DeleteOne(l *Location) *LocationDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LocationClient) DeleteOneID(id int) *LocationDeleteOne {
	builder := c.Delete().Where(location.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LocationDeleteOne{builder}
}

// Query returns a query builder for Location.
func (c *LocationClient) Query() *LocationQuery {
	return &LocationQuery{config: c.config}
}

// Get returns a Location entity by its id.
func (c *LocationClient) Get(ctx context.Context, id int) (*Location, error) {
	return c.Query().Where(location.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LocationClient) GetX(ctx context.Context, id int) *Location {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Location.
func (c *LocationClient) QueryParent(l *Location) *LocationQuery {
	query := &LocationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, id),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, location.ParentTable, location.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Location.
func (c *LocationClient) QueryChildren(l *Location) *LocationQuery {
	query := &LocationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, id),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, location.ChildrenTable, location.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LocationClient) Hooks() []Hook {
	return c.hooks.Location
}
