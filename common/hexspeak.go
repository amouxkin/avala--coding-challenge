package common

import (
	"github.com/dlclark/regexp2"
	"strconv"
	"strings"
)

var regexAlwaysIncreasing = regexp2.MustCompile(`^(?=[1234567890abcdef]{3,8}$)1*2*3*4*5*6*7*8*9*a*b*c*d*e*f*$`, 0)
var regexAlwaysDecreasing = regexp2.MustCompile(`^(?=[1234567890abcdef]{3,8}$)f*e*d*c*b*a*9*8*7*6*5*4*3*2*1*$`, 0)

func checkIsAlwaysIncreasing(hex string) bool {

	isMatch, _ := regexAlwaysIncreasing.MatchString(hex)
	return isMatch
}

func checkIsAlwaysDecreasing(hex string) bool {
	isMatch, _ := regexAlwaysDecreasing.MatchString(hex)
	return isMatch
}

func checkIfHexSpeak(hex string) bool {
	for _, hexSpeak := range hexSpeakList {
		if strings.Contains(hex, hexSpeak) {
			return true
		}
	}
	return false
}

func CheckIfOddLooking(num int) bool {
	var bs []byte
	hex := string(strconv.AppendUint(bs, uint64(num), 16))
	paddings := "00000000"
	hex = paddings[len(hex):] + hex
	return checkIsAlwaysIncreasing(hex) || checkIsAlwaysDecreasing(hex) || checkIfHexSpeak(hex)
}
