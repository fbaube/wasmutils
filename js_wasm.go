//go:build js && wasm

// (should probably test js only)

package wasmutils

import (
	"syscall/js"
)

// IsBad is compiled only for JS+WASM.
func IsBad(v js.Value) bool {
	// return v == js.Undefined() || v == js.Null()
	return v.Equal(js.Undefined()) || v.Equal(js.Null())
}

// IsWasm is compiled only for JS+WASM.
func IsWasm() bool {
	return js.Global != nil
}

// IsBrowser is compiled only for JS+WASM.
func IsBrowser() bool {
	return js.Global != nil && !js.Global().Get("document").Equal(js.Undefined())
}
