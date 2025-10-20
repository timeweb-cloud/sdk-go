# \AIAgentsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdditionalTokenPackage**](AIAgentsAPI.md#AddAdditionalTokenPackage) | **Post** /api/v1/cloud-ai/agents/{id}/add-additional-token-package | Добавление дополнительного пакета токенов
[**CreateAgent**](AIAgentsAPI.md#CreateAgent) | **Post** /api/v1/cloud-ai/agents | Создание AI агента
[**DeleteAgent**](AIAgentsAPI.md#DeleteAgent) | **Delete** /api/v1/cloud-ai/agents/{id} | Удаление AI агента
[**GetAgent**](AIAgentsAPI.md#GetAgent) | **Get** /api/v1/cloud-ai/agents/{id} | Получение AI агента
[**GetAgentStatistics**](AIAgentsAPI.md#GetAgentStatistics) | **Get** /api/v1/cloud-ai/agents/{id}/statistic | Получение статистики использования токенов агента
[**GetAgents**](AIAgentsAPI.md#GetAgents) | **Get** /api/v1/cloud-ai/agents | Получение списка AI агентов
[**UpdateAgent**](AIAgentsAPI.md#UpdateAgent) | **Patch** /api/v1/cloud-ai/agents/{id} | Обновление AI агента



## AddAdditionalTokenPackage

> AddAdditionalTokenPackage(ctx, id).AddTokenPackage(addTokenPackage).Execute()

Добавление дополнительного пакета токенов



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
    id := int32(1) // int32 | ID агента
    addTokenPackage := *openapiclient.NewAddTokenPackage() // AddTokenPackage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentsAPI.AddAdditionalTokenPackage(context.Background(), id).AddTokenPackage(addTokenPackage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AddAdditionalTokenPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAdditionalTokenPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addTokenPackage** | [**AddTokenPackage**](AddTokenPackage.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgent

> CreateAgent201Response CreateAgent(ctx).CreateAgent(createAgent).Execute()

Создание AI агента



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
    createAgent := *openapiclient.NewCreateAgent("Мой AI помощник", "private", float32(1), float32(1), *openapiclient.NewAgentSettings(*openapiclient.NewAgentModelSettings(), "SystemPrompt_example", false)) // CreateAgent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.CreateAgent(context.Background()).CreateAgent(createAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.CreateAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgent`: CreateAgent201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.CreateAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAgent** | [**CreateAgent**](CreateAgent.md) |  | 

### Return type

[**CreateAgent201Response**](CreateAgent201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgent

> DeleteAgent(ctx, id).Execute()

Удаление AI агента



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
    id := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentsAPI.DeleteAgent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.DeleteAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgent

> CreateAgent201Response GetAgent(ctx, id).Execute()

Получение AI агента



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
    id := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.GetAgent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgent`: CreateAgent201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAgent201Response**](CreateAgent201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentStatistics

> GetAgentStatistics200Response GetAgentStatistics(ctx, id).StartTime(startTime).EndTime(endTime).Interval(interval).Execute()

Получение статистики использования токенов агента



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(1) // int32 | ID агента
    startTime := time.Now() // time.Time | Начало временного диапазона (ISO 8601) (optional)
    endTime := time.Now() // time.Time | Конец временного диапазона (ISO 8601) (optional)
    interval := float32(60) // float32 | Интервал в минутах (по умолчанию 60) (optional) (default to 60)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.GetAgentStatistics(context.Background(), id).StartTime(startTime).EndTime(endTime).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentStatistics`: GetAgentStatistics200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **time.Time** | Начало временного диапазона (ISO 8601) | 
 **endTime** | **time.Time** | Конец временного диапазона (ISO 8601) | 
 **interval** | **float32** | Интервал в минутах (по умолчанию 60) | [default to 60]

### Return type

[**GetAgentStatistics200Response**](GetAgentStatistics200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgents

> GetAgents200Response GetAgents(ctx).Execute()

Получение списка AI агентов



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
    resp, r, err := apiClient.AIAgentsAPI.GetAgents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgents`: GetAgents200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRequest struct via the builder pattern


### Return type

[**GetAgents200Response**](GetAgents200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgent

> CreateAgent201Response UpdateAgent(ctx, id).UpdateAgent(updateAgent).Execute()

Обновление AI агента



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
    id := int32(1) // int32 | ID агента
    updateAgent := *openapiclient.NewUpdateAgent() // UpdateAgent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.UpdateAgent(context.Background(), id).UpdateAgent(updateAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.UpdateAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgent`: CreateAgent201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.UpdateAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAgent** | [**UpdateAgent**](UpdateAgent.md) |  | 

### Return type

[**CreateAgent201Response**](CreateAgent201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

