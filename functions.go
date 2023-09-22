package main

import "fmt"

func PrintAllNumbers(n int) {
	for i := range n {
		fmt.Println(i)
	}
}

func PrintReversNumber(n int) {
	for i := n; i > -1; i-- {
		fmt.Println(i)
	}
}
