# Phenix EdgeAuth Digest Tokens for Go

Easily generate secure digest tokens to use with the Phenix platform without requiring any networking activity.

## Installation

To install Phenix Edge Authorization Digest Token:

```shell script
$ go install github.com/PhenixRTS/EdgeAuth/go-edgeauth/...
```

## Testing

```shell script
$ go test -v
```

## Example

Create a Go test program:

```shell script
$ mkdir edgeauth-test
$ cd edgeauth-test
$ go mod init example.com/edgeauth-test
```

Copy the following text into a `main.go` file:

```go
package main

import (
	"fmt"
	edgeauth "github.com/PhenixRTS/EdgeAuth/go-edgeauth"
)

func main() {
	// Create a token to access a channel
	token, err := edgeauth.NewTokenBuilder().
		WithApplicationID("my-application-id").
		WithSecret("my-secret").
		ExpiresInSeconds(3600).
		ForChannel("us-northeast#my-application-id#my-channel.1345").
		Build()
    
	if err != nil {
		fmt.Println("Token encoding error:", err.Error())
	} else {
		fmt.Println(*token)
	}
}
```

Build and run the example:

```shell script
$ go build
$ ./edgeauth-test
```

## Command Line Examples

Display the help information:
```shell script
$ edgeauth -help
```

Create a token for channel access:
```shell script
$ edgeauth -application_id "my-application-id" -secret "my-secret" -expires_in_seconds 3600 -channel "us-northeast#my-application-id#my-channel.1345"
```
