package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) UpdateNameInformation(newNameInformation string) {
	p.FirstName = newNameInformation
	fmt.Println("UpdateNameInformation:", p)
}

func (p *Person) UpdateNameInformationWithPointer(newNameInformation string) {
	p.FirstName = newNameInformation
	fmt.Println("UpdateNameInformationWithPointer:", p)
}

func main() {

	person := Person{
		FirstName: "sertunc",
		LastName:  "selen",
	}

	fmt.Println("main:", person)

	//person.FirstName = "cahit"
	//person.UpdateNameInformation("cahit")
	person.UpdateNameInformationWithPointer("cahit")

	fmt.Println("sonuc:", person.FirstName, person.LastName)
}
