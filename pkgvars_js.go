// +build js,wasm

package wasmutils

import (
	"syscall/js"
)

// G, Doc, DocBody are (immutable?) global
// singletons, so let's go ahead and set
// them up as easy-to-use package variables.
//
// We initialize them in func init, but maybe users of
// this package want to re-call func init themselves
// by calling func Init()] 
// (but only from a single thread, to avoid a race).
// .
var G, Doc, DocBody js.Value

