package randutils

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRandString(t *testing.T) {
	assertor := assert.New(t)

	tests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

	for _, testLen := range tests {
		name := strconv.FormatInt(int64(testLen), 10)
		name = "Len:" + name
		t.Run(name, func(t *testing.T) {
			randString := RandString(testLen)
			assertor.Equal(len(randString), testLen, testLen)
		})
	}

}
