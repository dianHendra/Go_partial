//semua data yang belum diinisiasi biasanya default null
// di golang ada katakunci "nil" yang hanya bisa digunakan pada type data
// interface, function, map, slice, pointer dan channel

package main
import "fmt"

func NewMap(name string) map[string] string {
	if name == ""{
		return nil
	} else {
		return map[string]string{
			"name" : name,
		}
	}
}

func main(){
	//biasanya 
	var person map[string] string = nil
	fmt.Println(person)

	orang := NewMap("dian")
	fmt.Println(orang)

	datakosong := NewMap("")
	if datakosong == nil {
		fmt.Println("Data kosong")
	}
}