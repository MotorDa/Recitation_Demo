package main

import "fmt"

func PrintAllNumbers(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func PrintReversNumber(n int) {
	for i := n; i > -1; i-- {
		fmt.Println(i)
	}
}
