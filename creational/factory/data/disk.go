package data

import "fmt"

type diskStorage struct {
}

func newDiskStorage() *diskStorage {
	return &diskStorage{}
}

func (m *diskStorage) Open() {
	fmt.Println("DiskStorage open")
}
