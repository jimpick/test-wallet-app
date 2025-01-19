include .env

deploy:
	forge create --rpc-url ${CALIBRATIONNET_RPC_URL} \
		--private-key ${PRIVATE_KEY} \
		--contracts src/SimpleCoin.sol SimpleCoin \
		--broadcast

solc-output:
	cd out; \
		cat SimpleCoin.sol/SimpleCoin.json | jq .abi > SimpleCoin.abi

abigen:
	cd scripts; ./abigen.sh

cast-balance:
	cast --from-wei `cast call ${SIMPLE_COIN} 'getBalance(address)' ${OWNER} --rpc-url ${CALIBRATIONNET_RPC_URL}`

cast-send-coin:
	cast send ${SIMPLE_COIN} 'sendCoin(address, uint)' ${TEST1} 10 --rpc-url ${CALIBRATIONNET_RPC_URL} --private-key ${PRIVATE_KEY}

cast-send-coin-2:
	cast send ${SIMPLE_COIN} 'sendCoin(address, uint)' ${TEST1} 10 --rpc-url ${CALIBRATIONNET_RPC_URL} --private-key ${PRIVATE_KEY} --priority-gas-price 200000 --gas-price 200000

.PHONY: abigen
