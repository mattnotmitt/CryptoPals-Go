package util

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type VerifyPaddingSuite struct {
	suite.Suite
}

func (suite *VerifyPaddingSuite) TestDefault() {
	// Test padding stripping
	input0 := "ICE ICE BABY\x04\x04\x04\x04"
	expected0 := "ICE ICE BABY"
	resp0, err0 := VerifyPadding(input0, 16)
	suite.Nil(err0)
	suite.Equal(expected0, resp0)
	// Test valid non-padded string
	input1 := "ICE ICE BABY...."
	expected1 := "ICE ICE BABY...."
	resp1, err1 := VerifyPadding(input1, 16)
	suite.Nil(err1)
	suite.Equal(expected1, resp1)
}

func (suite *VerifyPaddingSuite) TestInvalidPadChars() {
	input0 := "ICE ICE BABY\x05\x05\x05\x05"
	input1 := "ICE ICE BABY\x01\x02\x03\x04"
	_, err0 := VerifyPadding(input0, 16)
	if suite.NotNil(err0) {
		suite.Equal(err0.Error(), "non-printable characters other than padding in text")
	}
	_, err1 := VerifyPadding(input1, 16)
	if suite.NotNil(err1) {
		suite.Equal(err1.Error(), "non-printable characters other than padding in text")
	}
}

func (suite *VerifyPaddingSuite) TestShortPadding() {
	input := "ICE ICE BABY\x04\x04\x04"
	_, err := VerifyPadding(input, 16)
	if suite.NotNil(err) {
		suite.Equal(err.Error(), "padded string must be multiple of blocksize")
	}
}

func TestUtils(t *testing.T) {
	suite.Run(t, new(VerifyPaddingSuite))
}
