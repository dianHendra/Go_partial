//function as value
package main

import "fmt"


func getGoodby(name string) string{
	return "goodby " + name
}

func main(){
	sayGoodby := getGoodby
	result := sayGoodby("dian")
	fmt.Println(result)
	fmt.Println(getGoodby("eko"))
}