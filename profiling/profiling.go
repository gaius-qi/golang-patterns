package main

import (
	"fmt"
	"time"

	"github.com/gaius-qi/golang-patterns/profiling/profile"
)

func main() {
	// Arguments to a defer statement is immediately evaluated and stored.
	// The deferred function receives the pre-evaluated values when its invoked.
	defer profile.Duration(time.Now(), "IntFactorial")

	fmt.Println("working...")
}
