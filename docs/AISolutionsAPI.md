# \AISolutionsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAiSolutionTelegramBot**](AISolutionsAPI.md#AddAiSolutionTelegramBot) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/telegram/bots | Подключение Telegram-бота к AI-решению
[**ApproveAiSolutionTelegramPairing**](AISolutionsAPI.md#ApproveAiSolutionTelegramPairing) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/telegram/bots/{account_key}/pairing/approve | Подтверждение владельца Telegram-бота AI-решения
[**CreateAiSolution**](AISolutionsAPI.md#CreateAiSolution) | **Post** /api/v1/cloud-ai/solutions | Создание AI-решения
[**DeleteAiSolution**](AISolutionsAPI.md#DeleteAiSolution) | **Delete** /api/v1/cloud-ai/solutions/{solution_id} | Удаление AI-решения
[**GetAiSolution**](AISolutionsAPI.md#GetAiSolution) | **Get** /api/v1/cloud-ai/solutions/{solution_id} | Получение AI-решения
[**GetAiSolutionChatFullHistory**](AISolutionsAPI.md#GetAiSolutionChatFullHistory) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/history | Получение всей истории чата AI-решения
[**GetAiSolutionChatHistory**](AISolutionsAPI.md#GetAiSolutionChatHistory) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/conversations/{conversation_id}/messages | Получение истории разговора AI-решения
[**GetAiSolutionChatStatus**](AISolutionsAPI.md#GetAiSolutionChatStatus) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/status | Получение статуса чата AI-решения
[**GetAiSolutionMainChatHistory**](AISolutionsAPI.md#GetAiSolutionMainChatHistory) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/messages | Получение истории основного разговора AI-решения
[**ListAiSolutionConversations**](AISolutionsAPI.md#ListAiSolutionConversations) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/chat/conversations | Получение списка разговоров AI-решения
[**ListAiSolutionTelegramBots**](AISolutionsAPI.md#ListAiSolutionTelegramBots) | **Get** /api/v1/cloud-ai/solutions/{solution_id}/telegram/bots | Получение списка Telegram-ботов AI-решения
[**ListAiSolutionTypes**](AISolutionsAPI.md#ListAiSolutionTypes) | **Get** /api/v1/cloud-ai/solutions/types | Получение списка типов AI-решений
[**ListAiSolutions**](AISolutionsAPI.md#ListAiSolutions) | **Get** /api/v1/cloud-ai/solutions | Получение списка AI-решений
[**RegenerateAiSolutionAiKey**](AISolutionsAPI.md#RegenerateAiSolutionAiKey) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/ai-key/regenerate | Перевыпуск ключа AI Gateway решения
[**RemoveAiSolutionTelegramBot**](AISolutionsAPI.md#RemoveAiSolutionTelegramBot) | **Delete** /api/v1/cloud-ai/solutions/{solution_id}/telegram/bots/{account_key} | Отключение Telegram-бота от AI-решения
[**ResetAiSolutionChat**](AISolutionsAPI.md#ResetAiSolutionChat) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/chat/reset | Сброс контекста разговора AI-решения
[**SendAiSolutionMessage**](AISolutionsAPI.md#SendAiSolutionMessage) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/chat/messages | Отправка сообщения агенту AI-решения
[**ShutdownAiSolution**](AISolutionsAPI.md#ShutdownAiSolution) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/shutdown | Выключение AI-решения
[**StartAiSolution**](AISolutionsAPI.md#StartAiSolution) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/start | Запуск AI-решения
[**StopAiSolutionChat**](AISolutionsAPI.md#StopAiSolutionChat) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/chat/stop | Остановка хода агента AI-решения
[**StreamAiSolutionChat**](AISolutionsAPI.md#StreamAiSolutionChat) | **Post** /api/v1/cloud-ai/solutions/{solution_id}/chat/stream | Отправка сообщения агенту AI-решения с потоковым ответом
[**UpdateAiSolution**](AISolutionsAPI.md#UpdateAiSolution) | **Patch** /api/v1/cloud-ai/solutions/{solution_id} | Изменение AI-решения



## AddAiSolutionTelegramBot

> AddAiSolutionTelegramBot201Response AddAiSolutionTelegramBot(ctx, solutionId).ConnectTelegram(connectTelegram).Execute()

Подключение Telegram-бота к AI-решению



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
    solutionId := int32(15) // int32 | ID AI-решения
    connectTelegram := *openapiclient.NewConnectTelegram("123456789:AAExampleTokenWithLettersAndDigits0") // ConnectTelegram | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.AddAiSolutionTelegramBot(context.Background(), solutionId).ConnectTelegram(connectTelegram).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.AddAiSolutionTelegramBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAiSolutionTelegramBot`: AddAiSolutionTelegramBot201Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.AddAiSolutionTelegramBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAiSolutionTelegramBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectTelegram** | [**ConnectTelegram**](ConnectTelegram.md) |  | 

### Return type

[**AddAiSolutionTelegramBot201Response**](AddAiSolutionTelegramBot201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveAiSolutionTelegramPairing

> ApproveAiSolutionTelegramPairing(ctx, solutionId, accountKey).ApproveTelegramPairing(approveTelegramPairing).Execute()

Подтверждение владельца Telegram-бота AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    accountKey := "b123456789" // string | ID Telegram-бота в решении (`account_key` из списка ботов) в формате `b<число>`, например `b123456789`
    approveTelegramPairing := *openapiclient.NewApproveTelegramPairing("A1B2C3D4") // ApproveTelegramPairing | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AISolutionsAPI.ApproveAiSolutionTelegramPairing(context.Background(), solutionId, accountKey).ApproveTelegramPairing(approveTelegramPairing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.ApproveAiSolutionTelegramPairing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 
**accountKey** | **string** | ID Telegram-бота в решении (&#x60;account_key&#x60; из списка ботов) в формате &#x60;b&lt;число&gt;&#x60;, например &#x60;b123456789&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveAiSolutionTelegramPairingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **approveTelegramPairing** | [**ApproveTelegramPairing**](ApproveTelegramPairing.md) |  | 

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


## CreateAiSolution

> CreateAiSolution201Response CreateAiSolution(ctx).CreateAiSolution(createAiSolution).Execute()

Создание AI-решения



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
    createAiSolution := *openapiclient.NewCreateAiSolution("openclaw", "Мой агент", int32(12), int32(4531)) // CreateAiSolution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.CreateAiSolution(context.Background()).CreateAiSolution(createAiSolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.CreateAiSolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiSolution`: CreateAiSolution201Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.CreateAiSolution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiSolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAiSolution** | [**CreateAiSolution**](CreateAiSolution.md) |  | 

### Return type

[**CreateAiSolution201Response**](CreateAiSolution201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiSolution

> DeleteAiSolution(ctx, solutionId).Execute()

Удаление AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AISolutionsAPI.DeleteAiSolution(context.Background(), solutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.DeleteAiSolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiSolutionRequest struct via the builder pattern


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


## GetAiSolution

> CreateAiSolution201Response GetAiSolution(ctx, solutionId).Execute()

Получение AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.GetAiSolution(context.Background(), solutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.GetAiSolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiSolution`: CreateAiSolution201Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.GetAiSolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAiSolution201Response**](CreateAiSolution201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiSolutionChatFullHistory

> GetAiSolutionMainChatHistory200Response GetAiSolutionChatFullHistory(ctx, solutionId).Limit(limit).Offset(offset).Execute()

Получение всей истории чата AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    limit := int32(100) // int32 | Количество сообщений на странице, от 1 до 10 000 (optional) (default to 100)
    offset := int32(0) // int32 | Количество сообщений, которые нужно пропустить (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.GetAiSolutionChatFullHistory(context.Background(), solutionId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.GetAiSolutionChatFullHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiSolutionChatFullHistory`: GetAiSolutionMainChatHistory200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.GetAiSolutionChatFullHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSolutionChatFullHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Количество сообщений на странице, от 1 до 10 000 | [default to 100]
 **offset** | **int32** | Количество сообщений, которые нужно пропустить | [default to 0]

### Return type

[**GetAiSolutionMainChatHistory200Response**](GetAiSolutionMainChatHistory200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiSolutionChatHistory

> GetAiSolutionMainChatHistory200Response GetAiSolutionChatHistory(ctx, solutionId, conversationId).Limit(limit).Offset(offset).Execute()

Получение истории разговора AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    conversationId := "solution-15" // string | ID разговора. У основного разговора решения — `solution-<solution_id>`
    limit := int32(100) // int32 | Количество сообщений на странице, от 1 до 10 000 (optional) (default to 100)
    offset := int32(0) // int32 | Количество сообщений, которые нужно пропустить (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.GetAiSolutionChatHistory(context.Background(), solutionId, conversationId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.GetAiSolutionChatHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiSolutionChatHistory`: GetAiSolutionMainChatHistory200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.GetAiSolutionChatHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 
**conversationId** | **string** | ID разговора. У основного разговора решения — &#x60;solution-&lt;solution_id&gt;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSolutionChatHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Количество сообщений на странице, от 1 до 10 000 | [default to 100]
 **offset** | **int32** | Количество сообщений, которые нужно пропустить | [default to 0]

### Return type

[**GetAiSolutionMainChatHistory200Response**](GetAiSolutionMainChatHistory200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiSolutionChatStatus

> GetAiSolutionChatStatus200Response GetAiSolutionChatStatus(ctx, solutionId).Execute()

Получение статуса чата AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.GetAiSolutionChatStatus(context.Background(), solutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.GetAiSolutionChatStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiSolutionChatStatus`: GetAiSolutionChatStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.GetAiSolutionChatStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSolutionChatStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAiSolutionChatStatus200Response**](GetAiSolutionChatStatus200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiSolutionMainChatHistory

> GetAiSolutionMainChatHistory200Response GetAiSolutionMainChatHistory(ctx, solutionId).Limit(limit).Offset(offset).Execute()

Получение истории основного разговора AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    limit := int32(100) // int32 | Количество сообщений на странице, от 1 до 10 000 (optional) (default to 100)
    offset := int32(0) // int32 | Количество сообщений, которые нужно пропустить (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.GetAiSolutionMainChatHistory(context.Background(), solutionId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.GetAiSolutionMainChatHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiSolutionMainChatHistory`: GetAiSolutionMainChatHistory200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.GetAiSolutionMainChatHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSolutionMainChatHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Количество сообщений на странице, от 1 до 10 000 | [default to 100]
 **offset** | **int32** | Количество сообщений, которые нужно пропустить | [default to 0]

### Return type

[**GetAiSolutionMainChatHistory200Response**](GetAiSolutionMainChatHistory200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAiSolutionConversations

> ListAiSolutionConversations200Response ListAiSolutionConversations(ctx, solutionId).Execute()

Получение списка разговоров AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.ListAiSolutionConversations(context.Background(), solutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.ListAiSolutionConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAiSolutionConversations`: ListAiSolutionConversations200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.ListAiSolutionConversations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAiSolutionConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAiSolutionConversations200Response**](ListAiSolutionConversations200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAiSolutionTelegramBots

> ListAiSolutionTelegramBots200Response ListAiSolutionTelegramBots(ctx, solutionId).Execute()

Получение списка Telegram-ботов AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.ListAiSolutionTelegramBots(context.Background(), solutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.ListAiSolutionTelegramBots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAiSolutionTelegramBots`: ListAiSolutionTelegramBots200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.ListAiSolutionTelegramBots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAiSolutionTelegramBotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAiSolutionTelegramBots200Response**](ListAiSolutionTelegramBots200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAiSolutionTypes

> ListAiSolutionTypes200Response ListAiSolutionTypes(ctx).Execute()

Получение списка типов AI-решений



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
    resp, r, err := apiClient.AISolutionsAPI.ListAiSolutionTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.ListAiSolutionTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAiSolutionTypes`: ListAiSolutionTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.ListAiSolutionTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAiSolutionTypesRequest struct via the builder pattern


### Return type

[**ListAiSolutionTypes200Response**](ListAiSolutionTypes200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAiSolutions

> ListAiSolutions200Response ListAiSolutions(ctx).Execute()

Получение списка AI-решений



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
    resp, r, err := apiClient.AISolutionsAPI.ListAiSolutions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.ListAiSolutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAiSolutions`: ListAiSolutions200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.ListAiSolutions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAiSolutionsRequest struct via the builder pattern


### Return type

[**ListAiSolutions200Response**](ListAiSolutions200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateAiSolutionAiKey

> RegenerateAiSolutionAiKey200Response RegenerateAiSolutionAiKey(ctx, solutionId).Execute()

Перевыпуск ключа AI Gateway решения



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
    solutionId := int32(15) // int32 | ID AI-решения

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.RegenerateAiSolutionAiKey(context.Background(), solutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.RegenerateAiSolutionAiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegenerateAiSolutionAiKey`: RegenerateAiSolutionAiKey200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.RegenerateAiSolutionAiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateAiSolutionAiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegenerateAiSolutionAiKey200Response**](RegenerateAiSolutionAiKey200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAiSolutionTelegramBot

> RemoveAiSolutionTelegramBot(ctx, solutionId, accountKey).Execute()

Отключение Telegram-бота от AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    accountKey := "b123456789" // string | ID Telegram-бота в решении (`account_key` из списка ботов) в формате `b<число>`, например `b123456789`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AISolutionsAPI.RemoveAiSolutionTelegramBot(context.Background(), solutionId, accountKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.RemoveAiSolutionTelegramBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 
**accountKey** | **string** | ID Telegram-бота в решении (&#x60;account_key&#x60; из списка ботов) в формате &#x60;b&lt;число&gt;&#x60;, например &#x60;b123456789&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAiSolutionTelegramBotRequest struct via the builder pattern


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


## ResetAiSolutionChat

> ResetAiSolutionChat200Response ResetAiSolutionChat(ctx, solutionId).ResetAiSolutionChat(resetAiSolutionChat).Execute()

Сброс контекста разговора AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    resetAiSolutionChat := *openapiclient.NewResetAiSolutionChat() // ResetAiSolutionChat |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.ResetAiSolutionChat(context.Background(), solutionId).ResetAiSolutionChat(resetAiSolutionChat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.ResetAiSolutionChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetAiSolutionChat`: ResetAiSolutionChat200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.ResetAiSolutionChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetAiSolutionChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetAiSolutionChat** | [**ResetAiSolutionChat**](ResetAiSolutionChat.md) |  | 

### Return type

[**ResetAiSolutionChat200Response**](ResetAiSolutionChat200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendAiSolutionMessage

> SendAiSolutionMessage200Response SendAiSolutionMessage(ctx, solutionId).SendAiSolutionMessage(sendAiSolutionMessage).Execute()

Отправка сообщения агенту AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    sendAiSolutionMessage := *openapiclient.NewSendAiSolutionMessage() // SendAiSolutionMessage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.SendAiSolutionMessage(context.Background(), solutionId).SendAiSolutionMessage(sendAiSolutionMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.SendAiSolutionMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendAiSolutionMessage`: SendAiSolutionMessage200Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.SendAiSolutionMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendAiSolutionMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendAiSolutionMessage** | [**SendAiSolutionMessage**](SendAiSolutionMessage.md) |  | 

### Return type

[**SendAiSolutionMessage200Response**](SendAiSolutionMessage200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShutdownAiSolution

> ShutdownAiSolution(ctx, solutionId).Execute()

Выключение AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AISolutionsAPI.ShutdownAiSolution(context.Background(), solutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.ShutdownAiSolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownAiSolutionRequest struct via the builder pattern


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


## StartAiSolution

> StartAiSolution(ctx, solutionId).Execute()

Запуск AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AISolutionsAPI.StartAiSolution(context.Background(), solutionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.StartAiSolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartAiSolutionRequest struct via the builder pattern


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


## StopAiSolutionChat

> StopAiSolutionChat(ctx, solutionId).StopAiSolutionChat(stopAiSolutionChat).Execute()

Остановка хода агента AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    stopAiSolutionChat := *openapiclient.NewStopAiSolutionChat() // StopAiSolutionChat |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AISolutionsAPI.StopAiSolutionChat(context.Background(), solutionId).StopAiSolutionChat(stopAiSolutionChat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.StopAiSolutionChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAiSolutionChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stopAiSolutionChat** | [**StopAiSolutionChat**](StopAiSolutionChat.md) |  | 

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


## StreamAiSolutionChat

> string StreamAiSolutionChat(ctx, solutionId).SendAiSolutionMessage(sendAiSolutionMessage).Execute()

Отправка сообщения агенту AI-решения с потоковым ответом



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
    solutionId := int32(15) // int32 | ID AI-решения
    sendAiSolutionMessage := *openapiclient.NewSendAiSolutionMessage() // SendAiSolutionMessage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.StreamAiSolutionChat(context.Background(), solutionId).SendAiSolutionMessage(sendAiSolutionMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.StreamAiSolutionChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StreamAiSolutionChat`: string
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.StreamAiSolutionChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamAiSolutionChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendAiSolutionMessage** | [**SendAiSolutionMessage**](SendAiSolutionMessage.md) |  | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAiSolution

> CreateAiSolution201Response UpdateAiSolution(ctx, solutionId).UpdateAiSolution(updateAiSolution).Execute()

Изменение AI-решения



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
    solutionId := int32(15) // int32 | ID AI-решения
    updateAiSolution := *openapiclient.NewUpdateAiSolution() // UpdateAiSolution | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AISolutionsAPI.UpdateAiSolution(context.Background(), solutionId).UpdateAiSolution(updateAiSolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AISolutionsAPI.UpdateAiSolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAiSolution`: CreateAiSolution201Response
    fmt.Fprintf(os.Stdout, "Response from `AISolutionsAPI.UpdateAiSolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **int32** | ID AI-решения | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAiSolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAiSolution** | [**UpdateAiSolution**](UpdateAiSolution.md) |  | 

### Return type

[**CreateAiSolution201Response**](CreateAiSolution201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

