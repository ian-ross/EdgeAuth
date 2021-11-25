const TokenBuilder = require('../../src/TokenBuilder');
const { Given, When, Then } = require("@cucumber/cucumber");
const assert = require("assert").strict;

// INITIALIZATION STEPS

Given("I have a bad token", function () {
  this.setToken("DIGEST:bad-token");
});

Given("I have a good token", function () {
	this.builder = this.builder
    .withApplicationId("my-application-id")
		.withSecret("my-secret")
		.expiresAt(new Date(1000))
});

Given("I have a good token with URI {string}", function (uri) {
	this.builder = this.builder
    .withApplicationId("my-application-id")
		.withSecret("my-secret")
    .withUri(uri)
		.expiresAt(new Date(1000))
});

// VERIFICATION STEPS

When("I try to verify it with a bad secret", function () {
	this.buildToken()
	this.verifyAndDecode("bad-secret")
});

When("I try to verify it with a good secret", function () {
	this.buildToken()
  if (this.correctToken) {
    assert.equal(this.token, this.correctToken, "token does not match expected value")
  }
	this.verifyAndDecode("my-secret")
});

Given("The correct token is {string}", function (arg1) {
	this.correctToken = arg1
});

// TOKEN SETUP STEPS

Given("The token is for a channel {string}", function (arg1) {
	this.builder = this.builder.forChannel(arg1);
})

Given("The token is for a channel alias {string}", function (arg1) {
	this.builder = this.builder.forChannelAlias(arg1);
})

Given("The token is for a room {string}", function (arg1) {
	this.builder = this.builder.forRoom(arg1);
})

Given("The token is for a room alias {string}", function (arg1) {
	this.builder = this.builder.forRoomAlias(arg1);
})

Given("The token is for a remote address {string}", function (arg1) {
	this.builder = this.builder.forRemoteAddress(arg1);
})

Given("The token is for a session {string}", function (arg1) {
	this.builder = this.builder.forSession(arg1);
})

Given("The token is for streaming only", function () {
	this.builder = this.builder.forStreamingOnly();
})

Given("The token is for publishing only", function () {
	this.builder = this.builder.forPublishingOnly();
})

Given("The token is for tag {string}", function (arg1) {
	this.builder = this.builder.forTag(arg1);
})

Given("The token has a {string} tag applied", function (arg1) {
	this.builder = this.builder.applyTag(arg1);
})

Given("The token has capability {string}", function (arg1) {
	this.builder = this.builder.withCapability(arg1);
})

// FIELD TESTING STEPS

Then("The remote address field should be {string}", function (arg1) {
  this.checkField('remoteAddress', arg1);
});

Then("The URI field should be {string}", function (arg1) {
	this.checkField('uri', arg1);
});

Then("The session field should be {string}", function (arg1) {
	this.checkField('sessionId', arg1);
});

Then("The type field should be {string}", function (arg1) {
	this.checkField('type', arg1);
});

Then("The tag field should be {string}", function (arg1) {
	this.checkField('requiredTag', arg1);
});

Then("The applied tags field should be {string}", function (arg1) {
	this.checkArrayField('applyTags', arg1);
});

Then("The capabilities field should be {string}", function (arg1) {
	this.checkArrayField('capabilities', arg1);
});

// RESULT CHECKING STEPS

Then("Verification should fail with error {string}", function (arg1) {
  assert(this.result, "verification result is not set");
	assert(!this.result.verified, "token did not fail to verify")
	assert.equal(this.result.code, arg1, 
		           "result Code should be '" + arg1 + "', but is '" + this.result.code + "'")
  assert.equal(this.result.message, undefined,
		           "result message should be undefined, but is '" + this.result.message + "'")
  assert.equal(this.result.value, undefined,
		           "result value should be undefined, but is not")
});

Then("Verification should pass", function () {
  assert(this.result, "verification result is not set");
	assert(this.result.verified, "token failed to verify");
  assert.equal(this.result.code, "verified",
		           "result code should be 'verified', but is '" + this.result.code + "'")
	assert(this.result.value, "result value is nil, and should not be");
});
