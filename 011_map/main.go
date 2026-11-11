package main

import (
	"fmt"
)

func main(){
	//map[keyType]value; keyvalue datatype

	ages := map[string]int{
		"rushabh" : 24,
		"rush" : 25,

	}

	fmt.Println(ages["rush"], len(ages));
	delete(ages, "rush")
	fmt.Println(ages) // map[rushabh:24]

	// you can also use make to create map 
	var scores map[string]int // nil map  also declaring a map doesn't automatically initialize it.
	
	scores = make(map[string]int)

	scores["math"]=99;
	fmt.Println(scores);

	


}