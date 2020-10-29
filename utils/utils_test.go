package utils_test

import (
	"testing"

	"github.com/incrypt0/cokut-server/utils"
)

func Test(t *testing.T) {
	t.Log("hmm")
	t.Log(utils.Distance(10.3476757, 76.2071317, 10.5113799, 76.1532094) * 1.2)
}
