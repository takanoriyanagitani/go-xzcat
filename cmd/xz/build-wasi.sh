#!/bin/sh

tinygo \
	build \
	-o ./xz.wasm \
	-target=wasip1 \
	-opt=z \
	-no-debug \
	./main.go
