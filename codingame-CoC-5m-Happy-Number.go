package main

import "fmt"
import "strconv"

func main() {

	orig := "86"
	n := orig
	sum := 0

	for true {
		
		sum = 0
		
		for i := 0; i < len(n); i++ {
			d, _ := strconv.Atoi(n[i:i+1])
			sum += d*d
		}
		
		if sum < 10 {
			break
		}
		
		n = strconv.Itoa(sum)
	}

	
	if sum == 1 {
		fmt.Println(orig,"IS HAPPY")
	} else {
		fmt.Println(orig,"IS UNHAPPY")
	}
}
