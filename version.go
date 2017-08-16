package web3

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

// Version of the various protocols and apis
type Version struct {
	API string // The ethereum JS api version.
	js  *js.Object
}

// GetNode takes a callback with node version and error as a params
// Returns the client/node version.
func (v *Version) GetNode(callback func(string, error)) {
	v.js.Get("version").Call("getNode", func(err, node *js.Object) {
		if err != nil {
			callback(node.String(), errors.New(err.String()))
		} else {
			callback(node.String(), nil)
		}
	})
}

// GetNetwork takes a callback with network version and error as a params
// Returns the network protocol version.
func (v *Version) GetNetwork(callback func(string, error)) {
	v.js.Get("version").Call("getNode", func(err, network *js.Object) {
		if err != nil {
			callback(network.String(), errors.New(err.String()))
		} else {
			callback(network.String(), nil)
		}
	})
}

// GetEthereum takes a callback with ethereum version and error as a params
// Returns the ethereum protocol version.
func (v *Version) GetEthereum(callback func(string, error)) {
	v.js.Get("version").Call("getEthereum", func(err, ethereum *js.Object) {
		if err != nil {
			callback(ethereum.String(), errors.New(err.String()))
		} else {
			callback(ethereum.String(), nil)
		}
	})
}

// GetWhisper takes a callback with whisper version and error as a params
// Returns the whisper protocol version.
func (v *Version) GetWhisper(callback func(string, error)) {
	v.js.Get("version").Call("getWhisper", func(err, whisper *js.Object) {
		if err != nil {
			callback(whisper.String(), errors.New(err.String()))
		} else {
			callback(whisper.String(), nil)
		}
	})
}
