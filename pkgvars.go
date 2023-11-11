// +build js,wasm

package wasmutils

import (
	"syscall/js"
)

// G, Doc, DocBody are (immutable?) global
// singletons, so let's name them that way.
var G, Doc, DocBody js.Value

