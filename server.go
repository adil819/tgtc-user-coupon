package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/configs/database"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/graph"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/graph/generated"
	"github.com/kelompok-1-tgtc/tgtc-user-coupon/internal/models"
)

const defaultPort = "8080"

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Coupon{})
	database.DB.AutoMigrate(&models.UserHasCoupon{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
