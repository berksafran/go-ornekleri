package main

import (
	// Tanımladığımız models/config.go'yu import etmemiz gerekiyor.
	"io/ioutil"
	"log"
	// . "./models" // models içindeki tüm nesneleri import et.
	// "gopkg.in/yaml.v2"
)

func main() {

	// YAML Dosyamızı tanımlıyoruz.
	fileName := "./config.yaml"

	/* Eğer kullanıcı tarafından alınmasını isteseydik
	   os.Args[1] methodunu kullanabilirdik.
	   Çalıştırırken: go run main.go config.yaml
	*/

	// var config Config
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	print(string(source))

	/*
		Config model'ını "gopkg.in/yaml.v2" kütüphanesindeki yaml.Unmarshal(..) komutunu kullanmak için ekledik.
		Bu kodda yer almadığı için hem paketi, hem de model'i comment altına aldık.
	*/
}
