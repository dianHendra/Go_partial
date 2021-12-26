//struct merupakan template data yang digunakan untuk menggabungkan nol atau lebih type data lainnya dalam satu kesatuan
//struct itu ibarat class model data
//Struct Literals 

package main

import "fmt"

type Customer struct{
	Name, Address string
	Age int
}

func main(){
	var Eko Customer
	Eko.Name = "Eko Kurniawan"
	Eko.Address = "Indonesia"
	Eko.Age = 13

	//Struct Literals
	Joko := Customer{
		Name : "Jokko",
		Address : "Amerika",
		Age : 15,
	}

	fmt.Println(Eko.Address)
	fmt.Println(Eko)
	fmt.Println(Joko)

}

