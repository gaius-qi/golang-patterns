package main

import (
	"fmt"

	"github.com/gaius-qi/golang-patterns/behavioral/strategy/operator"
)

func main() {
	add := operator.Operation{operator.Addition{}}
	fmt.Printf("Addition operator: %v\n", add.Operate(3, 5))

	mult := operator.Operation{operator.Multiplication{}}
	fmt.Printf("Multiplication operator: %v\n", mult.Operate(3, 5))
}
