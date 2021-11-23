package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/ian-ross/EdgeAuth/golang"
)

func TestWhenVerifyingATokenForPublishing(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.TypeField, "publish")
	checkVerifyWithBadSecret(t, token)
}
