package main

import (
	"fmt"
)

func main() {

	total,index := 0,1

	for index <= 100 {
		total += index
		index++
	}

	fmt.Println(total)

}