/*
Package web3 provides a wrapper for ethereum's web3.js for use with gopherjs.

Getting Started:

$ go get -u github.com/jaynagpaul/go-web3
$ go get -u github.com/gopherjs/gopherjs

	import "github.com/jaynagpaul/go-web3"

	func main() {
		// Check if injected by browser
		w3 := web3.NewWeb3()
		w3.Version.GetEthereum(func(ver, err) {
			if err != nil {
				println(err)
			} else {
				println(ver) // Prints the version to the developer console.
			}
		})
		// TODO
	}

$ gopherjs build
*/
package web3
