package cryptoEncrypt

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

func triasCBCEncrypt(key []byte, plaintext []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    origData := PKCS5Padding(plaintext, block.BlockSize())
    ciphertext := make([]byte, aes.BlockSize+len(origData))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], origData)

    return ciphertext
}

func triasCBCDecrypt(key []byte, ciphertext []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    iv := ciphertext[:aes.BlockSize]
    plaintext = ciphertext[aes.BlockSize:]
    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(plaintext, plaintext)

    return PKCS5UnPadding(plaintext)
}

