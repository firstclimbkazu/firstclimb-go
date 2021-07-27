package main

import (
	"firstclimb-go/external"
	"firstclimb-go/graph"
	"firstclimb-go/graph/generated"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const defaultPort = "5000"

// Defining the Graphql handler
func graphqlHandler(db *gorm.DB) gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// 環境変数読み込み
	err := godotenv.Load(fmt.Sprintf("./.%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		return
	}
	// ポート番号設定
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// gorm(DBORマッパー)設定
	db, err := external.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	// ginルーティング設定（gqlgen構成追加）
	r := gin.Default()
	r.GET("/", playgroundHandler())
	r.POST("/query", graphqlHandler(db))
	r.Run()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
