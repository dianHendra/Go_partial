//secara default data di passing by value bukan reference
//jadi ketika kita mengirim variable ke tempat lain artinya itu adalah duplikasi dari value tersebut
//pointer merupakan kebalikan dari pass by value, seperti reference di bahasa lain
//untuk membuat variable pointer ke variable lain dapat menggunakan "&variable"

//adapun penggunaan operator *

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	address1 := Address{"Subang", "jabar", "indo"}
	address2 := address1
	fmt.Println("address1 = ",address1)

	address2.City = "Bandung"
	fmt.Println("------address2 diubah jadi Bandung----------------")
	fmt.Println("address1 = ",address1.City)
	fmt.Println("address1 = ",address2.City)
	
	address3 := &address1
	address3.City = "SUrabaya"
	//jika menggunkan operator & untuk merubah semua itemnya, tiadk adakan berhasil, malah buat object baru
	address2 = Address{"Jakarta", "DKI", "indo"}
	fmt.Println("-----address 3 reference ke address 1 & address2 malah buat new struct----")
	fmt.Println("address1 = ",address1)
	fmt.Println("address2 = ",address2)
	fmt.Println("address3 = ",address3)
	//coba address3 diubah semua itemnya
	
	address3 = &Address{"JKT", "DKI", "indo"}
	fmt.Println("-----sudah di ubah ke pointer dan address2----")
	fmt.Println("address1 tetap = ",address1)
	fmt.Println("address3 = ",address3 )

	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")
	fmt.Println("------penggunaan operator *------")

	var address4 *Address = &address1
	*address4 = Address{"Malang", "Jatim", "indonesia"}
	//atau kalo mau bikin baru
	//var address4 *Address = &Address{"Malang", "Jatim", "indonesia"}

	fmt.Println("address1 = ",address1)
	fmt.Println("address2 = ",address2)
	fmt.Println("address3 udah terpisah = ",address3 )
	fmt.Println("address4 = ",address4 )

	//golang juga ada new untk membuat pointer denga isi data kosong
	var addressNew *Address = new(Address)
	fmt.Println("addressNew = ",addressNew )

}