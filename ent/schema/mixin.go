// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// TimestampedMixin for all schemas in the graph.
type TimestampedMixin struct {
	mixin.Schema
}

func (TimestampedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// VehicleMixin for all schemas in the graph.
type VehicleMixin struct {
	mixin.Schema
}

func (VehicleMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").Optional(),
	}
}

// FolderOrFileMixin for all schemas in the graph.
type FolderOrFileMixin struct {
	mixin.Schema
}

func (FolderOrFileMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("parentIds").NotEmpty(),
	}
}

type DatapointMixin struct {
	mixin.Schema
}

func (DatapointMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("producedBy").NotEmpty(),
	}
}

func (DatapointMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("producedBy"),
	}
}

type MagneticFieldScanMarkupMixin struct {
	mixin.Schema
}

func (MagneticFieldScanMarkupMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("markupSource").Values("USER", "PREDICTION", "COMPUTATION"),
	}
}

// We cannot put the edges in here, because it is a ref edge and need to be done for all implementing types of the schema
func (MagneticFieldScanMarkupMixin) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (MagneticFieldScanMarkupMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("markupSource"),
	}
}
