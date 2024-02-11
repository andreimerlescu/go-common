package fibonacci

type CachedResults map[int]int

var cached = make(CachedResults) // depth=value

var freeCostDepth = 25
var expensiveCostDepthLimit = 33

func SetMaxDepth(n int) bool {
	freeCostDepth = n
	return freeCostDepth == n
}

func GetMaxDepth() int {
	return freeCostDepth
}

func Cached() CachedResults {
	return cached
}

func SetCap(n int) bool {
	expensiveCostDepthLimit = n
	return expensiveCostDepthLimit == n
}

func GetCap() int {
	return expensiveCostDepthLimit
}

func AtDepth(n int) int {
	if n > freeCostDepth {
		n = freeCostDepth
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
		cached = make(CachedResults)
	}
	if len(cached) == freeCostDepth {
		return
	}
	if len(cached) > freeCostDepth {
		for k, _ := range cached {
			if k > expensiveCostDepthLimit {
				delete(cached, k)
			}
		}
	}
	for i := 0; i < expensiveCostDepthLimit; i++ {
		cached[i] = recursive(i)
	}
}

func recursive(n int) int {
	if n <= 1 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}
