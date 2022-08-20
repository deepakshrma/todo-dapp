mkdir -p golang/contracts
solcjs --abi --include-path node_modules/ --base-path contracts contracts/Todo.sol  -o build
solcjs --bin --include-path node_modules/ --base-path contracts contracts/Todo.sol  -o build
abigen --bin=./build/Todo_sol_Todo.bin  --abi=./build/Todo_sol_Todo.abi --pkg=todo --out=golang/contracts/Todo.go


