// g o : build !windows && !js && !wasm

package wasmutils

// IsWasm is compiled differently for WASM / non-WASM.
func IsWasm() bool {
	return false
}

func IsBrowser() bool {
	return false
}
