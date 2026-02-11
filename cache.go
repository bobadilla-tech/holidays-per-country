package holidays

import (
	"fmt"
	"sync"

	"github.com/bobadilla-tech/holidays-per-country/common"
)

func cacheKey(countryCode string, year int) string {
	return fmt.Sprintf("%s:%d", countryCode, year)
}

var (
	holidaysCache = map[string][]common.Holiday{}
	cacheMutex    sync.RWMutex
)
