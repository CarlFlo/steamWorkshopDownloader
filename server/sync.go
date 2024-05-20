package server

import "sync"

type workshopMutex struct {
	wMu         *sync.Mutex
	activeCount uint
}

var (
	mu         sync.Mutex
	processing = make(map[string]*workshopMutex)
)

func getOrCreateMutex(workshopID string) *workshopMutex {
	mu.Lock()
	defer mu.Unlock()

	// The request is already being handeled so increment activeCount
	if wm, exists := processing[workshopID]; exists {
		wm.activeCount++
		return wm
	}

	// Create the entry if missing
	wm := &workshopMutex{
		wMu:         &sync.Mutex{},
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
