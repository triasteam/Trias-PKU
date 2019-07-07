package triasHash

import(
    "crypto/sha512"
    "encoding/hex"
)

func TriasSha512384(data string) string {
    hashop := sha512.New384()
    hashop.Write([]byte(data))
    resultByte := hashop.Sum(nil)
    resultString := hex.EncodeToString(resultByte)

    return resultString
}
