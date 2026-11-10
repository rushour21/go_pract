package main

import (
	"fmt"
)

func main(){
	const num int = 15;

	if num <= 12{
		fmt.Println("its less that 12 ", num)
	} else{
		fmt.Println("it is more than 12", num)
	}
}