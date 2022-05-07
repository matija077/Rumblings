package main

func main() {
	var cache = createLRU_Cache(10)
	cache.set("2", 2)
	cache.set("3", 3)

	//fmt.Println(cache.get("2"))
	Iterate(&cache)
}
