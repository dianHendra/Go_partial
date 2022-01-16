// untuk menentukan access modifier cukup dengan nama function atau variable
// jika dimulai dengan huruf besar : bisa diakses dari package lain
// jika huruf kecil tidak dapat diakses dari package lain
package main

//path harus dimulai dari child gopath
import (
	"PartialTutorial/Go_partial/helper"
	"PartialTutorial/Go_partial/other"
)

func main(){
	helper.SayHello("Dian")
	other.SayOther("Dian")
	//helper.sayGoodbay("Dian") //ini akan error jika di uncoment
}