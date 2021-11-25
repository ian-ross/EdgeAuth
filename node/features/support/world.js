const TokenBuilder = require('../../src/TokenBuilder');
const DigestTokens = require('../../src/DigestTokens');
const { setWorldConstructor } = require("@cucumber/cucumber");
const assert = require("assert").strict;

class CustomWorld {
  constructor() {
    this.token = undefined;
    this.correctToken = undefined;
    this.builder = new TokenBuilder();
    this.result = undefined;
  }

  setToken(token) {
    this.token = token;
  }

  buildToken() {
    if (!this.token && this.builder) {
		  this.token = this.builder.build()
    }
  }

  verifyAndDecode(secret) {
    this.result = new DigestTokens().verifyAndDecode(secret, this.token);
  }

  checkField(expectedField, expectedValue) {
	  assert(this.result && this.result.value, "the verification value is not set");
	  var check = this.result.value[expectedField];
	  assert.equal(check, expectedValue,
		             "required " + expectedField + " in value does not match");
	}

  checkArrayField(expectedField, expectedValues) {
	  assert(this.result && this.result.value, "the verification value is not set");
	  var expectedValue = expectedValues.split(",");
	  var check = this.result.value[expectedField];
	  assert.deepEqual(check, expectedValue,
		                 "required " + expectedField + " in value does not match");
  }
}

setWorldConstructor(CustomWorld);
