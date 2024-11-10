#!/bin/bash

function main() {
  solc --abi Store.sol  -o ./

  solc --bin Store.sol  -o ./
mkdir -p ./store/
  abigen --bin=Store.bin --abi=Store.abi --pkg=store --out=./store/Store.go
}

main
