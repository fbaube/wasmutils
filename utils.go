// +build js,wasm

package wasmutils

import (
       "syscall/js"
       "strconv"
       // "time"
       )

// JSfunc is a typedef for a JS function that can be called from Go.
type JSfunc func(this js.Value, args []js.Value) interface{}

var (
  // Doc is THE global HTML document.
	Doc  = js.Global().Get("document")
  // Body is THE global HTML <body> tag.
	Body = GetElmByID("body") // Doc.Call("getElementById", "body")
)

/*
  <input type="text" id="val1"/>
  <input type="text" id="val2"/>
  <button onClick="add('val1', 'val2', 'result');"
  	  id="addButton">Add</button>
  <button onClick="subtract('val1', 'val2', 'result');"
  	  id="subtractButton">Subtract</button>
  NOTE that this next one is an "input" even tho we use it as write-only!
     <input type="text" id="result">
*/

// GetElmByID is TBS.
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

// SetCallback returns a wrapped function.
// See https://golang.org/pkg/syscall/js/#FuncOf
//   func FuncOf(fn func(this Value, args []Value) interface{}) Func
// Note then that fn is:
//   func(this Value, args []Value) interface{}
//
// Invoking the JS function will synchronously call the Go function
// fn with the value of JS's "this" keyword and the arguments of the
// invocation. The return value of the invocation is the result of
// the Go function mapped back to JS according to ValueOf.
// A wrapped function triggered during a call from Go to JS gets
// executed on the same goroutine. A wrapped function triggered by
// JS's event loop gets executed on an extra goroutine. Blocking
// operations in the wrapped function will block the event loop.
// As a consequence, if one wrapped function blocks, other wrapped
// funcs will not be processed. A blocking function should therefore
// explicitly start a new goroutine.
// Func.Release must be called to free up resources when the function
// will not be used any more.
//
// Usage example:
//  var cb js.Func
//  cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
//     fmt.Println("button clicked")
//     cb.Release() // release the function if button won't be clicked again
//     return nil
//     })
//  js.Global().Get("document")
//  .Call("getElementById", "myButton")
//  .Call("addEventListener", "click", cb)
//
func SetCallback(funcName string, funcBody JSfunc) {
	// OBS/prev.: js.Global().Set(funcName, js.NewCallback(funcBody))
	js.Global().Set(funcName, js.FuncOf(funcBody))
}

// StayInMemory blocks WASM from exiting.
func StayInMemory() {
	c := make(chan struct{}, 0)
	<-c
}

// Alert displays a modal(?) dialog.
func Alert(s string) {
  	js.Global().Get("alert").Invoke(s)
}
