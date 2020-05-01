package util

import "fmt"

func Print(str string) string {
	fmt.Println("i/p", str)
	fmt.Println("o/p", "Util: "+str)
	return ("Util: " + str)
}
