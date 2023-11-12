package wasmutils

import (
	"fmt"
	"syscall/js"
)

// FuncOf takes a first class function with a js.Value and []js.Value,
// and returns an "any".
// The function passed to FuncOf will be called synchronously from JS.
// Arg 1 of the passed function is JS’s "this", i.e. JS’s global object.
// Arg 2 of the passed function is a []js.Value with the args for the JS
// function call: the unformatted JSON input string.
// It returns the function to be called from JS.
// 
func jsonWrapper() js.Func {
        jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
                if len(args) != 1 {
                        return "Bad arg count"
                }
		jsonOuputTextArea := Doc.Call(
			"getElementById", "id_outtext_json") // "jsonoutput")
		if !jsonOuputTextArea.Truthy() {
			return AsMappedError("Can't get output text area")
		}
                inputJSON := args[0].String()
                fmt.Printf("input %s\n", inputJSON)
                pretty, err := AsPrettyJson(inputJSON)
                if err != nil {
		       return AsMappedError(
		       "Can't to parse JSON, got error: \n" + err.Error() + "\n")
                }
		jsonOuputTextArea.Set("value", pretty)
                return nil
        })
	// When we return a value from Go to JS, the compiler auto-uses 
	// [js.ValueOf] to convert the Go value to a JS value. Here Go
	// is returning a string to JS, which the compiler will convert 
	// to JS's string type .. by using js.ValueOf()
	// 
	// i.e. https://pkg.go.dev/syscall/js#ValueOf
        return jsonFunc
}

