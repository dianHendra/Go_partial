package main

//path harus dimulai dari child gopath
import (
	"PartialTutorial/Go_partial/helper"
	"PartialTutorial/Go_partial/other"
)

func main(){
	helper.SayHello("Dian")
	other.SayOther("Dian")
}