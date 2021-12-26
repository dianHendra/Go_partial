//struct dapat juga digunakan sebagai parameter function
//method adalah function.
// jadi didalam struct ditambahkan function degnan membuat function yang seolah menemepel di struct
package main

import "fmt"

type Customer struct{
	Name, Address string
	Age int
}

//functio dengan parameter struct
func sayHi(cus Customer, name string){
	fmt.Println("Hello", name, "My Name is", cus.Name)
}

// struct method
func (cus Customer) sayHay(name string){
	fmt.Println("Hello", name, "My Name is", cus.Name)
}


func (cus Customer) sayGoodby(name string){
	fmt.Println("goodby", name, "My Name is", cus.Name)
}

func main(){
	var Eko Customer
	Eko.Name = "Eko Kurniawan"
	Eko.Address = "Indonesia"
	Eko.Age = 13

	sayHi(Eko, "dian")

	Eko.sayHay("dian")
	Eko.sayGoodby("dian")

	//Struct Literals
	// Joko := Customer{
	// 	Name : "Jokko",
	// 	Address : "Amerika",
	// 	Age : 15,
	// }

	// fmt.Println(Eko.Address)
	// fmt.Println(Eko)
	// fmt.Println(Joko)

}