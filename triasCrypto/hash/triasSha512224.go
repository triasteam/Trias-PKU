package triasHash

import(
    "crypto/sha512"
    "encoding/hex"
)

func TriasSha512224(data string) string {
    hashop := sha512.New512_224()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
