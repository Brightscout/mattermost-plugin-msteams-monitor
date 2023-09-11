package store

import (
	"crypto/sha512"
	"fmt"
)

// TODO: remove below line later
// nolint:all
func hashKey(prefix, hashableKey string) string {
	if hashableKey == "" {
		return prefix
	}

	h := sha512.New()
	_, _ = h.Write([]byte(hashableKey))
	return fmt.Sprintf("%s%x", prefix, h.Sum(nil))
}
