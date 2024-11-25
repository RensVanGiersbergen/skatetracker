package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrue_IsTrue(t *testing.T) {
	// Arrange
	var trueValue = true
	// Act
	// Assert
	assert.Equal(t, true, trueValue)
}
