package main

import (
	//"flag"
	"log"
)

func main() {

	//port := flag.String("port", "8080", "Port to listen to")
	//flag.Parse()

	listeningPort := ":8080"
	log.Println(listeningPort)

	httpServer := NewHTTPServer(listeningPort)

	if err := httpServer.Open(); err != nil {
		log.Fatal("could not open httpServer", err)
	}

}