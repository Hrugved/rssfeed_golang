package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Hrugved/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviornment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the enviornment")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	db := database.New(conn)
	apiCgf := apiConfig {
		DB: db,
	}

	go startScrapping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCgf.handlerCreateUser)
	v1Router.Get("/users", apiCgf.middlewareAuth(apiCgf.handleGetUser))
	v1Router.Post("/feeds", apiCgf.middlewareAuth(apiCgf.handlerCreateFeed))
	v1Router.Get("/feeds", apiCgf.handlerGetFeeds)
	v1Router.Post("/feed_follows", apiCgf.middlewareAuth(apiCgf.handlerFeedFollow))
	v1Router.Get("/feed_follows", apiCgf.middlewareAuth(apiCgf.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCgf.middlewareAuth(apiCgf.handlerDeleteFeedFollows))

	router.Mount("/v1",v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


}