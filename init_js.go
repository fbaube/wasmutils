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
// and should be idempotent. There are two basic use cases:
//  - If the document (i.e. web page) is modified such that
//    the highest-level tags html/head/body are modified, 
//    this should be called.
//  - If other packages' init() func's turn out to need this
//    initialization, it may be necessary to call this from
//    those other init() func's. Again, it should be idempotent.
// .
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

