Feature: Verifying a token for a URI and a channel alias

  Scenario: Verifying a token for a URI and a channel alias (good secret)
    When I try to verify a token with a good secret
    Given I have a good token with URI "https://my-custom-backend.example.org"
    And The token is for a channel alias "my-channel"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJLUjJIb0xDbXJTZTRQWktpbXZDZ2dDWWJxOEprdG5iQlJGWDJuRTR3WVl3SUdleGdacUR3MGZLUDNZbEM1aFpLbi9ZRTFzYWFlUE9lR040U0ZOTWMzdz09IiwidG9rZW4iOiJ7XCJ1cmlcIjpcImh0dHBzOi8vbXktY3VzdG9tLWJhY2tlbmQuZXhhbXBsZS5vcmdcIixcImV4cGlyZXNcIjoxMDAwLFwicmVxdWlyZWRUYWdcIjpcImNoYW5uZWxBbGlhczpteS1jaGFubmVsXCIsXCJ0eXBlXCI6XCJzdHJlYW1cIn0ifQ=="
    Then Verification should pass
    And The URI field should be "https://my-custom-backend.example.org"
    And The tag field should be "channelAlias:my-channel"

  Scenario: Verifying a token for a URI and a channel alias (bad secret)
    When I try to verify a token with a bad secret
    Given I have a good token with URI "https://my-custom-backend.example.org"
    And The token is for a channel alias "my-channel"
    And The token is for streaming only
    Then Verification should fail with error "bad-digest"
