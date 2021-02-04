package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Önce xml dosyamızı açalım.
	file, err := os.Open("sites.xml")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	// Her zamanki gibi açtığımız dosyayı kapatmayı unutmuyoruz.
	defer file.Close()

	// Açtığımız dosyadaki verileri okuyalım.
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Ürettiğimiz ObjectSites (genel yapı bu olduğu için)
	// struct nesne tipini bir değişkene atayalım.
	v := ObjectSites{}

	// XML verilerimizi ObjectSites nesnemize aktaralım.
	// ObjectSites xml dosyası ile birebir aynıdır. Buna dikkat!

	/*
	 Pointer olarak vermemizin sebebi, parametre içerisinde bir kopyası alınarak; kopya üzerinde işlem yapılır.
	 Biz ürettiğimiz asıl nesne üzerinde işlem yapmak istiyorsak, & ile pointerını vermemiz lazım.
	*/
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// fmt.Println(v) // Tüm ObjectSites'ı yazdırmak için.
	fmt.Println(v.Sites[1].Name) // www.berksafran.com
	v.Sites[1].Name = "www.testdeneme.com.tr"
	fmt.Println(v.Sites[1].Name) // www.testdeneme.com.tr

}

// ObjectSites contains sites from XML
// Tüm web sitelerini site nesnesini kullanarak dizi içerisinde tutalım.
type ObjectSites struct {
	XMLName     xml.Name `xml:"sites"`
	Version     string   `xml:"version,attr"`
	Sites       []site   `xml:"site"`
	Description string   `xml:",innerxml"`
}

type site struct {
	XMLName     xml.Name `xml:"site"`
	Name        string   `xml:"Name"`
	Description string   `xml:"Description"`
	Category    string   `xml:"Category"`
}
