package triasHash

import(
    "golang.org/x/crypto/ripemd160"
    "encoding/hex"
)

func TriasRipemd160(data string) string {
    hashop := ripemd160.New()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
