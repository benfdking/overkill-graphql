package shared

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bendfking/overkill-graphql/packages/users/graph"
	"github.com/bendfking/overkill-graphql/packages/users/graph/generated"
)

func ServerPrep(defaultPort string) http.Handler {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(&debug.Tracer{})

	handler := http.NewServeMux()

	handler.Handle("/", playground.Handler("GraphQL playground", "/query"))
	handler.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return handler
}
