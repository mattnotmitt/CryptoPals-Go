package set2

import (
	"CryptoPals/util"
	"regexp"
	"sort"
	"strings"
	"sync"
)

func kvParser(kvStr string) map[string]string {
	kvMap := make(map[string]string)
	pairs := strings.Split(kvStr, "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			panic("Error when parsing key-value string")
		}
		kvMap[kv[0]] = kv[1]
	}
	return kvMap
}

func kvEncoder(kvMap map[string]string, profile bool) string {
	kvStr := ""
	// Iterating through maps in go is *deliberately* random,
	// so we sort the keys alphabetically so the same output
	// is received each time
	if profile {
		return "email=" + kvMap["email"] + "&uid=10&role=user"
	}
	ks := make([]string, 0, len(kvMap))
	for k := range kvMap {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		kvStr += k + "=" + kvMap[k] + "&"
	}
	return strings.TrimRight(kvStr, "&")
}

func profileFor(email string) map[string]string {
	re := regexp.MustCompile(`[&=]`)
	emailClean := re.ReplaceAllString(email, "")
	profile := map[string]string{
		"email": emailClean,
		"uid":   "10",
		"role":  "user",
	}
	return profile
}

func encryptProfile(profile string, key []byte) []byte {
	profileByte := []byte(profile)
	encryptedProf := util.AESECBEncrypt(profileByte, key)
	return encryptedProf
}

var keySetup13 sync.Once
var key13 []byte

func login(email string) []byte {
	keySetup13.Do(func() { key13 = util.RandBytes(16) }) // Generate key on first run of program and persist
	// Generate profile and encode to KV string
	ptProf := kvEncoder(profileFor(email), true)
	return encryptProfile(ptProf, key13)
}

func verifyCookie(cookie []byte) map[string]string {
	kvPt := util.AESECBDecrypt(cookie, key13)
	return kvParser(string(kvPt))
}
