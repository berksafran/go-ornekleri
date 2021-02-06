package main

import (
	"fmt"
	"log"
	"net/http"
)

// messageHandler nesnemizi oluşturalım.
type messageHandler struct {
	message string
}

// Message nesnemize ait ServeHTTP methodumuzu ekleyelim mux kullanımı için.
// NOT!: Buraya func (m *messageHandler) şeklinde nesnenin pointer'lı halini de verebiliyoruz.
// Pointer verdiğimizde tek bir nesneyi kullanıyoruz. Diğer türlü yarattığımız nesneyi aslında kopyalıyoruz.
func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.message)
}

func main() {
	mux := http.NewServeMux()

	// Yarattığımız Custom Handler'larını tanımlıyoruz.
	msg1 := &messageHandler{"Hi, it is first message!"}
	msg2 := &messageHandler{"Wow, second message is here!"}

	mux.Handle("/bir", msg1)
	mux.Handle("/iki", msg2)

	log.Println("Listening on 8080...")

	// mux'ı ikinci parametre olarak vermeyi unutmayalım.
	http.ListenAndServe(":8080", mux)
}
