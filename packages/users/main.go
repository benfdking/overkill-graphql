package main

import (
	"log"

	"github.com/bendfking/overkill-graphql/lib/shared"
	"github.com/bendfking/overkill-graphql/packages/users/graph"
	"github.com/bendfking/overkill-graphql/packages/users/graph/generated"
)

//go:generate go run github.com/99designs/gqlgen generate

const defaultPort = "4002"

func main() {
	es := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	s := shared.ServerPrep(defaultPort, es)
	log.Fatal(s.ListenAndServe())
}
