package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/samarthps29/rss-aggregator/internal/database"
)

type apiConfig struct {
	// DB is a pointer type variable pointing to a Queries Struct Object
	DB *database.Queries
}

func main() {
	godotenv.Load()
	// loading the environment variables
	portString := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("PORT variable was not found in the environment")
	}
	if dbURL == "" {
		log.Fatal("DB_URL variable was not found in the environment")
	}

	// opening a new sql connection using appropriate database driver and database url
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to the database: ", err)
	}

	// we now a connection but we want it in a different format, so...
	// creating a new pointer variable pointing to a Queries struct object, created by passing the 'conn' variable to the 'new' method of the database package
	db := database.New(conn)

	// creating a new apiConfig object with a Queries poointer variable
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()
	v1router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Mount("/v1", v1router)
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/error", handlerError)
	v1router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	v1router.Post("/users", apiCfg.handlerCreateUser)
	v1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))

	v1router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server started and running on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
