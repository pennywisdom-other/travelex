// travelex-test documentation
//   This is a simple service that supports 2 endpoints
//     GET /v1/countries?target=source
//     GET /v1/countries?target=destination
package main

import (
"context"
"log"
"net/http"
"os"
"os/signal"
"time"

	"github.com/gin-gonic/gin"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/countries", requestAuthorizerMiddleware, handleCountryRequest)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}


