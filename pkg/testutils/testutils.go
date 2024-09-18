package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMessageContains(t *testing.T, err error, expectedMessages ...string) {
	t.Helper()
	assert.Error(t, err)
	for _, msg := range expectedMessages {
		assert.Contains(t, err.Error(), msg)
	}
}
