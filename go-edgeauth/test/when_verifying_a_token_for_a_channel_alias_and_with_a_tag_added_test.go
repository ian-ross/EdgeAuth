package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/ian-ross/EdgeAuth/go-edgeauth"
)

func TestWhenVerifyingATokenForAChannelAliasAndWithATagAdded(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.RequiredTagField, "channelAlias:my-channel")
	checkArrayField(t, value, edgeauth.ApplyTagsField, []string{"customer1"})
	checkVerifyWithBadSecret(t, token)
}
