package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/PhenixRTS/EdgeAuth/go-edgeauth"
)

func TestWhenVerifyingATokenForAChannelAliasAndSession(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.RequiredTagField, "channelAlias:my-channel")
	checkField(t, value, edgeauth.SessionIDField, "session-id")
	checkVerifyWithBadSecret(t, token)
}
