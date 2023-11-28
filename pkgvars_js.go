// +build js,wasm

package wasmutils

import (
	"syscall/js"
)

// G, Doc, DocBody are (immutable?) global
// singletons, so let's go ahead and set
// them up as easy-to-use package variables.
// .
var G, Doc, DocBody js.Value


