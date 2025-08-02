package main

import "encoding/hex"

func fixedXor(string1 string, string2 string) string {
	hexValue1, e := hex.DecodeString(string1)
	if e != nil {
		panic(e)
	}

	hexValue2, e := hex.DecodeString(string2)
	if e != nil {
		panic(e)
	}

	resultString := make([]byte, len(hexValue1))

	for i := 0; i < len(hexValue1); i++ {
		resultString[i] = hexValue1[i] ^ hexValue2[i]
	}

	return hex.EncodeToString(resultString)
}
