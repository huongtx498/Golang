package main

import (
	"GOLANG/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/hash", api.HashUrl).Methods("GET")

	router.HandleFunc("/api/v1/user/find", api.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/getall", api.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/user/create", api.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", api.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete", api.Delete).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
}
