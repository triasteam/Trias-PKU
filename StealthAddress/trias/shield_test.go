package stealthaddress

import (
    "crypto/ecdsa"
    "crypto/rand"
    "fmt"
    "testing"
)

//generate public and private key pairs
var privateView, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
var privateSpend, _ = ecdsa.GenerateKey(curveP256, rand.Reader)

func TestCreateVerifyShieldAddr(test *testing.T) {
    publicView := privateView.Public().(*ecdsa.PublicKey)
    publicSpend := privateSpend.Public().(*ecdsa.PublicKey)
    publicShieldP, publicShieldR := CreateShieldAddr(publicView, publicSpend)
    isAddress := VerifyShieldAddr(privateView, publicSpend, publicShieldP, publicShieldR)
    privateStealth := GetprivateStealth(privateView, privateSpend, publicShieldR)

    fmt.Println("View ", privateView.D, "\n", publicView)
    fmt.Println("Spend", privateSpend.D, "\n", publicSpend)
    fmt.Println("Stealth", publicShieldP, "\nR", publicShieldR)
    fmt.Println(isAddress)
    fmt.Println("Stealth", privateStealth)
}
