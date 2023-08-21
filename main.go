package main

import (
	"go-testcontainers/customer"
	"go-testcontainers/database"
	"go-testcontainers/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pseidemann/finish"
)

func main() {
	database.ConnectDatabase()
	utils.Migrate()

	routes := gin.New()
	customer.NewCustomerController(routes)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: routes,
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
