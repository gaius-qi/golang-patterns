package data

import "fmt"

type memoryStorage struct {
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{}
}

func (m *memoryStorage) Open() {
	fmt.Println("MemoryStorage open")
}
