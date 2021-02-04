package main

import (
	"log"
	"os"
)

// Dosyaya veri yazmak.

func main() {

	// İlgili dosyayı writable olarak açalım.
	// Eğer dosya olmasaydı, flag şöyle olacaktı:
	// file, err := os.OpenFile("juno.txt", os.O_WRONLY | os.O_CREATE, 0666)

	file, err := os.OpenFile("juno.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Veriyi oluşturalım.
	byteSlice := []byte("Hello world! I'm a cat.\n") // Sonuna da alt satıra geçmesini söyleyelim.

	// Veriyi ilgili dosyaya yazalım.
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Dosyaya yazma işlemi tamamlandı. Toplam yazılan %d bytes.", bytesWritten)
}
