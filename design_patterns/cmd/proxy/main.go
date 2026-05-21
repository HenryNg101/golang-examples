package main

import "fmt"

// The common interface for both the service, and the proxy/proxies
type UserRepository interface {
	GetUser(id int) string
}

// Real DB service
type Database struct{}

func (d Database) GetUser(id int) string {
	fmt.Println("Fetching from DB")
	return "user"
}

// Proxy service, this is still the same as DB, but with extra behavior of caching
type CacheProxy struct {
	repo  UserRepository
	cache map[int]string
}

func (p *CacheProxy) GetUser(id int) string {
	if val, ok := p.cache[id]; ok {
		fmt.Println("Returning from cache")
		return val
	}

	result := p.repo.GetUser(id)
	p.cache[id] = result
	return result
}

func main() {
	db := Database{}
	proxy := &CacheProxy{
		repo:  db,
		cache: make(map[int]string),
	}

	// You only interact through the proxy object, not the actual DB inside
	proxy.GetUser(1) // DB
	proxy.GetUser(1) // cache
}
