package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Oluşturulacak dosyanın adı? ")
	file, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	text := getText()

	writeTextOnDocument(text, file)
}

func writeTextOnDocument(text, file string) {
	createFile(file)
	doc, err := os.OpenFile(file, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	byteSlice := []byte(text)

	bytesWritten, err := doc.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Dosyaya yazma işlemi tamamlandı. Toplam yazılan %d bytes.", bytesWritten)

	defer doc.Close()

}

// getData from User on CLI
func getText() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Eklemek istediğiniz yazıyı yazın: ")
	text, err := reader.ReadString('\n')
	// Sağ ve solundaki boşlukları trimle.
	text = strings.TrimSpace(text)
	if err != nil {
		log.Fatal(err)
	}

	return text
}

// createFile Yazılacak hedef dosyayı oluştur
func createFile(fileName string) {
	file, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Dosya bulunamadı, oluşturuluyor..")
		if os.IsNotExist(err) {
			newFile, err := os.Create(fileName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Dosya oluşturuldu. Dosya adı: %s \n", newFile.Name())
			return
		}
	}
	fmt.Println("Dosya zaten varmış.", file.Name())
}
