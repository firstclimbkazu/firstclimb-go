package main

import (
	"firstclimb-go/external"
	"firstclimb-go/graph"
	"firstclimb-go/graph/generated"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "5000"

func main() {
	err := godotenv.Load(fmt.Sprintf("./.%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := external.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
