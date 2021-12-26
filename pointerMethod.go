// sama seperti di parameter func, di struct disarankan untuk menggunakan pointer
package main

import "fmt"

type Man struct{
	Name string
}
//mrhod biasa
func (man Man) Married(){
	man.Name = "Mr. " + man.Name
	
}
func (man *Man) MarriedP(){
	man.Name = "Mr. " + man.Name
	fmt.Println("Check in method name is", man.Name)
}

func main(){
	Eko := Man{"eko"}
	Eko.Married()
	fmt.Println(Eko.Name)
	Eko.MarriedP()
	fmt.Println("setelah di kasih pointer ", Eko.Name)

}