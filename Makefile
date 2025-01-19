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

cast-get-balance-owner:
	cast --to-dec `cast call ${SIMPLE_COIN} 'getBalance(address)' ${OWNER} --rpc-url ${CALIBRATIONNET_RPC_URL}`

cast-get-balance-test1:
	cast --to-dec `cast call ${SIMPLE_COIN} 'getBalance(address)' ${TEST1} --rpc-url ${CALIBRATIONNET_RPC_URL}`

cast-send-coin:
	cast send ${SIMPLE_COIN} 'sendCoin(address, uint)' ${TEST1} 1 --rpc-url ${CALIBRATIONNET_RPC_URL} --private-key ${PRIVATE_KEY}

cast-send-coin-2:
	cast send ${SIMPLE_COIN} 'sendCoin(address, uint)' ${TEST1} 1 --rpc-url ${CALIBRATIONNET_RPC_URL} --private-key ${PRIVATE_KEY} --priority-gas-price 200000 --gas-price 200000

cast-send-coin-async:
	cast send ${SIMPLE_COIN} 'sendCoin(address, uint)' ${TEST1} 1 --rpc-url ${CALIBRATIONNET_RPC_URL} --private-key ${PRIVATE_KEY} --async

cast-send-coin-async-replace:	NONCE=$(shell cast nonce ${OWNER} --rpc-url ${CALIBRATIONNET_RPC_URL})
cast-send-coin-async-replace:
	@echo Nonce: ${NONCE}
	cast send ${SIMPLE_COIN} 'sendCoin(address, uint)' ${TEST1} 1 --rpc-url ${CALIBRATIONNET_RPC_URL} --private-key ${PRIVATE_KEY} --async --nonce ${NONCE}
	sleep 10
	cast send ${SIMPLE_COIN} 'sendCoin(address, uint)' ${TEST1} 1 --rpc-url ${CALIBRATIONNET_RPC_URL} --private-key ${PRIVATE_KEY} --priority-gas-price 200000 --gas-price 200000 --async --nonce ${NONCE}

nonce:
	cast nonce ${OWNER} --rpc-url ${CALIBRATIONNET_RPC_URL}

cast-send-coin-3:
	./test-wallet-app send-coin owner test1 1

cast-send-coin-3-replace:	NONCE=$(shell cast nonce ${OWNER} --rpc-url ${CALIBRATIONNET_RPC_URL})
cast-send-coin-3-replace:
	@echo Nonce: ${NONCE}
	./test-wallet-app send-coin owner test1 1
	sleep 10
	./test-wallet-app send-coin owner test1 1 --gas-multiply 1.5 --nonce ${NONCE}


.PHONY: abigen
