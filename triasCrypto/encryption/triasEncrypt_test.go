package cryptoEncrypt

import (
    "fmt"
    "testing"
    "encoding/hex"
    "crypto/aes"
)

var passward, _  = hex.DecodeString("6368616e676520746869732070617373")
var nonce = []byte("123")
var plaintext = []byte("exampleplaintext")
var ciphertext = []byte("exampleciphertext")

func TestTriasEncTools(test *testing.T) {
    fmt.Println("plaintext      - ", plaintext)

    padData := PKCS5Padding(plaintext, aes.BlockSize)
    unpadData := PKCS5UnPadding(padData)
    fmt.Println("PKCS5Padding   - ", padData, len(padData))
    fmt.Println("PKCS5unPadding - ", unpadData, len(unpadData))
}

func TestTriasCBCEncrypt(test *testing.T) {
    ciphertext = triasCBCEncrypt(passward, plaintext)
    fmt.Println("CBC ciphertext - ", ciphertext, len(ciphertext))
}

func TestTriasCBCDecrypt(test *testing.T) {
    plaintext = triasCBCDecrypt(passward, ciphertext)
    fmt.Println("CBC plaintext  - ", plaintext)
}

func TestTriasCFBEncrypt(test *testing.T) {
    ciphertext = triasCFBEncrypt(passward, plaintext)
    fmt.Println("CFB ciphertext - ", ciphertext, len(ciphertext))
}

func TestTriasCFBDecrypt(test *testing.T) {
    plaintext = triasCFBDecrypt(passward, ciphertext)
    fmt.Println("CFB plaintext  - ", plaintext)
}

func TestTriasCTREncrypt(test *testing.T) {
    ciphertext = triasCTREncrypt(passward, plaintext)
    fmt.Println("CTR ciphertext - ", ciphertext, len(ciphertext))
}

func TestTriasCTRDecrypt(test *testing.T) {
    plaintext = triasCTRDecrypt(passward, ciphertext)
    fmt.Println("CTR plaintext  - ", plaintext)
}

func TestTriasOFBEncrypt(test *testing.T) {
    ciphertext = triasOFBEncrypt(passward, plaintext)
    fmt.Println("OFB ciphertext - ", ciphertext, len(ciphertext))
}

func TestTriasOFBDecrypt(test *testing.T) {
    plaintext = triasOFBDecrypt(passward, ciphertext)
    fmt.Println("OFB plaintext  - ", plaintext)
}

func TestTriasGCMEncrypt(test *testing.T) {
    ciphertext, nonce = triasGCMEncrypt(passward, plaintext)
    fmt.Println("GCM ciphertext - ", ciphertext, len(ciphertext))
    fmt.Println("GCM nonce - ", nonce)
}

func TestTriasGCMDecrypt(test *testing.T) {
    plaintext = triasGCMDecrypt(passward, ciphertext, nonce)
    fmt.Println("GCM plaintext  - ", plaintext)
}

