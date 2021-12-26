// function yang memanggil dirinya sendiri

//contoh kasus factorial (pohon faktor)

package main 

import "fmt"

//looping
func factorialLoop(val int) int{
	result := 1
	for i := val; i > 0; i-- {
		result *= i
	}
	return result
}

//recursive function
func factorialFuncRecursive(value int) int{
	if	value == 1{
		return 1
	} else {
		return value * factorialFuncRecursive(value -1)
	}
}

func main(){
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(factorialLoop(10))
	
	recursive := factorialFuncRecursive(10)
	fmt.Println(recursive)
}