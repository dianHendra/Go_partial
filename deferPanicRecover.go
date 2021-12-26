//defer merupakan funciton yang akan dieksekusi setelah sebuah funciton di eksekusi
//panic merupakan function yang bisa digunakan untuk menghentikan program
//panic dipanggil ketika terjadi error biasannya
//recover adalah function yan dapat digunakan untuk menangkap data panic
//dengan recover panic tidak akan menghentikan applikasinya

package main

import "fmt"

//defer
func logging(){
	fmt.Println("DEFER___loging funciton...")
}

//panic
func endApp(){
	fmt.Println("DEFER___Apps selesai")
	//supaya dijalankan di akhir dapat disimpan ke func defer
	
	message := recover()
	if message != ""{
		fmt.Println("Error ", message)
	}
}

func runApp(error bool){
	defer endApp()
	if	error{
		panic("Applikasi ERROR...")//functio panic
	}
	fmt.Println("Applikasi Berjalan...")
	
}

func runApplication(num int){
	defer logging()//ini akan dieksekusi setelah runApplication selesai exec
	result := 12/ num
	fmt.Println("Run Application result ", result)
}

func main(){
	runApplication(2)
	runApp(true)
	fmt.Println("checking app is end__")
}