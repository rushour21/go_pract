package main 

import (
	"fmt"
)

func main(){
	var marks [3] int

	marks[0] = 99;
	marks[1] = 98;
	marks[2] = 97;

	fmt.Println(marks)

	// array literal
	rollNo := [5] int {1,2,3,4,5};

	fmt.Println(len(rollNo)); // len array method like in in js array.lenght
}