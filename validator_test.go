package validator

import (
	"testing"

	"github.com/golib/assert"
)

func TestIsValidSN(t *testing.T) {
	assertion := assert.New(t)

	validSN := "513826199104253418"
	ok, err := IsValidSN(validSN)
	assertion.Nil(err)
	assertion.True(ok)
}
