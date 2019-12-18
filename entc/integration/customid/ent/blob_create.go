// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/blob"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// BlobCreate is the builder for creating a Blob entity.
type BlobCreate struct {
	config
	id   *uuid.UUID
	uuid *uuid.UUID
}

// SetUUID sets the uuid field.
func (bc *BlobCreate) SetUUID(u uuid.UUID) *BlobCreate {
	bc.uuid = &u
	return bc
}

// SetID sets the id field.
func (bc *BlobCreate) SetID(u uuid.UUID) *BlobCreate {
	bc.id = &u
	return bc
}

// Save creates the Blob in the database.
func (bc *BlobCreate) Save(ctx context.Context) (*Blob, error) {
	if bc.uuid == nil {
		v := blob.DefaultUUID()
		bc.uuid = &v
	}
	return bc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlobCreate) SaveX(ctx context.Context) *Blob {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BlobCreate) sqlSave(ctx context.Context) (*Blob, error) {
	var (
		b    = &Blob{config: bc.config}
		spec = &sqlgraph.CreateSpec{
			Table: blob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		}
	)
	if value := bc.id; value != nil {
		b.ID = *value
		spec.ID.Value = *value
	}
	if value := bc.uuid; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  *value,
			Column: blob.FieldUUID,
		})
		b.UUID = *value
	}
	if err := sqlgraph.CreateNode(ctx, bc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}