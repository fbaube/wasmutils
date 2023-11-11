// +build js,wasm

package wasmutils

import (
	"syscall/js"
)

func init() {
     initPkgVars()
}

// initPkgVars sets up some basic global (package) variables
// related to JS. Don't stress if a "<body>" tag is not found.
// 
// This can be called explicitly from other init() func's,
// in the event that it turns out that
//  - other packages' init() funcs need this initialization, and
//  - the system has not been doing the various init()s in the
//    "right order"
//  - note that it should be idempotent
// 
func initPkgVars() {
     	G = js.Global() 
	Doc = G.Get("document")
	if !(G.Truthy() && Doc.Truthy())  {
		// return errors.New("Unable to get document object")
		panic("CAN'T GET JS Global and/or webpage DOM")
	}
	DocBody = Doc.Get("body")
	if !DocBody.Truthy() {
		// panic("CAN'T GET webpage <body>")
		println("Document has no <body>")
	}
}

