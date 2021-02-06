package main

import (
	"fmt"
	"log"
	"net/http"
)

// Query Params
// www.google.com?name=berk&surname=safran&age=30&username=brksfrn

func main() {
	http.HandleFunc("/", Search)
	http.HandleFunc("/search", Search)

	log.Println("Listening on 8080..")
	// mux vb. bir router nesnesi olmadığı için 2. parametreye nil yazabiliriz.
	http.ListenAndServe(":8080", nil)
}

// Search is handler of "/"
func Search(w http.ResponseWriter, r *http.Request) {
	// Gelen Query Parametrelerini karşılıyoruz.
	paramName := r.FormValue("name")       // "/?name=berk"
	paramSurname := r.FormValue("surname") // "/?surname=safran"

	params := r.URL.Query()

	fmt.Println("Params:", params)             // map tipinde tüm paramlar dönecektir.
	fmt.Println("Param Name:", params["name"]) // sadece name param'ı dönecektir.

	// w.Write([]byte(string)) işlevi görür.
	fmt.Fprintf(w, "Name Surname: %s %s\n", paramName, paramSurname)

	fmt.Println(paramName, paramSurname)

}
