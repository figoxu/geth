#!/bin/bash

function main() {
  solcjs --abi erc20.sol
  abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go
}

main
