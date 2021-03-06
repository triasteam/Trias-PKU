package triasHash

import (
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
)

func TriasBlake2b_384(data string) string {
	hashop, _ := blake2b.New384(nil)
	hashop.Write([]byte(data))
	resultByte := hashop.Sum(nil)
	resultString := hex.EncodeToString(resultByte)

	return resultString
}
