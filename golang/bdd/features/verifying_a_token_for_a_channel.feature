Feature: Verifying a token for a channel

  Scenario: Verifying a token for a channel (good secret)
    When I try to verify a token with a good secret
    Given I have a good token
    And The token is for a channel "us-northeast#my-application-id#my-channel.134566"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJZNGM3Tmp6eDVhalkzLzRWK3pwTVliNTBBU1ZCUXc0NlAvS2dwc3JrTnpDdFAzZWM5NzVzblorN3lJNzZiM0wrTmNtb2FoL3hOTUhQZ00vNEExaDI4UT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsSWQ6dXMtbm9ydGhlYXN0I215LWFwcGxpY2F0aW9uLWlkI215LWNoYW5uZWwuMTM0NTY2XCIsXCJ0eXBlXCI6XCJzdHJlYW1cIn0ifQ=="
    Then Verification should pass
    And The tag field should be "channelId:us-northeast#my-application-id#my-channel.134566"

  Scenario: Verifying a token for a channel (bad secret)
    When I try to verify a token with a bad secret
    Given I have a good token
    And The token is for a channel "us-northeast#my-application-id#my-channel.134566"
    And The token is for streaming only
    Then Verification should fail with error "bad-digest"
