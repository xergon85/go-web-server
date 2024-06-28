package main

import (
	"log"

	"github.com/xergon85/go-web-server/internal/db"
)

func main() {
	if err := db.Run(); err != nil {
		log.Fatal(err)
	}
}
