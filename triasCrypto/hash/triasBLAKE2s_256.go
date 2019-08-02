package triasHash

import (
	"encoding/hex"
	"golang.org/x/crypto/blake2s"
)

func TriasBlake2s_256(data string) string {
	hashop, _ := blake2s.New256(nil)
	hashop.Write([]byte(data))
	resultByte := hashop.Sum(nil)
	resultString := hex.EncodeToString(resultByte)

	return resultString
}
