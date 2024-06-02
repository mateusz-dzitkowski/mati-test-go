package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard log")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	mylog := slog.New(jsonHandler)
	mylog.Info("hello from info")
	mylog.Info("hello again from debug", "key", "val", "age", 25)
}
