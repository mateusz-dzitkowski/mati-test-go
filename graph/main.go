package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	db "graph/db/sqlc"
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

	labels := []string{"a", "b", "c"}
	nodes := make([]db.Node, len(labels))
	for i, label := range labels {
		node, err := queries.CreateNode(ctx, label)
		if err != nil {
			fmt.Println(err)
		}
		nodes[i] = node
		fmt.Println(node)
	}

	for _, params := range []db.CreateEdgeParams{
		{
			Src: nodes[0].ID,
			Dst: nodes[1].ID,
		},
		{
			Src: nodes[1].ID,
			Dst: nodes[2].ID,
		},
		{
			Src: nodes[2].ID,
			Dst: nodes[0].ID,
		},
	} {
		edge, err := queries.CreateEdge(ctx, params)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(edge)
	}
}
