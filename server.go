package main

import (
	"log"
	"net/http"
	"os"
	"parts/graph"
	"parts/graph/generated"
	"parts/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	svc, err := service.Create()

	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Svc: svc}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil /* handlers.CombinedLoggingHandler(os.Stdout, http.DefaultServeMux)*/))
}
