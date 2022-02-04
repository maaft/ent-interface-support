package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"backend/ent"
	"backend/ent/schema/pulid"
	"backend/server/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) AddPlane(ctx context.Context, input []*model.AddPlaneInput) (*model.AddPlanePayload, error) {
	client := ent.FromContext(ctx)

	builders := []*ent.PlaneCreate{}
	for _, I := range input {
		desc := ""
		if I.Description != nil {
			desc = *I.Description
		}

		builders = append(builders, client.Plane.Create().
			SetName(I.Name).
			SetDescription(desc).
			SetAltitude(I.Altitude))
	}
	cars, err := client.Plane.CreateBulk(builders...).Save(ctx)

	if err != nil {
		return nil, err
	}

	ret := &model.AddPlanePayload{
		AddPlane: cars,
	}

	return ret, nil
}

func (r *mutationResolver) UpdatePlane(ctx context.Context, input model.UpdatePlaneInput) (*model.UpdatePlanePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePlane(ctx context.Context, filter *ent.PlaneWhereInput) (*model.DeletePlanePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryPlane(ctx context.Context, filter *ent.PlaneWhereInput) ([]*ent.Plane, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPlane(ctx context.Context, id pulid.ID) (*ent.Plane, error) {
	panic(fmt.Errorf("not implemented"))
}
