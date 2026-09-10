package main

import (
	"fmt"
)

func main() {
	i1,i2 := 1,1
	for i1 <= 9 {
		for i2 <= 9 {
			fmt.Println(i1*i2)
			i2++
		}
		i1++
		i2 = 1
	}
}