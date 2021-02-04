package main

import (
	"fmt"
	"sync" //WaitGroup'u kullammak için
	"time" //bekleme işlemleri için
)

/*
* waitgroup nesnesini işaretçi olarak parametre veriyoruz.
* işaretçi olarak vermemizin sebebi, programın bekleme işlemi için
* asıl waitgroup nesnesini kontrol etmesidir.
 */
func fonksiyon1(wg *sync.WaitGroup) {

	//fonksiyonun 2 sn beklemesini istiyoruz.
	time.Sleep(2 * time.Second)
	fmt.Println("Fonk1 tamamlandı")

	//wg.Done() fonksiyonu ile waitgroup nesnesini -1 azalttık.
	wg.Done()
}

//bu fonksiyonumuza da wg nesnesini işaretçi ile parametre olarak verdik.
func fonksiyon2(wg *sync.WaitGroup) {
	//fonksiyonu 3 sn uyuttuk.
	time.Sleep(3 * time.Second)
	fmt.Println("Fonk2 tamamlandı")

	//-1 daha eksilttik.
	wg.Done()
}

func main() {
	/*
	* Öncelikle waitgroup'u kullanabilmek için  bir waitgroup
	* nesnesi oluşturuyoruz.
	 */
	var wg sync.WaitGroup

	/*
	* waitgroup'a 2 ekliyoruz. Yani 2 tane işlemden yanıt gelmesini
	* beklemesini istiyoruz. Aslında burada beklemeyecek. Sadece
	* işlem sayısını belirttik.
	 */
	wg.Add(2)

	/*
	* fonksiyon1 ve fonksiyon2'ye oluşturduğumuz wg örneğinin
	* bellekteki adresinin veriyoruz.
	 */
	go fonksiyon1(&wg)
	go fonksiyon2(&wg)
	fmt.Println("Merhaba Dünya!")

	/*
	* Burada wg.Wait() fonksiyonu ile asenkron işlemleri beklemesini
	* sağlıyoruz. yani waitgroup'un 0'a düşmesini bekliyoruz.
	* Eğer waitgroup olmadan yapsaydık. asenkron fonksiyonlarımızın tamamlanmasını
	* beklemeden program kendini sonlandırırdı.
	 */
	wg.Wait()

	//waitgroup tamamlandığında ekrana yazı bastıralım.
	fmt.Println("WaitGroup'lar tamamlandı.")
}
