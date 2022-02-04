package resolver

import (
	"backend/ent"
	"backend/server/graph/generated"

	"github.com/99designs/gqlgen/graphql"
	"github.com/minio/minio-go/v7"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct {
	Client      *ent.Client
	MinioClient *minio.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, minioClient *minio.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			Client: client,
		},
	})
}
