// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	kanal := make(chan string) //kanal oluşturuyoruz
// 	go func() {
// 		fmt.Println("Kanala değer gönderiliyor...")
// 		time.Sleep(time.Second * 2)    //2 saniye uyku
// 		kanal <- "Öylesine bir mesaj." //İletişime geçiriyoruz
// 		fmt.Println("Kanala değer gönderildi!")
// 	}()
// 	fmt.Println("Gelen Değer:", <-kanal) //kanaldan gelen veri bekleniyor
// }
