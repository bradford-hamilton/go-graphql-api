package gql

import (
	"github.com/bradford-hamilton/go-graphql-api/postgres"
	"github.com/graphql-go/graphql"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *postgres.Db) *Root {
	resolver := Resolver{db: db}
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"user": &graphql.Field{
						Type: User,
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}
	return &root
}
