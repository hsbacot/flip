package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const (
	launchAgentLabel = "com.hsbacot.flip"
	plistTemplate    = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>{{.Label}}</string>
    <key>ProgramArguments</key>
    <array>
        <string>{{.ExecutablePath}}</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>`
)

type plistData struct {
	Label          string
	ExecutablePath string
}

func getPlistPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library/LaunchAgents", launchAgentLabel+".plist"), nil
}

func isLaunchAtLoginEnabled() bool {
	path, err := getPlistPath()
	if err != nil {
		return false
	}
	_, err = os.Stat(path)
	return err == nil
}

func toggleLaunchAtLogin() error {
	if isLaunchAtLoginEnabled() {
		return disableLaunchAtLogin()
	}
	return enableLaunchAtLogin()
}

func enableLaunchAtLogin() error {
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not get executable path: %v", err)
	}

	// If we are running from a bundle, we want the bundle path if possible,
	// but os.Executable() usually gives the binary inside the bundle.
	// For LaunchAgents, the direct path to the binary is fine.

	path, err := getPlistPath()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.New("plist").Parse(plistTemplate)
	if err != nil {
		return err
	}

	data := plistData{
		Label:          launchAgentLabel,
		ExecutablePath: execPath,
	}

	return tmpl.Execute(f, data)
}

func disableLaunchAtLogin() error {
	path, err := getPlistPath()
	if err != nil {
		return err
	}
	return os.Remove(path)
}
