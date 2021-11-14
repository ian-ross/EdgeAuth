package edgeauth

import (
	"testing"
	"time"
)

//- HELPERS --------------------------------------------------------------------

func checkOK(t *testing.T, result *VerifyAndDecodeResult) *Token {
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

func checkFail(t *testing.T, result *VerifyAndDecodeResult, err string) {
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

func checkVerifyWithGoodSecret(t *testing.T, token *string) *Token {
	var retval *Token
	t.Run("The token successfully verifies with the correct secret", func(t *testing.T) {
		result := VerifyAndDecode("my-secret", *token)
		retval = checkOK(t, result)
	})
	return retval
}

func checkField(t *testing.T, value *Token, expectedField string, expectedValue string) {
	t.Run("The token successfully verifies with the correct secret", func(t *testing.T) {
		check, exists := value.Get(expectedField)
		if !exists || check != expectedValue {
			t.Errorf("required %s in value does not match", expectedField)
		}
	})
}

func checkArrayField(t *testing.T, value *Token, expectedField string, expectedValue []string) {
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
		result := VerifyAndDecode("bad-secret", *token)
		checkFail(t, result, "bad-digest")
	})
}

//------------------------------------------------------------------------------

func TestWhenVerifyingABadToken(t *testing.T) {
	token := "DIGEST:bad-token"

	t.Run("The token fails to verify", func(t *testing.T) {
		result := VerifyAndDecode("bad-secret", token)
		checkFail(t, result, "bad-token")
	})
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForAChannelAliasAndRemoteAddress(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForChannelAlias("my-channel").
		ForRemoteAddress("10.1.2.3").
		ForStreamingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiI4MitYd1dITVRUc0xWYThKcnFPUmdjYlRXL2g2clFBTlF1MjgvRytQeHllQ09qSHEyb2xDYzVacUJ1MktqN0tGYmYyTC84TDZyaE9xTTZPMjNBR29HUT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwicmVtb3RlQWRkcmVzc1wiOlwiMTAuMS4yLjNcIixcInR5cGVcIjpcInN0cmVhbVwifSJ9")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, RequiredTagField, "channelAlias:my-channel")
	checkField(t, value, RemoteAddressField, "10.1.2.3")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForAChannelAliasAndSession(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForChannelAlias("my-channel").
		ForSession("session-id").
		ForStreamingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJBQi9Nanp2a1lnMGRTODF6aU1SVDZ3OUtwWmtjMU42U3VMTW56V09CQVJQZWJuenRHZTlmM2ZNS1FURXVqaHpVTkY0TWVsNkpMekFiWlZ3TFBSbEN4QT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwic2Vzc2lvbklkXCI6XCJzZXNzaW9uLWlkXCIsXCJ0eXBlXCI6XCJzdHJlYW1cIn0ifQ==")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, RequiredTagField, "channelAlias:my-channel")
	checkField(t, value, SessionIDField, "session-id")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForAChannelAliasAndWithATagAdded(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForChannelAlias("my-channel").
		ForStreamingOnly().
		ApplyTag("customer1").
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJMU0VnS2dGTy9aRUdxdEFLazVZb0F6cFJuTnQ4enhwUjNsdEJ3cWtOR3E1VWdjWWZpcnZKTDk3NHhpangyNS9XbHpqaUg1dk5ZMHdaYklFSkE2MzJqdz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwidHlwZVwiOlwic3RyZWFtXCIsXCJhcHBseVRhZ3NcIjpbXCJjdXN0b21lcjFcIl19In0=")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, RequiredTagField, "channelAlias:my-channel")
	checkArrayField(t, value, ApplyTagsField, []string{"customer1"})
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForAChannelAlias(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForChannelAlias("my-channel").
		ForStreamingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJPMk90R1ZBMlErTGlhRkdjSjZ0cnlXZWE4L2l2dWFQR2gzcFJpcVd3ZlJPVWdBSSs0dFdaYXdBc011Y2MyMHNRTjZpaGZtVGVDNFVubXVoWko5aHBxUT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwidHlwZVwiOlwic3RyZWFtXCJ9In0=")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, RequiredTagField, "channelAlias:my-channel")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForAChannel(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForChannel("us-northeast#my-application-id#my-channel.134566").
		ForStreamingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJZNGM3Tmp6eDVhalkzLzRWK3pwTVliNTBBU1ZCUXc0NlAvS2dwc3JrTnpDdFAzZWM5NzVzblorN3lJNzZiM0wrTmNtb2FoL3hOTUhQZ00vNEExaDI4UT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsSWQ6dXMtbm9ydGhlYXN0I215LWFwcGxpY2F0aW9uLWlkI215LWNoYW5uZWwuMTM0NTY2XCIsXCJ0eXBlXCI6XCJzdHJlYW1cIn0ifQ==")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, RequiredTagField, "channelId:us-northeast#my-application-id#my-channel.134566")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForARoomAlias(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForRoomAlias("my-room").
		ForStreamingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiI1UkN3a0FrdFdJTDNWNllXN0V0dE14ejhpZXJvMWZkcXF0dEdRVFdaUDVCZ1k0OFhIUGltYmx3dDl1QUgyQWI3bHVVcWs0OG1DQktveE10WkhpaHNoQT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJyb29tQWxpYXM6bXktcm9vbVwiLFwidHlwZVwiOlwic3RyZWFtXCJ9In0=")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, RequiredTagField, "roomAlias:my-room")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForARoom(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForRoom("my-room.123456").
		ForStreamingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiI2WWdud09qWkx4Mk8zQXJjd29CUlVKU0UyYkRVNWVGY0FIYjI3OEJxVlMvcmplMXlsRU51bE5BSTVqakd2Mjc3VnZTTEtkYk1jTW1HenA3Nm9wNkNmZz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJyb29tSWQ6bXktcm9vbS4xMjM0NTZcIixcInR5cGVcIjpcInN0cmVhbVwifSJ9")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, RequiredTagField, "roomId:my-room.123456")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForATag(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForTag("my-tag=awesome").
		ForStreamingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJGUGRrTFFyVGlsS0toRDduc2QzeDZoNWV1aXVsaDVCYy9lNEtmQWY0THB5Qno4N2trK2lrQWN5ZUppcFk3alo4clpTN1N0bWw1aERMWEJIZXkrbmw2QT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJteS10YWc9YXdlc29tZVwiLFwidHlwZVwiOlwic3RyZWFtXCJ9In0=")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, RequiredTagField, "my-tag=awesome")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForAUriAndAChannelAlias(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		WithURI("https://my-custom-backend.example.org").
		ExpiresAt(time.UnixMilli(1000)).
		ForChannelAlias("my-channel").
		ForStreamingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJLUjJIb0xDbXJTZTRQWktpbXZDZ2dDWWJxOEprdG5iQlJGWDJuRTR3WVl3SUdleGdacUR3MGZLUDNZbEM1aFpLbi9ZRTFzYWFlUE9lR040U0ZOTWMzdz09IiwidG9rZW4iOiJ7XCJ1cmlcIjpcImh0dHBzOi8vbXktY3VzdG9tLWJhY2tlbmQuZXhhbXBsZS5vcmdcIixcImV4cGlyZXNcIjoxMDAwLFwicmVxdWlyZWRUYWdcIjpcImNoYW5uZWxBbGlhczpteS1jaGFubmVsXCIsXCJ0eXBlXCI6XCJzdHJlYW1cIn0ifQ==")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, URIField, "https://my-custom-backend.example.org")
	checkField(t, value, RequiredTagField, "channelAlias:my-channel")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForPublishing(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForPublishingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJrVElBcDh4ZUlqRXBxU2p0R3Zha3JOR2FFWnl5S1hMdmRMdmpBTHpJYkhYQmtqVXg2eU9hOHNmTGVoMFJydnNHaDJFbHF5OE5MMVBFVG51QjdQR3Z6dz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInR5cGVcIjpcInB1Ymxpc2hcIn0ifQ==")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, TypeField, "publish")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForPublishingToAChannelAlias(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForChannelAlias("my-channel").
		ForPublishingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJIREJPRzdiOFRuV0ZoNVMrR0Y5Z1lWQkNrM1J4WlhXNWh6UUN0bk9raXZLNlY0K1AxcDVKcHJ2TTNIVElyTUFBclUxMkY5bkltNGRvRm5TWXVjSzloUT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwidHlwZVwiOlwicHVibGlzaFwifSJ9")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, TypeField, "publish")
	checkField(t, value, RequiredTagField, "channelAlias:my-channel")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForPublishingToAChannel(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForChannel("us-northeast#my-application-id#my-channel.134566").
		ForPublishingOnly().
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJVZ3hjTDVVMlAvZDVtTXI4N3NzM3M5ZDdNNHo1elNZRGZrN0duL1BHS1d4S3NRS2t0c2pkN0Y5QTlRRHVQNnRSaTMzTG00TlpDVTZvSDFjbzFIa2Nmdz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsSWQ6dXMtbm9ydGhlYXN0I215LWFwcGxpY2F0aW9uLWlkI215LWNoYW5uZWwuMTM0NTY2XCIsXCJ0eXBlXCI6XCJwdWJsaXNoXCJ9In0=")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, TypeField, "publish")
	checkField(t, value, RequiredTagField, "channelId:us-northeast#my-application-id#my-channel.134566")
	checkVerifyWithBadSecret(t, token)
}

//------------------------------------------------------------------------------

func TestWhenVerifyingATokenForPublishingWithCapabilities(t *testing.T) {
	builder := NewTokenBuilder()
	token, err := builder.
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresAt(time.UnixMilli(1000)).
		ForPublishingOnly().
		WithCapability("multi-bitrate").
		WithCapability("streaming").
		Build()
	if err != nil {
		t.Errorf("failed to build token: %s", err.Error())
	}

	checkToken(t, token, "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJFKytBK3EwWGhGQ09LT011RnZqcnRIOVNyeHpwZ0Q1VVZYb1B6Q1VPaGNLU3pHTGRQZmsyRVYzVkZOOWRyM2tBVGZtSWRUeCtSTlFodjJ3aVJGbUM1Zz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInR5cGVcIjpcInB1Ymxpc2hcIixcImNhcGFiaWxpdGllc1wiOltcIm11bHRpLWJpdHJhdGVcIixcInN0cmVhbWluZ1wiXX0ifQ==")

	value := checkVerifyWithGoodSecret(t, token)
	checkField(t, value, TypeField, "publish")
	checkArrayField(t, value, CapabilitiesField, []string{"multi-bitrate", "streaming"})
	checkVerifyWithBadSecret(t, token)
}
