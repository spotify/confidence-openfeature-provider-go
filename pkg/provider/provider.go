package provider

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/open-feature/go-sdk/openfeature"
)

type FlagProvider struct {
	Config        APIConfig
	ResolveClient resolveClient
}

var (
    SDK_ID = "SDK_ID_GO_PROVIDER"
    SDK_VERSION = "0.1.8" // x-release-please-version
)

func NewFlagProvider(config APIConfig) (*FlagProvider, error) {
	validationError := config.validate()
	if validationError != nil {
		return nil, validationError
	}
	return &FlagProvider{Config: config,
		ResolveClient: httpResolveClient{Client: &http.Client{}, Config: config}}, nil
}

func (e FlagProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{Name: "ConfidenceFlagProvider"}
}

func (e FlagProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool,
	evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	res := e.resolveFlag(ctx, flag, defaultValue, evalCtx, reflect.Bool)
	return toBoolResolutionDetail(res, defaultValue)
}

func (e FlagProvider) StringEvaluation(ctx context.Context, flag string, defaultValue string,
	evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	res := e.resolveFlag(ctx, flag, defaultValue, evalCtx, reflect.String)
	return toStringResolutionDetail(res, defaultValue)
}

func (e FlagProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64,
	evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	res := e.resolveFlag(ctx, flag, defaultValue, evalCtx, reflect.Float64)
	return toFloatResolutionDetail(res, defaultValue)
}

func (e FlagProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64,
	evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	res := e.resolveFlag(ctx, flag, defaultValue, evalCtx, reflect.Int64)
	return toIntResolutionDetail(res, defaultValue)
}

func (e FlagProvider) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{},
	evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return e.resolveFlag(ctx, flag, defaultValue, evalCtx, reflect.Map)
}

func (e FlagProvider) resolveFlag(ctx context.Context, flag string, defaultValue interface{},
	evalCtx openfeature.FlattenedContext, expectedKind reflect.Kind) openfeature.InterfaceResolutionDetail {
	flagName, propertyPath := splitFlagString(flag)

	requestFlagName := fmt.Sprintf("flags/%s", flagName)
	resp, err := e.ResolveClient.sendResolveRequest(ctx,
		resolveRequest{ClientSecret: e.Config.APIKey,
			Flags: []string{requestFlagName}, Apply: true, EvaluationContext: processTargetingKey(evalCtx),
			Sdk: sdk{Id: SDK_ID, Version: SDK_VERSION}})

	if err != nil {
		return processResolveError(err, defaultValue)
	}

	if len(resp.ResolvedFlags) == 0 {
		return openfeature.InterfaceResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: openfeature.NewFlagNotFoundResolutionError(fmt.Sprintf("no active flag '%s' was found", flagName)),
				Reason:          openfeature.ErrorReason}}
	}

	resolvedFlag := resp.ResolvedFlags[0]
	if resolvedFlag.Flag != requestFlagName {
		return openfeature.InterfaceResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: openfeature.NewFlagNotFoundResolutionError(fmt.Sprintf("unexpected flag '%s' from remote", strings.TrimPrefix(resolvedFlag.Flag, "flags/"))),
				Reason:          openfeature.ErrorReason}}
	}

	return processResolvedFlag(resolvedFlag, defaultValue, expectedKind, propertyPath)
}

func processTargetingKey(evalCtx openfeature.FlattenedContext) openfeature.FlattenedContext {
	newEvalContext := openfeature.FlattenedContext{}
	newEvalContext = evalCtx
	if targetingKey, exists := evalCtx["targetingKey"]; exists {
		newEvalContext["targeting_key"] = targetingKey
	}
	return newEvalContext
}

func (e FlagProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}
