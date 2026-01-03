package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	domain = "NSGlobalDomain"
	key    = "com.apple.swipescrolldirection"
)

// getScrollState reads the macOS user default: true for Natural, false for Standard
func getScrollState() bool {
	cmd := exec.Command("defaults", "read", domain, key)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Error reading defaults: %v", err)
		return false
	}

	val := strings.TrimSpace(string(out))
	return val == "1"
}

// setScrollState writes the new state and attempts to refresh the system
func setScrollState(natural bool) {
	val := "false"
	if natural {
		val = "true"
	}

	cmd := exec.Command("defaults", "write", domain, key, "-bool", val)
	if err := cmd.Run(); err != nil {
		log.Printf("Error writing defaults: %v", err)
	}

	// Notify system services
	activateCmd := "/System/Library/PrivateFrameworks/SystemAdministration.framework/Resources/activateSettings"
	if _, err := os.Stat(activateCmd); err == nil {
		exec.Command(activateCmd, "-u").Run()
	} else {
		exec.Command("killall", "cfprefsd").Run()
	}
}

// watchPreferences monitors the macOS Global Preferences PLIST for changes
func watchPreferences() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("Error creating watcher: %v", err)
		return
	}
	defer watcher.Close()

	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Error getting home directory: %v", err)
		return
	}

	plistPath := filepath.Join(home, "Library/Preferences/.GlobalPreferences.plist")
	parentDir := filepath.Dir(plistPath)
	err = watcher.Add(parentDir)
	if err != nil {
		log.Printf("Error adding watcher: %v", err)
		return
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if filepath.Base(event.Name) == ".GlobalPreferences.plist" {
				go burstPoll()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}

// burstPoll rapidly polls for state changes to catch lazy writes
func burstPoll() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	timeout := time.After(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			newState := getScrollState()
			if newState != getState() {
				updateState(newState)
				return
			}
		case <-timeout:
			return
		}
	}
}
