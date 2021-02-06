package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Path // "localhost:8080/merhaba" yazdığımızda "/merhaba" döner
	y := x[1:]      // "merhaba" döner.

	if len(y) > 0 {
		w.Write([]byte(y))
	} else {
		fmt.Fprintf(w, "Hello World! %s\n", time.Now()) // w.Write() işlevi görür.
	}
}

func main() {
	http.HandleFunc("/", greet)

	log.Println("Listening on 8080..")
	http.ListenAndServe(":8080", nil)
}
