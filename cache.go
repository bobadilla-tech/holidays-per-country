package holidays

import (
	"fmt"
	"sync"
)

func cacheKey(countryCode string, year int) string {
	return fmt.Sprintf("%s:%d", countryCode, year)
}

var (
	holidaysCache = map[string][]Holiday{}
	cacheMutex    sync.RWMutex
)
