package main

import (
	"io"
	"log"
	"net/http"
)

// ironman adında bir nesne oluşturuyoruz.
type ironman int

/*
Daha sonra bu ironman nesnesine ServeHTTP adında method ekliyoruz.
Bu method mux içerisinde bulunan aslında default bir method.
Biz bu methodu tanımlayarak nesnemizi mux'tan türettiğimiz zaman, mux'un içerisindeki
methodu overwrite edebiliyoruz.
*/
func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {

	io.WriteString(res, "Mr. Ironman!")
}

type wolverine int

// Volwerine nesnesine de ServeHTTP adında method ekliyoruz.
func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {

	// io paketini kullandık ama w.Write() ile aynı sonucu dönecek.
	io.WriteString(res, "Mr. Wolverine!")
}

func main() {

	var i ironman
	var w wolverine

	/*
		mux nesnesi, aslında bir web framework mantığında çalışıyor.
		Her seferinde http.handleFunc() yazmak yerine mux nesnesini kullanabiliriz.
	*/

	mux := http.NewServeMux()
	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)

	/*
		Daha önce http.handleFunc() kullandığımız için http.ListenAndServe(":8080", nil) yazabilirdik.
		Şimdi farklı bir nesne üzerinden çalıştığımız için aşağıdaki gibi 2. parametreyi
		"mux" olarak düzenliyoruz.
	*/
	log.Println("Listening on :8080..")
	http.ListenAndServe(":8080", mux)
}
