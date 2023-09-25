package main

import (
	"fmt"
	"github.com/elcruzo/mylibrary/lrucache"
)

func main() {
	cache := lrucache.NewLRUCache(3)

	cache.Put("A", 1)
	cache.Put("B", 2)
	cache.Put("C", 3)

	fmt.Println("A:", cache.Get("A"))
	fmt.Println("B:", cache.Get("B"))
	fmt.Println("C:", cache.Get("C"))

	cache.Put("D", 4)

	fmt.Println("A:", cache.Get("A"))
}
