package main

import (
	"sync"

	"github.com/caseymrm/menuet"
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

	updateLabel(natural)
	menuet.App().MenuChanged()
}

// getState returns the current scroll state in a thread-safe manner
func getState() bool {
	stateMutex.RLock()
	defer stateMutex.RUnlock()
	return currentIsNatural
}
