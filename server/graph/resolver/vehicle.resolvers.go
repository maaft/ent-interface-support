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

func (r *mutationResolver) UpdateVehicle(ctx context.Context, input model.UpdateVehicleInput) (*model.UpdateVehiclePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVehicle(ctx context.Context, filter *model.VehicleWhereInput) (*model.DeleteVehiclePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QueryVehicle(ctx context.Context, filter *model.VehicleWhereInput) ([]ent.Vehicler, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Vehicles(ctx context.Context, ids []pulid.ID) ([]ent.Vehicler, error) {
	return r.Client.Vehiclers(ctx, ids, ent.WithVehicleType(ent.IDToType))
}

func (r *queryResolver) Vehicle(ctx context.Context, id pulid.ID) (ent.Vehicler, error) {
	return r.Client.Vehicler(ctx, id, ent.WithVehicleType(ent.IDToType))
}
