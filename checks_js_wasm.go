// g o : build js || wasm

// (should probably test js only)

package wasmutils

import (
	"syscall/js"
)

// IsWasm is compiled only for JS+WASM.
func IsWasm() bool {
	return true // js.Global != nil // GOOS js && GOARCH wasm 
}

// IsBrowser is compiled only for JS+WASM.
func IsBrowser() bool {
	// return js.Global != nil && !js.Global().Get("document").Equal(js.Undefined())
	return js.Global().Truthy() && js.Global().Get("document").Truthy()
}
