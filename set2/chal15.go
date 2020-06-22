package set2

import (
	"CryptoPals/util"
	"errors"
	"strconv"
	"strings"
)

func verifyPadding(pt string, size int) (string, error) {
	// Invalid padding if not divisible by block size
	if len(pt)%size != 0 {
		return "", errors.New("padded string must be multiple of blocksize")
	}
	// Check if all runes are printable
	allPrintable := util.StringEvery(pt, func(r rune) bool {
		return strconv.IsPrint(r)
	})
	// Return if all runes printable
	if allPrintable {
		return pt, nil
	}
	trimPt := strings.TrimRight(pt, "\x04")
	trimPrintable := util.StringEvery(trimPt, func(r rune) bool {
		return strconv.IsPrint(r)
	})
	// If all characters printable after padding trimmed, valid
	if trimPrintable {
		return trimPt, nil
	}
	return "", errors.New("non-printable characters other than padding in text")
}
