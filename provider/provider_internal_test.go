package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-feature/go-sdk/openfeature"
	confidence "github.com/spotify/confidence-openfeature-provider-go/confidence"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type MockResolveClient struct {
	MockedResponse confidence.ResolveResponse
	MockedError    error
	TestingT       *testing.T
}

func (r MockResolveClient) SendResolveRequest(_ context.Context,
	request confidence.ResolveRequest) (confidence.ResolveResponse, error) {
	assert.Equal(r.TestingT, "user1", request.EvaluationContext["targeting_key"])
	return r.MockedResponse, r.MockedError
}

func TestResolveBoolValue(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.BooleanValueDetails(
		context.Background(), "test-flag.boolean-key", false, openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, true, evalDetails.Value)
	assert.Equal(t, openfeature.TargetingMatchReason, evalDetails.Reason)
	assert.Equal(t, "flags/test-flag/variants/treatment", evalDetails.Variant)
	assert.Equal(t, "test-flag.boolean-key", evalDetails.FlagKey)
}

func TestResolveIntValue(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.IntValueDetails(
		context.Background(), "test-flag.integer-key", 99, openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, int64(40), evalDetails.Value)
}

func TestResolveDoubleValue(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.FloatValueDetails(
		context.Background(), "test-flag.double-key", 99.99, openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, 20.203, evalDetails.Value)
}

func TestResolveStringValue(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.StringValueDetails(
		context.Background(), "test-flag.string-key", "default", openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, "treatment", evalDetails.Value)
}

func TestResolveObjectValue(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.ObjectValueDetails(
		context.Background(), "test-flag.struct-key", "default", openfeature.NewEvaluationContext(
			"user1",
			attributes))

	_, ok := evalDetails.Value.(map[string]interface{})
	assert.True(t, ok)
}

func TestResolveNestedValue(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.BooleanValueDetails(
		context.Background(), "test-flag.struct-key.boolean-key", true, openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, false, evalDetails.Value)
}

func TestResolveDoubleNestedValue(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.BooleanValueDetails(
		context.Background(), "test-flag.struct-key.nested-struct-key.nested-boolean-key", true, openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, false, evalDetails.Value)
}

func TestResolveWholeFlagAsObject(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.ObjectValueDetails(
		context.Background(), "test-flag", "default", openfeature.NewEvaluationContext(
			"user1",
			attributes))

	_, ok := evalDetails.Value.(map[string]interface{})
	assert.True(t, ok)
}

func TestResolveWholeFlagAsObjectWithInts(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.ObjectValueDetails(
		context.Background(), "test-flag", "default", openfeature.NewEvaluationContext(
			"user1",
			attributes))

	value, _ := evalDetails.Value.(map[string]interface{})
	rootIntValue := value["integer-key"]

	assert.Equal(t, reflect.Int64, reflect.ValueOf(rootIntValue).Kind())
	assert.Equal(t, int64(40), rootIntValue)

	nestedIntValue := value["struct-key"].(map[string]interface{})["integer-key"]

	assert.Equal(t, reflect.Int64, reflect.ValueOf(nestedIntValue).Kind())
	assert.Equal(t, int64(23), nestedIntValue)
}

func TestResolveWithWrongType(t *testing.T) {
	client := client(t, templateResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.BooleanValueDetails(
		context.Background(), "test-flag.integer-key", false, openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, false, evalDetails.Value)
	assert.Equal(t, openfeature.ErrorReason, evalDetails.Reason)
	assert.Equal(t, openfeature.TypeMismatchCode, evalDetails.ErrorCode)
}

func TestResolveWithUnexpectedFlag(t *testing.T) {
	client := client(t, templateResponseWithFlagName("wrong-flag"), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.BooleanValueDetails(
		context.Background(), "test-flag.boolean-key", true, openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, true, evalDetails.Value)
	assert.Equal(t, openfeature.ErrorReason, evalDetails.Reason)
	assert.Equal(t, openfeature.FlagNotFoundCode, evalDetails.ErrorCode)
	assert.Equal(t, "unexpected flag 'wrong-flag' from remote", evalDetails.ErrorMessage)
}

func TestResolveWithNonExistingFlag(t *testing.T) {
	client := client(t, emptyResponse(), nil)
	attributes := make(map[string]interface{})

	evalDetails, _ := client.BooleanValueDetails(
		context.Background(), "test-flag.boolean-key", true, openfeature.NewEvaluationContext(
			"user1",
			attributes))

	assert.Equal(t, true, evalDetails.Value)
	assert.Equal(t, openfeature.ErrorReason, evalDetails.Reason)
	assert.Equal(t, openfeature.FlagNotFoundCode, evalDetails.ErrorCode)
	assert.Equal(t, "Flag not found", evalDetails.ErrorMessage)
}

func client(t *testing.T, response confidence.ResolveResponse, errorToReturn error) *openfeature.Client {
	resolveClient := MockResolveClient{MockedResponse: response, MockedError: errorToReturn, TestingT: t}
	conf := confidence.NewConfidenceBuilder().SetAPIConfig(confidence.APIConfig{APIKey: "apiKey"}).SetResolveClient(resolveClient).Build()
	provider := FlagProvider{
		confidence: conf,
	}
	openfeature.SetProvider(provider)
	return openfeature.NewClient("testApp")
}

func templateResponse() confidence.ResolveResponse {
	return templateResponseWithFlagName("test-flag")
}

func templateResponseWithFlagName(flagName string) confidence.ResolveResponse {
	templateResolveResponse := fmt.Sprintf(`
{
"resolvedFlags": [
{
"flag": "flags/%[1]s",
"variant": "flags/%[1]s/variants/treatment",
"value": {
"struct-key": {
"boolean-key": false,
"string-key": "treatment-struct",
"double-key": 123.23,
"integer-key": 23,
	"nested-struct-key": {
		"nested-boolean-key": false
	}
},
"boolean-key": true,
"string-key": "treatment",
"double-key": 20.203,
"integer-key": 40
},
"flagSchema": {
"schema": {
"struct-key": {
"structSchema": {
"schema": {
"boolean-key": {
"boolSchema": {}
},
"string-key": {
"stringSchema": {}
},
"double-key": {
"doubleSchema": {}
},
"integer-key": {
"intSchema": {}
},
	"nested-struct-key": {
		"structSchema": {
			"schema": {
				"nested-boolean-key": {
					"boolSchema": {}
				}
			}
		}
	}
}
}
},
"boolean-key": {
"boolSchema": {}
},
"string-key": {
"stringSchema": {}
},
"double-key": {
"doubleSchema": {}
},
"integer-key": {
"intSchema": {}
}
}
},
"reason": "RESOLVE_REASON_MATCH"
}],
"resolveToken": ""
}
`, flagName)
	var result confidence.ResolveResponse
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(templateResolveResponse)))
	decoder.UseNumber()
	_ = decoder.Decode(&result)
	return result
}

func emptyResponse() confidence.ResolveResponse {
	templateResolveResponse :=
		`
{
"resolvedFlags": [],
"resolveToken": ""
}
`
	var result confidence.ResolveResponse
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(templateResolveResponse)))
	decoder.UseNumber()
	_ = decoder.Decode(&result)
	return result
}
