package edgeauth

import (
	"testing"
	"time"

	edgeauth "github.com/ian-ross/EdgeAuth/go-edgeauth"
)

func TestWhenVerifyingATokenForARoom(t *testing.T) {
	builder := edgeauth.NewTokenBuilder()
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
	checkField(t, value, edgeauth.RequiredTagField, "roomId:my-room.123456")
	checkVerifyWithBadSecret(t, token)
}
