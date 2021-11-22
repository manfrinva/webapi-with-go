package main

import (
	"github.com/manfrinva/webapi-with-go/database"
	"github.com/manfrinva/webapi-with-go/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
