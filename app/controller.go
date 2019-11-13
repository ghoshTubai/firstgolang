package app

import (
	"fmt"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Print("My first program")
	w.Write([]byte("Hi from server"))
}
