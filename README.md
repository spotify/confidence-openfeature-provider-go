# Confidence OpenFeature Go Provider

This repo contains the OpenFeature Go flag provider for [Confidence](https://confidence.spotify.com/).

## OpenFeature

Before starting to use the provider, it can be helpful to read through the general [OpenFeature docs](https://docs.openfeature.dev/)
and get familiar with the concepts. 

## Adding the dependency
<!---x-release-please-start-version-->
```
require (
	github.com/spotify/confidence-openfeature-provider-go v0.1.8
)
```
<!---x-release-please-end-->
It's also important to add the underlying OpenFeature SDK dependency:
```
require (
	github.com/open-feature/go-sdk v1.7.0
)
```

### Creating and using the flag provider

Below is an example for how to create a OpenFeature client using the Confidence flag provider, and then resolve
a flag with a boolean attribute. The provider is configured with an api key and a region, which will determine
where it will send the resolving requests. 

The flag will be applied immediately, meaning that Confidence will count the targeted user as having received the treatment. 

You can retrieve attributes on the flag variant using property dot notation, meaning `test-flag.boolean-key` will retrieve
the attribute `boolean-key` on the flag `test-flag`. 

You can also use only the flag name `test-flag` and retrieve all values as a map with `client.ObjectValue()`. 

The flag's schema is validated against the requested data type, and if it doesn't match it will fall back to the default value. 

```go
import (
    "github.com/open-feature/go-sdk/openfeature"
    confidence "github.com/spotify/confidence-openfeature-provider-go/pkg/provider"
)

provider, err := confidence.NewFlagProvider(confidence.NewAPIConfig("clientSecret"))

if err != nil {
    // handle error	
}

openfeature.SetProvider(provider)
client := openfeature.NewClient("testApp")
	
attributes := make(map[string]interface{})
attributes["country"] = "SE"
attributes["plan"] = "premium"
attributes["user_id"] = "user1"

boolValue, error := client.BooleanValue(context.Background(), "test-flag.boolean-key", false, 
	openfeature.NewEvaluationContext("", attributes))
```
## Demo app

To run the demo app, replace the `CLIENT_SECRET` with client secret setup in the 
[Confidence](https://confidence.spotify.com/) console, the flags with existing flags and execute 
the app with `cd demo && go run GoDemoApp.go`.
