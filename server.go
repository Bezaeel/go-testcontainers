package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pseidemann/finish"
)

func NewServer(routes *gin.Engine) {
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
