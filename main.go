package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/izzulhaziq/attendance-service/resources"
)

func main() {
	log.Printf("Starting HTTP server on port :%d\n", 8080)

	server := createServer(8080)
	waitForStop()
	if err := server.Shutdown(nil); err != nil {
		panic(err)
	}

	log.Println("HTTP server has shutdown gracefully")
}

func waitForStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
}

func createServer(port int) *http.Server {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Setup middleware
	r.Use(gin.Recovery())

	// Register all resources here
	resources.NewStudentResource(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	return srv
}
