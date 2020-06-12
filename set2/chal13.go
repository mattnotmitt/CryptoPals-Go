package set2

import (
	"CryptoPals/util"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"sync"
)

func genKey () {
	key := util.RandBytes(16)
	err :=  ioutil.WriteFile("data/13.asc", key, 0644)
	util.Check(err)
}

func KVParser (kvStr string) map[string]string {
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

func KVEncoder (kvMap map[string]string, profile bool) string {
	kvStr := ""
	// Iterating through maps in go is *deliberately* random,
	// so we sort the keys alphabetically so the same output
	// is received each time
	if profile {
		return "email=" + kvMap["email"] + "&uid=10&role=user"
	}
	ks := make([]string, 0, len(kvMap))
	for k := range kvMap  {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for _, k := range ks {
		kvStr += k + "=" + kvMap[k] + "&"
	}
	return strings.TrimRight(kvStr, "&")
}

func ProfileFor (email string) map[string]string {
	re := regexp.MustCompile(`[&=]`)
	emailClean := re.ReplaceAllString(email, "")
	profile := map[string]string{
		"email": emailClean,
		"uid": "10",
		"role": "user",
	}
	return profile
}


func EncryptProfile (profile string, key []byte) []byte {
	profileByte := []byte(profile)
	encryptedProf := util.AESECBEncrypt(profileByte, key)
	return encryptedProf
}

var keySetup sync.Once
func Login (email string) []byte {
	keySetup.Do(func() { genKey() }) // Generate key on first run of program and persist
	// Load key from file
	key, err := ioutil.ReadFile("data/13.asc")
	util.Check(err)
	// Generate profile and encode to KV string
	ptProf := KVEncoder(ProfileFor(email), true)
	return EncryptProfile(ptProf, key)
}

func VerifyCookie (cookie []byte) map[string]string {
	key, err := ioutil.ReadFile("data/13.asc")
	util.Check(err)
	kvPt := util.AESECBDecrypt(cookie, key)
	return KVParser(string(kvPt))
}