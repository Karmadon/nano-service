package cache_common

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

func Hash(arr interface{}) string {
	var arrBytes []byte
	jsonBytes, _ := json.Marshal(arr)
	arrBytes = append(arrBytes, jsonBytes...)

	return fmt.Sprintf("%x", md5.Sum(arrBytes))
}
