//berisiakn fungsional untuk mengakses fitur sistem secata independen(untuk semua OS)

package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	fmt.Println("Argiment ", args)
	fmt.Println(args[1]) //Cob arun dengan go run packageOS.go Dia Hendra Saputra (untuk mengisi argument)
	fmt.Println(args[2])

	name , error := os.Hostname()

	if	error == nil {
		fmt.Println("hostname , ", name)
	} else {
		fmt.Println(error)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	//coba di cmd export username dan passwordnya dengan terminal,

	fmt.Println(username, password)
}