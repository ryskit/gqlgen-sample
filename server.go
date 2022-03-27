package main

//go:generate sqlboiler --wipe psql

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/ryskit/gqlgen-sample/graph"
	"github.com/ryskit/gqlgen-sample/graph/generated"
)

const defaultPort = "8080"

func connDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", `dbname=gqlgen_sample host=localhost user=postgres password=password sslmode=disable`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := connDB()
	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
