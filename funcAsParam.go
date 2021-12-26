//function as param
package main

import "fmt"

type Filter func(string) string

func sayHelloWIthFilter(name string, filter Filter){
	nameFiletered := filter(name)
	fmt.Println("Hello ", nameFiletered)
}

func spamFilter(name string) string {
	if name == "anjing"{
		return "..."
	} else{
		return name
	}
}

func main(){
	sayHelloWIthFilter("eko", spamFilter)
	//atau dapat juga seperti ini
	filter := spamFilter
	sayHelloWIthFilter("anjing", filter)
}