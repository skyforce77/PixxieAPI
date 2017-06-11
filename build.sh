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

go build -buildmode=plugin -o build/bindings/terminal.so TerminalBinding/terminal.go
go build -buildmode=plugin -o build/plugins/binaryclock.so BinaryClockPlugin/binaryclock.go
go build -buildmode=plugin -o build/plugins/youtube.so YoutubePlugin/youtube.go
go build -o build/pixxie Main/main.go
