# go-web3
[![GoDoc](https://godoc.org/github.com/jaynagpaul/go-web3?status.svg)](https://godoc.org/github.com/jaynagpaul/go-web3)

A wrapper for [web3](https://github.com/ethereum/web3.js/) to be used with [GopherJS](https://github.com/gopherjs/gopherjs).
Tries to enforce best practices. Most methods are called asynchronously to maintain support for metamask

## Documentation
[GoDoc](https://godoc.org/github.com/jaynagpaul/go-web3)

## Getting Started

`$ go get -u github.com/jaynagpaul/go-web3`
`$ go get -u github.com/gopherjs/gopherjs`

```Go

import "github.com/jaynagpaul/go-web3"

func main() {
    // Check if injected by browser
    w3 := web3.NewWeb3()
    w3.Version.GetVersion(func(ver, err) {
        if err != nil {
            println(err)
        } else {
            println(ver) // Prints the version to the developer console.
        }
    })
    // TODO
}
```

`gopherjs build`

[More Examples](https://godoc.org/github.com/jaynagpaul/go-web3#pkg-examples)
[One More](https://github.com/jaynagpaul/go-web3/tree/master/example)


## TODO
* Framework for testing.
* Finalize API.