package main

import (
	"log"
	"os"
)

// Dosyanın varlığını kontrol etme

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	filePath := "hello.txt"

	fileInfo, err = os.Stat(filePath)
	if err != nil {
		// fmt.Println(err) --> stat jello.txt: no such file or directory
		// Dikkat! os.IsNotExist() parametre olarak hata nesnesi alıyor.
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}

	log.Println("File exists!")
	log.Println("File Info:", fileInfo.Size())
}
