# \AIAgentsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdditionalTokenPackage**](AIAgentsAPI.md#AddAdditionalTokenPackage) | **Post** /api/v1/cloud-ai/agents/{id}/add-additional-token-package | Добавление дополнительного пакета токенов
[**ChangeAgentModel**](AIAgentsAPI.md#ChangeAgentModel) | **Patch** /api/v1/cloud-ai/agents/{id}/model | Смена модели AI агента
[**CreateAgent**](AIAgentsAPI.md#CreateAgent) | **Post** /api/v1/cloud-ai/agents | Создание AI агента
[**CreateAgentEntityLimit**](AIAgentsAPI.md#CreateAgentEntityLimit) | **Post** /api/v1/cloud-ai/agents/{id}/entity-limits | Создание лимита агента
[**CreateAgentV2**](AIAgentsAPI.md#CreateAgentV2) | **Post** /api/v2/cloud-ai/agents | Создание AI агента
[**DeleteAgent**](AIAgentsAPI.md#DeleteAgent) | **Delete** /api/v1/cloud-ai/agents/{id} | Удаление AI агента
[**DeleteAgentEntityLimit**](AIAgentsAPI.md#DeleteAgentEntityLimit) | **Delete** /api/v1/cloud-ai/agents/{id}/entity-limits/{limit_id} | Удаление лимита агента
[**ExportAgentChatsCsv**](AIAgentsAPI.md#ExportAgentChatsCsv) | **Get** /api/v1/cloud-ai/agents/{id}/chats/docs | Выгрузка чатов агента в CSV
[**GetAgent**](AIAgentsAPI.md#GetAgent) | **Get** /api/v1/cloud-ai/agents/{id} | Получение AI агента
[**GetAgentChatList**](AIAgentsAPI.md#GetAgentChatList) | **Get** /api/v1/cloud-ai/agents/{id}/chat-list | Получение списка чатов агента
[**GetAgentChatMessages**](AIAgentsAPI.md#GetAgentChatMessages) | **Get** /api/v1/cloud-ai/agents/{id}/chat-list/{chat_id}/messages | Получение сообщений чата агента
[**GetAgentEntityLimit**](AIAgentsAPI.md#GetAgentEntityLimit) | **Get** /api/v1/cloud-ai/agents/{id}/entity-limits/{limit_id} | Получение лимита агента
[**GetAgentEntityLimits**](AIAgentsAPI.md#GetAgentEntityLimits) | **Get** /api/v1/cloud-ai/agents/{id}/entity-limits | Получение лимитов агента
[**GetAgentStatistics**](AIAgentsAPI.md#GetAgentStatistics) | **Get** /api/v1/cloud-ai/agents/{id}/statistic | Получение статистики использования токенов агента
[**GetAgentV3**](AIAgentsAPI.md#GetAgentV3) | **Get** /api/v3/cloud-ai/agents/{id} | Получение AI агента (v3)
[**GetAgents**](AIAgentsAPI.md#GetAgents) | **Get** /api/v1/cloud-ai/agents | Получение списка AI агентов
[**GetAgentsV3**](AIAgentsAPI.md#GetAgentsV3) | **Get** /api/v3/cloud-ai/agents | Получение списка AI агентов (v3)
[**SyncAgentEntityLimits**](AIAgentsAPI.md#SyncAgentEntityLimits) | **Put** /api/v1/cloud-ai/agents/{id}/entity-limits/sync | Синхронизация лимитов агента за период
[**UpdateAgent**](AIAgentsAPI.md#UpdateAgent) | **Patch** /api/v1/cloud-ai/agents/{id} | Обновление AI агента
[**UpdateAgentEntityLimit**](AIAgentsAPI.md#UpdateAgentEntityLimit) | **Patch** /api/v1/cloud-ai/agents/{id}/entity-limits/{limit_id} | Изменение лимита агента



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


## ChangeAgentModel

> CreateAgent201Response ChangeAgentModel(ctx, id).ChangeAgentModel(changeAgentModel).Execute()

Смена модели AI агента



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
    changeAgentModel := *openapiclient.NewChangeAgentModel(int32(2)) // ChangeAgentModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.ChangeAgentModel(context.Background(), id).ChangeAgentModel(changeAgentModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.ChangeAgentModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeAgentModel`: CreateAgent201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.ChangeAgentModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeAgentModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeAgentModel** | [**ChangeAgentModel**](ChangeAgentModel.md) |  | 

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
    createAgent := *openapiclient.NewCreateAgent("Мой AI помощник", "private", int32(1), *openapiclient.NewAgentSettings(*openapiclient.NewAgentModelSettings(), "Ты — вежливый помощник службы поддержки.", false), int32(1)) // CreateAgent | 

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


## CreateAgentEntityLimit

> CreateAgentEntityLimit201Response CreateAgentEntityLimit(ctx, id).CreateAgentEntityLimit(createAgentEntityLimit).Execute()

Создание лимита агента



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
    createAgentEntityLimit := *openapiclient.NewCreateAgentEntityLimit("day") // CreateAgentEntityLimit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.CreateAgentEntityLimit(context.Background(), id).CreateAgentEntityLimit(createAgentEntityLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.CreateAgentEntityLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentEntityLimit`: CreateAgentEntityLimit201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.CreateAgentEntityLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentEntityLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAgentEntityLimit** | [**CreateAgentEntityLimit**](CreateAgentEntityLimit.md) |  | 

### Return type

[**CreateAgentEntityLimit201Response**](CreateAgentEntityLimit201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgentV2

> CreateAgent201Response CreateAgentV2(ctx).CreateAgentV2(createAgentV2).Execute()

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
    createAgentV2 := *openapiclient.NewCreateAgentV2("Мой AI помощник", "private", int32(1), *openapiclient.NewAgentSettings(*openapiclient.NewAgentModelSettings(), "Ты — вежливый помощник службы поддержки.", false)) // CreateAgentV2 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.CreateAgentV2(context.Background()).CreateAgentV2(createAgentV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.CreateAgentV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentV2`: CreateAgent201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.CreateAgentV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAgentV2** | [**CreateAgentV2**](CreateAgentV2.md) |  | 

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

> DeleteAgent(ctx, id).DeleteAgent(deleteAgent).Execute()

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
    deleteAgent := *openapiclient.NewDeleteAgent() // DeleteAgent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentsAPI.DeleteAgent(context.Background(), id).DeleteAgent(deleteAgent).Execute()
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

 **deleteAgent** | [**DeleteAgent**](DeleteAgent.md) |  | 

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


## DeleteAgentEntityLimit

> DeleteAgentEntityLimit(ctx, id, limitId).Execute()

Удаление лимита агента



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
    limitId := "550e8400-e29b-41d4-a716-446655440000" // string | ID лимита агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentsAPI.DeleteAgentEntityLimit(context.Background(), id, limitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.DeleteAgentEntityLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 
**limitId** | **string** | ID лимита агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentEntityLimitRequest struct via the builder pattern


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


## ExportAgentChatsCsv

> ExportAgentChatsCsv200Response ExportAgentChatsCsv(ctx, id).Offset(offset).Limit(limit).Order(order).From(from).To(to).Locale(locale).Execute()

Выгрузка чатов агента в CSV



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
    offset := int32(0) // int32 | Смещение относительно начала списка (optional) (default to 0)
    limit := int32(100) // int32 | Количество сообщений в ответе (optional) (default to 100)
    order := "desc" // string | Сортировка по дате сообщения (optional) (default to "desc")
    from := time.Now() // time.Time | Начало периода (дата и время в формате ISO 8601) (optional)
    to := time.Now() // time.Time | Конец периода (дата и время в формате ISO 8601) (optional)
    locale := "ru" // string | Язык заголовков и значений CSV (optional) (default to "ru")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.ExportAgentChatsCsv(context.Background(), id).Offset(offset).Limit(limit).Order(order).From(from).To(to).Locale(locale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.ExportAgentChatsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAgentChatsCsv`: ExportAgentChatsCsv200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.ExportAgentChatsCsv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAgentChatsCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Смещение относительно начала списка | [default to 0]
 **limit** | **int32** | Количество сообщений в ответе | [default to 100]
 **order** | **string** | Сортировка по дате сообщения | [default to &quot;desc&quot;]
 **from** | **time.Time** | Начало периода (дата и время в формате ISO 8601) | 
 **to** | **time.Time** | Конец периода (дата и время в формате ISO 8601) | 
 **locale** | **string** | Язык заголовков и значений CSV | [default to &quot;ru&quot;]

### Return type

[**ExportAgentChatsCsv200Response**](ExportAgentChatsCsv200Response.md)

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


## GetAgentChatList

> GetAgentChatList200Response GetAgentChatList(ctx, id).Offset(offset).Limit(limit).Order(order).From(from).To(to).ReactionType(reactionType).ResponseTimeRanges(responseTimeRanges).Execute()

Получение списка чатов агента



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
    offset := int32(0) // int32 | Смещение относительно начала списка (optional) (default to 0)
    limit := int32(100) // int32 | Количество чатов в ответе (optional) (default to 100)
    order := "asc" // string | Сортировка по дате первого сообщения (optional) (default to "asc")
    from := time.Now() // time.Time | Начало периода (дата и время в формате ISO 8601) (optional)
    to := time.Now() // time.Time | Конец периода (дата и время в формате ISO 8601) (optional)
    reactionType := "likes" // string | Фильтр по оценкам: `likes` — есть положительная, `dislikes` — есть отрицательная, `with` — есть любая, `none` — без оценок (optional)
    responseTimeRanges := []string{"Inner_example"} // []string | Фильтр по среднему времени ответа в секундах: диапазоны `мин-макс` или `мин-` («от мин и больше»). Можно передать несколько диапазонов — они объединяются по ИЛИ. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.GetAgentChatList(context.Background(), id).Offset(offset).Limit(limit).Order(order).From(from).To(to).ReactionType(reactionType).ResponseTimeRanges(responseTimeRanges).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentChatList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentChatList`: GetAgentChatList200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentChatList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentChatListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Смещение относительно начала списка | [default to 0]
 **limit** | **int32** | Количество чатов в ответе | [default to 100]
 **order** | **string** | Сортировка по дате первого сообщения | [default to &quot;asc&quot;]
 **from** | **time.Time** | Начало периода (дата и время в формате ISO 8601) | 
 **to** | **time.Time** | Конец периода (дата и время в формате ISO 8601) | 
 **reactionType** | **string** | Фильтр по оценкам: &#x60;likes&#x60; — есть положительная, &#x60;dislikes&#x60; — есть отрицательная, &#x60;with&#x60; — есть любая, &#x60;none&#x60; — без оценок | 
 **responseTimeRanges** | **[]string** | Фильтр по среднему времени ответа в секундах: диапазоны &#x60;мин-макс&#x60; или &#x60;мин-&#x60; («от мин и больше»). Можно передать несколько диапазонов — они объединяются по ИЛИ. | 

### Return type

[**GetAgentChatList200Response**](GetAgentChatList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentChatMessages

> GetAgentChatMessages200Response GetAgentChatMessages(ctx, id, chatId).Offset(offset).Limit(limit).Order(order).From(from).To(to).Execute()

Получение сообщений чата агента



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
    chatId := "ai-agent-1-0b1c2d3e-4f50-4a6b-8c7d-9e0f1a2b3c4d" // string | ID чата (диалога) из поля `id` списка чатов агента
    offset := int32(0) // int32 | Смещение относительно начала списка (optional) (default to 0)
    limit := int32(100) // int32 | Количество сообщений в ответе (optional) (default to 100)
    order := "desc" // string | Сортировка по дате сообщения (optional) (default to "desc")
    from := time.Now() // time.Time | Начало периода (дата и время в формате ISO 8601) (optional)
    to := time.Now() // time.Time | Конец периода (дата и время в формате ISO 8601) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.GetAgentChatMessages(context.Background(), id, chatId).Offset(offset).Limit(limit).Order(order).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentChatMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentChatMessages`: GetAgentChatMessages200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentChatMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 
**chatId** | **string** | ID чата (диалога) из поля &#x60;id&#x60; списка чатов агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentChatMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | Смещение относительно начала списка | [default to 0]
 **limit** | **int32** | Количество сообщений в ответе | [default to 100]
 **order** | **string** | Сортировка по дате сообщения | [default to &quot;desc&quot;]
 **from** | **time.Time** | Начало периода (дата и время в формате ISO 8601) | 
 **to** | **time.Time** | Конец периода (дата и время в формате ISO 8601) | 

### Return type

[**GetAgentChatMessages200Response**](GetAgentChatMessages200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentEntityLimit

> CreateAgentEntityLimit201Response GetAgentEntityLimit(ctx, id, limitId).Execute()

Получение лимита агента



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
    limitId := "550e8400-e29b-41d4-a716-446655440000" // string | ID лимита агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.GetAgentEntityLimit(context.Background(), id, limitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentEntityLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentEntityLimit`: CreateAgentEntityLimit201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentEntityLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 
**limitId** | **string** | ID лимита агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentEntityLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateAgentEntityLimit201Response**](CreateAgentEntityLimit201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentEntityLimits

> GetAgentEntityLimits200Response GetAgentEntityLimits(ctx, id).LimitType(limitType).Execute()

Получение лимитов агента



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
    limitType := "day" // string | Фильтр по периоду лимита (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.GetAgentEntityLimits(context.Background(), id).LimitType(limitType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentEntityLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentEntityLimits`: GetAgentEntityLimits200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentEntityLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentEntityLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limitType** | **string** | Фильтр по периоду лимита | 

### Return type

[**GetAgentEntityLimits200Response**](GetAgentEntityLimits200Response.md)

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
    startTime := time.Now() // time.Time | Начало периода (ISO 8601). По умолчанию — за 24 часа до `endTime`. (optional)
    endTime := time.Now() // time.Time | Конец периода (ISO 8601). По умолчанию — текущий момент. (optional)
    interval := int32(60) // int32 | Длина интервала агрегации в минутах (optional) (default to 60)

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

 **startTime** | **time.Time** | Начало периода (ISO 8601). По умолчанию — за 24 часа до &#x60;endTime&#x60;. | 
 **endTime** | **time.Time** | Конец периода (ISO 8601). По умолчанию — текущий момент. | 
 **interval** | **int32** | Длина интервала агрегации в минутах | [default to 60]

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


## GetAgentV3

> GetAgentV3200Response GetAgentV3(ctx, id).Execute()

Получение AI агента (v3)



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
    resp, r, err := apiClient.AIAgentsAPI.GetAgentV3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentV3`: GetAgentV3200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentV3200Response**](GetAgentV3200Response.md)

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


## GetAgentsV3

> GetAgentsV3200Response GetAgentsV3(ctx).Execute()

Получение списка AI агентов (v3)



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
    resp, r, err := apiClient.AIAgentsAPI.GetAgentsV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentsV3`: GetAgentsV3200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentsV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsV3Request struct via the builder pattern


### Return type

[**GetAgentsV3200Response**](GetAgentsV3200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAgentEntityLimits

> GetAgentEntityLimits200Response SyncAgentEntityLimits(ctx, id).SyncAgentEntityLimits(syncAgentEntityLimits).Execute()

Синхронизация лимитов агента за период



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
    syncAgentEntityLimits := *openapiclient.NewSyncAgentEntityLimits("day", []openapiclient.SyncAgentEntityLimitsLimitsInner{*openapiclient.NewSyncAgentEntityLimitsLimitsInner(NullableInt32(1))}) // SyncAgentEntityLimits | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.SyncAgentEntityLimits(context.Background(), id).SyncAgentEntityLimits(syncAgentEntityLimits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.SyncAgentEntityLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncAgentEntityLimits`: GetAgentEntityLimits200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.SyncAgentEntityLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAgentEntityLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncAgentEntityLimits** | [**SyncAgentEntityLimits**](SyncAgentEntityLimits.md) |  | 

### Return type

[**GetAgentEntityLimits200Response**](GetAgentEntityLimits200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
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


## UpdateAgentEntityLimit

> CreateAgentEntityLimit201Response UpdateAgentEntityLimit(ctx, id, limitId).UpdateAgentEntityLimit(updateAgentEntityLimit).Execute()

Изменение лимита агента



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
    limitId := "550e8400-e29b-41d4-a716-446655440000" // string | ID лимита агента
    updateAgentEntityLimit := *openapiclient.NewUpdateAgentEntityLimit() // UpdateAgentEntityLimit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentsAPI.UpdateAgentEntityLimit(context.Background(), id, limitId).UpdateAgentEntityLimit(updateAgentEntityLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.UpdateAgentEntityLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentEntityLimit`: CreateAgentEntityLimit201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.UpdateAgentEntityLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID агента | 
**limitId** | **string** | ID лимита агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentEntityLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAgentEntityLimit** | [**UpdateAgentEntityLimit**](UpdateAgentEntityLimit.md) |  | 

### Return type

[**CreateAgentEntityLimit201Response**](CreateAgentEntityLimit201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

