package main

import (
	"log"

	"github.com/xergon85/go-web-server/internal/web"
)

func main() {

	if err := web.Run(); err != nil {
		log.Fatal(err)
	}
}
