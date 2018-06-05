package main

import (
	"net/http"

	"github.com/urfave/negroni"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "test.wasm")
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/test.wasm", wasmHandler)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)

}
