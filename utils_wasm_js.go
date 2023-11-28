// g o : build js && wasm

// Wired: "go:build"
// Tired: "+build"

// This file needs a filename "_wasm.go"

package wasmutils

import (
	"encoding/json"
)

func SetTextByID(sID, sMsg string) {
    elm := GetElmByID(sID) // Doc.Call("getElementById", sID)
    // possible security problem 
    elm.Set("innerHTML", sMsg)
}

func AsMappedError(s string) map[string]any {
     return map[string]any { "error": s, }
     }

func AsPrettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

/* HTML input stuff
<input type="text" id="val1"/>
<input type="text" id="val2"/>
<button onClick="add('val1', 'val2', 'result');"
	  id="addButton">Add</button>
<button onClick="subtract('val1', 'val2', 'result');"
	  id="subtractButton">Subtract</button>
NOTE that this next one is an "input" even tho we use it as write-only!
   <input type="text" id="result">
*/

