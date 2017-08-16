package main

import (
	"github.com/gopherjs/gopherjs/js"
	web3 "github.com/jaynagpaul/go-web3"
)

// gopherjs serve
// Navigate to localhost:8080/jaynagpaul/go-web3/example
func main() {
	web3.NewWeb3().Version.GetNetwork(func(net string, err error) {
		if err != nil {
			println("Error:", err) // Log to console
		}
		println("Network Version:", net) // i.e Metamask/v3.9.7
		js.Global.Call("alert", net)     // Alert net version
	})
}
