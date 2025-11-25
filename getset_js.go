// g o : build js && wasm

// Wired: "go:build"
// Tired: "+build"

// This file needs a filename "_wasm.go" or "_js.go"

package wasmutils

import (
	"syscall/js"
	"strconv"
)

// GetElmByTag returns the first(?) XML element
// whose tag matches (e.g. "body", "p").
func GetElmByTag(s string) js.Value {
	return Doc.Get(s)
}

// GetElmByID returns the XML element
// whose "id=" attribute matches.
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

func GetAsString(name string) string {
     return GetElmByID(name).Get("value").String()
}

func NewElm(tagName string) js.Value {
     return Doc.Call("createElement", tagName)
}

