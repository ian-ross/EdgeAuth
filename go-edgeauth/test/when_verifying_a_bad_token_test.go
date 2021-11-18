package edgeauth

import (
	"testing"

	edgeauth "github.com/PhenixRTS/EdgeAuth/go-edgeauth"
)

func TestWhenVerifyingABadToken(t *testing.T) {
	token := "DIGEST:bad-token"

	t.Run("The token fails to verify", func(t *testing.T) {
		result := edgeauth.VerifyAndDecode("bad-secret", token)
		checkFail(t, result, "bad-token")
	})
}
