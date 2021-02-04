package main

import (
	"io"
	"log"
	"os"
)

// Dosya kopyalamak.

func main() {

	// Dosya kopyalamadan önce dosyayı açmak gerekir.
	originalPath, err := os.Open("tello.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Açtığımız dosyaları kapatmayı unutmuyoruz!
	defer originalPath.Close()

	// Yeni bir dosya oluştur.
	newFile, err := os.Create("./new_folder/tello_new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Verileri kaynaktan yeni dosyaya kopyala. (bytes türünde)
	bytesWritten, err := io.Copy(newFile, originalPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Toplam kopyalanan bytes: %d", bytesWritten)

	// Kopyalama işleme bittikten sonra kopyalanan veriler bellekte tutulur.
	// Bu verilerin tüm işlemler bittikten sonra bellekten boşaltılması gerekir.
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
