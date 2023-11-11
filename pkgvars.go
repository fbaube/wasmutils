// +build js,wasm

package wasmutils

import (
	"syscall/js"
)

// G, Doc, DocBody are (immutable?) global
// singletons, so let's name them that way.
var G, Doc, DocBody js.Value

// JSfunc is a typedef for a JS function that Go can call.
type JSfunc func(this js.Value, args []js.Value) interface{}

