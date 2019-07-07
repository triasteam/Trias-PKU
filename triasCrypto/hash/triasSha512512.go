package triasHash

import(
    "crypto/sha512"
    "encoding/hex"
)

func TriasSha512512(data string) string {
    hashop := sha512.New()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
