# Flip

**Flip** is a tiny, native macOS menu bar utility that allows you to instantly toggle between "Natural" and "Standard" scrolling. 

It's perfect for users who frequently switch between a Trackpad (where Natural scrolling feels best) and a Mouse (where Standard scrolling is often preferred).

## Features

- **Instant Toggle**: Switch scroll direction directly from the menu bar with zero lag.
- **Visual Indicators**: Dynamic "Template" icons change to reflect the current mode.
- **Raycast Integration**: Control Flip via keyboard shortcuts using the built-in `--toggle` CLI flag.
- **Optimized Performance**: High-speed icon caching ensures an instantaneous UI response.
- **Launch at Login**: Easily toggle the background service to start when you log in.

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
2. Look for the toggle icon in your menu bar.
3. Click the icon to switch modes or manage settings.

### Raycast Integration

You can toggle Flip using Raycast (or any other macro tool) by calling the binary with the `--toggle` flag:

```bash
/Applications/Flip.app/Contents/MacOS/Flip --toggle
```

## Technical Details

- **Language**: Go (Golang)
- **Framework**: [Wails v3](https://v3alpha.wails.io/)
- **Settings**: Modifies `com.apple.swipescrolldirection` in `NSGlobalDomain`.
- **UI**: Native macOS Menu Bar integration with Template Icon support.
