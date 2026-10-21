# \AIModelsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentsTokenPackages**](AIModelsAPI.md#GetAgentsTokenPackages) | **Get** /api/v1/cloud-ai/token-packages/agents | Получение списка пакетов токенов для агентов
[**GetAiModelParameters**](AIModelsAPI.md#GetAiModelParameters) | **Get** /api/v3/cloud-ai/model-parameters | Получение справочника характеристик моделей
[**GetAiModelTokenPackages**](AIModelsAPI.md#GetAiModelTokenPackages) | **Get** /api/v1/cloud-ai/models/{id}/token-packages | Получение пакетов токенов модели
[**GetAiProvider**](AIModelsAPI.md#GetAiProvider) | **Get** /api/v1/cloud-ai/providers/{id} | Получение провайдера моделей
[**GetAiProviders**](AIModelsAPI.md#GetAiProviders) | **Get** /api/v1/cloud-ai/providers | Получение списка провайдеров моделей
[**GetKnowledgebasesTokenPackages**](AIModelsAPI.md#GetKnowledgebasesTokenPackages) | **Get** /api/v1/cloud-ai/token-packages/knowledge-bases | Получение списка пакетов токенов для баз знаний
[**GetModels**](AIModelsAPI.md#GetModels) | **Get** /api/v1/cloud-ai/models | Получение списка моделей (v1)
[**GetModelsV3**](AIModelsAPI.md#GetModelsV3) | **Get** /api/v3/cloud-ai/models | Получение списка моделей



## GetAgentsTokenPackages

> GetAgentsTokenPackages200Response GetAgentsTokenPackages(ctx).Execute()

Получение списка пакетов токенов для агентов



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIModelsAPI.GetAgentsTokenPackages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetAgentsTokenPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentsTokenPackages`: GetAgentsTokenPackages200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetAgentsTokenPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsTokenPackagesRequest struct via the builder pattern


### Return type

[**GetAgentsTokenPackages200Response**](GetAgentsTokenPackages200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiModelParameters

> GetAiModelParameters200Response GetAiModelParameters(ctx).Execute()

Получение справочника характеристик моделей



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIModelsAPI.GetAiModelParameters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetAiModelParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiModelParameters`: GetAiModelParameters200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetAiModelParameters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiModelParametersRequest struct via the builder pattern


### Return type

[**GetAiModelParameters200Response**](GetAiModelParameters200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiModelTokenPackages

> GetAiModelTokenPackages200Response GetAiModelTokenPackages(ctx, id).Execute()

Получение пакетов токенов модели



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(17) // int32 | ID модели

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIModelsAPI.GetAiModelTokenPackages(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetAiModelTokenPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiModelTokenPackages`: GetAiModelTokenPackages200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetAiModelTokenPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID модели | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiModelTokenPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAiModelTokenPackages200Response**](GetAiModelTokenPackages200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiProvider

> GetAiProvider200Response GetAiProvider(ctx, id).Execute()

Получение провайдера моделей



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(1) // int32 | ID провайдера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIModelsAPI.GetAiProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetAiProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiProvider`: GetAiProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetAiProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID провайдера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAiProvider200Response**](GetAiProvider200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiProviders

> GetAiProviders200Response GetAiProviders(ctx).Execute()

Получение списка провайдеров моделей



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIModelsAPI.GetAiProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetAiProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiProviders`: GetAiProviders200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetAiProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiProvidersRequest struct via the builder pattern


### Return type

[**GetAiProviders200Response**](GetAiProviders200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgebasesTokenPackages

> GetKnowledgebasesTokenPackages200Response GetKnowledgebasesTokenPackages(ctx).Execute()

Получение списка пакетов токенов для баз знаний



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIModelsAPI.GetKnowledgebasesTokenPackages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetKnowledgebasesTokenPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKnowledgebasesTokenPackages`: GetKnowledgebasesTokenPackages200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetKnowledgebasesTokenPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebasesTokenPackagesRequest struct via the builder pattern


### Return type

[**GetKnowledgebasesTokenPackages200Response**](GetKnowledgebasesTokenPackages200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModels

> GetModels200Response GetModels(ctx).Execute()

Получение списка моделей (v1)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIModelsAPI.GetModels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetModels`: GetModels200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsRequest struct via the builder pattern


### Return type

[**GetModels200Response**](GetModels200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsV3

> GetModelsV3200Response GetModelsV3(ctx).Execute()

Получение списка моделей



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIModelsAPI.GetModelsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetModelsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetModelsV3`: GetModelsV3200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetModelsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsV3Request struct via the builder pattern


### Return type

[**GetModelsV3200Response**](GetModelsV3200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

