//merupakak kemampuan untuk merubah type dara menjadi type data yang diinginkan
//biasanya sering digunakan ketika bertemu dengan interface kosong
//pastikan ketika mengubah tipe data return dari func interface tersebut sesuai dengan apa yang di merupakan ubahan kita

package main

import "fmt"


func random() interface {} {
	return "string true"
}
func main(){
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	//resultNumber := result.(int)//jika tidak sesuai dengan return, akan terjadi panic

	//atau bisa juga menggunakan switch
	switch value := result.(type){
	case string:
		fmt.Println("string ",  value)
	case int:
		fmt.Println("int ", value)
	default:
		fmt.Println("unknow")
	}

}