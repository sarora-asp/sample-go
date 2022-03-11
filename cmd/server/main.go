package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sample/twirp/internal/helper"
	"sample/twirp/internal/hooks"
	usersvc "sample/twirp/internal/service/user"
	model "sample/twirp/model"
	user "sample/twirp/rpc/user"

	"github.com/joho/godotenv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/twitchtv/twirp"
)

func main() {
	// support for the env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("JWT_SECRET"))
	// create a db connection
	db, err := sqlx.Connect(os.Getenv("DB_TYPE"), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to db")
	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	db.MustExec(model.Schema)

	router := http.NewServeMux()
	var r model.Repository
	var h helper.Helper
	hook := hooks.LoggingHooks(os.Stderr)
	usersvr := usersvc.New(*db, r, h)
	userHandler := user.NewUserServiceServer(usersvr, hook, twirp.WithServerPathPrefix("/api/user"))
	fmt.Println("SERVICE", userHandler.PathPrefix())
	router.Handle(userHandler.PathPrefix(), userHandler)
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

// this is a comment
