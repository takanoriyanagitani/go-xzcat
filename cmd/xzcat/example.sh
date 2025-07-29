#!/bin/sh

ls -l |
	wazero run ../xz/xz.wasm |
	wazero run ./xzcat.wasm
