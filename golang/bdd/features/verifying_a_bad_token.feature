Feature: Verifying a bad token

  Scenario: Verifying a bad token (good secret)
    When I try to verify a token with a good secret
    Given I have a bad token
    Then Verification should fail with error "bad-token"

  Scenario: Verifying a bad token (bad secret)
    When I try to verify a token with a bad secret
    Given I have a bad token
    Then Verification should fail with error "bad-token"
