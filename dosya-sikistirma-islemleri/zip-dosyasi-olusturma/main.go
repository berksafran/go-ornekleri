package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

// Dosyaları ZIP olarak sıkıştırma
// .zip dosyası oluşturma

var fileFolderPath = "./files/"

// files klasörünün içindeki dosyalarımızı dizi içerisine alalım.
var files = []string{fileFolderPath + "hello.txt", fileFolderPath + "example.go"}

func addFile(fileName string, zw *zip.Writer) error {
	// Önce ilgili dosyayı açıyoruz.
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya açılırken hata (%s - %s): ", fileName, err)
	}
	// Dosyayı kapatmayı unutmuyoruz!
	defer file.Close()

	// zip dosyasını oluşturuyoruz.
	wr, err := zw.Create(fileName)
	if err != nil {
		msg := "%s zip dosyası içerisine yeni bir dosya oluşturulurken bir hata meydana geldi. %s: %s"
		return fmt.Errorf(msg, fileName, err)
	}

	// Kopyalama işlemine başlıyoruz.
	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("%s dosyası zip'e yazılırken bir hata oluştu. : %s", fileName, err)
	}

	// Eğer tüm işlemler başarılı ise fonksiyon da error dönmesi gerektiğinden
	// nil döndürüyoruz.
	return nil
}

// .zip dosyamızı oluşturalım.
func createArchiveZipFile(archiveFileName string) int {
	// Örnek bir parametre kontrolü.
	// String'in length'i 0 ise func int döndürdüğünden -1 döndürüyoruz.
	if len(archiveFileName) == 0 {
		return -1
	}

	// Dosyamızı oluştururken izinlerimizi de ayarlayalım.
	// Write izni, dosya yoksa oluşturması için CREATE izni,
	// var olan dosyanın sonuna ekleme yapmak izni TRUNC.
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveFileName+".zip", flags, 0644)
	if err != nil {
		log.Fatalf(".zip dosyasına yazmak için açılırken hata meydana geldi! : %s", err)
		return -1
	}

	// Açılan dosyayı her zaman kapatmalıyız!
	defer file.Close()

	zw := zip.NewWriter(file)
	defer zw.Close()

	// files klasöründe bulunan tüm dosyaları ZIP'e ekliyoruz.
	for _, fileName := range files {
		if err := addFile(fileName, zw); err != nil {
			log.Fatalf("%s dosyası ZIP dosyasına eklenirken hata meydana geldi! %s", fileName, err)
		}
	}

	// Tüm işlemleri başarılı ise fonksiyon hata döndüreceği için -1 gönderiyoruz.
	// 1 hata yok anlamına gelmektedir.
	return 1
}

func main() {
	result := createArchiveZipFile("yeni_zip_file")
	if result > 0 {
		fmt.Println("İşlem başarılı! : ", result)
	} else {
		fmt.Println("İşlem başarısız oldu! : ", result)
	}
}
