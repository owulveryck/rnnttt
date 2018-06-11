package main

import (
	"log"
	"mime"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mime.AddExtensionType(".wasm", "application/wasm")
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	log.Fatal(http.ListenAndServe(":3000", n))

}
