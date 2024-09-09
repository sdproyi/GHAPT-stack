package main

import (
	"app/templates/components"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	staticDir   = "./static"
	staticPath  = "/static/"
	portEnv     = "PORT"
	defaultPort = "8080"
)

func serveStaticFiles(folder string) {
	path := staticPath + folder + "/"
	fileServer := http.FileServer(http.Dir(staticDir + "/" + folder))
	http.Handle(path, http.StripPrefix(path, fileServer))
}

func getPort() string {
	port := os.Getenv(portEnv)
	if port == "" {
		port = defaultPort
	}
	return ":" + port
}

func main() {
	components.Routes()

	serveStaticFiles("html")
	serveStaticFiles("font")
	serveStaticFiles("style")
	serveStaticFiles("images")

	port := getPort()
	fmt.Printf("Listening on %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
