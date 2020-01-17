/**
 * Copyright 2019 Phenix Real Time Solutions, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
const TokenBuilder = require('../src/TokenBuilder');
const DigestTokens = require('../src/DigestTokens');

describe('When verifying a token for publishing', () => {
  var token;

  beforeEach(() => {
    token = new TokenBuilder()
      .withApplicationId('my-application-id')
      .withSecret(('my-secret'))
      .expiresAt(new Date(1000))
      .forStreamingOnly()
      .withCapability('multi-bitrate')
      .withCapability('streaming')
      .build();
  });

  test('The token matches the expected value', () => {
    expect(token).toBe('DIGEST:eyJhcHBsaWNhdGlvbklkIjoibXktYXBwbGljYXRpb24taWQiLCJkaWdlc3QiOiJMQU5xV1d2TWZvMmNxMzM2cEZEZU11VTFHa25YWnhCdEpTNnc1dE9VRXdCK1pmaTA1dWFwaFowUmNpZGFhNmFaUm4rSHkzMUF1eDNqUFlubE9pTnowUT09IiwidG9rZW4iOiJ7XCJleHBpcmVzXCI6MTAwMCxcInR5cGVcIjpcInN0cmVhbVwiLFwiY2FwYWJpbGl0aWVzXCI6W1wibXVsdGktYml0cmF0ZVwiLFwic3RyZWFtaW5nXCJdfSJ9');
  });

  test('The token successfully verifies with the correct secret', () => {
    const result = new DigestTokens().verifyAndDecode('my-secret', token);

    expect(result.verified).toBe(true);
    expect(result.code).toBe('verified');
    expect(result.value).not.toBe(undefined);
    expect(result.value.capabilities).toStrictEqual(['multi-bitrate', 'streaming']);
  });

  test('The token fails to verify with a bad secret', () => {
    const result = new DigestTokens().verifyAndDecode('bad-secret', token);

    expect(result.verified).toBe(false);
    expect(result.code).toBe('bad-digest');
    expect(result.value).toBe(undefined);
  });
});