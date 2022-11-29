package main

import (
	"fmt"
)

func main() {
	start, finish := 1, 5
	def := finish - start

	fmt.Println(int(def/3) + def%3)

}
