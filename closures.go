// kemampuan funciton berinteraksi dengan data disektarnya dalam scope yang sama
// harap bijak dalam pengguaanya jika tidak dapat mengubah data yang seharusnya tidak diubah

package main

import "fmt"

func main(){
	counter := 0
	name := "eko"
	fmt.Println("counter is ",counter)
	fmt.Println("name is ", name)
	

	increment := func() {
		fmt.Println("increment")
		counter++
		//misal ini tidak sengaja malah counter di ++ sehinggal tiap jalan function increment 
		//counter di dalam  func main terus ditambah
		name := "dian" 
		//dapat ditasi dengan mendeclare variable jika tidak ingin bersinggungan dengna variable yang sma di dalam funcmain
		fmt.Println("name increment is ", name)
	}

	increment()
	fmt.Println("counter now is ",counter)
	increment()
	fmt.Println("counter now is ",counter)
	fmt.Println("name still ", name)
	
	
}