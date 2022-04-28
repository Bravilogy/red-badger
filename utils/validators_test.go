package utils

import (
	"testing"
)

func TestValidateInputFormat(t *testing.T) {
	input := `5 3`
	_, err := ValidateInputFormat(input)
	if err.Error() == invalidInputFormatError.Error() {
		t.Errorf("ValidateInputFormat(%s); want invalidInputFormatError; got %s", input, err)
	}
}
