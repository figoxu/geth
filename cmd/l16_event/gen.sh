#!/bin/bash

function main() {
  solc --abi exchange.sol  -o ./
  solc --bin exchange.sol  -o ./
  mkdir -p ./pkg/abi/exchange/
  abigen --bin=exchange.bin --abi=exchange.abi --pkg=exchange --out=./pkg/abi/exchange/exchange.go


  solc --abi store.sol  -o ./
  solc --bin store.sol  -o ./
  mkdir -p ./pkg/abi/store/
  abigen --bin=store.bin --abi=store.abi --pkg=store --out=./pkg/abi/store/store.go


  solc --abi erc20.sol  -o ./
  solc --bin erc20.sol  -o ./
  mkdir -p ./pkg/abi/ecr20/
  abigen --bin=ERC20.bin --abi=ERC20.abi --pkg=store --out=./pkg/abi/ecr20/ecr20.go
}

main
