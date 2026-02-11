package holidays

import "sync"

// registry holds all registered country holiday providers.
var registry = map[string]provider{}

// registryMutex protects concurrent access to the registry map.
var registryMutex sync.RWMutex
