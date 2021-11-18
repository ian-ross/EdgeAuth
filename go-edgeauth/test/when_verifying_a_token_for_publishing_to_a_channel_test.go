package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/PhenixRTS/EdgeAuth/go-edgeauth"
)

func TestWhenVerifyingATokenForPublishingToAChannel(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.TypeField, "publish")
	checkField(t, value, edgeauth.RequiredTagField, "channelId:us-northeast#my-application-id#my-channel.134566")
	checkVerifyWithBadSecret(t, token)
}
