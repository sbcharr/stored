package main

import (
	"net/http"

	"github.com/gorilla/mux"
	svc "github.com/sbcharr/stored/service"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/key/{key}", svc.KVPutHandler).Methods("PUT")
	router.HandleFunc("/v1/key/{key}", svc.KVGetHandler).Methods("GET")
	router.HandleFunc("/v1/key/{key}", svc.KVDeleteHandler).Methods("DELETE")
	panic(http.ListenAndServe(":8080", router))
}
