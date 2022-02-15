//go:build !js || !wasm

package wasmutils

// IsWasm is compiled differently for WASM / non-WASM.
func IsWasm() bool {
	return false
}
