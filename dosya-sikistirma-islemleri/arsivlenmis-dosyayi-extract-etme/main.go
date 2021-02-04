package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Arşivlenmiş (sıkıştırılmış) dosyayı extract etmek
// Arşiv dosyası türü = .zip

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Klasör oluşturuluyor:", dir)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Şu klasör zaten var:", dir)
	}
}

func main() {
	// İlgili .zip dosyasını açıyoruz.
	zr, err := zip.OpenReader("yeni_zip_file.zip")
	if err != nil {
		log.Fatal(err)
	}

	// Dosyayı açtığımız için kapatmayı unutmuyoruz!
	defer zr.Close()

	// Zip dosyası içerisindeki dosyaları geziyoruz.
	// zr.Reader.File bir dizi. İçerisindeki nesneler ise birer struct.
	for _, file := range zr.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(targetDir, file.Name)

		// Gelen dosya isimlerinde klasör ve dosya ismini ayırmamız gerekiyor.
		dirName := strings.Split(file.Name, "/")

		createDirIfNotExist(dirName[1])

		// fmt.Println(extractedFilePath)
		/*
			Output:
			files/hello.txt
			files/example.go
		*/

		if file.FileInfo().IsDir() {
			log.Println("Klasör oluşturuluyor..", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("Dosya çıkarılma işlemine başlandı.. Dosya:", file.Name)

			// İlgili dosya içerisindeki verileri kopyalacağız.
			// Önce ilgili dosyayı açmamız gerekiyor.
			outFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}

			defer outFile.Close()

			// Kopyalama işlemine başlıyoruz.
			// ZippedFile'dan extract ettiğimiz dosyalara verileri de aktarıyoruz.
			_, err = io.Copy(outFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
