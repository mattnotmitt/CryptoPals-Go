package set2

import (
	"errors"
	"strings"
)

func VerifyPadding(pt string, size int) (string, error) {
	// Invalid padding if not divisible by block size
	if len(pt)%size != 0 {
		return "", errors.New("padded string must be multiple of blocksize")
	}
	// Get last padding byte
	padByte := pt[len(pt)-1]
	trimPt := strings.TrimRight(pt, string(padByte))
	if len(pt) - len(trimPt) != int(padByte) {
		return "", errors.New("invalid length after removing number of bytes specified in last padding char")
	}
	return trimPt, nil
}
