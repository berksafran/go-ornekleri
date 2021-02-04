package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Önce geçici klasörümüzü oluşturalım.
	// Oluşturulacağı yer işletim sistemimizdeki geçici klasör olacaktır.
	tempDirPath, err := ioutil.TempDir("", "geciciKlasor")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçici klasör oluşturuldu: ", tempDirPath)

	// Geçici dosyamızı, geçici klasörümüzün altına oluşturuyoruz.
	tempFile, err := ioutil.TempFile(tempDirPath, "geciciDosya.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçici dosya oluşturuldu: ", tempFile.Name())

	// Geçici dosyamızı kapatmayı unutmayalım.
	defer tempFile.Close()

	// Geçici dosya ve dizini silerken içerden dışarıya doğru silmemiz gerekir.
	// Önce Geçici dosyamızı silelim.
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Geçici dosya silindi.")

	// ve en son olarak Geçici dizini (klasörümüzü) silelim.

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Geçici klasör silindi.")
}
