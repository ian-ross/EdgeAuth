Feature: Verifying tokens

  Scenario: Verifying a bad token (good secret)
    Given I have a bad token
    When I try to verify it with a good secret
    Then Verification should fail with error "bad-token"

  Scenario: Verifying a bad token (bad secret)
    Given I have a bad token
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-token"


  Scenario: Verifying a token for a channel alias and remote address (good secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for a remote address "10.1.2.3"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiI4MitYd1dITVRUc0xWYThKcnFPUmdjYlRXL2g2clFBTlF1MjgvRytQeHllQ09qSHEyb2xDYzVacUJ1MktqN0tGYmYyTC84TDZyaE9xTTZPMjNBR29HUT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwicmVtb3RlQWRkcmVzc1wiOlwiMTAuMS4yLjNcIixcInR5cGVcIjpcInN0cmVhbVwifSJ9"
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "channelAlias:my-channel"
    And The remote address field should be "10.1.2.3"

  Scenario: Verifying a token for a channel alias and remote address (bad secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for a remote address "10.1.2.3"
    And The token is for streaming only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for a channel alias and session (good secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for a session "session-id"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJBQi9Nanp2a1lnMGRTODF6aU1SVDZ3OUtwWmtjMU42U3VMTW56V09CQVJQZWJuenRHZTlmM2ZNS1FURXVqaHpVTkY0TWVsNkpMekFiWlZ3TFBSbEN4QT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwic2Vzc2lvbklkXCI6XCJzZXNzaW9uLWlkXCIsXCJ0eXBlXCI6XCJzdHJlYW1cIn0ifQ=="
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "channelAlias:my-channel"
    And The session field should be "session-id"

  Scenario: Verifying a token for a channel alias and session (bad secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for a session "session-id"
    And The token is for streaming only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for a channel alias and with a tag added (good secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for streaming only
    And The token has a "customer1" tag applied
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJMU0VnS2dGTy9aRUdxdEFLazVZb0F6cFJuTnQ4enhwUjNsdEJ3cWtOR3E1VWdjWWZpcnZKTDk3NHhpangyNS9XbHpqaUg1dk5ZMHdaYklFSkE2MzJqdz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwidHlwZVwiOlwic3RyZWFtXCIsXCJhcHBseVRhZ3NcIjpbXCJjdXN0b21lcjFcIl19In0="
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "channelAlias:my-channel"
    And The applied tags field should be "customer1"

  Scenario: Verifying a token for a channel alias and with a tag added (bad secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for streaming only
    And The token has a "customer1" tag applied
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for a channel alias (good secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJPMk90R1ZBMlErTGlhRkdjSjZ0cnlXZWE4L2l2dWFQR2gzcFJpcVd3ZlJPVWdBSSs0dFdaYXdBc011Y2MyMHNRTjZpaGZtVGVDNFVubXVoWko5aHBxUT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwidHlwZVwiOlwic3RyZWFtXCJ9In0="
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "channelAlias:my-channel"

  Scenario: Verifying a token for a channel alias (bad secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for streaming only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for a channel (good secret)
    Given I have a good token
    And The token is for a channel "us-northeast#my-application-id#my-channel.134566"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJZNGM3Tmp6eDVhalkzLzRWK3pwTVliNTBBU1ZCUXc0NlAvS2dwc3JrTnpDdFAzZWM5NzVzblorN3lJNzZiM0wrTmNtb2FoL3hOTUhQZ00vNEExaDI4UT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsSWQ6dXMtbm9ydGhlYXN0I215LWFwcGxpY2F0aW9uLWlkI215LWNoYW5uZWwuMTM0NTY2XCIsXCJ0eXBlXCI6XCJzdHJlYW1cIn0ifQ=="
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "channelId:us-northeast#my-application-id#my-channel.134566"

  Scenario: Verifying a token for a channel (bad secret)
    Given I have a good token
    And The token is for a channel "us-northeast#my-application-id#my-channel.134566"
    And The token is for streaming only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for a room alias (good secret)
    Given I have a good token
    And The token is for a room alias "my-room"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiI1UkN3a0FrdFdJTDNWNllXN0V0dE14ejhpZXJvMWZkcXF0dEdRVFdaUDVCZ1k0OFhIUGltYmx3dDl1QUgyQWI3bHVVcWs0OG1DQktveE10WkhpaHNoQT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJyb29tQWxpYXM6bXktcm9vbVwiLFwidHlwZVwiOlwic3RyZWFtXCJ9In0="
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "roomAlias:my-room"

  Scenario: Verifying a token for a room alias (bad secret)
    Given I have a good token
    And The token is for a room alias "my-room"
    And The token is for streaming only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for a room (good secret)
    Given I have a good token
    And The token is for a room "my-room.123456"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiI2WWdud09qWkx4Mk8zQXJjd29CUlVKU0UyYkRVNWVGY0FIYjI3OEJxVlMvcmplMXlsRU51bE5BSTVqakd2Mjc3VnZTTEtkYk1jTW1HenA3Nm9wNkNmZz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJyb29tSWQ6bXktcm9vbS4xMjM0NTZcIixcInR5cGVcIjpcInN0cmVhbVwifSJ9"
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "roomId:my-room.123456"

  Scenario: Verifying a token for a room (bad secret)
    Given I have a good token
    And The token is for a room "my-room.123456"
    And The token is for streaming only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for a tag (good secret)
    Given I have a good token
    And The token is for tag "my-tag=awesome"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJGUGRrTFFyVGlsS0toRDduc2QzeDZoNWV1aXVsaDVCYy9lNEtmQWY0THB5Qno4N2trK2lrQWN5ZUppcFk3alo4clpTN1N0bWw1aERMWEJIZXkrbmw2QT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJteS10YWc9YXdlc29tZVwiLFwidHlwZVwiOlwic3RyZWFtXCJ9In0="
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "my-tag=awesome"

  Scenario: Verifying a token for a tag (bad secret)
    Given I have a good token
    And The token is for tag "my-tag=awesome"
    And The token is for streaming only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for a URI and a channel alias (good secret)
    Given I have a good token with URI "https://my-custom-backend.example.org"
    And The token is for a channel alias "my-channel"
    And The token is for streaming only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJLUjJIb0xDbXJTZTRQWktpbXZDZ2dDWWJxOEprdG5iQlJGWDJuRTR3WVl3SUdleGdacUR3MGZLUDNZbEM1aFpLbi9ZRTFzYWFlUE9lR040U0ZOTWMzdz09IiwidG9rZW4iOiJ7XCJ1cmlcIjpcImh0dHBzOi8vbXktY3VzdG9tLWJhY2tlbmQuZXhhbXBsZS5vcmdcIixcImV4cGlyZXNcIjoxMDAwLFwicmVxdWlyZWRUYWdcIjpcImNoYW5uZWxBbGlhczpteS1jaGFubmVsXCIsXCJ0eXBlXCI6XCJzdHJlYW1cIn0ifQ=="
    When I try to verify it with a good secret
    Then Verification should pass
    And The URI field should be "https://my-custom-backend.example.org"
    And The tag field should be "channelAlias:my-channel"

  Scenario: Verifying a token for a URI and a channel alias (bad secret)
    Given I have a good token with URI "https://my-custom-backend.example.org"
    And The token is for a channel alias "my-channel"
    And The token is for streaming only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for publishing (good secret)
    Given I have a good token
    And The token is for publishing only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJrVElBcDh4ZUlqRXBxU2p0R3Zha3JOR2FFWnl5S1hMdmRMdmpBTHpJYkhYQmtqVXg2eU9hOHNmTGVoMFJydnNHaDJFbHF5OE5MMVBFVG51QjdQR3Z6dz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInR5cGVcIjpcInB1Ymxpc2hcIn0ifQ=="
    When I try to verify it with a good secret
    Then Verification should pass
    And The type field should be "publish"

  Scenario: Verifying a token for publishing (bad secret)
    Given I have a good token
    And The token is for publishing only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for publishing to a channel alias (good secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for publishing only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJIREJPRzdiOFRuV0ZoNVMrR0Y5Z1lWQkNrM1J4WlhXNWh6UUN0bk9raXZLNlY0K1AxcDVKcHJ2TTNIVElyTUFBclUxMkY5bkltNGRvRm5TWXVjSzloUT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsQWxpYXM6bXktY2hhbm5lbFwiLFwidHlwZVwiOlwicHVibGlzaFwifSJ9"
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "channelAlias:my-channel"
    And The type field should be "publish"

  Scenario: Verifying a token for publishing to a channel alias (bad secret)
    Given I have a good token
    And The token is for a channel alias "my-channel"
    And The token is for publishing only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for publishing to a channel (good secret)
    Given I have a good token
    And The token is for a channel "us-northeast#my-application-id#my-channel.134566"
    And The token is for publishing only
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJVZ3hjTDVVMlAvZDVtTXI4N3NzM3M5ZDdNNHo1elNZRGZrN0duL1BHS1d4S3NRS2t0c2pkN0Y5QTlRRHVQNnRSaTMzTG00TlpDVTZvSDFjbzFIa2Nmdz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInJlcXVpcmVkVGFnXCI6XCJjaGFubmVsSWQ6dXMtbm9ydGhlYXN0I215LWFwcGxpY2F0aW9uLWlkI215LWNoYW5uZWwuMTM0NTY2XCIsXCJ0eXBlXCI6XCJwdWJsaXNoXCJ9In0="
    When I try to verify it with a good secret
    Then Verification should pass
    And The tag field should be "channelId:us-northeast#my-application-id#my-channel.134566"
    And The type field should be "publish"

  Scenario: Verifying a token for publishing to a channel (bad secret)
    Given I have a good token
    And The token is for a channel "us-northeast#my-application-id#my-channel.134566"
    And The token is for publishing only
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"


  Scenario: Verifying a token for publishing with capabilities (good secret)
    Given I have a good token
    And The token is for publishing only
    And The token has capability "multi-bitrate"
    And The token has capability "streaming"
    And The correct token is "DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJFKytBK3EwWGhGQ09LT011RnZqcnRIOVNyeHpwZ0Q1VVZYb1B6Q1VPaGNLU3pHTGRQZmsyRVYzVkZOOWRyM2tBVGZtSWRUeCtSTlFodjJ3aVJGbUM1Zz09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInR5cGVcIjpcInB1Ymxpc2hcIixcImNhcGFiaWxpdGllc1wiOltcIm11bHRpLWJpdHJhdGVcIixcInN0cmVhbWluZ1wiXX0ifQ=="
    When I try to verify it with a good secret
    Then Verification should pass
    And The type field should be "publish"
    And The capabilities field should be "multi-bitrate,streaming"

  Scenario: Verifying a token for publishing with capabilities (bad secret)
    Given I have a good token
    And The token is for publishing only
    And The token has capability "multi-bitrate"
    And The token has capability "streaming"
    When I try to verify it with a bad secret
    Then Verification should fail with error "bad-digest"
