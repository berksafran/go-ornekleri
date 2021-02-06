package main

import (
	"fmt"
	"net/http"
)

// '/about' Handler
func aboutPage(w http.ResponseWriter, r *http.Request) {
	/*
		Status kodlarını da gönderebiliriz.
		"Status Code: 400 Bad Request" dönecektir.
		Dikkat!! WriteHeader'ı Write'tan önce yazman gerek geçerli olabilmesi için.
	*/

	w.WriteHeader(http.StatusBadRequest)

	// Byte tipinde response dönüyoruz.
	w.Write([]byte("All is well!"))
}

/*
	Her dosya save'inde otomatik tekrar compile olabilmesi için:
	nodemon --exec go run main.go --signal SIGTERM
*/

func main() {

	// "localhost:8080/"e gelen requestler için handler yazıyoruz.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*
			Write() methodu ile istediğiniz türde veri gönderebilirsiniz.
			Fakat, network üzerinden veriler byte olarak gönderilir.
			(Programlarda kendi aralarında byte ile konuşur. Bu durum, farkedilmeden arka planda hazırlanır.)
		*/
		w.Write([]byte("Merhaba Dünya!"))
	})

	// Farklı bir handler yazım stili istersek. (Fonksiyonu dışarı alıyoruz.)
	http.HandleFunc("/about", aboutPage)

	// Sunucumuzu 8080 portunda başlatıyoruz.
	fmt.Println("Server is listening on 8080..")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error on listening..")
	}

}
