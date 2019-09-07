package cryptoEncrypt

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/sha256"
    "crypto/rand"
)

func triasDHShareKey(publicPartner *ecdsa.PublicKey, privateLocal *ecdsa.PrivateKey) []byte {
    Mrx, Mry := curveP256.ScalarMult(publicPartner.X, publicPartner.Y, privateLocal.D.Bytes())
    HashP := sha256.Sum256(append(Mrx.Bytes(), Mry.Bytes()...))

    return HashP[:]
}

func triasDHEncrypt(publicPartner *ecdsa.PublicKey, plaintext []byte) (ciphertext []byte, publicLocal *ecdsa.PublicKey) {
    var privateLocal, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
    publicLocal = privateLocal.Public().(*ecdsa.PublicKey)

    secretKey := triasDHShareKey(publicPartner, privateLocal)
    ciphertext = triasCTREncrypt(secretKey, plaintext)

    return
}

func triasDHDecrypt(publicPartner *ecdsa.PublicKey, ciphertext []byte, privateLocal *ecdsa.PrivateKey) (plaintext []byte) {
    secretKey := triasDHShareKey(publicPartner, privateLocal)
    plaintext = triasCTRDecrypt(secretKey, ciphertext)

    return
}

var curveP256 = elliptic.P256() // 椭圆曲线参数,公共参数
