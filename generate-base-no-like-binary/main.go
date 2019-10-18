package main

import "fmt"

func main() {
	for i:=0; i<5; i++{
		var input int64
		fmt.Scanf("%d", &input) // Reading input from STDIN
		fmt.Println(Calculate(input, 9))
	}
	
}

func Calculate(input, base int64) int64 {
	if input < base {
		return input

	}
	return input%base + 10*Calculate(input/base, base)
}
