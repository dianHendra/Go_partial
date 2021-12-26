package main

import "fmt"

func sumAll(numbers ...int) int{
	total := 0
	for _, value := range numbers{
		total += value
	}
	return total
}


func main(){
	total := sumAll(10, 10, 20)
	slicenum := []int{20, 90, 80, 12}
	
	jumlah := sumAll(slicenum...)
	fmt.Println(total)
	fmt.Println(jumlah)
}