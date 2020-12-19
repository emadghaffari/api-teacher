package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	str := Generate(10)
	assert.Equal(t, 10, len(str))
}
