package hackonpkg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

//TestOriginStealthAdressGenerate
func TestOriginStealthAdressGenerate(test *testing.T) {
	//global params
	var curveP256 = elliptic.P256()
	//generate public and private key pairs
	var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)
}

//TestOriginStealthAdressVerify
func TestOriginStealthAdressVerify(test *testing.T) {
	//global params
	var curveP256 = elliptic.P256()
	//generate public and private key pairs
	var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)

	P, r := OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)
	OriginStealthAdressVerify(privateKeyA, privateKeyB, P, r, curveP256)
}

//BenchOriginStealthAdressGenerate
func BenchmarkOriginStealthAdressGenerate(b *testing.B) {
	//global params
	var curveP256 = elliptic.P256()
	//generate public and private key pairs
	var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)

	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()
	var n int
	for i := 0; i < b.N; i++ {
		n++
		OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)
	}
}

//BenchOriginStealthAdressVerify
func BenchmarkOriginStealthAdressVerify(b *testing.B) {
	//global params
	var curveP256 = elliptic.P256()
	//generate public and private key pairs
	var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	P, r := OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)

	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()
	var n int
	for i := 0; i < b.N; i++ {
		n++
		OriginStealthAdressVerify(privateKeyA, privateKeyB, P, r, curveP256)
	}
}
