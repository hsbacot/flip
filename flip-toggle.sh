#!/bin/bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title Flip Scroll Direction
# @raycast.mode silent

# Optional parameters:
# @raycast.icon 🔄
# @raycast.packageName Flip

# Note: This path assumes you have Flip.app in your /Applications folder.
# Adjust the path if you have it elsewhere (e.g., $HOME/dev/flip/Flip.app)
APP_PATH="/Applications/Flip.app"

if [ ! -d "$APP_PATH" ]; then
  # Fallback to local dev path for current session
  APP_PATH="./Flip.app"
fi

"$APP_PATH/Contents/MacOS/Flip" --toggle
echo "Scrolling Toggled"
