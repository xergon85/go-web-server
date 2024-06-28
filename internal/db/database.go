package db

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"reflect"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xergon85/go-web-server/internal/db/user"
)

//go:embed schema.sql
var dd1 string

func Run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// create tables
	if _, err := db.ExecContext(ctx, dd1); err != nil {
		return err
	}

	// list all users
	queries := user.New(db)

	users, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}

	log.Println(users)

	// create user
	insertUser, err := queries.CreateUser(ctx, user.CreateUserParams{
		Username:  "Bengt",
		Password:  "abcd",
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertUser)

	// get the user we just inserted
	fetchedUser, err := queries.GetUser(ctx, insertUser.ID)
	if err != nil {
		return err
	}

	log.Println(reflect.DeepEqual(insertUser, fetchedUser))

	return nil
}
