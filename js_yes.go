// +build js,wasm

package wasmutils

import (
	"syscall/js"
)

// IsBad is compiled only for WASM.
func IsBad(v js.Value) bool {
	return v == js.Undefined() || v == js.Null()
}

// IsWasm is compiled only for WASM.
func IsWasm() bool {
	return true
}
