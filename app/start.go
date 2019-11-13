package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Start (){
	fmt.Println("starting server!!")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello",greetings).Methods("GET")
	router.HandleFunc("/hello",greetings).Methods("POST")
	http.ListenAndServe("0.0.0.0:5555",router)
}