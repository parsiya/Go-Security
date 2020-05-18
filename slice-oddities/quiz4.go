package main

import "fmt"

func printSlice(s string, a []int) {
	fmt.Printf("%p - %v\tlen:%d\tcap:%d\t%s\n", a, a, len(a), cap(a), s)
}

func surprise(a []int) {
	printSlice("Inside surprise, before append", a)
	a = append(a, 5)
	printSlice("Inside surprise, after append", a)
	printSlice("Inside surprise, before assignment", a)
	for i := range a {
		a[i] = 5
	}
	printSlice("Inside surprise, after assignment", a)
}

// Quiz #4
func main() {
	a := []int{1, 2, 3, 4}
	printSlice("Inside main, before append", a)
	a = append(a, 5)
	printSlice("Inside main, after append", a)
	printSlice("Inside main, before surprise", a)
	surprise(a)
	printSlice("Inside main, after surprise", a)
}
