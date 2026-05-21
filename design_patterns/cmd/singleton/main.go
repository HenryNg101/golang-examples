package main

import (
	"fmt"
	"sync"
)

type Config struct{}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		fmt.Println("Config is initialized")
	})
	return instance
}

func main() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		GetConfig()
	}
}
