package snvalidator

import (
	"testing"
)

func TestIsValidSN(t *testing.T) {
	validSN := "513826199104253418"
	isValid, err := IsValidSN(validSN)
	if err != nil {
		t.Errorf("%s check error should be nil", validSN)
	}
	if !isValid {
		t.Errorf("%s check result should be true", validSN)
	}

	invalidSN := "513826199104253417"
	isValid, err = IsValidSN(invalidSN)
	if err != nil {
		t.Errorf("%s check error should be nil", invalidSN)
	}
	if isValid {
		t.Errorf("%s check result should be false", invalidSN)
	}
}
