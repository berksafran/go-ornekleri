package main

// Dosya oluşturmak.

import (
	"fmt"
	"log"
	"os"
)

var (
	// pointer vermezsek şöyle bir hata alırız
	// cannot assign *os.File to newFile (type os.File) in multiple assignment
	newFile *os.File
	err     error
)

func main() {
	// text.txt adında bir dosya oluşturalım.
	newFile, err = os.Create("text.txt")

	fmt.Println("Dosya oluşturuldu.", newFile)

	// Error handling'imizi de ekleyelim.
	if err != nil {
		log.Fatal(err)
	}
}
