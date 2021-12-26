package main

import "fmt"

func main(){
	var name = "hendra"

	if	name == "dian"{
		fmt.Println("nama saya dian")
	} else	{
		fmt.Println("saya bukan dian")
	}

	switch name{
	case "hendra":
			fmt.Println("saya adalah hendra")
			break
	case "dian":
		fmt.Println("saya adalah dian")
		break
	default:
		fmt.Println("saya orang")
		break
	}
}