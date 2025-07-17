package main

import (
	"context"
	"github.com/gnomedevreact/CombatIntel/internal/database"
	"github.com/gnomedevreact/CombatIntel/internal/routes"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiCfg := apiConfig{}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	} else {
		conn, err := pgx.Connect(ctx, dbUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close(ctx)

		queries := database.New(conn)
		apiCfg.db = queries
		log.Println("Connected to database")
	}

	mux := http.NewServeMux()
	routes.RegisterRouter(mux)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Println("Listening on http://localhost" + srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
