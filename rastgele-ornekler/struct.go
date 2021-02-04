package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	age       int
	isAdmin   bool
}

type insan struct {
	isim string
	yas  int
}

// Struct'a ait fonksiyonlar oluşturmak:
func (kisi insan) tanimla() {
	fmt.Printf("Tanımlandı: %s %d", kisi.isim, kisi.yas)
}

func main() {

	// struct nesne oluşturmamızı sağlar.

	// Normal struct
	user1 := user{"Berk", "Safran", 31, true}

	// Sonradan property ekleyebiliyoruz.
	user2 := user{}
	user2.lastname = "Yilan"
	user2.age = 35
	user2.isAdmin = false
	user2.firstname = "Ahmet"

	// Parametre yerinde propertyleri belirtip tanımlayabiliyoruz.
	user3 := user{isAdmin: false, lastname: "Ördek", firstname: "Fahrik", age: 65}

	// Anonim structlar (Hemen çalışırlar. Aynı anonim func'lar gibi)
	user4 := struct {
		name       string
		surname    string
		age        int
		percentage float64
		isAdmin    bool
	}{
		"Ayşe", "Fatma", 23, 32.4, true,
	}

	// Struct'a fonksiyon ekleme (methodlar)
	user5 := insan{"Berk", 25}
	user5.tanimla()

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
	fmt.Println(user4)
	fmt.Println(user5)
}
