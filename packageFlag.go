//package flag bersisi funcsionalitas untuk memparsing comand line argumant

package main

import (
	"fmt"
	"flag"
)

func main(){
	var host *string = flag.String("host", "localhost", "put your database host") 
	var user *string = flag.String("user", "root", "put your database user") 
	var pass *string = flag.String("password", "root", "put your database password") 
	var number *int = flag.Int("number", 0, "put your database number") 

	flag.Parse()//jika tidak dipanggil flag tidak adan dapat diuah dari terminal

	fmt.Println("Host ", host)
	fmt.Println("User name ", user)
	fmt.Println("Password ", pass)
	//untuk membedakan cara menagambil data dari pointer coba run dengan seperti ini
	fmt.Println("Host ", *host)
	fmt.Println("User name ", *user)
	fmt.Println("Password ", *pass)
	fmt.Println("number ", *number)
	//untuk pengetesan flag nya
	//coba run dengan go run packageFlag.go -user=Dian

}