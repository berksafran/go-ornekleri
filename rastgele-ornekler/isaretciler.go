package main

import "fmt"

type Person struct {
	Name string
}

// Introduce is İlişkilendirilmesini sağlıyor..
func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// // Introduce is İlişkilendirilmesini sağlıyor..
// func (s *Saiyan) Introduce() {
// 	fmt.Printf("Hello!, I'm %s\n", s.Name)
// }

type Saiyan struct {
	Person
	Power int
}

func main() {

	// and to use it:
	goku := Saiyan{
		Person: Person{"Goku"},
		Power:  9001,
	}
	goku.Introduce()
	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
	goku.Introduce()
	goku.Person.Introduce()
}
