package main

import "fmt"

func main() {
	fmt.Println(pc)
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}