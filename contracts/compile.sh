#!/bin/bash
solc --abi Poll.sol -o build --overwrite
abigen --abi=./build/Poll.abi -pkg=poll --out=Poll.go
solc --bin Poll.sol -o build --overwrite
abigen --bin=./build/Poll.bin --abi=./build/Poll.abi -pkg=poll --out=Poll.go
cp ./Poll.go ../backend/poll/

solc --abi Events.sol -o build --overwrite
abigen --abi=./build/Events.abi -pkg=events --out=Events.go
solc --bin Events.sol -o build --overwrite
abigen --bin=./build/Events.bin --abi=./build/Events.abi -pkg=events --out=Events.go
# mkdir -p ../backend/events/
cp ./Events.go ../backend/poll/