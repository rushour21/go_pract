package main 

import (
	"fmt"
)

func main(){
	// common collection
	// dynamic and can grow
	// []type {....}

	sub := [3] string{"math", "sci", "phy"}

	fmt.Println(sub);
	sub[0] = "chemi";
	fmt.Println(sub)


	var nums [] int;
	 nums = append(nums, 10);
	 nums = append(nums, 20, 30) // works like push method in js

	fmt.Println(nums)
}
