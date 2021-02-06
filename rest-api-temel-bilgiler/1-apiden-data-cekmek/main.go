package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// GET ile datayı alıyoruz.
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	// Datayı aldıktan sonra kapatma işlemini yapmamız gerekiyor.
	// Aynı bir dosyayı açtıktan sonra kapattığımız gibi.
	defer resp.Body.Close()

	// Gelen veriyi byte'a çeviriyoruz.
	body, err := ioutil.ReadAll(resp.Body)

	// Elde ettiğimiz byte verisini string'e çeviriyoruz.
	fmt.Println(string(body))
}
