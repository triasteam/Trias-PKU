package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "crypto/aes"
    "crypto/cipher"
    "math/big"
    "fmt"
    "bytes"
    "time"
)

func printKey(privateKey *ecdsa.PrivateKey) {
    fmt.Println("Private:", privateKey.D.String(), "Public:", privateKey.PublicKey.X, privateKey.PublicKey.Y)
}

//pks full bytes
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

//aes encrypt
func AesEncrypt(origData []byte, key []byte) ([]byte, error) {
    hashkey := sha256.Sum256(key)
    block, err := aes.NewCipher(hashkey[:])
    if err != nil {
        return nil, err
    }
    blockMode := cipher.NewCBCEncrypter(block, hashkey[:16])

    origData = PKCS5Padding(origData, block.BlockSize())
    crypted := make([]byte, len(origData))
    blockMode.CryptBlocks(crypted, origData)

    return crypted, nil
}

//aes decrypt
func AesDecrypt(crypted []byte, key []byte) ([]byte, error) {
    hashkey := sha256.Sum256(key)
    block, err := aes.NewCipher(hashkey[:])
    if err != nil {
        return nil, err
    }
    blockMode := cipher.NewCBCDecrypter(block, hashkey[:16])

    origData := make([]byte, len(crypted))
    blockMode.CryptBlocks(origData, crypted)
    origData = PKCS5UnPadding(origData)

    return origData, nil
}

func originStealthAdressGenerate() (*big.Int, *ecdsa.PrivateKey){
    privateKeyR, err := ecdsa.GenerateKey(curveP256, rand.Reader)
    if err != nil {
        panic(err)
    }
    //    printKey(privateKeyR)

    //generate Stealth Address
    Px, Py := curveP256.ScalarMult(privateKeyA.PublicKey.X, privateKeyA.PublicKey.Y, privateKeyR.D.Bytes()) //Mr
    //fmt.Println("Px", P, "Py", Py)

    HashP := sha256.Sum256(append(Px.Bytes(), Py.Bytes()...))
//    x, y := curveP256.ScalarBaseMult(P.Bytes()) //(Mr)G
    x, y := curveP256.ScalarBaseMult(HashP[:]) //(Mr)G
    P, _ := curveP256.Add(x, y, privateKeyB.PublicKey.X, privateKeyB.PublicKey.Y) //(Mr)G+N

    return P, privateKeyR
}


func originStealthAdressVerify(P *big.Int, privateKeyR *ecdsa.PrivateKey) {
    //verify Stealth Address
    px, py := curveP256.ScalarMult(privateKeyR.PublicKey.X, privateKeyR.PublicKey.Y, privateKeyA.D.Bytes()) //mR
    //fmt.Println("px", p.Bytes(), "py", py.Bytes())

    HashP := sha256.Sum256(append(px.Bytes(), py.Bytes()...))
    x, y := curveP256.ScalarBaseMult(HashP[:]) //(mR)G
//    x, y := curveP256.ScalarBaseMult(p.Bytes()) //(mR)G
    p, _ := curveP256.Add(x, y, privateKeyB.PublicKey.X, privateKeyB.PublicKey.Y) //(mR)G+N

    if P.Cmp(p) != 0 {
        fmt.Println("Stealth Address", false)
    }
}

func improveStealthAdressGenerate() (*big.Int, *ecdsa.PrivateKey, []byte){
    privateKeyR, err := ecdsa.GenerateKey(curveP256, rand.Reader)
    if err != nil {
        panic(err)
    }
    //    printKey(privateKeyR)

    //generate Stealth Address
    P, Py := curveP256.ScalarMult(privateKeyA.PublicKey.X, privateKeyA.PublicKey.Y, privateKeyR.D.Bytes()) //Mr
    //fmt.Println("Px", P, "Py", Py)

    x, y := curveP256.ScalarBaseMult(P.Bytes()) //(Mr)G
    P, _ = curveP256.Add(x, y, privateKeyB.PublicKey.X, privateKeyB.PublicKey.Y) //(Mr)G+N

    //improve Stealth Address
    pen, _ := AesEncrypt(P.Bytes(), Py.Bytes())

    return P, privateKeyR, pen
}

func improveStealthAdressVerify(P *big.Int, privateKeyR *ecdsa.PrivateKey, pen []byte) {
    //verify Stealth Address
    i := 1
    px, py := privateKeyR.PublicKey.X, privateKeyR.PublicKey.Y
    valueA := privateKeyA.D.Bytes()

    tStart := time.Now()
    for i<1000 {
//        px, py = curveP256.ScalarMult(privateKeyR.PublicKey.X, privateKeyR.PublicKey.Y, privateKeyA.D.Bytes()) //mR
        px, py = curveP256.ScalarMult(privateKeyR.PublicKey.X, privateKeyR.PublicKey.Y, valueA) //mR
        i++
    }
    tCost := time.Since(tStart)
    fmt.Println("Phase 01 Cost Time", tCost)

    tStart = time.Now()
    i = 0
    for i<1000 {
        pde, _ := AesDecrypt(pen, py.Bytes())
        //fmt.Println("Pde", pde,"\nP", P.Bytes())
        result := bytes.Equal(pde, P.Bytes())
        if !result {
            fmt.Println("pde", pde, "\npen", pen,"\nP", P.Bytes(), "\npy", py.Bytes())
        }
        i++
    }
    tCost = time.Since(tStart)
    fmt.Println("Phase 02 Cost Time", tCost)

    tStart = time.Now()
    i = 0
    for i<1000 {
        x, y := curveP256.ScalarBaseMult(px.Bytes()) //(mR)G
        p, _ := curveP256.Add(x, y, privateKeyB.PublicKey.X, privateKeyB.PublicKey.Y) //(mR)G+N
        if P.Cmp(p) != 0 {
            fmt.Println("Stealth Address", false)
        }
        i++
    }
    tCost = time.Since(tStart)
    fmt.Println("Phase 03 Cost Time", tCost)
}

//global params
var curveP256 = elliptic.P256()
//generate public and private key pairs
var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
/*if err != nil {
    panic(err)
}
printKey(privateKeyB)
*/

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

    tStart = time.Now()
    i = 0
    for i<1000 {
        improveStealthAdressGenerate()
        i++
    }
    tCost = time.Since(tStart)
    fmt.Println("Improve Generate Cost Time", tCost)

//    tStart = time.Now()
    i = 0
    for i<1 {
        P, R, M := improveStealthAdressGenerate()
        improveStealthAdressVerify(P, R, M)
        i++
    }
//    tCost = time.Since(tStart)
//    fmt.Println("Improve Generate and Verify Cost Time", tCost)

/*
    hash := sha256.Sum256([]byte(msg))

    r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
    if err != nil {
        panic(err)
    }
    fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)

    valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
    fmt.Println("signature verified:", valid)
*/

}
