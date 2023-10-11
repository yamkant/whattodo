package main

import (
	"log"
	"net/http"

	"example.com/m/api"
	"example.com/m/auth"
	"example.com/m/web"
)

func main() {
	apiMux := api.APIHttpHandler()
	webMux := web.WebHttpHandler()
	authMux := auth.AuthHttpHandler()

	log.Println("Starting server...")
	http.Handle("/api/", apiMux)
	http.Handle("/auth/", authMux)
	http.Handle("/", webMux)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
