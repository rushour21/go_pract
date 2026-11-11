package main

import (
	"fmt"
)

func main(){
	// make is use for create slice with lenght and capacity 
	// make([]T, Length, Capacity)
	// res := make([]int, 0, 5);

	// fmt.Println(res, len(res), cap(res))

	// in case we are exceding capacity by appending elements,  go grows the backing array 
	// (usually duble the capacity at certain size) 

	todo := []string{"do project", "do research"};
	now := [] string{"learn go"};

	todo = append(todo, now...) // spread and add the now in the todo
	fmt.Println(todo)
}