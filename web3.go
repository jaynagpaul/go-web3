package web3

import (
	"github.com/gopherjs/gopherjs/js"
)

// NewWeb3 creates a Web3 instance with args passed to Web3.providers.HttpProvider.
// HttpProvider takes 4 arguments (host, timeout, user, password)
func NewWeb3(args ...interface{}) *Web3 {
	w3 := js.Global.Get("Web3")

	provider := w3.
		Get("providers").
		Get("HttpProvider").
		New(args)

	if w := js.Global.Get("web3"); w != nil {
		provider = w.Get("currentProvider")
	}

	w3 = w3.New(provider)

	js.Global.Set("web3", w3)
	version := w3.Get("version")
	return &Web3{
		js: w3,
		Version: Version{
			API: version.Get("api").String(),
			js:  w3,
		},
	}
}

// Web3 instance
type Web3 struct {
	js      *js.Object
	Version Version
}

// IsConnected should be called to check if a connection to a node exists
func (w3 *Web3) IsConnected() bool {
	return w3.js.Call("isConnected").Bool()
}
