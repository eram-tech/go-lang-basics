package main

import (
	 
	"fmt"
 
)

func main() {
	

	totalInputs := 0
	ans := 0
	fmt.Scanf("%d", &totalInputs)
	//fmt.Println(totalInputs)

	newArr := make([]int, totalInputs)

	for j := 0; j < totalInputs; j++ {
		//fmt.Print()
		fmt.Scanf("%d : ", &newArr[j])

	}
	for k := 0; k < len(newArr); k++ {
		ans = 0
		for j := 1; j < newArr[k]; j++ {
			if newArr[k]%j == 0 {
				ans = ans + j
			}
		}
		fmt.Println(ans)
	}

}
