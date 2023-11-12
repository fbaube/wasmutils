// g o : build js && wasm

// Wired: "go:build"
// Tired: "+build"

// This file needs a filename "_wasm.go"

package wasmutils

import (
	"syscall/js"
)

// JSfunc is a typedef for a Go function that can be registered with JS,
// so that a JS function can call back into Go.
// A function passed to [FuncOf] will be called synchronously from JS.
// Arg 1 is JS's "this", i.e. the [js.Global] object.
// Arg 2 is a []js.Value with the args for the JS function call.
// . 
type JSfunc func(this js.Value, args []js.Value) any // interface{}

// SetCallback wraps a Go function to return a function callable from JS. 
// See https://golang.org/pkg/syscall/js/#FuncOf
//
//	func FuncOf(fn func(this Value, args []Value) interface{}) Func
//
// Note then that fn is an instance of [JSfunc]: 
//
//	func(this Value, args []Value) interface{}
//
// Then execution is as follows:
//  - In the web oage's JS code, invoking the JS function will
//    synchronously call the Go function fn with the value of
//    JS's "this" keyword and the arguments of the invocation.
//  - The return value of the invocation is the result of the
//    Go function mapped back to JS according to ValueOf(..).
// Regarding execution threads:
//  - If a call from Go into JS then calls the wrapped Go 
//    function, it is executed on the same goroutine.
//  - If the wrapped Go function is called from the JS event
//    loop, the Go function is executed on an extra goroutine.
//    In this case, blocking operations in the wrapped Go 
//    function will block the event loop. As a consequence, 
//    if one wrapped function blocks, other wrapped funcs will 
//    not be processed. A blocking function should therefore 
//    explicitly start a new goroutine.
//  - [Func.Release] must be called to free up resources when 
//    the function will not be used any more. It may be called
//    while the function is still executing a call. 
// 
// FuncOf takes a first class function with a js.Value and []js.Value,
// and returns an "any".
//  - The function passed to FuncOf will be called synchronously from JS
//  - Arg 1 of the passed function is JS's "this", i.e. JS's global object
//  - Arg 2 of the passed function is a []js.Value with the args for the 
//    JS function call: the unformatted JSON input string
//  - It returns the function to be called from JS
//
// Usage example
//
//	var cb js.Func
//	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
//	   fmt.Println("button clicked")
//	   cb.Release() // release the function if button won't be clicked again
//	   return nil
//	   })
//	js.Global().Get("document")
//	.Call("getElementById", "myButton")
//	.Call("addEventListener", "click", cb)
// . 
func SetCallback(funcName string, funcBody JSfunc) {
	// The old obsolete way:
	// js.Global().Set(funcName, js.NewCallback(funcBody))
	// The new improved way: 
	js.Global().Set(funcName, js.FuncOf(funcBody))
}

