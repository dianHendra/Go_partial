//perulangan di golang hanya ada forloop
package main

import "fmt"

func main(){
	//long statement
	//counter := 1
	
	// for counter <=10{
	// 	fmt.Println("loop ", counter)
	// 	counter++
	// }

	//short statment
	// for counter := 1; counter <= 10; counter++{
	// 	fmt.Println("loop ", counter)
	// }

	//range statement
	slice := []string{"dian", "hendra", "Saputra"}
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	//bisa seperti ini juga
	for i, val:= range slice{
		fmt.Println("index: ", i, " => ", val)
	}

	//contoh untuk type data map
	person := make(map[string]string)
	person["name"] = "dian"
	person["title"] = "programmer"

	for key, val := range person{
		fmt.Println(key, " adalah ", val)
	}

}