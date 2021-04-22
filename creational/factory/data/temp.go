package data

import "fmt"

type tempStorage struct {
}

func newTempStorage() *tempStorage {
	return &tempStorage{}
}

func (m *tempStorage) Open() {
	fmt.Println("TempStorage open")
}
