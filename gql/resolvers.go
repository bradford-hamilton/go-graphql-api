package gql

import (
	"github.com/bradford-hamilton/go-graphql-api/postgres"
	"github.com/graphql-go/graphql"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *postgres.Db
}

// UserResolver resolves our person query through a db call to GetUserByName
func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)
	if ok {
		block := r.db.GetUserByName(name)
		return block, nil
	}

	return nil, nil
}
