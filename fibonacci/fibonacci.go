package fibonacci

var cached = make(map[int]int) // depth=value

var nearFreeDepth = 25

func SetMaxDepth(n int) bool {
	nearFreeDepth = n
	return nearFreeDepth == n
}

func AtDepth(n int) int {
	if n > nearFreeDepth {
		n = nearFreeDepth
	}
	value, ok := cached[n]
	if !ok {
		cache()
		return AtDepth(n)
	}

	if ok {
		return value
	}
	return 0
}

func cache() {
	if cached == nil {
		cached = make(map[int]int)
	}
	if len(cached) == nearFreeDepth {
		return
	}
	if len(cached) > nearFreeDepth {
		for k, _ := range cached {
			if k > 33 {
				delete(cached, k)
			}
		}
	}
	for i := 0; i < 33; i++ {
		cached[i] = recursive(i)
	}
}

func recursive(n int) int {
	if n <= 1 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}
