package main

import (
	"go-testcontainers/database"
	"go-testcontainers/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pseidemann/finish"
)

var routes *gin.Engine

func main() {
	dbOptions := database.DbOptions{
		Host:     "",
		Port:     0,
		User:     "",
		Password: "",
		Name:     "",
	}
	database.ConnectDatabase(dbOptions)
	utils.Migrate()

	routes = gin.New()

	routes.Run()

	Init()

	srv := &http.Server{
		Addr:        ":8000",
		Handler:     routes,
		ReadTimeout: finish.DefaultTimeout,
	}

	fin := finish.New()
	fin.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fin.Wait()

}
