package main

import (
	"log"

	"github.com/bendfking/overkill-graphql/lib/shared"
	"github.com/bendfking/overkill-graphql/packages/products/graph"
	"github.com/bendfking/overkill-graphql/packages/products/graph/generated"
)

//go:generate go run github.com/99designs/gqlgen generate

const defaultPort = "4003"

func main() {
	es := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	s := shared.ServerPrep(defaultPort, es)
	log.Fatal(s.ListenAndServe())
}
