
package cache

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/karmadon/nano-service/internal/pkg/nano-service_lib/managers/cache/cache_common"
)

//GetCacheTagName returns Prefix with ID in one string
func GetCacheTagName(prefix cache_common.TagPrefix, id string) string {
	return string(prefix) + ":" + id
}

func GetCacheTagStr(prefix string, id string) string {
	return prefix + ":" + id
}

//GetCacheTagName returns Prefix with ID in one string
func GetCacheTagNameForObject(prefix cache_common.TagPrefix, obj ...interface{}) string {
	return string(prefix) + ":" + Hash(obj)
}

//Hash creates Hash string for object
func Hash(arr interface{}) string {
	var arrBytes []byte
	jsonBytes, _ := json.Marshal(arr)
	arrBytes = append(arrBytes, jsonBytes...)

	return fmt.Sprintf("%x", md5.Sum(arrBytes))
}
