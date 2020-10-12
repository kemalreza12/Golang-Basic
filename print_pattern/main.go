package main

import (
	"fmt"
)

func cetakGambar(num int) {
	if num%2 == 1 {
		count := num + 1
		for row := 1; row < count; row++ {
			str := ""
			for col := 1; col < count; col++ {
				if col == 1 || col == num || row == (num/2)+1 {
					str += " * "
				} else {
					str += " = "
				}
			}
			fmt.Println(str)
		}
	} else {
		fmt.Println("Parameter harus merupakan bilangan ganjil")
	}
}

func main() {
	cetakGambar(5)
}
