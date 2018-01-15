package main

import (
	"crypto/sha1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit2(t *testing.T) {
	input := "some random string"
	hash := sha1.New()
	hash.Write([]byte(input))
	secret := hash.Sum(nil)

	parts, err := Split2(secret[:])
	assert.NoError(t, err)

	restored := Combine(parts)
	assert.Equal(t, secret, restored)

}
