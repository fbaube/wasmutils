// +build js,wasm

package wasmutils

import (
	"syscall/js"
	"fmt"
	"os"
)

func Init() {
     init()
}

func init() {
     initPkgVars()
}

// initPkgVars sets up some basic global (package) variables
// related to JS. Don't panic if a "<body>" tag is not found.
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
		fmt.Fprintf(os.Stderr, "CAN'T GET JS Global and/or webpage DOM")
		return
	}
	DocBody = Doc.Get("body")
	if !DocBody.Truthy() { // Don't panic !
		// panic("CAN'T GET webpage <body>")
		fmt.Fprintf(os.Stderr, "Document has no <body> tag")
	}
}

