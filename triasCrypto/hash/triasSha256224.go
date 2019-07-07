package triasHash

import(
    "crypto/sha256"
    "encoding/hex"
)

func TriasSha224(data string) string {
    hashop := sha256.New224()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
