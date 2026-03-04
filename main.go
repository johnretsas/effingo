package main

import (
	"effingo/effingo"
	"log"
)

func main() {

	mapResult := effingo.Map([]int{1, 2, 3, 4, 5, 6}, func(x int) int { return 2 * x })

	log.Println("Map result is: ", mapResult)
}
