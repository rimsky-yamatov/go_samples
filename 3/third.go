package main

import (
	"fmt"
)

func main() {
	a,b := 0,0
	fmt.Print("Input number : ")
	fmt.Scanf("%d %d",&a,&b)

	fmt.Printf("Add : %d\n",a+b)
	fmt.Printf("Sub : %d\n",a-b)
	fmt.Printf("Mul : %d\n",a*b)
	fmt.Printf("Div : %d\n",a/b)
}