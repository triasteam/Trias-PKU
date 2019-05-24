package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "math/big"
    "fmt"
    "time"
)

func printKey(privateKey *ecdsa.PrivateKey) {
    fmt.Println("Private:", privateKey.D.String(), "Public:", privateKey.PublicKey.X, privateKey.PublicKey.Y)
}

func originStealthAdressGenerate() (*big.Int, *ecdsa.PrivateKey){
    //generate Stealth Address
    privateKeyR, err := ecdsa.GenerateKey(curveP256, rand.Reader)
    if err != nil {
        panic(err)
    }

    Px, Py := curveP256.ScalarMult(privateKeyA.PublicKey.X, privateKeyA.PublicKey.Y, privateKeyR.D.Bytes()) //Mr
    HashP := sha256.Sum256(append(Px.Bytes(), Py.Bytes()...))
    x, y := curveP256.ScalarBaseMult(HashP[:]) //(Mr)G
    P, _ := curveP256.Add(x, y, privateKeyB.PublicKey.X, privateKeyB.PublicKey.Y) //(Mr)G+N

    return P, privateKeyR
}

func originStealthAdressVerify(P *big.Int, privateKeyR *ecdsa.PrivateKey) {
    //verify Stealth Address
    px, py := curveP256.ScalarMult(privateKeyR.PublicKey.X, privateKeyR.PublicKey.Y, privateKeyA.D.Bytes()) //mR
    HashP := sha256.Sum256(append(px.Bytes(), py.Bytes()...))
    x, y := curveP256.ScalarBaseMult(HashP[:]) //(mR)G
    p, _ := curveP256.Add(x, y, privateKeyB.PublicKey.X, privateKeyB.PublicKey.Y) //(mR)G+N

    if P.Cmp(p) != 0 {
        fmt.Println("Stealth Address", false)
    }
}

//global params
var curveP256 = elliptic.P256()
//generate public and private key pairs
var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)

func main() {
    tStart := time.Now()
    i := 0
    for i<1000 {
        originStealthAdressGenerate()
        i++
    }
    tCost := time.Since(tStart)
    fmt.Println("Origin Generate Cost Time", tCost)

    tStart = time.Now()
    i = 0
    for i<1000 {
        P, R := originStealthAdressGenerate()
        originStealthAdressVerify(P, R)
        i++
    }
    tCost = time.Since(tStart)
    fmt.Println("Origin Generate and Verify Cost Time", tCost)

}
