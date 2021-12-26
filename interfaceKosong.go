//adalah interface yang tidak memiliki method atau fuc satupun
//otomatis semua type data akan menjadi implementasinya
//beberapa contoh peggunaan interface kosong seperti pada 
// fmt.Println(a... interface{})
// panic(v interface{})
// recover() interface{} dan lain2

package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1{
		return 1
	} else if i == 2{
		return 2
	} else {
		return "ups"
	}
}

func main(){
	var data interface{} = Ups(3)
	fmt.Println(data)
}