package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/ian-ross/EdgeAuth/golang"
)

func TestWhenVerifyingATokenForPublishingToAChannelAlias(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.TypeField, "publish")
	checkField(t, value, edgeauth.RequiredTagField, "channelAlias:my-channel")
	checkVerifyWithBadSecret(t, token)
}
