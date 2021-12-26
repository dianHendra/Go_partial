package main

import "fmt"

func sayHeloTo(name string){
	fmt.Println("Hello ", name)
}

//funtion with return value
func getFullName() (string, string){
	return "john", "kenedi"
}

//named return value
func GetCompleteName() (firstName, middleName, lastName string){
	firstName = "dian"
	middleName = "hendra"
	lastName = "saputra"

	return
}
func main(){
	sayHeloTo("--------------function------------------")
	sayHeloTo("din syamsudin")

	//exec funtion with return
	//firstName, lastName := getFullName()

	//jika tidak butuh dengan lastname (return kedua)
	sayHeloTo("--------------function with parameter------------------")
	firstName, _ := getFullName()	

	fmt.Println(firstName)
	//fmt.Println(lastName)

	//exec named  parameter
	a, b, c := GetCompleteName()
	sayHeloTo("--------------return func with named parameter------------------")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)



}