package main

import (
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var (
	wailsApp *application.App
	tray     *application.SystemTray

	iconCache = make(map[string][]byte)
)

// setupUI initializes the Wails v3 application and system tray
func setupUI() {
	wailsApp = application.New(application.Options{
		Name:        "Flip",
		Description: "Toggle Natural Scrolling",
		Mac: application.MacOptions{
			ActivationPolicy: application.ActivationPolicyAccessory,
		},
	})

	preloadIcons()

	tray = wailsApp.SystemTray.New()

	refreshUI()
}

// refreshUI updates both the icon and the menu checkmarks
func refreshUI() {
	if tray == nil {
		return
	}

	isNatural := getState()
	updateLabel(isNatural)

	menu := wailsApp.NewMenu()
	menu.Add("Scroll Direction").SetEnabled(false)

	menu.AddRadio("Natural Scrolling", isNatural).OnClick(func(ctx *application.Context) {
		toggleScroll(true)
	})

	menu.AddRadio("Standard Scrolling (Mouse)", !isNatural).OnClick(func(ctx *application.Context) {
		toggleScroll(false)
	})

	menu.AddSeparator()

	menu.AddCheckbox("Launch at Login", isLaunchAtLoginEnabled()).OnClick(func(ctx *application.Context) {
		toggleLaunchAtLogin()
		refreshUI()
	})

	menu.AddSeparator()

	menu.Add("Quit").OnClick(func(ctx *application.Context) {
		wailsApp.Quit()
	})

	tray.SetMenu(menu)
}

// updateLabel updates the Menu Bar icon from cache
func updateLabel(isNatural bool) {
	iconName := "standardTemplate.png"
	if isNatural {
		iconName = "naturalTemplate.png"
	}

	if data, ok := iconCache[iconName]; ok {
		tray.SetTemplateIcon(data)
	}
}

// preloadIcons loads icons into memory to avoid disk I/O lag
func preloadIcons() {
	execPath, _ := os.Executable()
	contentsDir := filepath.Dir(filepath.Dir(execPath)) // Contents folder

	icons := []string{"naturalTemplate.png", "standardTemplate.png"}
	for _, name := range icons {
		path := filepath.Join(contentsDir, "Resources", name)
		// Fallback for development
		if _, err := os.Stat(path); err != nil {
			path = filepath.Join("Resources", name)
		}

		data, err := os.ReadFile(path)
		if err == nil {
			iconCache[name] = data
		}
	}
}

// toggleScroll performs an optimistic UI update and system change
func toggleScroll(natural bool) {
	if getState() == natural {
		return
	}
	updateState(natural) // Optimistic UI update
	go setScrollState(natural)
}
