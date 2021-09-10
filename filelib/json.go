package filelib

import (
	"bytes"
)

func jsonIsArray(b []byte) bool {
    d := bytes.TrimLeft(b, " \t\r\n")
    return len(d) > 0 && d[0] == '['
}

func jsonIsObject(b []byte) bool {
    d := bytes.TrimLeft(b, " \t\r\n")
    return len(d) > 0 && d[0] == '{'
}