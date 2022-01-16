//conversi type data yang berbeda (int to string, atau sebaliknya)
package main

import (
	"fmt"
	"strconv"
)

func main(){
	boolean, err := strconv.ParseBool("salah")
	number, err := strconv.ParseInt("200", 10, 64)
	bin := strconv.FormatInt(200, 2)

	if err == nil{
		fmt.Println(boolean)
		fmt.Println(number)
		fmt.Println(bin)
	} else {
		fmt.Println("error", err.Error())
	}
}