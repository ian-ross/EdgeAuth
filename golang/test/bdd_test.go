package edgeauth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cucumber/godog"
	edgeauth "github.com/ian-ross/EdgeAuth/golang"
)

// --  STATE  ------------------------------------------------------------------

var token *string
var correctToken *string
var builder *edgeauth.TokenBuilder
var result *edgeauth.VerifyAndDecodeResult

// --  STEPS  ------------------------------------------------------------------

// INITIALIZATION STEPS

func iHaveABadToken() error {
	tmp := "DIGEST:bad-token"
	token = &tmp
	return nil
}

func iHaveAGoodToken() error {
	builder = edgeauth.NewTokenBuilder()
	builder = builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000))
	return nil
}

func iHaveAGoodTokenWithURI(arg1 string) error {
	builder = edgeauth.NewTokenBuilder()
	builder = builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		WithURI(arg1).
		ExpiresAt(time.UnixMilli(1000))
	return nil
}

// VERIFICATION STEPS

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

// TOKEN SETUP STEPS

func theTokenIsForAChannel(arg1 string) error {
	builder = builder.ForChannel(arg1)
	return nil
}

func theTokenIsForAChannelAlias(arg1 string) error {
	builder = builder.ForChannelAlias(arg1)
	return nil
}

func theTokenIsForARoom(arg1 string) error {
	builder = builder.ForRoom(arg1)
	return nil
}

func theTokenIsForARoomAlias(arg1 string) error {
	builder = builder.ForRoomAlias(arg1)
	return nil
}

func theTokenIsForARemoteAddress(arg1 string) error {
	builder = builder.ForRemoteAddress(arg1)
	return nil
}

func theTokenIsForASession(arg1 string) error {
	builder = builder.ForSession(arg1)
	return nil
}

func theTokenIsForStreamingOnly() error {
	builder = builder.ForStreamingOnly()
	return nil
}

func theTokenIsForPublishingOnly() error {
	builder = builder.ForPublishingOnly()
	return nil
}

func theTokenIsForTag(arg1 string) error {
	builder = builder.ForTag(arg1)
	return nil
}

func theTokenHasATagApplied(arg1 string) error {
	builder = builder.ApplyTag(arg1)
	return nil
}

func theTokenHasCapability(arg1 string) error {
	builder = builder.WithCapability(arg1)
	return nil
}

// FIELD TESTING STEPS

func theRemoteAddressFieldShouldBe(arg1 string) error {
	return bddCheckField(edgeauth.RemoteAddressField, arg1)
}

func theURIFieldShouldBe(arg1 string) error {
	return bddCheckField(edgeauth.URIField, arg1)
}

func theSessionFieldShouldBe(arg1 string) error {
	return bddCheckField(edgeauth.SessionIDField, arg1)
}

func theTypeFieldShouldBe(arg1 string) error {
	return bddCheckField(edgeauth.TypeField, arg1)
}

func theTagFieldShouldBe(arg1 string) error {
	return bddCheckField(edgeauth.RequiredTagField, arg1)
}

func theAppliedTagsFieldShouldBe(arg1 string) error {
	return bddCheckArrayField(edgeauth.ApplyTagsField, arg1)
}

func theCapabilitiesFieldShouldBe(arg1 string) error {
	return bddCheckArrayField(edgeauth.CapabilitiesField, arg1)
}

// RESULT CHECKING STEPS

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

// --  INITIALIZATION  ---------------------------------------------------------

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		token = nil
		correctToken = nil
		builder = nil
		result = nil
		return ctx, nil
	})

	ctx.Step(`^I have a bad token$`, iHaveABadToken)
	ctx.Step(`^I have a good token$`, iHaveAGoodToken)
	ctx.Step(`^I have a good token with URI "([^"]*)"$`, iHaveAGoodTokenWithURI)
	ctx.Step(`^I try to verify it with a good secret$`, iTryToVerifyItWithAGoodSecret)
	ctx.Step(`^I try to verify it with a bad secret$`, iTryToVerifyItWithABadSecret)
	ctx.Step(`^The correct token is "([^"]*)"$`, theCorrectTokenIs)
	ctx.Step(`^The token is for a channel "([^"]*)"$`, theTokenIsForAChannel)
	ctx.Step(`^The token is for a channel alias "([^"]*)"$`, theTokenIsForAChannelAlias)
	ctx.Step(`^The token is for a room "([^"]*)"$`, theTokenIsForARoom)
	ctx.Step(`^The token is for a room alias "([^"]*)"$`, theTokenIsForARoomAlias)
	ctx.Step(`^The token is for a remote address "([^"]*)"$`, theTokenIsForARemoteAddress)
	ctx.Step(`^The token is for a session "([^"]*)"$`, theTokenIsForASession)
	ctx.Step(`^The token is for tag "([^"]*)"$`, theTokenIsForTag)
	ctx.Step(`^The token has a "([^"]*)" tag applied$`, theTokenHasATagApplied)
	ctx.Step(`^The token is for streaming only$`, theTokenIsForStreamingOnly)
	ctx.Step(`^The token is for publishing only$`, theTokenIsForPublishingOnly)
	ctx.Step(`^The token has capability "([^"]*)"$`, theTokenHasCapability)
	ctx.Step(`^Verification should fail with error "([^"]*)"$`, verificationShouldFailWithError)
	ctx.Step(`^Verification should pass$`, verificationShouldPass)
	ctx.Step(`^The tag field should be "([^"]*)"$`, theTagFieldShouldBe)
	ctx.Step(`^The remote address field should be "([^"]*)"$`, theRemoteAddressFieldShouldBe)
	ctx.Step(`^The session field should be "([^"]*)"$`, theSessionFieldShouldBe)
	ctx.Step(`^The URI field should be "([^"]*)"$`, theURIFieldShouldBe)
	ctx.Step(`^The applied tags field should be "([^"]*)"$`, theAppliedTagsFieldShouldBe)
	ctx.Step(`^The type field should be "([^"]*)"$`, theTypeFieldShouldBe)
	ctx.Step(`^The capabilities field should be "([^"]*)"$`, theCapabilitiesFieldShouldBe)
}

// --  HELPERS  ----------------------------------------------------------------

func buildToken() error {
	if builder != nil && token == nil {
		var err error
		token, err = builder.Build()
		if err != nil {
			return fmt.Errorf("token builder failed: %v", err)
		}
	}
	return nil
}

func bddCheckField(expectedField string, expectedValue string) error {
	if result == nil || result.Value == nil {
		return fmt.Errorf("the verification value is not set")
	}
	check, exists := result.Value.Get(expectedField)
	if !exists || check != expectedValue {
		return fmt.Errorf("required %s in value does not match", expectedField)
	}
	return nil
}

func bddCheckArrayField(expectedField string, expectedValues string) error {
	if result == nil || result.Value == nil {
		return fmt.Errorf("the verification value is not set")
	}
	expectedValue := strings.Split(expectedValues, ",")
	check, exists := result.Value.Get(expectedField)
	if !exists {
		return fmt.Errorf("required %s in value does not match", expectedField)
	}
	values, ok := check.([]interface{})
	if !ok || len(values) != len(expectedValue) {
		return fmt.Errorf("required %s in value does not match", expectedField)
	}
	for i := range values {
		s, ok := values[i].(string)
		if !ok || s != expectedValue[i] {
			return fmt.Errorf("required %s in value does not match", expectedField)
		}
	}
	return nil
}
