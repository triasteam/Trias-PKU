package cryptoEncrypt

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

func triasGCMEncrypt(key []byte, plaintext []byte) ([]byte, []byte) {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err.Error())
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }

    origData := PKCS5Padding(plaintext, block.BlockSize())
    ciphertext := aesgcm.Seal(nil, nonce, origData, nil)

    return ciphertext, nonce
}

func triasGCMDecrypt(key []byte, ciphertext []byte, nonce []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err.Error())
    }
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }

    plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        panic(err.Error())
    }

    return PKCS5UnPadding(plaintext)
}

