package triasHash

import (
	"fmt"
	"testing"
)

var testString string = "Trias Hash Test String"

func TestTriasMd5(test *testing.T) {
	fmt.Println("Md5      ", TriasMd5(testString))
}

func BenchmarkTriasMd5(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasMd5(testString)
	}
}

func TestTriasSha1(test *testing.T) {
	fmt.Println("sha1     ", TriasSha1(testString))
}

func BenchmarkTriasSha1(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSha1(testString)
	}
}

func TestTriasSha224(test *testing.T) {
	fmt.Println("sha256224", TriasSha224(testString))
}

func BenchmarkTriasSha224(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSha224(testString)
	}
}

func TestTriasSha256(test *testing.T) {
	fmt.Println("sha256256", TriasSha256(testString))
}

func BenchmarkTriasSha256(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSha256(testString)
	}
}

func TestTriasSha512224(test *testing.T) {
	fmt.Println("sha512224", TriasSha512224(testString))
}

func BenchmarkTriasSha512224(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSha512224(testString)
	}
}

func TestTriasSha512256(test *testing.T) {
	fmt.Println("sha512256", TriasSha512256(testString))
}

func BenchmarkTriasSha512256(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSha512256(testString)
	}
}

func TestTriasSha512384(test *testing.T) {
	fmt.Println("sha512384", TriasSha512384(testString))
}

func BenchmarkTriasSha512384(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSha512384(testString)
	}
}

func TestTriasSha512512(test *testing.T) {
	fmt.Println("sha512512", TriasSha512512(testString))
}

func BenchmarkTriasSha512512(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSha512512(testString)
	}
}

func TestTriasRipemd160(test *testing.T) {
	fmt.Println("TriasRipemd160", TriasRipemd160(testString))
}

func BenchmarkTriasRipemd160(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasRipemd160(testString)
	}
}

func TestTriasSHA3_224(test *testing.T) {
	fmt.Println("TriasSHA3_224", TriasSHA3_224(testString))
}

func BenchmarkTriasSHA3_224(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSHA3_224(testString)
	}
}

func TestTriasSHA3_256(test *testing.T) {
	fmt.Println("TriasSHA3_256", TriasSHA3_256(testString))
}

func BenchmarkTriasSHA3_256(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSHA3_256(testString)
	}
}

func TestTriasSHA3_384(test *testing.T) {
	fmt.Println("TriasSHA3_384", TriasSHA3_384(testString))
}

func BenchmarkTriasSHA3_384(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSHA3_384(testString)
	}
}

func TestTriasSHA3_512(test *testing.T) {
	fmt.Println("TriasSHA3_512", TriasSHA3_512(testString))
}

func BenchmarkTriasSHA3_512(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasSHA3_512(testString)
	}
}

func TestTriasBlake2s_256(test *testing.T) {
	fmt.Println("TriasBlake2s_256", TriasBlake2s_256(testString))
}

func BenchmarkTriasBlake2s_256(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasBlake2s_256(testString)
	}
}

func TestTriasBlake2b_256(test *testing.T) {
	fmt.Println("TriasBlake2b_256", TriasBlake2b_256(testString))
}

func BenchmarkTriasBlake2b_256(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasBlake2b_256(testString)
	}
}

func TestTriasBlake2b_384(test *testing.T) {
	fmt.Println("TriasBlake2b_384", TriasBlake2b_384(testString))
}

func BenchmarkTriasBlake2b_384(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasBlake2b_384(testString)
	}
}

func TestTriasBlake2b_512(test *testing.T) {
	fmt.Println("TriasBlake2b_512", TriasBlake2b_512(testString))
}

func BenchmarkTriasBlake2b_512(b *testing.B) {
	//reset timer
	b.ResetTimer()
	//stop timer
	b.StopTimer()
	//start timer
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		TriasBlake2b_512(testString)
	}
}
