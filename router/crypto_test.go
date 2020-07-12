package router

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSHA256(t *testing.T) {
	assert.Equal(t, "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b", SHA256("1"))
	assert.Equal(t, "d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35", SHA256("2"))
}
