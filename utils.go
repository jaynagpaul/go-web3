package web3

// Reset can be called to reset state of web3. Resets everything except manager. Uninstalls all filters. Stops polling.
// If keepIsSyncing is true it will uninstall all filters, but will keep the web3.eth.isSyncing() polls
func (w3 *Web3) Reset(keepIsSyncing bool) {
	w3.js.Call("reset", keepIsSyncing)
}
