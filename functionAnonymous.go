// Function Anonymous
package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist){
	if	blacklist(name){
		fmt.Println("you are blocked ", name)
	} else {
		fmt.Println("welcome ", name)
	}
}

//sebelumnya kita harus buat fungsi dulu


func main(){
	bl := func(name string) bool{
		return name == "admin"
	}

	registerUser("admin", bl)
	registerUser("eko", bl)

	//atau juga ada bisa langsung dimasukan ke dalam parameter
	registerUser("root", func(name string) bool{
		return name == "root"
	})
	
}