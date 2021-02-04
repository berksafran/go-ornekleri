package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

// Dosyaları TAR olarak sıkıştırma
// .tar dosyası oluşturma

var fileFolderPath = "./files/"

// files klasörünün içindeki dosyalarımızı dizi içerisine alalım.
var files = []string{fileFolderPath + "hello.txt", fileFolderPath + "example.go"}

// addFile ... Bu fonksiyon ile var olan bir .tar dosyasına eklemeler yapıyoruz.
func addFile(fileName string, tw *tar.Writer) error {
	// Önce ilgili dosyayı açıyoruz.
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya açılırken hata (%s - %s): ", fileName, err)
	}
	// Dosyayı kapatmayı unutmuyoruz!
	defer file.Close()

	// Dosya ile ilgili bilgileri alıyoruz.
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Dosya bilgisi alınırken hata (%s - %s): ", fileName, err)
	}

	// Header'lar ile oluşturacağımız tar dosyasının temel bilgileri üzerinde çalışabiliyoruz.
	header := &tar.Header{
		ModTime: stat.ModTime(),
		Name:    fileName,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
	}

	// Oluşturduğumuz header'ı .tar dosyasının headerına ekliyoruz.
	if err := tw.WriteHeader(header); err != nil {
		msg := "TAR header yazılırken bir hata meydana geldi. %s: %s"
		return fmt.Errorf(msg, fileName, err)
	}

	// Kopyalama işlemine başlıyoruz.
	copied, err := io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("%s dosyası TAR'a yazılırken bir hata oluştu. : %s", fileName, err)
	}

	// Sıkıştırılan dosyanın boyutu her zaman, normal dosyadan küçük olmalıdır.
	if copied < stat.Size() {
		msg := "%s dosyasına %d kadar veri yazıldı. Ama beklenen veri %d kadardır."
		return fmt.Errorf(msg, fileName, copied, stat.Size())
	}

	// Eğer tüm işlemler başarılı ise fonksiyon da error dönmesi gerektiğinden
	// nil döndürüyoruz.
	return nil
}

// .tar dosyamızı oluşturalım.
func createArchiveTarFile(archiveFileName string) int {
	// Örnek bir parametre kontrolü.
	// String'in length'i 0 ise func int döndürdüğünden -1 döndürüyoruz.
	if len(archiveFileName) == 0 {
		return -1
	}

	// Dosyamızı oluştururken izinlerimizi de ayarlayalım.
	// Write izni, dosya yoksa oluşturması için CREATE izni,
	// var olan dosyanın sonuna ekleme yapmak izni TRUNC.
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveFileName+".tar", flags, 0644)
	if err != nil {
		log.Fatalf(".tar dosyasına yazmak için açılırken hata meydana geldi! : %s", err)
		return -1
	}

	// Açılan dosyayı her zaman kapatmalıyız!
	defer file.Close()

	tw := tar.NewWriter(file)
	defer tw.Close()

	// files klasöründe bulunan tüm dosyaları TAR'a ekliyoruz.
	for _, fileName := range files {
		if err := addFile(fileName, tw); err != nil {
			log.Fatalf("%s dosyası TAR dosyasına eklenirken hata meydana geldi! %s", fileName, err)
		}
	}

	// Tüm işlemleri başarılı ise fonksiyon hata döndüreceği için -1 gönderiyoruz.
	// 1 hata yok anlamına gelmektedir.
	return 1
}

/*
	Önemli NOT: Eğer bu dosyayı bir package haline getirmek istersek,
	-> Fonksiyonların baş harflerini büyük yaparak public haline getirmemiz,
	-> Gelen parametreleri doğru mu değil mi diye kontrol etmemiz gerekir.

	Biz burada package haline getirmediğimiz için gerek duymadık.
*/
func main() {
	result := createArchiveTarFile("dosyaX")
	if result > 0 {
		fmt.Println("İşlem başarılı! : ", result)
	} else {
		fmt.Println("İşlem başarısız oldu! : ", result)
	}
}
