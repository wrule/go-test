package main

import "fmt"

func main() {
	var list []int
	list = append(list, 2)
	fmt.Println(list[0], len(list), cap(list))
}
