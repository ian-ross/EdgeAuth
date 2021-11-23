package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/PhenixRTS/EdgeAuth/golang"
)

func TestWhenVerifyingATokenForAUriAndAChannelAlias(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.URIField, "https://my-custom-backend.example.org")
	checkField(t, value, edgeauth.RequiredTagField, "channelAlias:my-channel")
	checkVerifyWithBadSecret(t, token)
}
