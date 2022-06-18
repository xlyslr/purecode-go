package q0146

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Printf("%#v\n", cache.Get(1))
	cache.Put(3, 3)
	fmt.Printf("%#v\n", cache.Get(2))
	cache.Put(4, 4)
	fmt.Printf("%#v\n", cache.Get(1))
	fmt.Printf("%#v\n", cache.Get(3))
	fmt.Printf("%#v\n", cache.Get(4))
}
