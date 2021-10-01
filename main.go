package main

import (
	"log"
	"net/http"
	"os"
)

const DefaultPort = "8080"

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

func EchoHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

	writer.Header().Set("Access-Control-Allow-Origin", "*")

	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	err := request.Write(writer)
	if err != nil {
		return
	}
}

func main() {
	log.Println("Starting server, listening on port " + getServerPort())

	http.HandleFunc("/", EchoHandler)
	err := http.ListenAndServe(":"+getServerPort(), nil)
	if err != nil {
		return
	}
}
