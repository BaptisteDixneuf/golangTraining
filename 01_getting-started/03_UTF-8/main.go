package main

import "fmt"

func main() {
	for i := 0; i < 128; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
