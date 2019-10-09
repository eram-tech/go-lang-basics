//WordCount
package main

import (
	//"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	retMap := make(map[string]int)
	
	for _,word := range strings.Fields(s) {
	retMap[word]++
	fmt.Println("word : ", word)
	fmt.Println(retMap[word])
	}
	return retMap //map[string]int{"x": 1}
}

func main() {
	//wc.Test(WordCount)
	WordCount("I am learning Go! I")
}
