package main

import (
	"fmt"
	"github.com/pascalallen/auth"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8080...")

	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/register", auth.Register)

	http.ListenAndServe(":8080", nil)
}
