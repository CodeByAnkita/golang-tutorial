# use this command for macos first time solidity with go

https://docs.soliditylang.org/en/latest/installing-solidity.html#macos-packages

macOS Packages
We distribute the Solidity compiler through Homebrew as a build-from-source version. Pre-built bottles are currently not supported.

```bash
brew update
brew upgrade
brew tap ethereum/ethereum
brew install solidity
```

check compiler install or not for that

```bash
solc
```

# after go here and download and use this command

https://github.com/protocolbuffers/protobuf/releases

# for instal protobuf compiler

go to same folder trminal

```bash
 brew install protobuf
protoc --version
protoc

```

# for build contract with make buid folder with bin and abi file

```bash
solc --bin --abi 07-todo-smartcontract/todo.sol -o build
 cd Documents/code/Golang/Golang-Ethereum/protoc-28.3-osx-aarch_64\ \(2\)
ls
mv bin/protoc $GOBIN
 protoc
ankitakapadiya@Ankitas-MacBook-Pro go % cd pkg/mod/github.com/ethereum/go-ethereum@v1.14.12/
make devtools
 abigen --help
abigen --bin=build/Todo.bin --abi=build/Todo.abi --pkg=todo --out=gen/todo.go
```
