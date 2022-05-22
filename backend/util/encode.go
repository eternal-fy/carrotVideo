package util

import "encoding/base64"

func StringEncode(name string) string {
	str := base64.StdEncoding.EncodeToString([]byte(name))
	return str

}

func StringDecode(code string) string {
	decodeString, _ := base64.StdEncoding.DecodeString(code)
	return string(decodeString)
}
