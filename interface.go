// interface adalah type data abstract, dia tidak memiliki  implementasi langsing
// sebuah interface berisikan definisi2 method
// biasanya digunakan sebagai kontrak
// ketika ingin membuat interface sesuai kontrak, harus buat struct yng sama deklrasi dengan interfacenya

// di golang serba otomatis dengan menyebutkan funciton yang di dalam interface untuk sebuah struct method,
// sudah menjalin knotrak dengan interface tersebut

package main

import "fmt"

type HasName interface {
	GetName() string
	GetAge() int
}

func sayHello(hasName HasName){
	fmt.Println("Hello ", hasName.GetName(), "iam ", hasName.GetAge(), "Years old")
}

type Person struct {
	Name string
	Age int
}

// struct method GetName() dari Person ini langsung terikat kontrak dengan interface HasName
func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}

func main() {
	var eko Person
	eko.Name = "Eko"
	eko.Age = 12
	sayHello(eko)
}