package main

import (
	"fmt"
	"log"
	"os"
)

// Dosya açma ve kapama

func main() {
	// OpenReadOnly()
	OpenWriteData()
}

// OpenWriteData ... Bu fonksiyon dosyayı açar ve veri ekleyebilir.
func OpenWriteData() {

	/*
		İlk parametre dosya ismi, ikincisi flag ve üçüncü permission alır.
		-
		flag: Dosya açılış amacını ayarlar.
		permission: Dosya izinlerini belirler.
	*/

	/*
		flag çeşitleri:
		os.O_RDONLY: Sadece okumak için.
		os.O_WRONLY: Sadece yazmak için.
		os.O_RDWR: Okuma ve yazma yapmak için.
		os.O_APPEND: Veriyi dosyanın sonuna eklemek için.
		os.O_CREATE: Dosya yoksa oluşturmak için. (Üstüne yazar dikkat!)
		os.O_TRUNC: Açılırken dosyayı kesmek için.

		Aynı flag içerisinde birden fazla ayar kullanılabilir:
		os.O_CREATE|os.O_APPEND
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	*/

	file, err := os.OpenFile("hello.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

// OpenReadOnly ... Bu fonksiyon dosyayı sadece read-only olarak açar.
func OpenReadOnly() {
	// Dikkat! Bu yöntem dosyayı salt okunur (read-only) olarak açar.
	// Herhangi bir veri üzerine yazamazsınız.
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file) // Output: &{0xc00004e180}

	// defer bu scope'ta en son çalışır. İşlem bittiğinde dosya kapatılır.
	defer file.Close()

}
