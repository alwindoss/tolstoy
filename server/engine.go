package server

import "log"
import "net/http"

func Run(port string) {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
