package main

import (
	"context"
	"fmt"
	"graphql/database"
	"graphql/graph"
	"graphql/repository"
	"graphql/services"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/zerolog/log"
)

const defaultPort = "8080"

func main() {
	sv, err := Start()
	if err != nil {
		log.Info().Err(err).Msg("error in Start App")

	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{S: sv}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal().Err(http.ListenAndServe(":"+port, nil))
}
func Start() (services.UserService, error) {
	db, err := database.Open()
	if err != nil {
		return &services.Service{}, fmt.Errorf("error in the connecting the database")
	}
	pg, err := db.DB()
	if err != nil {
		return &services.Service{}, fmt.Errorf("error in getting the instance:%w", err)
	}
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	err = pg.PingContext(ctx)
	if err != nil {
		return &services.Service{}, fmt.Errorf("error in connecting to database: %w", err)
	}
	repo, err := repository.NewRepo(db)
	if err != nil {
		return &services.Service{}, fmt.Errorf("repository not initiated: %w", err)
	}

	svc, err := services.NewServices(repo)
	if err != nil {
		return &services.Service{}, fmt.Errorf("service is not initiated: %w", err)
	}
	return svc, nil
}
