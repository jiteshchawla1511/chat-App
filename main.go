package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jiteshchawla1511/chat-app/pkg/httpserver"
	"github.com/jiteshchawla1511/chat-app/pkg/sockets"

	"github.com/joho/godotenv"
)

func init() {
	// Load the environment file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load the env file.", err)
	}
}

func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()

	if *server == "http" {
		fmt.Println("http server is starting on :8080")
		httpserver.StartHttpServer()
	} else if *server == "websocket" {
		fmt.Println("websocket server is starting on :8081")
		sockets.StartWebsocketServer()
	} else {
		fmt.Println("invalid server Available server: http or websocket")
	}
}
