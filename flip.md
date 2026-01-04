# **PRD: Flip (Native macOS)**

## **1. Overview**

**Product Name:** Flip  
**Platform:** macOS (Exclusive)  
**Framework:** [Wails v3](https://v3alpha.wails.io/) (Go + Native macOS APIs)  
**Objective:** A minimalist macOS menu bar utility to instantly toggle the system-wide "Natural Scrolling" setting.

## **2. Architecture**

### **2.1 Technical Stack**
- **Language:** Go (Golang)
- **Framework:** Wails v3 (utilizing `SystemTrayManager` and `MenuManager`)
- **System Interaction:** macOS `defaults` command via the `NSGlobalDomain`.

### **2.2 Core Components**
- **System Watcher:** Uses `fsnotify` to monitor `~/Library/Preferences/.GlobalPreferences.plist`, ensuring the UI stays in sync if settings are changed elsewhere.
- **Icon Cache:** An in-memory cache of icon bytes to eliminate disk I/O lag during state changes.
- **CLI Interface:** Supports a `--toggle` flag for external integration (e.g., Raycast).

## **3. Functional Requirements**

### **3.1 Menu Bar UI**
- **Dynamic Icons:** The menu bar icon dynamically swaps between `naturalTemplate.png` (Trackpad/Natural) and `standardTemplate.png` (Mouse/Standard) assets.
- **Menu Items:**
    - **Natural Scrolling:** (Radio Item) Toggles state to 1.
    - **Standard Scrolling:** (Radio Item) Toggles state to 0.
    - **Launch at Login:** (Checkbox) Manages a LaunchAgent plist.
    - **Quit:** Standard app termination.

### **3.2 External Integration**
- **CLI Toggle:** The binary accepts a `--toggle` argument to switch the current state and exit immediately. Useful for Raycast or automation.

## **4. Performance & UX**
- **Optimistic Updates:** The UI reflects the state change immediately path before the system command completes.
- **Zero-Latency Icons:** All graphical assets are preloaded into memory on startup.
- **Native Feel:** Uses standard macOS menu behaviors and radio button patterns.

## **5. Build & Distribution**
- Packed as a standard macOS `.app` bundle.
- Requires `MACOSX_DEPLOYMENT_TARGET=15.0` (or applicable version) for modern linker compatibility.
