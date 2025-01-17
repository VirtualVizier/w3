# w3

[![Go Reference](https://pkg.go.dev/badge/github.com/lmittmann/w3.svg)](https://pkg.go.dev/github.com/lmittmann/w3)
[![Go Report Card](https://goreportcard.com/badge/github.com/lmittmann/w3)](https://goreportcard.com/report/github.com/lmittmann/w3)

<img src="https://user-images.githubusercontent.com/3458786/153202258-24bf253e-5ab0-4efd-a0ed-43dc1bf093c9.png" align="right" alt="W3 Gopher" width="158" height="224">

Package `w3` implements a modular and fast Ethereum JSON RPC client with
first-class ABI support.

* **Modular** API allows to create custom RPC method integrations that can be
  used alongside the methods implemented by the package.
* **Batch request** support significantly reduces the duration of requests to
  both remote and local endpoints.
* **ABI** bindings are specified for individual functions using Solidity syntax.
  No need for `abigen` and ABI JSON files.

`w3` is closely linked to [go-ethereum](https://github.com/ethereum/go-ethereum)
and uses a variety of its types, such as [`common.Address`](https://pkg.go.dev/github.com/ethereum/go-ethereum/common#Address)
or [`types.Transaction`](https://pkg.go.dev/github.com/ethereum/go-ethereum/core/types#Transaction).


## Install

```
go get github.com/lmittmann/w3@latest
```


## Getting Started

Connect to an RPC endpoint via HTTP, WebSocket, or IPC using [`Dial`](https://pkg.go.dev/github.com/lmittmann/w3#Dial)
or [`MustDial`](https://pkg.go.dev/github.com/lmittmann/w3#MustDial).

```go
// Connect (or panic on error)
client := w3.MustDial("https://cloudflare-eth.com")
defer client.Close()
```

## Batch Requests

Batch request support in the [`Client`](https://pkg.go.dev/github.com/lmittmann/w3#Client)
allows to send multiple RPC requests in a single HTTP request. The speed gains
to remote endpoints are huge. Fetching 100 blocks in a single batch request
with `w3` is ~80x faster compared to sequential requests with `ethclient`.

Example: Request the nonce and balance of an address in a single request

```go
var (
	addr = w3.A("0x000000000000000000000000000000000000c0Fe")

	nonce   uint64
	balance big.Int
)

err := client.Call(
	eth.Nonce(addr).Returns(&nonce),
	eth.Balance(addr).Returns(&balance),
)
```


## ABI Bindings

ABI bindings in `w3` are specified for individual functions using Solidity
syntax and are usable for any contract that supports that function.

Example: ABI binding for the ERC20-function `balanceOf`

```go
funcBalanceOf := w3.MustNewFunc("balanceOf(address)", "uint256")
```

A [`Func`](https://pkg.go.dev/github.com/lmittmann/w3#Func) can be used to

* encode arguments to the contracts input data ([`Func.EncodeArgs`](https://pkg.go.dev/github.com/lmittmann/w3#Func.EncodeArgs)),
* decode arguments from the contracts input data ([`Func.DecodeArgs`](https://pkg.go.dev/github.com/lmittmann/w3#Func.DecodeArgs)), and
* decode returns form the contracts output data ([`Func.DecodeReturns`](https://pkg.go.dev/github.com/lmittmann/w3#Func.DecodeReturns)).

### Reading Contracts

[`Func`](https://pkg.go.dev/github.com/lmittmann/w3#Func)'s can be used with
[`eth.CallFunc`](https://pkg.go.dev/github.com/lmittmann/w3/module/eth#CallFunc)
in the client to read contract data.

```go
var (
	weth9 = w3.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	dai   = w3.A("0x6B175474E89094C44Da98b954EedeAC495271d0F")

	weth9Balance big.Int
	daiBalance   big.Int
)

err := client.Call(
	eth.CallFunc(funcBalanceOf, weth9, addr).Returns(&weth9Balance),
	eth.CallFunc(funcBalanceOf, dai, addr).Returns(&daiBalance),
)
```


## Custom RPC Methods

Custom RPC methods can be called with the `w3` client by creating a
[`core.Caller`](https://pkg.go.dev/github.com/lmittmann/w3/core#Caller)
implementation.
The `w3/module/eth` package can be used as implementation reference.


## RPC Methods

List of supported RPC methods.

### `eth`

Method                      | Go Code
:---------------------------|:--------
`eth_blockNumber`           | `eth.BlockNumber().Returns(blockNumber *big.Int)`
`eth_call`                  | `eth.Call(msg ethereum.CallMsg).Returns(output *[]byte)`<br>`eth.CallFunc(fn core.Func, contract common.Address, args ...interface{}).Returns(returns ...interface{})`
`eth_chainId`               | `eth.ChainID().Returns(chainID *uint64)`
`eth_gasPrice`              | `eth.GasPrice().Returns(gasPrice *big.Int)`
`eth_getBalance`            | `eth.Balance(addr common.Address).Returns(balance *big.Int)`
`eth_getBlockByHash`        | `eth.BlockByHash(hash common.Hash).Returns(block *types.Block)`<br>`eth.BlockByHash(hash common.Hash).ReturnsRAW(block *eth.RPCBlock)` <br>`eth.HeaderByHash(hash common.Hash).Returns(header *types.Header)`<br>`eth.HeaderByHash(hash common.Hash).ReturnsRAW(header *eth.RPCHeader)`
`eth_getBlockByNumber`      | `eth.BlockByNumber(number *big.Int).Returns(block *types.Block)`<br>`eth.BlockByNumber(number *big.Int).ReturnsRAW(block *eth.RPCBlock)`<br>`eth.HeaderByNumber(number *big.Int).Returns(header *types.Header)`<br>`eth.HeaderByNumber(number *big.Int).ReturnsRAW(header *eth.RAWHeader)`
`eth_getCode`               | `eth.Code(addr common.Address).Returns(code *[]byte)`
`eth_getLogs`               | `eth.Logs(q ethereum.FilterQuery).Returns(logs *[]types.Log)`
`eth_getStorageAt`          | `eth.StorageAt(addr common.Address, slot common.Hash).Returns(storage *common.Hash)`
`eth_getTransactionByHash`  | `eth.TransactionByHash(hash common.Hash).Returns(tx *types.Transaction)`<br>`eth.TransactionByHash(hash common.Hash).ReturnsRAW(tx *eth.RPCTransaction)`
`eth_getTransactionCount`   | `eth.Nonce(addr common.Address).Returns(nonce *uint64)`
`eth_getTransactionReceipt` | `eth.TransactionReceipt(hash common.Hash).Returns(receipt *types.Receipt)`<br>`eth.TransactionReceipt(hash common.Hash).ReturnsRAW(receipt *eth.RPCReceipt)`
`eth_sendRawTransaction`    | `eth.SendTransaction(tx *types.Transaction).Returns(hash *common.Hash)`<br>`eth.SendRawTransaction(rawTx []byte).Returns(hash *common.Hash)`


### Third Party RPC Method Packages

Package                                                                  | Description
:------------------------------------------------------------------------|:------------
[github.com/lmittmann/flashbots](https://github.com/lmittmann/flashbots) | Package `flashbots` implements RPC API bindings for the Flashbots relay and mev-geth.
