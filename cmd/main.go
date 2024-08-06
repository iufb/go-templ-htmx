package main

import (
	"log"

	"github.com/iufb/go-templ-htmx/cmd/api"
	"github.com/iufb/go-templ-htmx/db"
)

func main() {
	db := db.SetupDatabase()
	server := api.NewAPIServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
