package main

import (
	"fmt"
)

func main(){
	var city string;
	city = "pune";

	var name = "rushabh"  //inferred to string

	// :=

	likes := 5000;

	comments, share := 100, 4000

	fmt.Println(city, name, "has", likes, "likes with", comments, "and", share, "comments and share" )
}