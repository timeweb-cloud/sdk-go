# \AIGatewayAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAiKey**](AIGatewayAPI.md#CreateAiKey) | **Post** /api/v1/cloud-ai/ai-keys | Создание ключа AI Gateway
[**CreateAiKeyLimit**](AIGatewayAPI.md#CreateAiKeyLimit) | **Post** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits | Создание лимита ключа AI Gateway
[**DeleteAiKey**](AIGatewayAPI.md#DeleteAiKey) | **Delete** /api/v1/cloud-ai/ai-keys/{key_id} | Удаление ключа AI Gateway
[**DeleteAiKeyLimit**](AIGatewayAPI.md#DeleteAiKeyLimit) | **Delete** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits/{limit_id} | Удаление лимита ключа AI Gateway
[**GetAiKeyCallsStatistic**](AIGatewayAPI.md#GetAiKeyCallsStatistic) | **Get** /api/v1/cloud-ai/ai-keys/statistic/calls | Получение статистики запросов ключа AI Gateway
[**GetAiKeyLimit**](AIGatewayAPI.md#GetAiKeyLimit) | **Get** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits/{limit_id} | Получение лимита ключа AI Gateway
[**GetAiKeyLimits**](AIGatewayAPI.md#GetAiKeyLimits) | **Get** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits | Получение списка лимитов ключа AI Gateway
[**GetAiKeyModelsStatistic**](AIGatewayAPI.md#GetAiKeyModelsStatistic) | **Get** /api/v1/cloud-ai/ai-keys/statistic/models | Получение списка моделей, использованных ключом AI Gateway
[**GetAiKeyTokenUsageStatistic**](AIGatewayAPI.md#GetAiKeyTokenUsageStatistic) | **Get** /api/v1/cloud-ai/ai-keys/statistic/token-usage | Получение статистики расхода токенов ключа AI Gateway
[**GetAiKeys**](AIGatewayAPI.md#GetAiKeys) | **Get** /api/v1/cloud-ai/ai-keys | Получение списка ключей AI Gateway
[**RegenerateAiKey**](AIGatewayAPI.md#RegenerateAiKey) | **Put** /api/v1/cloud-ai/ai-keys/{key_id}/regenerate | Перевыпуск ключа AI Gateway
[**SyncAiKeyLimits**](AIGatewayAPI.md#SyncAiKeyLimits) | **Put** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits/sync | Синхронизация лимитов ключа AI Gateway
[**UpdateAiKey**](AIGatewayAPI.md#UpdateAiKey) | **Patch** /api/v1/cloud-ai/ai-keys/{key_id} | Изменение ключа AI Gateway
[**UpdateAiKeyLimit**](AIGatewayAPI.md#UpdateAiKeyLimit) | **Patch** /api/v1/cloud-ai/ai-keys/{key_id}/entity-limits/{limit_id} | Изменение лимита ключа AI Gateway



## CreateAiKey

> CreateAiKey201Response CreateAiKey(ctx).CreateAiKey(createAiKey).Execute()

Создание ключа AI Gateway



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
    createAiKey := *openapiclient.NewCreateAiKey("Ключ для бота", time.Now()) // CreateAiKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.CreateAiKey(context.Background()).CreateAiKey(createAiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.CreateAiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiKey`: CreateAiKey201Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.CreateAiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAiKey** | [**CreateAiKey**](CreateAiKey.md) |  | 

### Return type

[**CreateAiKey201Response**](CreateAiKey201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiKeyLimit

> CreateAiKeyLimit201Response CreateAiKeyLimit(ctx, keyId).CreateAiKeyLimit(createAiKeyLimit).Execute()

Создание лимита ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)
    createAiKeyLimit := *openapiclient.NewCreateAiKeyLimit("aikey", "day") // CreateAiKeyLimit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.CreateAiKeyLimit(context.Background(), keyId).CreateAiKeyLimit(createAiKeyLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.CreateAiKeyLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiKeyLimit`: CreateAiKeyLimit201Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.CreateAiKeyLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiKeyLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAiKeyLimit** | [**CreateAiKeyLimit**](CreateAiKeyLimit.md) |  | 

### Return type

[**CreateAiKeyLimit201Response**](CreateAiKeyLimit201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiKey

> DeleteAiKey(ctx, keyId).Execute()

Удаление ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIGatewayAPI.DeleteAiKey(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.DeleteAiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiKeyRequest struct via the builder pattern


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


## DeleteAiKeyLimit

> DeleteAiKeyLimit(ctx, keyId, limitId).Execute()

Удаление лимита ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)
    limitId := "0f6b7d8e-1c2a-4b3d-9e8f-7a6b5c4d3e2f" // string | ID лимита (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIGatewayAPI.DeleteAiKeyLimit(context.Background(), keyId, limitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.DeleteAiKeyLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 
**limitId** | **string** | ID лимита (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiKeyLimitRequest struct via the builder pattern


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


## GetAiKeyCallsStatistic

> GetAiKeyCallsStatistic200Response GetAiKeyCallsStatistic(ctx).EntityId(entityId).EntityType(entityType).ModelName(modelName).StartTime(startTime).EndTime(endTime).Interval(interval).Execute()

Получение статистики запросов ключа AI Gateway



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
    entityId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа — поле `id` из списка ключей.
    entityType := "aikey" // string | Тип ключа — поле `entity_type` из списка ключей: `aikey` — ключ, созданный пользователем, `k8s` — ключ кластера Kubernetes. Если тип не совпадает с типом ключа, вернётся 404.
    modelName := "openai/gpt-5.4-mini,openai/gpt-4.1-mini" // string | Фильтр по моделям: ID модели или несколько ID через запятую. Значение другого формата (например, с заглавными буквами) игнорируется — статистика возвращается по всем моделям. (optional)
    startTime := time.Now() // time.Time | Начало периода в формате ISO 8601. По умолчанию — за 24 часа до `end_time`. (optional)
    endTime := time.Now() // time.Time | Конец периода в формате ISO 8601. По умолчанию — текущий момент. (optional)
    interval := int32(60) // int32 | Шаг агрегации в минутах. (optional) (default to 60)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.GetAiKeyCallsStatistic(context.Background()).EntityId(entityId).EntityType(entityType).ModelName(modelName).StartTime(startTime).EndTime(endTime).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.GetAiKeyCallsStatistic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiKeyCallsStatistic`: GetAiKeyCallsStatistic200Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.GetAiKeyCallsStatistic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAiKeyCallsStatisticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | ID ключа — поле &#x60;id&#x60; из списка ключей. | 
 **entityType** | **string** | Тип ключа — поле &#x60;entity_type&#x60; из списка ключей: &#x60;aikey&#x60; — ключ, созданный пользователем, &#x60;k8s&#x60; — ключ кластера Kubernetes. Если тип не совпадает с типом ключа, вернётся 404. | 
 **modelName** | **string** | Фильтр по моделям: ID модели или несколько ID через запятую. Значение другого формата (например, с заглавными буквами) игнорируется — статистика возвращается по всем моделям. | 
 **startTime** | **time.Time** | Начало периода в формате ISO 8601. По умолчанию — за 24 часа до &#x60;end_time&#x60;. | 
 **endTime** | **time.Time** | Конец периода в формате ISO 8601. По умолчанию — текущий момент. | 
 **interval** | **int32** | Шаг агрегации в минутах. | [default to 60]

### Return type

[**GetAiKeyCallsStatistic200Response**](GetAiKeyCallsStatistic200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiKeyLimit

> CreateAiKeyLimit201Response GetAiKeyLimit(ctx, keyId, limitId).Execute()

Получение лимита ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)
    limitId := "0f6b7d8e-1c2a-4b3d-9e8f-7a6b5c4d3e2f" // string | ID лимита (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.GetAiKeyLimit(context.Background(), keyId, limitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.GetAiKeyLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiKeyLimit`: CreateAiKeyLimit201Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.GetAiKeyLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 
**limitId** | **string** | ID лимита (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiKeyLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateAiKeyLimit201Response**](CreateAiKeyLimit201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiKeyLimits

> GetAiKeyLimits200Response GetAiKeyLimits(ctx, keyId).EntityType(entityType).LimitType(limitType).Execute()

Получение списка лимитов ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)
    entityType := "aikey" // string | Фильтр по типу сущности лимита (`entity_type`). (optional)
    limitType := "day" // string | Фильтр по периоду лимита. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.GetAiKeyLimits(context.Background(), keyId).EntityType(entityType).LimitType(limitType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.GetAiKeyLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiKeyLimits`: GetAiKeyLimits200Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.GetAiKeyLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiKeyLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityType** | **string** | Фильтр по типу сущности лимита (&#x60;entity_type&#x60;). | 
 **limitType** | **string** | Фильтр по периоду лимита. | 

### Return type

[**GetAiKeyLimits200Response**](GetAiKeyLimits200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiKeyModelsStatistic

> GetAiKeyModelsStatistic200Response GetAiKeyModelsStatistic(ctx).EntityId(entityId).EntityType(entityType).ModelName(modelName).StartTime(startTime).EndTime(endTime).Execute()

Получение списка моделей, использованных ключом AI Gateway



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
    entityId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа — поле `id` из списка ключей.
    entityType := "aikey" // string | Тип ключа — поле `entity_type` из списка ключей: `aikey` — ключ, созданный пользователем, `k8s` — ключ кластера Kubernetes. Если тип не совпадает с типом ключа, вернётся 404.
    modelName := "openai/gpt-5.4-mini,openai/gpt-4.1-mini" // string | Фильтр по моделям: ID модели или несколько ID через запятую. Значение другого формата (например, с заглавными буквами) игнорируется — статистика возвращается по всем моделям. (optional)
    startTime := time.Now() // time.Time | Начало периода в формате ISO 8601. По умолчанию — за 24 часа до `end_time`. (optional)
    endTime := time.Now() // time.Time | Конец периода в формате ISO 8601. По умолчанию — текущий момент. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.GetAiKeyModelsStatistic(context.Background()).EntityId(entityId).EntityType(entityType).ModelName(modelName).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.GetAiKeyModelsStatistic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiKeyModelsStatistic`: GetAiKeyModelsStatistic200Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.GetAiKeyModelsStatistic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAiKeyModelsStatisticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | ID ключа — поле &#x60;id&#x60; из списка ключей. | 
 **entityType** | **string** | Тип ключа — поле &#x60;entity_type&#x60; из списка ключей: &#x60;aikey&#x60; — ключ, созданный пользователем, &#x60;k8s&#x60; — ключ кластера Kubernetes. Если тип не совпадает с типом ключа, вернётся 404. | 
 **modelName** | **string** | Фильтр по моделям: ID модели или несколько ID через запятую. Значение другого формата (например, с заглавными буквами) игнорируется — статистика возвращается по всем моделям. | 
 **startTime** | **time.Time** | Начало периода в формате ISO 8601. По умолчанию — за 24 часа до &#x60;end_time&#x60;. | 
 **endTime** | **time.Time** | Конец периода в формате ISO 8601. По умолчанию — текущий момент. | 

### Return type

[**GetAiKeyModelsStatistic200Response**](GetAiKeyModelsStatistic200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiKeyTokenUsageStatistic

> GetAiKeyTokenUsageStatistic200Response GetAiKeyTokenUsageStatistic(ctx).EntityId(entityId).EntityType(entityType).ModelName(modelName).StartTime(startTime).EndTime(endTime).Interval(interval).Execute()

Получение статистики расхода токенов ключа AI Gateway



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
    entityId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа — поле `id` из списка ключей.
    entityType := "aikey" // string | Тип ключа — поле `entity_type` из списка ключей: `aikey` — ключ, созданный пользователем, `k8s` — ключ кластера Kubernetes. Если тип не совпадает с типом ключа, вернётся 404.
    modelName := "openai/gpt-5.4-mini,openai/gpt-4.1-mini" // string | Фильтр по моделям: ID модели или несколько ID через запятую. Значение другого формата (например, с заглавными буквами) игнорируется — статистика возвращается по всем моделям. (optional)
    startTime := time.Now() // time.Time | Начало периода в формате ISO 8601. По умолчанию — за 24 часа до `end_time`. (optional)
    endTime := time.Now() // time.Time | Конец периода в формате ISO 8601. По умолчанию — текущий момент. (optional)
    interval := int32(60) // int32 | Шаг агрегации в минутах. (optional) (default to 60)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.GetAiKeyTokenUsageStatistic(context.Background()).EntityId(entityId).EntityType(entityType).ModelName(modelName).StartTime(startTime).EndTime(endTime).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.GetAiKeyTokenUsageStatistic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiKeyTokenUsageStatistic`: GetAiKeyTokenUsageStatistic200Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.GetAiKeyTokenUsageStatistic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAiKeyTokenUsageStatisticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | ID ключа — поле &#x60;id&#x60; из списка ключей. | 
 **entityType** | **string** | Тип ключа — поле &#x60;entity_type&#x60; из списка ключей: &#x60;aikey&#x60; — ключ, созданный пользователем, &#x60;k8s&#x60; — ключ кластера Kubernetes. Если тип не совпадает с типом ключа, вернётся 404. | 
 **modelName** | **string** | Фильтр по моделям: ID модели или несколько ID через запятую. Значение другого формата (например, с заглавными буквами) игнорируется — статистика возвращается по всем моделям. | 
 **startTime** | **time.Time** | Начало периода в формате ISO 8601. По умолчанию — за 24 часа до &#x60;end_time&#x60;. | 
 **endTime** | **time.Time** | Конец периода в формате ISO 8601. По умолчанию — текущий момент. | 
 **interval** | **int32** | Шаг агрегации в минутах. | [default to 60]

### Return type

[**GetAiKeyTokenUsageStatistic200Response**](GetAiKeyTokenUsageStatistic200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiKeys

> GetAiKeys200Response GetAiKeys(ctx).Execute()

Получение списка ключей AI Gateway



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
    resp, r, err := apiClient.AIGatewayAPI.GetAiKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.GetAiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiKeys`: GetAiKeys200Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.GetAiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiKeysRequest struct via the builder pattern


### Return type

[**GetAiKeys200Response**](GetAiKeys200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateAiKey

> CreateAiKey201Response RegenerateAiKey(ctx, keyId).Execute()

Перевыпуск ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.RegenerateAiKey(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.RegenerateAiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegenerateAiKey`: CreateAiKey201Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.RegenerateAiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateAiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAiKey201Response**](CreateAiKey201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAiKeyLimits

> GetAiKeyLimits200Response SyncAiKeyLimits(ctx, keyId).SyncAiKeyLimits(syncAiKeyLimits).Execute()

Синхронизация лимитов ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)
    syncAiKeyLimits := *openapiclient.NewSyncAiKeyLimits([]openapiclient.SyncAiKeyLimitsLimitsInner{*openapiclient.NewSyncAiKeyLimitsLimitsInner(NullableInt32(1))}, "day") // SyncAiKeyLimits | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.SyncAiKeyLimits(context.Background(), keyId).SyncAiKeyLimits(syncAiKeyLimits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.SyncAiKeyLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncAiKeyLimits`: GetAiKeyLimits200Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.SyncAiKeyLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAiKeyLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncAiKeyLimits** | [**SyncAiKeyLimits**](SyncAiKeyLimits.md) |  | 

### Return type

[**GetAiKeyLimits200Response**](GetAiKeyLimits200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAiKey

> UpdateAiKey200Response UpdateAiKey(ctx, keyId).UpdateAiKey(updateAiKey).Execute()

Изменение ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)
    updateAiKey := *openapiclient.NewUpdateAiKey() // UpdateAiKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.UpdateAiKey(context.Background(), keyId).UpdateAiKey(updateAiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.UpdateAiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAiKey`: UpdateAiKey200Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.UpdateAiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAiKey** | [**UpdateAiKey**](UpdateAiKey.md) |  | 

### Return type

[**UpdateAiKey200Response**](UpdateAiKey200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAiKeyLimit

> CreateAiKeyLimit201Response UpdateAiKeyLimit(ctx, keyId, limitId).UpdateAiKeyLimit(updateAiKeyLimit).Execute()

Изменение лимита ключа AI Gateway



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
    keyId := "550e8400-e29b-41d4-a716-446655440000" // string | ID ключа AI Gateway (UUID)
    limitId := "0f6b7d8e-1c2a-4b3d-9e8f-7a6b5c4d3e2f" // string | ID лимита (UUID)
    updateAiKeyLimit := *openapiclient.NewUpdateAiKeyLimit() // UpdateAiKeyLimit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIGatewayAPI.UpdateAiKeyLimit(context.Background(), keyId, limitId).UpdateAiKeyLimit(updateAiKeyLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIGatewayAPI.UpdateAiKeyLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAiKeyLimit`: CreateAiKeyLimit201Response
    fmt.Fprintf(os.Stdout, "Response from `AIGatewayAPI.UpdateAiKeyLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | ID ключа AI Gateway (UUID) | 
**limitId** | **string** | ID лимита (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAiKeyLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAiKeyLimit** | [**UpdateAiKeyLimit**](UpdateAiKeyLimit.md) |  | 

### Return type

[**CreateAiKeyLimit201Response**](CreateAiKeyLimit201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

