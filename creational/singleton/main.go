package main

import (
	"fmt"

	"github.com/gaius-qi/golang-patterns/creational/singleton/config"
)

func main() {
	c := config.New()

	c["this"] = "that"

	c2 := config.New()

	fmt.Println("This is ", c2["this"])
}
