// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory/_examples/basic/ent/character"
	"github.com/flume/enthistory/_examples/basic/ent/friendship"
	"github.com/flume/enthistory/_examples/basic/ent/residence"
)

// CharacterCreate is the builder for creating a Character entity.
type CharacterCreate struct {
	config
	mutation *CharacterMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CharacterCreate) SetCreatedAt(t time.Time) *CharacterCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CharacterCreate) SetNillableCreatedAt(t *time.Time) *CharacterCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CharacterCreate) SetUpdatedAt(t time.Time) *CharacterCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CharacterCreate) SetNillableUpdatedAt(t *time.Time) *CharacterCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetAge sets the "age" field.
func (cc *CharacterCreate) SetAge(i int) *CharacterCreate {
	cc.mutation.SetAge(i)
	return cc
}

// SetName sets the "name" field.
func (cc *CharacterCreate) SetName(s string) *CharacterCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNicknames sets the "nicknames" field.
func (cc *CharacterCreate) SetNicknames(s []string) *CharacterCreate {
	cc.mutation.SetNicknames(s)
	return cc
}

// SetInfo sets the "info" field.
func (cc *CharacterCreate) SetInfo(m map[string]interface{}) *CharacterCreate {
	cc.mutation.SetInfo(m)
	return cc
}

// AddFriendIDs adds the "friends" edge to the Character entity by IDs.
func (cc *CharacterCreate) AddFriendIDs(ids ...int) *CharacterCreate {
	cc.mutation.AddFriendIDs(ids...)
	return cc
}

// AddFriends adds the "friends" edges to the Character entity.
func (cc *CharacterCreate) AddFriends(c ...*Character) *CharacterCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddFriendIDs(ids...)
}

// SetResidenceID sets the "residence" edge to the Residence entity by ID.
func (cc *CharacterCreate) SetResidenceID(id uuid.UUID) *CharacterCreate {
	cc.mutation.SetResidenceID(id)
	return cc
}

// SetNillableResidenceID sets the "residence" edge to the Residence entity by ID if the given value is not nil.
func (cc *CharacterCreate) SetNillableResidenceID(id *uuid.UUID) *CharacterCreate {
	if id != nil {
		cc = cc.SetResidenceID(*id)
	}
	return cc
}

// SetResidence sets the "residence" edge to the Residence entity.
func (cc *CharacterCreate) SetResidence(r *Residence) *CharacterCreate {
	return cc.SetResidenceID(r.ID)
}

// AddFriendshipIDs adds the "friendships" edge to the Friendship entity by IDs.
func (cc *CharacterCreate) AddFriendshipIDs(ids ...string) *CharacterCreate {
	cc.mutation.AddFriendshipIDs(ids...)
	return cc
}

// AddFriendships adds the "friendships" edges to the Friendship entity.
func (cc *CharacterCreate) AddFriendships(f ...*Friendship) *CharacterCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cc.AddFriendshipIDs(ids...)
}

// Mutation returns the CharacterMutation object of the builder.
func (cc *CharacterCreate) Mutation() *CharacterMutation {
	return cc.mutation
}

// Save creates the Character in the database.
func (cc *CharacterCreate) Save(ctx context.Context) (*Character, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CharacterCreate) SaveX(ctx context.Context) *Character {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CharacterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CharacterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CharacterCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := character.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := character.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CharacterCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Character.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Character.updated_at"`)}
	}
	if _, ok := cc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Character.age"`)}
	}
	if v, ok := cc.mutation.Age(); ok {
		if err := character.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "Character.age": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Character.name"`)}
	}
	return nil
}

func (cc *CharacterCreate) sqlSave(ctx context.Context) (*Character, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CharacterCreate) createSpec() (*Character, *sqlgraph.CreateSpec) {
	var (
		_node = &Character{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(character.Table, sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(character.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(character.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Age(); ok {
		_spec.SetField(character.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Nicknames(); ok {
		_spec.SetField(character.FieldNicknames, field.TypeJSON, value)
		_node.Nicknames = value
	}
	if value, ok := cc.mutation.Info(); ok {
		_spec.SetField(character.FieldInfo, field.TypeJSON, value)
		_node.Info = value
	}
	if nodes := cc.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.FriendsTable,
			Columns: character.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &FriendshipCreate{config: cc.config, mutation: newFriendshipMutation(cc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ResidenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   character.ResidenceTable,
			Columns: []string{character.ResidenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(residence.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.residence_occupants = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.FriendshipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.FriendshipsTable,
			Columns: []string{character.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CharacterCreateBulk is the builder for creating many Character entities in bulk.
type CharacterCreateBulk struct {
	config
	err      error
	builders []*CharacterCreate
}

// Save creates the Character entities in the database.
func (ccb *CharacterCreateBulk) Save(ctx context.Context) ([]*Character, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Character, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CharacterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CharacterCreateBulk) SaveX(ctx context.Context) []*Character {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CharacterCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CharacterCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
