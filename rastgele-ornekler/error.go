package main

import (
	"fmt"
	"os"

	"github.com/martinlindhe/notify"
)

func main() {
	if d := "isaretciler.go"; dosyaVarmı(d) {
		fmt.Println(d, "bulunuyor")
		notify.Notify("Berk App", "Başlık", "Hello guys!", "")
	} else {
		fmt.Println(d, "bulunmuyor!")
	}
}

func dosyaVarmı(isim string) bool {
	bilgi, hata := os.Stat(isim)
	fmt.Println(bilgi)
	fmt.Println(hata)
	if os.IsNotExist(hata) {
		return false
	}
	return !bilgi.IsDir()
}
