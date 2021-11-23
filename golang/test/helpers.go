package edgeauth

import (
	"testing"

	edgeauth "github.com/ian-ross/EdgeAuth/golang"
)

func checkOK(t *testing.T, result *edgeauth.VerifyAndDecodeResult) *edgeauth.Token {
	if !result.Verified {
		t.Error("token failed to verify")
	}
	if result.Code != "verified" {
		t.Errorf("result Code should be 'verified', but is '%s'", result.Code)
	}
	if result.Value == nil {
		t.Error("result Value is nil, and should not be")
	}
	return result.Value
}

func checkFail(t *testing.T, result *edgeauth.VerifyAndDecodeResult, err string) {
	if result.Verified {
		t.Error("token did not fail to verify")
	}
	if result.Code != err {
		t.Errorf("result Code should be '%s', but is '%s'", err, result.Code)
	}
	if result.Message != "" {
		t.Errorf("result Message should be empty, but is '%s'", result.Message)
	}
	if result.Value != nil {
		t.Error("result Value should be nil, but is not")
	}
}

func checkToken(t *testing.T, token *string, expected string) {
	t.Run("The token matches the expected value", func(t *testing.T) {
		if *token != expected {
			t.Error("token does not match expected value")
		}
	})
}

func checkVerifyWithGoodSecret(t *testing.T, token *string) *edgeauth.Token {
	var retval *edgeauth.Token
	t.Run("The token successfully verifies with the correct secret", func(t *testing.T) {
		result := edgeauth.VerifyAndDecode("my-secret", *token)
		retval = checkOK(t, result)
	})
	return retval
}

func checkField(t *testing.T, value *edgeauth.Token, expectedField string, expectedValue string) {
	t.Run("The token successfully verifies with the correct secret", func(t *testing.T) {
		check, exists := value.Get(expectedField)
		if !exists || check != expectedValue {
			t.Errorf("required %s in value does not match", expectedField)
		}
	})
}

func checkArrayField(t *testing.T, value *edgeauth.Token, expectedField string, expectedValue []string) {
	t.Run("The token successfully verifies with the correct secret", func(t *testing.T) {
		check, exists := value.Get(expectedField)
		if !exists {
			t.Errorf("required %s in value does not match", expectedField)
		}
		values, ok := check.([]interface{})
		if !ok || len(values) != len(expectedValue) {
			t.Errorf("required %s in value does not match", expectedField)
		}
		for i := range values {
			s, ok := values[i].(string)
			if !ok || s != expectedValue[i] {
				t.Errorf("required %s in value does not match", expectedField)
			}
		}
	})
}

func checkVerifyWithBadSecret(t *testing.T, token *string) {
	t.Run("The token fails to verify with a bad secret", func(t *testing.T) {
		result := edgeauth.VerifyAndDecode("bad-secret", *token)
		checkFail(t, result, "bad-digest")
	})
}
