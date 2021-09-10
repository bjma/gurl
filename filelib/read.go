package filelib

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// For our use case, we want to be able to parse
// JSON files, but also read regular text from files.
// So, we should always check the file extension first.
func ReadFile(filepath string) []byte {
	var d []byte

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// If not JSON, simply read into byte array
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic("gurl: unable to read file")
	}

	// We want to be structure-agnostic, meaning that
	// the structure/schema of the JSON data does
	// not necessarily have to be uniform. We simply
	// want to read the data into a data structure,
	// and send it via HTTP.
	// see: https://www.sohamkamani.com/golang/parsing-json/#unstructured-data-decoding-json-to-maps

	// We also want to be able to read JSON arrays, as `curl` currently does not do this.
	// see: https://www.sohamkamani.com/golang/parsing-json/#json-arrays
	if GetFileExtension(filepath) == "json" {
		var jsonData map[string]interface{}
		json.Unmarshal(bytes, &jsonData)
		d, err = json.Marshal(jsonData)
	} else {
		d = bytes
	}
	return d
}
