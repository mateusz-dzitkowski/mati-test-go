package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"thing/thing"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	databaseUrl, _ := os.LookupEnv("DATABASE_URL")
	fmt.Println(databaseUrl)
	conn, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	thingQueries := thing.New(conn)

	inserted, err := thingQueries.CreateThing(ctx, thing.CreateThingParams{
		Name:   "thing name",
		Amount: 10,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(inserted)
}
