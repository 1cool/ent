// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/privacy/ent/planet"
	"github.com/facebook/ent/schema/field"
)

// PlanetCreate is the builder for creating a Planet entity.
type PlanetCreate struct {
	config
	mutation *PlanetMutation
	hooks    []Hook
}

// SetName sets the name field.
func (pc *PlanetCreate) SetName(s string) *PlanetCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetAge sets the age field.
func (pc *PlanetCreate) SetAge(u uint) *PlanetCreate {
	pc.mutation.SetAge(u)
	return pc
}

// SetNillableAge sets the age field if the given value is not nil.
func (pc *PlanetCreate) SetNillableAge(u *uint) *PlanetCreate {
	if u != nil {
		pc.SetAge(*u)
	}
	return pc
}

// AddNeighborIDs adds the neighbors edge to Planet by ids.
func (pc *PlanetCreate) AddNeighborIDs(ids ...int) *PlanetCreate {
	pc.mutation.AddNeighborIDs(ids...)
	return pc
}

// AddNeighbors adds the neighbors edges to Planet.
func (pc *PlanetCreate) AddNeighbors(p ...*Planet) *PlanetCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddNeighborIDs(ids...)
}

// Mutation returns the PlanetMutation object of the builder.
func (pc *PlanetCreate) Mutation() *PlanetMutation {
	return pc.mutation
}

// Save creates the Planet in the database.
func (pc *PlanetCreate) Save(ctx context.Context) (*Planet, error) {
	var (
		err  error
		node *Planet
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlanetCreate) SaveX(ctx context.Context) *Planet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlanetCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := planet.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (pc *PlanetCreate) sqlSave(ctx context.Context) (*Planet, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PlanetCreate) createSpec() (*Planet, *sqlgraph.CreateSpec) {
	var (
		_node = &Planet{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: planet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: planet.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: planet.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: planet.FieldAge,
		})
		_node.Age = value
	}
	if nodes := pc.mutation.NeighborsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   planet.NeighborsTable,
			Columns: planet.NeighborsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlanetCreateBulk is the builder for creating a bulk of Planet entities.
type PlanetCreateBulk struct {
	config
	builders []*PlanetCreate
}

// Save creates the Planet entities in the database.
func (pcb *PlanetCreateBulk) Save(ctx context.Context) ([]*Planet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Planet, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlanetMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pcb *PlanetCreateBulk) SaveX(ctx context.Context) []*Planet {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
