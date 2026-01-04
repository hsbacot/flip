package main

import (
	"sync"
)

var (
	currentIsNatural bool
	stateMutex       sync.RWMutex
)

// updateState updates the global state and triggers UI refreshes
func updateState(natural bool) {
	stateMutex.Lock()
	if currentIsNatural == natural {
		stateMutex.Unlock()
		return
	}
	currentIsNatural = natural
	stateMutex.Unlock()

	refreshUI()
}

// getState returns the current scroll state in a thread-safe manner
func getState() bool {
	stateMutex.RLock()
	defer stateMutex.RUnlock()
	return currentIsNatural
}
