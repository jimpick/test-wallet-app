include .env

deploy:
	forge create --rpc-url https://api.calibration.node.glif.io/rpc/v1 \
		--private-key ${PRIVATE_KEY} \
		--contracts src/SimpleCoin.sol SimpleCoin \
		--broadcast

solc-output:
	cd out; \
		cat SimpleCoin.sol/SimpleCoin.json | jq .abi > SimpleCoin.abi

abigen:
	cd scripts; ./abigen.sh

.PHONY: abigen