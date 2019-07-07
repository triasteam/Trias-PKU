package triasHash

import(
    "crypto/sha256"
    "encoding/hex"
)

func TriasSha256(data string) string {
    hashop := sha256.New()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
