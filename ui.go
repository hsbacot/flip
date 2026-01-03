package main

import (
	"github.com/caseymrm/menuet"
)

// updateLabel updates the Menu Bar icon using native Template images
func updateLabel(isNatural bool) {
	icon := "standardTemplate.png"
	if isNatural {
		icon = "naturalTemplate.png"
	}
	menuet.App().SetMenuState(&menuet.MenuState{
		Image: icon,
	})
}

// toggleScroll performs an optimistic UI update and system change
func toggleScroll(natural bool) {
	if getState() == natural {
		return
	}
	updateState(natural) // Optimistic UI update
	go setScrollState(natural)
}

// menuItems returns the current menu structure
func menuItems() []menuet.MenuItem {
	isNatural := getState()
	return []menuet.MenuItem{
		{
			Text: "Scroll Direction",
			Type: menuet.Regular,
		},
		{
			Text:  "Natural Scrolling",
			State: isNatural,
			Clicked: func() {
				toggleScroll(true)
			},
		},
		{
			Text:  "Standard Scrolling (Mouse)",
			State: !isNatural,
			Clicked: func() {
				toggleScroll(false)
			},
		},
		{Type: menuet.Separator},
		{
			Text:  "Launch at Login",
			State: isLaunchAtLoginEnabled(),
			Clicked: func() {
				toggleLaunchAtLogin()
				menuet.App().MenuChanged()
			},
		},
	}
}
