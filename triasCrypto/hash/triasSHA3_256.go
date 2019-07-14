package triasHash

import(
    "golang.org/x/crypto/sha3"
    "encoding/hex"
)

func TriasSHA3_256(data string) string {
    hashop := sha3.New256()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
