package main

import (
	log "log"
	"time"

	"github.com/friday182/gql-api-server/internal/orm"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	// "github.com/friday182/gql-api-server/internal/handlers"
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

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// AllowOriginFunc:  func(origin string) bool { return origin == "http://localhost:3000" },
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db, err := orm.InitOrm()
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}

	// GraphQL handlers
	r.POST("/graghql", handlers.GraphqlHandler(db))
	log.Println("GraphQL @ " + endpoint + "/graphql")

	// Run the server
	log.Fatal(r.Run(host + ":" + port))
}
