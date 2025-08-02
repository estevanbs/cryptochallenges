package main

import (
	"encoding/base64"
	"encoding/hex"
)

func hexToBase64(hexString string) string {
	hexValue, e := hex.DecodeString(hexString)
	if e != nil {
		panic(e)
	}

	converted := base64.StdEncoding.EncodeToString(hexValue)

	return converted
}
