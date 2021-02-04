package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// xml datalarımızı okumak için ilgili xml dosyasını açıyoruz.
	xmlFile, err := os.Open("Employees.xml")
	// Hata yönetimini unutmuyoruz.
	if err != nil {
		log.Fatal(err) // Bunu yazmazsak hata alamayız.
	}
	// Dosyayı açtığımız için kapatmayı da unutmuyoruz.
	defer xmlFile.Close()

	// xml datalarımızı okumaya başlıyoruz.
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	// Okuduğumuz dataları bir değişken tanımlayarak, değişkene atıyoruz.
	var c Company

	// Unmarshal operasyonu ile xmlData'daki xml verilerini değişkenimize değer olarak atıyoruz.
	err = xml.Unmarshal(xmlData, &c)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// Sonucu ekrana yansıtalım.
	fmt.Println(c.People)

	// BONUS: XML verilerimizi JSON'a çevirelim.

	var person jsonPerson
	var people []jsonPerson

	for _, value := range c.People {
		person.ID = value.ID
		person.FirstName = value.FirstName
		person.LastName = value.LastName
		person.UserName = value.UserName

		people = append(people, person)
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		os.Exit(1) // Uygulamayı kapatmanın farklı bir yolu.
	}

	// Ekrana yazdıralım.
	fmt.Println(string(jsonData))

	// Sonrasında da bir JSON dosyası oluşturup yazdıralım.
	jsonFile, err := os.Create("./Employees.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
}

// Person ... BONUS: JSON Person nesnemiz
type jsonPerson struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
}

// Person nesnemiz
type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	UserName  string   `xml:"username"`
}

// Company is ... Person nesnelerimizi içeren Company nesnesi
type Company struct {
	XMLName xml.Name `xml:"company"`
	People  []Person `xml:"person"`
}

// String ... Person nesnesine bağlı String() methodu tanımlıyoruz.
// TRICK: Bu String methodunu tanımlayarak, Person nesnesini yazdırmak istediğimizde bu method çalışır.
func (p Person) String() string {
	// string'ler için %s, integer'lar için %s kullanabiliriz.
	return fmt.Sprintf("\t ID: %d - Firstname: %s - Lastname: %s - Username: %s \n", p.ID, p.FirstName, p.LastName, p.UserName)
}
