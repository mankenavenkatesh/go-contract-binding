# go-contract-binding

- Bindings are used for Server side to interact with blockchain easily.
- Bindings for a smart contract is generated from the abi json file.
- abi json file can be generated by compiling smart contract from Remix or using solc compiler.

 For Example, 
- we can generate Go Binding for any smart contract automatically. 
- Invoking the api's in generated binding will invoke the smart contract.

## Step by Step
- Let's take sample smart contract "token contract" and generate a binding
https://gist.github.com/karalabe/08f4b780e01c8452d989
https://ethereum.org/developers/

- Token Contract implements a custom token which can be deployed on Ethereum

- Install abigen 
- Install solc compiler
- Compile Contract and Generate the abi json 
     solc -o outputs/ --abi --pretty-json --overwrite storage.sol
-  Generate Binding
    abigen --abi abi/StorageContract.abi --pkg main -type Storage --out go-bindings/storage.go
- Deploy Contract on Testnet and get the contract address (0x692a70D2e424a56D2C6C27aA97D1a86395877b3A)
- 




References - 
- https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts
- https://ethereum.org/developers/#getting-started
