package edgeauth

import (
	edgeauth "github.com/ian-ross/EdgeAuth/golang"
)

var token *string
var correctToken *string
var builder *edgeauth.TokenBuilder
var result *edgeauth.VerifyAndDecodeResult
