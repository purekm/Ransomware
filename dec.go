package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	cipherFile := "C:/Users/Alpha/Desktop/FLAG.txt.ryk" // file path
	enc_data, err := os.ReadFile(cipherFile)            // encrypted_data file
	if err != nil {
		panic(err)
	}

	hexkey := "133a985d25765d4af3c84fcb1f8296f888d5d8fa028697e186939dbaf283108e" // AES 32 Bytes Key

	key, err := hex.DecodeString(hexkey) // hex to bytes
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block) // AES-GCM
	if err != nil {
		panic(err)
	}

	nonce := enc_data[:gcm.NonceSize()] // NonceSize = 12
	data := enc_data[gcm.NonceSize():]  // enc_data + tag = 16

	plaintext, err := gcm.Open(nil, nonce, data, nil) // decrypt + tag verify
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", plaintext)
}
