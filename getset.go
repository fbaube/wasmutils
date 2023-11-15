// g o : build js && wasm

// Wired: "go:build"
// Tired: "+build"

// This file needs a filename "_wasm.go"

package wasmutils

import (
	"syscall/js"
	"strconv"
)

// GetElmByTag is TBS.
//
// func (v Value) Get(p string) Value :: 
//  - Get returns the JS property p of value v
//  - It panics if v is not a JS object
// .
func GetElmByTag(s string) js.Value {
	return Doc.Get(s)
}

// GetElmByID is fab.
//
// func (v Value) Call(m string, args ...any) Value
//  - Call does a JS call to the method m of value v with the given args
//  - It panics if v has no method m
//  - The arguments get mapped to JS values per the ValueOf function.
// . 
func GetElmByID(s string) js.Value {
	return Doc.Call("getElementById", s)
}

// GetElmIntValByID is TBS.
func GetElmIntValByID(s string) int {
	ss := GetElmByID(s).Get("value").String()
	ii, _ := strconv.Atoi(ss)
	return ii
}

// SetElmIntValByID is TBS.
func SetElmIntValByID(s string, i int) {
	GetElmByID(s).Set("value", i)
}

// SetElmTextValByID is TBS.
func SetElmTextValByID(s string, v string) {
	GetElmByID(s).Set("value", v)
}

// Alert displays a modal(?) dialog.
func Alert(s string) {
	js.Global().Get("alert").Invoke(s)
}

