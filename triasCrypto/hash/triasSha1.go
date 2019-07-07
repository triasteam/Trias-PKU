package triasHash

import(
    "crypto/sha1"
    "encoding/hex"
)

func TriasSha1(data string) string {
    hashop := sha1.New()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
