Feature: Verifying a token for a channel alias and remote address

  Scenario: Verifying a token for a channel alias and remote address (good secret)
    When I try to verify a token with a good secret
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for a remote address "10.1.2.3"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiI4MitYd1dITVRUc0xWYThKcnFPUmdjYlRXL2g2clFBTlF1MjgvRytQeHllQ09qSHEyb2xDYzVacUJ1MktqN0tGYmYyTC84TDZyaE9xTTZPMjNBR29HUT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwicmVtb3RlQWRkcmVzc1wiOlwiMTAuMS4yLjNcIixcInR5cGVcIjpcInN0cmVhbVwifSJ9"
    Then Verification should pass
    And The tag field should be "channelAlias:my-channel"
    And The remote address field should be "10.1.2.3"

  Scenario: Verifying a token for a channel alias and remote address (bad secret)
    When I try to verify a token with a bad secret
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for a remote address "10.1.2.3"
    And The token is for streaming only
    Then Verification should fail with error "bad-digest"
