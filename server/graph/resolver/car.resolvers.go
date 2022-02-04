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

func (r *mutationResolver) AddCar(ctx context.Context, input []*model.AddCarInput) (*model.AddCarPayload, error) {
	client := ent.FromContext(ctx)

	builders := []*ent.CarCreate{}
	for _, I := range input {
		desc := ""
		if I.Description != nil {
			desc = *I.Description
		}

		builders = append(builders, client.Car.Create().
			SetName(I.Name).
			SetDescription(desc).
			SetWheelPressure(I.WheelPressure))
	}
	cars, err := client.Car.CreateBulk(builders...).Save(ctx)

	if err != nil {
		return nil, err
	}

	ret := &model.AddCarPayload{
		AddCar: cars,
	}

	return ret, nil
}

func (r *mutationResolver) UpdateCar(ctx context.Context, input model.UpdateCarInput) (*model.UpdateCarPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCar(ctx context.Context, filter *ent.CarWhereInput) (*model.DeleteCarPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryCar(ctx context.Context, filter *ent.CarWhereInput) ([]*ent.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCar(ctx context.Context, id pulid.ID) (*ent.Car, error) {
	panic(fmt.Errorf("not implemented"))
}
