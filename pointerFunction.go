// saat memebuat parameter pada function secara default pass by value,
// jadi ketika data tsb diubah dalam funciton tidak akan merubah data aslina
// untuk dapat mengubah dat asli yang dijadikan parameter dan diubah dalam function 
// dapat menggunakan pointer
// untuk menjadikan parameter sebagai pointer kita dapat menggunakan operator * di parameter

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryTobeIndo(add *Address){
	add.Country = "indonesia"
}
func changeCountryTobeMalay(add *Address){
	add.Country = "Malaysia"
}
func main(){
	alamat := Address{
		City : "Subang",
		Province : "Jawa Barat",
		Country : "",
	}
	//dapat di direct seperti ini
	changeCountryTobeIndo(&alamat)
	//atau dengan seperti ini
	var alamat2 *Address = &alamat
	changeCountryTobeMalay(alamat2)
	fmt.Println(alamat)
}