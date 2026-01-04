package main

import (
	"log"
	"os"
)

func main() {
	// Handle CLI toggles
	if len(os.Args) > 1 && os.Args[1] == "--toggle" {
		current := getScrollState()
		setScrollState(!current)
		return
	}

	// 1. Initial State
	initialState := getScrollState()
	updateState(initialState)

	// 2. Start System Watcher
	go watchPreferences()

	// 3. Setup UI
	setupUI()

	// 4. Run loop
	err := wailsApp.Run()
	if err != nil {
		log.Fatal(err)
	}
}
