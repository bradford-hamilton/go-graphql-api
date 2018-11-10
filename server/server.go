package server

import (
	"encoding/json"
	"net/http"

	"github.com/bradford-hamilton/go-graphql-api/gql"
	"github.com/bradford-hamilton/go-graphql-api/postgres"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

// Server will hold connection to the db as well as handlers
type Server struct {
	Db        *postgres.Db
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

// GraphQL returns an http.HandlerFunc for our /graphql endpoint
func (s *Server) GraphQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Must provide graphql query in request body", 400)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema, s.Db)
		render.JSON(w, r, result)
	}
}

// RestfulEndpoint returns an http.HandlerFunc for our /restful/endpoint endpoint
func (s *Server) RestfulEndpoint() http.HandlerFunc {
	type response struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Profession string `json:"profession"`
		Friendly   bool   `json:"friendly"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := s.Db.RestQuery()
		render.JSON(w, r, res)
	}
}
