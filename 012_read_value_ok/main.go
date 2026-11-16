package main

import (
	"fmt"
)

func main(){
	
	// points := map[string] int{
	// 	"a" : 10,
	// 	"b" : 0,
	// }

	// fmt.Println("a", points["a"])
	// fmt.Println("b", points["b"])
	// fmt.Println("c", points["c"])
	//this will Println
	// 	a 10
	// b 0
	// c 0 even it is nor present in the mao 

	// so we can use val and oka 
	// comma-ok idiom in Go. It is commonly used when working with maps.
	// value, isExists := mapName["key"]

	// valB, okB :=  points["a"];
	// fmt.Println(valB) // 100
	// fmt.Println(okB)  // true

	// val, isExist := points["c"]
	// fmt.Println(val) // 0
	// fmt.Println(isExist)  // false

	// // how can we use it
	// if val, exist := points["b"]; exist{
	// 	fmt.Println("it is present", val)
	// }else{
	// 	fmt.Println(" is not present")
	// }

	prices := map[string] int{
		"a":100,
		"b":200,
		"c":300,
	}

	total := 0;

	// for item, price := range prices{
	// 	//  fmt.Println(item, price)
	// 	total = total + price
	// }
	// fmt.Println(total)

	// if you dont to use itme then 
	// go will give error declared and not used: item
	// you can use _ (underscore) when you don't need a variable.
	for _, price := range prices{
		total= total+ price
	}
	fmt.Println(total)
	
}