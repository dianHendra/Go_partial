// kita dapat membuat func yan bisa diakses ketika package diakses
// cocok untuk yang berhubungan dengan koneksi ke db
// cukup buat func dengan nama init
// Blank indentifier
// kadang kita ingin menjalankan init func di package tanpa harus mengeksekusi salah satu func nya
// secara default Go akan komplen ketika ada package yang di import tidak digunakan
//untuk menanganinya dapat menggunakan blank identifier(_) sebelum nama package ketka import 


package main

import (
	"PartialTutorial/Go_partial/database"
	"fmt"
	_"PartialTutorial/Go_partial/other"//initidak dipakai tapi di import
)


func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}