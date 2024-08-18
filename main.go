package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Alvphil/Blog-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func main() {
	godotenv.Load()
	const filepathRoot = "."
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErrors)
	v1Router.Post("/users", apiCfg.CreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.GetUserByApiKey))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetAllFeeds)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}
