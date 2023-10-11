package main

import (
	"fmt"
	"net/http"

	"github.com/fanialfi/golang-sql/routing"
)

var port = ":8080"

func main() {
	http.HandleFunc("/users", routing.HandleUsers)
	http.HandleFunc("/user", routing.HandleUser)
	http.HandleFunc("/prepare", routing.Prepare)

	fmt.Printf("server listening on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
