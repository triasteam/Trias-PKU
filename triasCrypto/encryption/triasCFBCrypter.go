package cryptoEncrypt

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

func triasCFBEncrypt(key []byte, plaintext []byte) []byte {
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

    mode := cipher.NewCFBEncrypter(block, iv)
    mode.XORKeyStream(ciphertext[aes.BlockSize:], origData)

    return ciphertext
}

func triasCFBDecrypt(key []byte, ciphertext []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    iv := ciphertext[:aes.BlockSize]
    plaintext = ciphertext[aes.BlockSize:]
    mode := cipher.NewCFBDecrypter(block, iv)
    mode.XORKeyStream(plaintext, plaintext)

    return PKCS5UnPadding(plaintext)
}

