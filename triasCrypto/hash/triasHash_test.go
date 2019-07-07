package triasHash

import(
    "fmt"
    "testing"
)

var testString string = "Trias Hash Test String"


func TestTriasMd5(test *testing.T) {
    fmt.Println("Md5      ", TriasMd5(testString))
}

func TestTriasSha1(test *testing.T) {
    fmt.Println("sha1     ", TriasSha1(testString))
}

func TestTriasSha224(test *testing.T) {
    fmt.Println("sha256224", TriasSha224(testString))
}

func TestTriasSha256(test *testing.T) {
    fmt.Println("sha256256", TriasSha256(testString))
}

func TestTriasSha512224(test *testing.T) {
    fmt.Println("sha512224", TriasSha512224(testString))
}

func TestTriasSha512256(test *testing.T) {
    fmt.Println("sha512256", TriasSha512256(testString))
}

func TestTriasSha512384(test *testing.T) {
    fmt.Println("sha512384", TriasSha512384(testString))
}

func TestTriasSha512512(test *testing.T) {
    fmt.Println("sha512512", TriasSha512512(testString))
}
