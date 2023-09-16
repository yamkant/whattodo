package main

import (
	"log"
	"net/http"

	"example.com/m/app"
	"github.com/urfave/negroni"
)

func main() {
	mux := app.MyHttpHandler()

	log.Println("Starting server...")
	n := negroni.Classic()
	n.UseHandler(mux)
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
