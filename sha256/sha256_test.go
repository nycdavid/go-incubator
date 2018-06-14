package sha256

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestSha256_returnsHash(t *testing.T) {
	h := sha256.New()
	fmt.Println(h)
}
