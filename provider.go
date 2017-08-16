package web3

// CurrentProvider can be used to check if mist, metamask etc. have already set a provider.
func (w3 *Web3) CurrentProvider() bool {
	return w3.js.Get("currentProvider") == nil
}

// SetProvider should be called to set provider.
// See NewWeb3 for more detailed information on args.
func (w3 *Web3) SetProvider(args ...interface{}) {
	w3.js.Call("setProvider",
		w3.js.
			Get("providers").
			Get("HttpProvider").
			New(args))

}
