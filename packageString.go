//Berisikanfunction untuk pengolahan string

package main

import (
	"fmt"
	"strings"
)

var kata = "Dian Hendrda Saputra" 
func main(){
	fmt.Println(strings.Contains(kata, "Dian"))
	fmt.Println(strings.ToLower(kata))
	fmt.Println(strings.ToUpper(kata))
	fmt.Println(strings.Split(kata," "))
	fmt.Println(strings.Trim("  ini harusnya di trim ", " "))
	fmt.Println(strings.ReplaceAll("dian dian andi", "dian", "Anda"))
	fmt.Println(strings.Replace("dian dian andi", "dian", "Anda", 1))
}