# \AIModelsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentsTokenPackages**](AIModelsAPI.md#GetAgentsTokenPackages) | **Get** /api/v1/cloud-ai/token-packages/agents | Получение списка пакетов токенов для агентов
[**GetKnowledgebasesTokenPackages**](AIModelsAPI.md#GetKnowledgebasesTokenPackages) | **Get** /api/v1/cloud-ai/token-packages/knowledge-bases | Получение списка пакетов токенов для баз знаний
[**GetModels**](AIModelsAPI.md#GetModels) | **Get** /api/v1/cloud-ai/models | Получение списка моделей



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


## GetKnowledgebasesTokenPackages

> GetAgentsTokenPackages200Response GetKnowledgebasesTokenPackages(ctx).Execute()

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
    // response from `GetKnowledgebasesTokenPackages`: GetAgentsTokenPackages200Response
    fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetKnowledgebasesTokenPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebasesTokenPackagesRequest struct via the builder pattern


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


## GetModels

> GetModels200Response GetModels(ctx).Execute()

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

