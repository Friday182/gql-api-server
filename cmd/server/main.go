package main

import (
	log "log"
	"time"

	"github.com/gin-contrib/cors"
	// "gql-api-server/internal/orm"

	"github.com/gin-gonic/gin"
	// "gql-api-server/internal/handlers"
)

var host, port string

func init() {
	// read configuration from configs/config.json
	host = "localhost"
	port = "8090"
}

// Config and start up the server
func main() {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// AllowOriginFunc:  func(origin string) bool { return origin == "http://localhost:3000" },
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Handlers
	// Simple keep-alive/ping handler
	//r.GET("/heartbeat", handlers.Heartbeat())

	// GraphQL handlers
	//r.POST("/graghql", handlers.GraphqlHandler(orm))
	log.Println("GraphQL @ " + endpoint + "/graphql")

	// Run the server
	// Inform the user where the server is listening
	log.Println("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatal(r.Run(host + ":" + port))
}
