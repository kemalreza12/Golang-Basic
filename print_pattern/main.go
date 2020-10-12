package main

import (
	"fmt"
)

func cetakGambar(num int) {
	if num%2 == 1 {
		value := num + 1
		for i := 1; i < value; i++ {
			str := ""
			for j := 1; j < value; j++ {
				if j == 1 || j == num || i == (num/2)+1 {
					str += " * "
				} else {
					str += " = "
				}
			}
			fmt.Println(str)
		}
	} else {
		fmt.Println("Parameter harus bilangan ganjil")
	}
}

func main() {
	cetakGambar(5)
}
