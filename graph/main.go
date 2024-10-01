package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	db "graph/db/sqlc"
	"net/http"
	"os"
)

func main() {
	fmt.Println("starting")

	ctx := context.Background()

	databaseUrl := os.Getenv("DATABASE_URL")
	conn, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}
	queries := db.New(conn)

	fmt.Println(queries.GetNodes(ctx))

	r := gin.Default()
	r.GET("/node", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
