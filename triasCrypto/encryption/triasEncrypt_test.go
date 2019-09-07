package cryptoEncrypt

import (
    "fmt"
    "testing"
    "encoding/hex"
    "crypto/aes"
    "crypto/rand"
    "crypto/ecdsa"
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

func BenchmarkTriasCBCEncrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext = triasCBCEncrypt(passward, plaintext)
    }
}

func TestTriasCBCDecrypt(test *testing.T) {
    plaintext = triasCBCDecrypt(passward, ciphertext)
    fmt.Println("CBC plaintext  - ", plaintext)
}

func BenchmarkTriasCBCDecrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext = triasCBCEncrypt(passward, plaintext)
        plaintext = triasCBCDecrypt(passward, ciphertext)
    }
}

func TestTriasCFBEncrypt(test *testing.T) {
    ciphertext = triasCFBEncrypt(passward, plaintext)
    fmt.Println("CFB ciphertext - ", ciphertext, len(ciphertext))
}

func BenchmarkTriasCFBEncrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext = triasCFBEncrypt(passward, plaintext)
    }
}

func TestTriasCFBDecrypt(test *testing.T) {
    plaintext = triasCFBDecrypt(passward, ciphertext)
    fmt.Println("CFB plaintext  - ", plaintext)
}

func BenchmarkTriasCFBDecrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext = triasCFBEncrypt(passward, plaintext)
        plaintext = triasCFBDecrypt(passward, ciphertext)
    }
}

func TestTriasCTREncrypt(test *testing.T) {
    ciphertext = triasCTREncrypt(passward, plaintext)
    fmt.Println("CTR ciphertext - ", ciphertext, len(ciphertext))
}

func BenchmarkTriasCTREncrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext = triasCTREncrypt(passward, plaintext)
    }
}

func TestTriasCTRDecrypt(test *testing.T) {
    plaintext = triasCTRDecrypt(passward, ciphertext)
    fmt.Println("CTR plaintext  - ", plaintext)
}

func BenchmarkTriasCTRDecrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext = triasCTREncrypt(passward, plaintext)
        plaintext = triasCTRDecrypt(passward, ciphertext)
    }
}

func TestTriasOFBEncrypt(test *testing.T) {
    ciphertext = triasOFBEncrypt(passward, plaintext)
    fmt.Println("OFB ciphertext - ", ciphertext, len(ciphertext))
}

func BenchmarkTriasOFBEncrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext = triasOFBEncrypt(passward, plaintext)
    }
}

func TestTriasOFBDecrypt(test *testing.T) {
    plaintext = triasOFBDecrypt(passward, ciphertext)
    fmt.Println("OFB plaintext  - ", plaintext)
}

func BenchmarkTriasOFBDecrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext = triasOFBEncrypt(passward, plaintext)
        plaintext = triasOFBDecrypt(passward, ciphertext)
    }
}

func TestTriasGCMEncrypt(test *testing.T) {
    ciphertext, nonce = triasGCMEncrypt(passward, plaintext)
    fmt.Println("GCM ciphertext - ", ciphertext, len(ciphertext))
    fmt.Println("GCM nonce - ", nonce)
}

func BenchmarkTriasGCMEncrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext, nonce = triasGCMEncrypt(passward, plaintext)
    }
}

func TestTriasGCMDecrypt(test *testing.T) {
    plaintext = triasGCMDecrypt(passward, ciphertext, nonce)
    fmt.Println("GCM plaintext  - ", plaintext)
}

func BenchmarkTriasGCMDecrypt(b *testing.B) {
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext, nonce = triasGCMEncrypt(passward, plaintext)
        plaintext = triasGCMDecrypt(passward, ciphertext, nonce)
    }
}

func TestTriasDHShareKey(test *testing.T) {
    var privateLocal, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    var privatePartner, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    publicPartner := privatePartner.Public().(*ecdsa.PublicKey)
    publicLocal := privateLocal.Public().(*ecdsa.PublicKey)
    secretKey := triasDHShareKey(publicPartner, privateLocal)
    fmt.Println("DH secretKey 1- ", secretKey, len(secretKey))
    secretKey = triasDHShareKey(publicLocal, privatePartner)
    fmt.Println("DH secretKey 2- ", secretKey, len(secretKey))
}

func BenchmarkTriasDHShareKey(b *testing.B) {
    var privateLocal, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    var privatePartner, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    publicPartner := privatePartner.Public().(*ecdsa.PublicKey)
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        triasDHShareKey(publicPartner, privateLocal)
    }
}

func TestTriasDHEncrypt(test *testing.T) {
    var privatePartner, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    publicPartner := privatePartner.Public().(*ecdsa.PublicKey)

    ciphertext, publicLocal := triasDHEncrypt(publicPartner, plaintext)
    fmt.Println("DH CTR ciphertext - ", ciphertext, len(ciphertext))
    fmt.Println("DH publicLocal - ", publicLocal)
}

func BenchmarkTriasDHEncrypt(b *testing.B) {
    var privatePartner, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    publicPartner := privatePartner.Public().(*ecdsa.PublicKey)
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        triasDHEncrypt(publicPartner, plaintext)
    }
}

func TestTriasDHDecrypt(test *testing.T) {
    var privatePartner, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    publicPartner := privatePartner.Public().(*ecdsa.PublicKey)

    ciphertext, publicLocal := triasDHEncrypt(publicPartner, plaintext)
    plaintext = triasDHDecrypt(publicLocal, ciphertext, privatePartner)
    fmt.Println("DH CTR plaintext - ", plaintext, len(plaintext))
}

func BenchmarkTriasDHDecrypt(b *testing.B) {
    var privatePartner, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    publicPartner := privatePartner.Public().(*ecdsa.PublicKey)
    //reset timer
    b.ResetTimer()
    //stop timer
    b.StopTimer()
    //start timer
    b.StartTimer()

    for i := 0; i < b.N; i++ {
        ciphertext, publicLocal := triasDHEncrypt(publicPartner, plaintext)
        plaintext = triasDHDecrypt(publicLocal, ciphertext, privatePartner)
    }
}
