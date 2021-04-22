package main

import "github.com/gaius-qi/golang-patterns/creational/factory/data"

func main() {
	d := data.NewStore(data.DiskStorage)
	d.Open()

	t := data.NewStore(data.TempStorage)
	t.Open()

	m := data.NewStore(data.MemoryStorage)
	m.Open()
}
