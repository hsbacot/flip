package main

import (
	"github.com/caseymrm/menuet"
)

func main() {
	// 1. Initial State
	menuet.App().HideStartup()
	initialState := getScrollState()
	updateState(initialState)

	// 2. Start System Watcher
	go watchPreferences()

	// 3. Configure Application
	menuet.App().Label = "Flip"
	menuet.App().Children = menuItems

	// 4. Run loop
	menuet.App().RunApplication()
}
