package set2

import "strings"

func KVParser (kvEnc string) map[string]string {
	kvs := make(map[string]string)
	pairs := strings.Split(kvEnc, "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			panic("Error when parsing key-value string")
		}
		kvs[kv[0]] = kv[1]
	}
	return kvs
}