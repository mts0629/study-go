package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 42

	fmt.Println("map:", m, "len:", len(m))
	// fmt.Println("len:", len(m))

	fmt.Println("key1:", m["key1"])
	fmt.Println("key2:", m["key2"])
	fmt.Println("key3:", m["key3"])

	delete(m, "key2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	value, prs := m["key2"] // 2nd retval: key was present in a map
	fmt.Println("value:", value, "prs:", prs)

	n := map[string]int{"height": 175, "weight": 70}
	fmt.Println("map:", n)

	n2 := map[string]int{"height": 175, "weight": 70}
	fmt.Println("map:", n)

	if maps.Equal(n, n2) {
		fmt.Println("n equals n2")
	}
}
