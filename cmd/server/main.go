package main

import (
	log "log"

	"github.com/friday182/gql-api-server/internal/orm"
	"github.com/gin-gonic/gin"
	"github.com/friday182/gql-api-server/internal/handlers"
	"github.com/99designs/gqlgen/handler"
)

var host, port string

func init() {
	// read configuration from configs/config.json
	host = "localhost"
	port = "8090"
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Config and start up the server
func main() {
	endpoint := "http://" + host + ":" + port

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	db := orm.ConnectDb()
	defer db.Close()

	// GraphQL handlers
	r.POST("/query", handlers.GraphqlHandler(db))
	r.POST("/graphql", handlers.GraphqlHandler(db))
	r.GET("/", playgroundHandler())
	log.Println("GraphQL @ " + endpoint + "/graphql")

	// Run the server
	log.Fatal(r.Run(host + ":" + port))
}
