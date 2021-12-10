package edgeauth

import (
	"fmt"
)

func verificationShouldFailWithError(arg1 string) error {
	if result.Verified {
		return fmt.Errorf("token did not fail to verify")
	}
	if result.Code != arg1 {
		return fmt.Errorf("result Code should be '%s', but is '%s'", arg1, result.Code)
	}
	if result.Message != "" {
		return fmt.Errorf("result Message should be empty, but is '%s'", result.Message)
	}
	if result.Value != nil {
		return fmt.Errorf("result Value should be nil, but is not")
	}
	return nil
}

func verificationShouldPass() error {
	if !result.Verified {
		return fmt.Errorf("token failed to verify")
	}
	if result.Code != "verified" {
		return fmt.Errorf("result Code should be 'verified', but is '%s'", result.Code)
	}
	if result.Value == nil {
		return fmt.Errorf("result Value is nil, and should not be")
	}
	return nil
}
