package config

import "sync"

type config map[string]string

var (
	once sync.Once

	instance config
)

func New() config {
	once.Do(func() {
		instance = make(config)
	})

	return instance
}
