package server

import "sync"

var (
	mu         sync.Mutex
	processing = make(map[string]*sync.Mutex)
)

func getOrCreateMutex(workshopID string) *sync.Mutex {
	mu.Lock()
	defer mu.Unlock()

	// Create the entry if missing
	if _, exists := processing[workshopID]; !exists {
		processing[workshopID] = &sync.Mutex{}
	}
	return processing[workshopID]
}

func clearMutex(workshopID string) {
	mu.Lock()
	delete(processing, workshopID)
	mu.Unlock()
}
