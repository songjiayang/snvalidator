package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSN(t *testing.T) {
	assertion := assert.New(t)

	validSN := "513826199104253418"
	isValid, err := IsValidSN(validSN)
	assertion.Nil(err)
	assertion.True(isValid)

	invalidSN := "513826199104253417"
	isValid, err = IsValidSN(invalidSN)
	assertion.Nil(err)
	assertion.False(isValid)
}
