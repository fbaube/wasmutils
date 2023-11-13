// +build js,wasm

package wasmutils

import (
	"fmt"
	S "strings"
	"syscall/js"
	"time"
)

func CurrentTime() string {
	sTime := time.Now().Local().Format(time.RFC3339)
	fmt.Println("Execution at", sTime)
	// Replace(str, oldstr, newstr string, m int) string
	sTime = S.Replace(sTime, "T", "&nbsp;", 1)
	sTime = S.Replace(sTime, "+", "&nbsp;<small>GMT+", 1)
	sTime = S.TrimSuffix(sTime, ":00")
	sTime = S.TrimPrefix(sTime, "20")
	sTime = "<font color=\"grey\">20</font>" + sTime + "</small>"
	return sTime
}

func DoDomDemo() {
	var p js.Value
	// 1. Add an <h1>
	p = Doc.Call("createElement", "h1")
	p.Set("innerHTML", "This is an H1")
	DocBody.Call("appendChild", p)
	p = Doc.Call("createElement", "h2")
	p.Set("innerHTML", "This is an H2")
	DocBody.Call("appendChild", p)
	p = Doc.Call("createElement", "h3")
	p.Set("innerHTML", "This is an H3")
	DocBody.Call("appendChild", p)
	p = Doc.Call("createElement", "h4")
	p.Set("innerHTML", "This is an H4")
	DocBody.Call("appendChild", p)
	p = Doc.Call("createElement", "h5")
	p.Set("innerHTML", "This is an H5")
	DocBody.Call("appendChild", p)
	p = Doc.Call("createElement", "h6")
	p.Set("innerHTML", "This is an H6")
	DocBody.Call("appendChild", p)
	// 2. Expose Go functions/values in JS variables.
	// js.Global().Set("goVar", "I am a variable set from Go")
	// js.Global().Set("sayHello", js.FuncOf(sayHello))
}

