package main

import (
	"log"
	"net/http"
	"progategolang/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HomeHandler) //root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/tugasprogate", handler.TugasprogateHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	fileServer := http.FileServer(http.Dir("asset"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting Web on https://anggitmaulana.github.io/progategolang/")
	err := http.ListenAndServe("https://anggitmaulana.github.io/progategolang/", mux)
	log.Fatal(err)
}
