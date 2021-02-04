package main

import (
	"fmt"
	"log"
	"os"
)

// Bir dosyanın okuma ve yazma izinlerini kontrol etmek.

func main() {

	// jello.txt dosyasından yazma izinlerini kaldırdık.
	file, err := os.OpenFile("jello.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			// log ve fmt arasındaki fark: zaman damgası.
			log.Println("Hata:", err) // Output: 2021/01/31 13:40:35 Hata: open jello.txt: permission denied
			fmt.Println("Hata:", err) // Output: Hata: open jello.txt: permission denied
		}
	}

	defer file.Close()
}
