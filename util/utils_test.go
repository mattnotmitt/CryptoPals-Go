package util

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AESCTRSuite struct {
	suite.Suite
}

func (suite *AESCTRSuite) TestCTREncDec() {
	key := []byte("YELLOW SUBMARINE")
	nonce := bytes.Repeat([]byte("\x01"), 8)
	pt := []byte("The quick brown fox jumps over the lazy dog.")
	enc, err := AESCTREncrypt(pt, key, nonce)
	suite.Nil(err)
	dec, err := AESCTRDecrypt(enc, key, nonce)
	suite.Nil(err)
	suite.Equal(pt, dec)
}

func (suite *AESCTRSuite) TestCTRShortNonce() {
	key := []byte("YELLOW SUBMARINE")
	nonce := bytes.Repeat([]byte("\x01"), 7)
	pt := []byte("The quick brown fox jumps over the lazy dog.")
	_, err := AESCTREncrypt(pt, key, nonce)
	if suite.NotNil(err) {
		suite.Equal(err.Error(), "random nonce must be half the keysize")
	}
	_, err = AESCTRDecrypt(pt, key, nonce)
	if suite.NotNil(err) {
		suite.Equal(err.Error(), "random nonce must be half the keysize")
	}

}

func TestUtils(t *testing.T) {
	suite.Run(t, new(AESCTRSuite))
}
