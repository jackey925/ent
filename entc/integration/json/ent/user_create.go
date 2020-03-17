// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/json/ent/user"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetURL sets the url field.
func (uc *UserCreate) SetURL(u *url.URL) *UserCreate {
	uc.mutation.SetURL(u)
	return uc
}

// SetRaw sets the raw field.
func (uc *UserCreate) SetRaw(jm json.RawMessage) *UserCreate {
	uc.mutation.SetRaw(jm)
	return uc
}

// SetDirs sets the dirs field.
func (uc *UserCreate) SetDirs(h []http.Dir) *UserCreate {
	uc.mutation.SetDirs(h)
	return uc
}

// SetInts sets the ints field.
func (uc *UserCreate) SetInts(i []int) *UserCreate {
	uc.mutation.SetInts(i)
	return uc
}

// SetFloats sets the floats field.
func (uc *UserCreate) SetFloats(f []float64) *UserCreate {
	uc.mutation.SetFloats(f)
	return uc
}

// SetStrings sets the strings field.
func (uc *UserCreate) SetStrings(s []string) *UserCreate {
	uc.mutation.SetStrings(s)
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldURL,
		})
		u.URL = value
	}
	if value, ok := uc.mutation.Raw(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldRaw,
		})
		u.Raw = value
	}
	if value, ok := uc.mutation.Dirs(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldDirs,
		})
		u.Dirs = value
	}
	if value, ok := uc.mutation.Ints(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldInts,
		})
		u.Ints = value
	}
	if value, ok := uc.mutation.Floats(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldFloats,
		})
		u.Floats = value
	}
	if value, ok := uc.mutation.Strings(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldStrings,
		})
		u.Strings = value
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
