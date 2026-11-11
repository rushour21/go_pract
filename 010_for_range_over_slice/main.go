package main

import (
	"fmt"
)

func main(){
	views := [] int{20, 40, 50, 43, 87};

	total := 0;

	// to loop slices we use range i is index and v is the value
	for i, v := range views{
		fmt.Println("day", i, "views", v);
		total = total + v;
	}
	fmt.Println(total)
}