# Flip 💻🖱️

**Flip** is a tiny, native macOS menu bar utility that allows you to instantly toggle between "Natural" and "Standard" scrolling. 

It's perfect for users who frequently switch between a Trackpad (where Natural scrolling feels best) and a Mouse (where Standard scrolling is often preferred).

## Features

- **One-Click Toggle**: Switch scroll direction directly from the menu bar.
- **Visual Indicators**: The menu bar icon changes to reflect the current mode (💻 for Natural, 🖱️ for Standard).
- **Native Experience**: Built with Go and the `menuet` library for a lightweight, native feel.
- **Fast & Reliable**: Updates system settings and refreshes preferences immediately.

## Installation

Currently, Flip is provided as a source project. You can build it yourself using the instructions below.

## Building from Source

### Prerequisites

- A Mac running macOS.
- [Go](https://go.dev/doc/install) installed.

### Build Steps

1. Clone the repository or download the source code.
2. Open a terminal in the project directory.
3. Run the build script:
   ```bash
   ./build.sh
   ```
4. This will create a `Flip.app` in the project directory.
5. Move `Flip.app` to your `/Applications` folder.

## Usage

1. Launch **Flip**.
2. Look for the 💻 or 🖱️ icon in your menu bar.
3. Click the icon to see the current state or switch to the other mode.

## Launch on Login

To have Flip start automatically when you log in to your Mac, you can use one of the following methods:

### Method 1: System Settings (easiest)
1. Open **System Settings**.
2. Go to **General** > **Login Items**.
3. Under **Open at Login**, click the **+** button.
4. Navigate to your `/Applications` folder and select **Flip.app**.

### Method 2: Launch Agent (command line)
If you want to use the provided `com.hsbacot.flip.plist`:
1. Copy the plist to your LaunchAgents directory:
   ```bash
   cp com.hsbacot.flip.plist ~/Library/LaunchAgents/
   ```
2. Load the agent:
   ```bash
   launchctl load ~/Library/LaunchAgents/com.hsbacot.flip.plist
   ```
   *Note: Ensure the path in the plist (`../Flip.app/Contents/MacOS/Flip`) matches your actual installation path.*

## Technical Details

- **Language**: Go (Golang)
- **Library**: [github.com/caseymrm/menuet](https://github.com/caseymrm/menuet)
- **Settings**: Modifies `com.apple.swipescrolldirection` in `NSGlobalDomain`.

---

[nsFlip.com](https://nsFlip.com)
