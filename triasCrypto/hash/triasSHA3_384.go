package triasHash

import(
    "golang.org/x/crypto/sha3"
    "encoding/hex"
)

func TriasSHA3_384(data string) string {
    hashop := sha3.New384()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
