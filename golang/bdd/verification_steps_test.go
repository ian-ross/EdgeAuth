package edgeauth

import (
	"fmt"

	edgeauth "github.com/ian-ross/EdgeAuth/golang"
)

func iTryToVerifyItWithABadSecret() error {
	err := buildToken()
	if err != nil {
		return err
	}
	result = edgeauth.VerifyAndDecode("bad-secret", *token)
	return nil
}

func iTryToVerifyItWithAGoodSecret() error {
	err := buildToken()
	if err != nil {
		return err
	}
	if token != nil && correctToken != nil && *token != *correctToken {
		return fmt.Errorf("token does not match expected value")
	}
	result = edgeauth.VerifyAndDecode("my-secret", *token)
	return nil
}

func theCorrectTokenIs(arg1 string) error {
	correctToken = &arg1
	return nil
}
