package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/PhenixRTS/EdgeAuth/golang"
)

func TestWhenVerifyingATokenForAChannelAliasAndRemoteAddress(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.RequiredTagField, "channelAlias:my-channel")
	checkField(t, value, edgeauth.RemoteAddressField, "10.1.2.3")
	checkVerifyWithBadSecret(t, token)
}
