package writebuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WriteBuffer_SafeWithLeakyPool(t *testing.T) {
	wb := New()

	n, err := wb.Write(make([]byte, 1))
	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	assert.NotPanics(t, func() {
		n, err = wb.Write(make([]byte, InitialSize+1))
		assert.Equal(t, InitialSize+1, n)
		assert.NoError(t, err)
	})
}
