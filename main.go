package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusOK, "SUCCESS")
	})

	log.Fatal(http.ListenAndServe(":8090", nil))
}
