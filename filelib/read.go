package filelib

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bjma/gurl/handler"
)

// Reads from file defined by `path` and returns bytes read.
func ReadFile(path string) []byte {
	var d []byte

	f, err := os.Open(path)
	if err != nil {
		handler.HandleError(err)
	}
	defer f.Close()

	// If not JSON, simply read into byte array
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
        // NOTE: cURL handles null file descriptors as empty bodies
		handler.HandleError(err)
	}
	// We also want to be able to read JSON arrays, as `curl` currently does not do this.
	// see: https://www.sohamkamani.com/golang/parsing-json/#json-arrays
	if GetFileExtension(path) == "json" {
		if jsonIsArray(bytes) {
			d = readJsonArray(bytes)
		} else {
			d = readJsonObj(bytes)
		}
	} else {
		d = bytes
	}
	return d
}

// Reads JSON object
func readJsonObj(b []byte) []byte {
    if !jsonIsObject(b) {
        err := handler.NewError("not a valid JSON object")
        handler.HandleError(err)
    }
	var jsonData map[string]interface{}
	json.Unmarshal(b, &jsonData)
	d, err := json.Marshal(jsonData)
	if err != nil {
		handler.HandleError(err)
	}
	return d
}

// Read JSON array
func readJsonArray(b []byte) []byte {
	var jsonData []map[string]interface{}
	json.Unmarshal(b, &jsonData)
	d, err := json.Marshal(jsonData)
	if err != nil {
		handler.HandleError(err)
	}
	return d
}
