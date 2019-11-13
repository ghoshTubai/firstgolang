package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Start (){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello",greetings).Methods("GET")
	router.HandleFunc("/hello",greetings).Methods("POST")
	http.ListenAndServe("0.0.0.0:5555",router)
}