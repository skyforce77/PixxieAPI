#!/bin/bash
if [ ! -d "build" ]; then
        mkdir build
fi
if [ ! -d "build/bindings" ]; then
	mkdir build/bindings
fi
if [ ! -d "build/plugins" ]; then
        mkdir build/plugins
fi

echo "Copying config"
cp config.json build/config.json
echo "Building terminal bindings..."
go build -buildmode=plugin -o build/bindings/terminal.so TerminalBinding/terminal.go
echo "Building binary clock plugin..."
go build -buildmode=plugin -o build/plugins/binaryclock.so BinaryClockPlugin/binaryclock.go
echo "Building youtube plugin..."
go build -buildmode=plugin -o build/plugins/youtube.so YoutubePlugin/youtube.go
echo "Building twitter plugin..."
go build -buildmode=plugin -o build/plugins/twitter.so TwitterPlugin/twitter.go
echo "Building core..."
go build -o build/pixxie Main/main.go
