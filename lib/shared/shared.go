package shared

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
)

func ServerPrep(defaultPort string, e graphql.ExecutableSchema) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(e)
	srv.Use(&debug.Tracer{})

	handler := http.NewServeMux()

	handler.Handle("/", playground.Handler("GraphQL playground", "/query"))
	handler.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}
