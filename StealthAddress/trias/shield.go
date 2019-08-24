package stealthaddress

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "math/big"
)

func CreateShieldAddr(publicView, publicSpend *ecdsa.PublicKey) (publicShieldP, publicShieldR *ecdsa.PublicKey){
    publicShieldP = new(ecdsa.PublicKey)
    publicShieldP.Curve = curveP256

    privateShieldR, err := ecdsa.GenerateKey(curveP256, rand.Reader)
    if err != nil {
        panic(err)
    }
    publicShieldR = privateShieldR.Public().(*ecdsa.PublicKey)

    //generate Stealth Address
    Mrx, Mry := curveP256.ScalarMult(publicView.X, publicView.Y, privateShieldR.D.Bytes()) //rM
    HashP := sha256.Sum256(append(Mrx.Bytes(), Mry.Bytes()...)) //H(rM)
    HGx, HGy := curveP256.ScalarBaseMult(HashP[:]) //H(rM)G
    publicShieldP.X, publicShieldP.Y = curveP256.Add(HGx, HGy, publicSpend.X, publicSpend.Y) //H(rM)G+N

    return
}

func VerifyShieldAddr(privateView *ecdsa.PrivateKey, publicSpend, publicShieldP, publicShieldR *ecdsa.PublicKey) (isAddress bool) {
    //generate Stealth Address
    mRx, mRy := curveP256.ScalarMult(publicShieldR.X, publicShieldR.Y, privateView.D.Bytes()) //mR
    HashP := sha256.Sum256(append(mRx.Bytes(), mRy.Bytes()...)) //H(mR)
    HGx, HGy := curveP256.ScalarBaseMult(HashP[:]) //H(mR)G
    verifyShieldPX, verifyShieldPY := curveP256.Add(HGx, HGy, publicSpend.X, publicSpend.Y) //H(rM)G+N

    isAddress = false
    if publicShieldP.X.Cmp(verifyShieldPX) == 0 && publicShieldP.Y.Cmp(verifyShieldPY) == 0 {
        isAddress = true
    }

    return
}

func GetprivateStealth(privateView, privateSpend *ecdsa.PrivateKey, publicShieldR *ecdsa.PublicKey) (privateStealth *ecdsa.PrivateKey) {
    mRx, mRy := curveP256.ScalarMult(publicShieldR.X, publicShieldR.Y, privateView.D.Bytes()) //mR
    HashP := sha256.Sum256(append(mRx.Bytes(), mRy.Bytes()...)) //H(mR)

    privateStealth = new(ecdsa.PrivateKey)
    privateStealth.D = new(big.Int).Add(new(big.Int).SetBytes(HashP[:]), privateSpend.D) //H(mR)+n
    privateStealth.PublicKey.Curve = curveP256
    privateStealth.PublicKey.X, privateStealth.PublicKey.Y = curveP256.ScalarBaseMult(privateStealth.D.Bytes()) //(H(mR)+n)G

    return
}

//generate public and private key pairs
var curveP256 = elliptic.P256() // 椭圆曲线参数,公共参数
