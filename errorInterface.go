//interface yang digunakan sebagia kontrak membuat error.
// nama interface nya adalah error

package main

import (
	"fmt"
	"errors"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi degan url")
	} else {
		result := nilai/pembagi
		return result, nil
	}
}

func main() {
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())
	
	hasil, err := pembagian(100, 5)

	if err == nil {
		fmt.Println("hasil = ", hasil)
	} else {
		fmt.Println("error => ", err.Error())
	}
}