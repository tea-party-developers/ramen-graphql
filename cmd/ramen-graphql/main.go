package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tea-party-developers/ramen-graphql/internal/graphql"
)

const (
	port  = "8080"
	debug = true
)

func main() {
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}))

	http.Handle("/graphql", srv)
	if debug {
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
