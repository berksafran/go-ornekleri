package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// JSON değişkenimiz.
	jsonStr := `
		{
			"data": {
				"object": "creditCard",
				"id": "card_23428347829",
				"firstName": "Berk",
				"lastName": "Safran",
				"balance": "12.211"
			}
		}	
	`

	/* JSON içerisinde "object içerisinde object" olduğundan
	   gelen datayı karşılamak için GO'da da "map içinde map"
	   tanımlamamız gerekiyor.
	*/
	var jsonMap map[string]map[string]interface{}

	// Gelen JSON datasını byte'a dönüştürüp, jsonMap değişkenine değerini atıyoruz.
	// json.Unmarshal = Byte -> JSON
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}

	// Ekrana bastırıyoruz.
	// Ekrana bastırdığımızda gelen "data" objesi yukarıdaki JSON string'inden geliyor.

	fmt.Println(jsonMap) // Bu tercih edilebilir.

	fmt.Println("******************")

	// Satır satır yazmak için.
	for key, value := range jsonMap["data"] {
		fmt.Printf("\t%s : %v \n", key, value)
	}

	// ********************
	fmt.Println("******************")

	// JSON'a çevirdiğimiz datayı, tekrar byte'a çevirelim.
	// json.Marshal = JSON -> Byte

	b, err := json.Marshal(jsonMap)
	if err != nil {
		log.Fatal(err)
	}

	// Byte verisini string'e çevirebiliriz.
	fmt.Println(string(b))
}
