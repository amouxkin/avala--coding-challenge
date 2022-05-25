package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfOddLooking(t *testing.T) {
	tests := []struct {
		name  string
		num   int
		isOdd bool
	}{
		{"should be false for repeated", 0xAA_AAA_AAA, true},
		{"should be false for repeated", 0x131_312, false},
		{"should be false for hexSpeak", 0xdeadd00d, true},
	}

	for _, test := range tests {
		assert.Equal(t, test.isOdd, CheckIfOddLooking(test.num))
	}
}
