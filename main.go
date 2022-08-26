package main

import (
	"InMemoryCache/database"
	"InMemoryCache/server"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {

	err := database.ConnectAndMigrate("localhost", "5433", "postgres", "local", "local", database.SSLModeDisable)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	fmt.Println("connected")

	srv := server.NewServer()
	err = srv.Run(":8080")
	if err != nil {
		logrus.Error(err)
	}

}
