package set1

import (
	"crypto/aes"
)

func Chal7 (enc, key []byte) []byte {
	ciph, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	decrypted := make([]byte, len(enc))
	size := len(key)

	for bs, be := 0, size; bs < len(enc); bs, be = bs + size, be + size {
		ciph.Decrypt(decrypted[bs:be], enc[bs:be])
	}
	paddingSize := int(decrypted[len(decrypted)-1])
	return decrypted[0:len(decrypted)-paddingSize]
}