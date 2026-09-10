package main

import (
	"fmt"
)

func main() {
	index := 1
	for index <= 50{
		if index % 2 == 0 {
			fmt.Println(index)
		}
		index++
	}	
}