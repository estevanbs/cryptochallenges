package main

import "encoding/hex"

func encryptRepeatingKeyXor(text string, key string) string {
	textBytes := []byte(text)
	keyBytes := []byte(key)

	encryptedString := []byte{}
	for i := 0; i < len(textBytes); i++ {
		j := i % len(keyBytes)

		encryptedString = append(encryptedString, textBytes[i]^keyBytes[j])
	}

	return hex.EncodeToString(encryptedString)
}
