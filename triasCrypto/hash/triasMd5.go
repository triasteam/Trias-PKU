package triasHash

import (
	"crypto/md5"
	"encoding/hex"
)

func TriasMd5(data string) string {
	hashop := md5.New()
	hashop.Write([]byte(data))
	resultByte := hashop.Sum(nil)
	resultString := hex.EncodeToString(resultByte)

	return resultString
}
