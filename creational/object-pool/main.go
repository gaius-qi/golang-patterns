package main

import (
	"github.com/gaius-qi/golang-patterns/creational/object-pool/pool"
)

func main() {
	p := pool.New(2)

	for {
		select {
		case obj := <-*p:
			obj.Do()

			*p <- obj
		default:
			// No more objects left — retry later or fail
			return
		}
	}
}
