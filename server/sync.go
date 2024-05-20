package server

import "sync"

type workshopMutex struct {
	mu          *sync.Mutex
	activeCount uint
}

var (
	mu         sync.Mutex
	processing = make(map[string]*workshopMutex)
)

func getOrCreateMutex(workshopID string) *workshopMutex {
	mu.Lock()
	defer mu.Unlock()

	// Create the entry if missing
	if wm, exists := processing[workshopID]; exists {
		wm.activeCount++
		return wm
	}

	wm := &workshopMutex{
		mu:          &sync.Mutex{},
		activeCount: 1,
	}
	processing[workshopID] = wm
	return wm
}

func releaseMutex(workshopID string) {
	mu.Lock()
	defer mu.Unlock()

	if wm, exists := processing[workshopID]; exists {
		wm.activeCount--
		if wm.activeCount == 0 {
			delete(processing, workshopID)
		}
	}
}
