package hackonpkg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

func TestOriginStealthAdressGenerate(test *testing.T) {
	//global params
	var curveP256 = elliptic.P256()
	//generate public and private key pairs
	var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)
}

func TestOriginStealthAdressVerify(test *testing.T) {
	//global params
	var curveP256 = elliptic.P256()
	//generate public and private key pairs
	var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)

	P, r := OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)
	OriginStealthAdressVerify(privateKeyA, privateKeyB, P, r, curveP256)
}

func BenchmarkOriginStealthAdressGenerate(b *testing.B) {
	//global params
	var curveP256 = elliptic.P256()
	//generate public and private key pairs
	var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)

	// 重置计时器
	b.ResetTimer()
	// 停止计时器
	b.StopTimer()
	// 开始计时器
	b.StartTimer()
	var n int
	for i := 0; i < b.N; i++ {
		n++
		OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)
	}
}

func BenchmarkOriginStealthAdressVerify(b *testing.B) {
	//global params
	var curveP256 = elliptic.P256()
	//generate public and private key pairs
	var privateKeyA, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	var privateKeyB, _ = ecdsa.GenerateKey(curveP256, rand.Reader)
	P, r := OriginStealthAdressGenerate(privateKeyA, privateKeyB, curveP256)

	// 重置计时器
	b.ResetTimer()
	// 停止计时器
	b.StopTimer()
	// 开始计时器
	b.StartTimer()
	var n int
	for i := 0; i < b.N; i++ {
		n++
		OriginStealthAdressVerify(privateKeyA, privateKeyB, P, r, curveP256)
	}
}
