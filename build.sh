#!/bin/bash

# 1. Build the binary
go build -o Flip .

# 2. Create the bundle structure
mkdir -p Flip.app/Contents/MacOS
mkdir -p Flip.app/Contents/Resources

# 3. Move files into place
mv Flip Flip.app/Contents/MacOS/
cp Info.plist Flip.app/Contents/
cp Resources/*Template*.png Flip.app/Contents/Resources/

echo "Build complete: Flip.app created."
