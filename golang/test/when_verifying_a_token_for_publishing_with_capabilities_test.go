package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/PhenixRTS/EdgeAuth/golang"
)

func TestWhenVerifyingATokenForPublishingWithCapabilities(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.TypeField, "publish")
	checkArrayField(t, value, edgeauth.CapabilitiesField, []string{"multi-bitrate", "streaming"})
	checkVerifyWithBadSecret(t, token)
}
