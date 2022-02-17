package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateJWT(t *testing.T) {
	actual, err := GenerateJWT()

	assert.Equal(t, 163, len(actual))
	assert.Nil(t, err)
}
