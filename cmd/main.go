package main

import (
	"devOps00Boots/handlers/rest"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"
	//Sets the port to listen on❶

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	//Registers the translation Handler❷

	log.Printf("listening on %s\n", addr)
	//Logs the listening port❸

	log.Fatal(http.ListenAndServe(addr, mux))
	//Runs the server and logs if it fails
}
