package wasmutils

import (
	"syscall/js"
)

// IsWasm is compiled only for JS+WASM.
func IsWasm() bool {
	return true 
}

// IsBrowser is compiled only for JS+WASM.
func IsBrowser() bool {
	return js.Global().Truthy() && js.Global().Get("document").Truthy()
}
