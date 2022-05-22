package test

import (
	"encoding/base64"
	"testing"
)

func TestEncode(t *testing.T) {
	name := "刘德afas"
	toString := base64.StdEncoding.EncodeToString([]byte(name))
	print(toString)
	decodeString, _ := base64.StdEncoding.DecodeString(toString)
	println(string(decodeString))
}
