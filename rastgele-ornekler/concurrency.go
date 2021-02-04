package main

import (
	"fmt"
	"time"
)

func main() {

	//2 string taşıyan bir kanal oluşturalım
	k := make(chan string, 3)
	//bu kanalımız string değer taşıyacak

	//asenkron bir iş parçacığı oluşturalım
	go func() {

		//bu iş parçacığı 5 sn beklesin
		time.Sleep(time.Second * 5)

		//k kanalına string ilk değeri gönderelim
		fmt.Println("Kanala 1. değer gönderiliyor.")
		k <- "Merhaba"

		// 2 sn bekletelim
		time.Sleep(time.Second * 2)

		fmt.Println("Kanala 2. değer gönderiliyor.")
		k <- "World!"

		time.Sleep(time.Second * 3)

		fmt.Println("Kanala 3. değer gönderiliyor.")
		k <- "Selam Dünyalı!"

		fmt.Println("Tüm değerler gönderildi.")
	}()

	//ana iş parçacığı k kanalına değer gelene kadar bekleyecek
	data, secondData, thirdData := <-k, <-k, <-k
	fmt.Println("Kanala 3 değer de geldi:", data, secondData, thirdData)
	fmt.Println("Program sonlandı.")
	//değer geldiğinde program sonlanacaktır.
}
