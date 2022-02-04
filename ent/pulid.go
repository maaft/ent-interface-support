package ent

import (
	"backend/ent/car"
	"backend/ent/plane"
	"backend/ent/schema/pulid"
	"context"
	"fmt"
	"strings"
)

// prefixMap maps PULID prefixes to table names.
var prefixMap = map[pulid.ID]string{
	"Car":   car.Table,
	"Plane": plane.Table,
}

// IDToType maps a pulid.ID to the underlying table.
func IDToType(ctx context.Context, id pulid.ID) (string, error) {
	for p := range prefixMap {
		if strings.Contains(string(id), string(p)) {
			return prefixMap[p], nil
		}
	}

	return "", fmt.Errorf("IDToType: could not map id '%s' to a type", id)
}
