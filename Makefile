contracts:
	solc --abi contracts/SankoPredicts.sol --base-path . --include-path node_modules -o build
	abigen --abi=./build/SankoPredicts.abi --pkg=sanko_predicts --out=pkg/sankopredicts/sanko_predicts.go
