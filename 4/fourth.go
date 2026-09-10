package main

import (
	"fmt"
)

func main() {
	input := 0
	fmt.Print("Input number : ")
	fmt.Scanf("%d",&input)

	fmt.Printf("%s\n",judge(input))
}

func judge(in int) string {
	if in == 0 {
		return "Zero"
	} else if in > 0 {
		return "Positive"
	} else {
		return "Negative"
	}
}