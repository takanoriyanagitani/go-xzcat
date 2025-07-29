#!/bin/sh

tinygo \
	build \
	-o ./xzcat.wasm \
	-target=wasip1 \
	-opt=z \
	-no-debug \
	./main.go
