// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"minimal/ent/water"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WaterCreate is the builder for creating a Water entity.
type WaterCreate struct {
	config
	mutation *WaterMutation
	hooks    []Hook
}

// SetLiters sets the "liters" field.
func (wc *WaterCreate) SetLiters(f float64) *WaterCreate {
	wc.mutation.SetLiters(f)
	return wc
}

// SetTopic sets the "topic" field.
func (wc *WaterCreate) SetTopic(s string) *WaterCreate {
	wc.mutation.SetTopic(s)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WaterCreate) SetCreatedAt(t time.Time) *WaterCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WaterCreate) SetNillableCreatedAt(t *time.Time) *WaterCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// Mutation returns the WaterMutation object of the builder.
func (wc *WaterCreate) Mutation() *WaterMutation {
	return wc.mutation
}

// Save creates the Water in the database.
func (wc *WaterCreate) Save(ctx context.Context) (*Water, error) {
	var (
		err  error
		node *Water
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WaterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Water)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WaterMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WaterCreate) SaveX(ctx context.Context) *Water {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WaterCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WaterCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WaterCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := water.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WaterCreate) check() error {
	if _, ok := wc.mutation.Liters(); !ok {
		return &ValidationError{Name: "liters", err: errors.New(`ent: missing required field "Water.liters"`)}
	}
	if _, ok := wc.mutation.Topic(); !ok {
		return &ValidationError{Name: "topic", err: errors.New(`ent: missing required field "Water.topic"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Water.created_at"`)}
	}
	return nil
}

func (wc *WaterCreate) sqlSave(ctx context.Context) (*Water, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wc *WaterCreate) createSpec() (*Water, *sqlgraph.CreateSpec) {
	var (
		_node = &Water{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: water.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: water.FieldID,
			},
		}
	)
	if value, ok := wc.mutation.Liters(); ok {
		_spec.SetField(water.FieldLiters, field.TypeFloat64, value)
		_node.Liters = value
	}
	if value, ok := wc.mutation.Topic(); ok {
		_spec.SetField(water.FieldTopic, field.TypeString, value)
		_node.Topic = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(water.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// WaterCreateBulk is the builder for creating many Water entities in bulk.
type WaterCreateBulk struct {
	config
	builders []*WaterCreate
}

// Save creates the Water entities in the database.
func (wcb *WaterCreateBulk) Save(ctx context.Context) ([]*Water, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Water, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WaterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WaterCreateBulk) SaveX(ctx context.Context) []*Water {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WaterCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WaterCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
