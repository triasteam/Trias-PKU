package hackonpkg

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/sha256"
    "crypto/rand"
    "math/big"
    "fmt"
)

func OriginStealthAdressGenerate(privateKeyA *ecdsa.PrivateKey, privateKeyB *ecdsa.PrivateKey,
                                 curveP256 elliptic.Curve) (*big.Int, *ecdsa.PrivateKey){
    //generate Stealth Address
    privateKeyR, err := ecdsa.GenerateKey(curveP256, rand.Reader) //r
    if err != nil {
        panic(err)
    }

    Px, Py := curveP256.ScalarMult(privateKeyA.PublicKey.X, privateKeyA.PublicKey.Y, privateKeyR.D.Bytes()) //rM

    HashP := sha256.Sum256(append(Px.Bytes(), Py.Bytes()...))
    x, y := curveP256.ScalarBaseMult(HashP[:]) //H(rM)G
    P, _ := curveP256.Add(x, y, privateKeyB.PublicKey.X, privateKeyB.PublicKey.Y) //H(rM)G+N

    return P, privateKeyR
}

func OriginStealthAdressVerify(privateKeyA *ecdsa.PrivateKey, privateKeyB *ecdsa.PrivateKey,
                               P *big.Int, privateKeyR *ecdsa.PrivateKey, curveP256 elliptic.Curve) {
    //verify Stealth Address
    px, py := curveP256.ScalarMult(privateKeyR.PublicKey.X, privateKeyR.PublicKey.Y, privateKeyA.D.Bytes()) //mR

    HashP := sha256.Sum256(append(px.Bytes(), py.Bytes()...))
    x, y := curveP256.ScalarBaseMult(HashP[:]) //H(mR)G
    p, _ := curveP256.Add(x, y, privateKeyB.PublicKey.X, privateKeyB.PublicKey.Y) //H(mR)G+N

    if P.Cmp(p) != 0 {
        fmt.Println("Stealth Address", false)
    }
}
