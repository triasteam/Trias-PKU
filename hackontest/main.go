package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "hackonpkg"
    "fmt"
    "time"
)

//global params
var curveP256 = elliptic.P256()
//generate public and private key pairs
var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)

func main() {
    tStart := time.Now()
    i := 0
    for i<1 {
        hackonpkg.OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)
        i++
    }
    tCost := time.Since(tStart)
    fmt.Println("Origin Generate Cost Time", tCost)

    tStart = time.Now()
    i = 0
    for i<1 {
        P, r := hackonpkg.OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)
        hackonpkg.OriginStealthAdressVerify(privateKeyA, privateKeyB, P, r, curveP256)
        i++
    }
    tCost = time.Since(tStart)
    fmt.Println("Origin Generate and Verify Cost Time", tCost)
}
