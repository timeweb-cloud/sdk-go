# \AIAgentIntegrationsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgentBotMaxBot**](AIAgentIntegrationsAPI.md#CreateAgentBotMaxBot) | **Post** /api/v1/cloud-ai/max-bots | Подключение бота MAX к агенту
[**CreateAgentBotTelegramBot**](AIAgentIntegrationsAPI.md#CreateAgentBotTelegramBot) | **Post** /api/v1/cloud-ai/bots | Подключение Telegram-бота к агенту
[**CreateAgentIntegrationCrisp**](AIAgentIntegrationsAPI.md#CreateAgentIntegrationCrisp) | **Post** /api/v1/cloud-ai/agent/{agent_id}/crisp | Подключение агента к Crisp
[**CreateAgentIntegrationJivo**](AIAgentIntegrationsAPI.md#CreateAgentIntegrationJivo) | **Post** /api/v1/cloud-ai/agent/{agent_id}/jivo | Подключение агента к Jivo
[**DeleteAgentBotMaxBot**](AIAgentIntegrationsAPI.md#DeleteAgentBotMaxBot) | **Delete** /api/v1/cloud-ai/max-bots/{bot_id} | Отключение бота MAX
[**DeleteAgentBotTelegramBot**](AIAgentIntegrationsAPI.md#DeleteAgentBotTelegramBot) | **Delete** /api/v1/cloud-ai/bots/{bot_id} | Отключение Telegram-бота
[**DeleteAgentIntegrationCrisp**](AIAgentIntegrationsAPI.md#DeleteAgentIntegrationCrisp) | **Delete** /api/v1/cloud-ai/agent/{agent_id}/crisp | Отключение агента от Crisp
[**DeleteAgentIntegrationJivo**](AIAgentIntegrationsAPI.md#DeleteAgentIntegrationJivo) | **Delete** /api/v1/cloud-ai/agent/{agent_id}/jivo | Отключение агента от Jivo
[**GetAgentBotMaxBot**](AIAgentIntegrationsAPI.md#GetAgentBotMaxBot) | **Get** /api/v1/cloud-ai/max-bots/{bot_id} | Получение информации о боте MAX
[**GetAgentBotMaxBots**](AIAgentIntegrationsAPI.md#GetAgentBotMaxBots) | **Get** /api/v1/cloud-ai/max-bots | Получение списка ботов MAX
[**GetAgentBotMaxBotsByAgent**](AIAgentIntegrationsAPI.md#GetAgentBotMaxBotsByAgent) | **Get** /api/v1/cloud-ai/agents/{agent_id}/max-bots | Получение списка ботов MAX агента
[**GetAgentBotTelegramBot**](AIAgentIntegrationsAPI.md#GetAgentBotTelegramBot) | **Get** /api/v1/cloud-ai/bots/{bot_id} | Получение информации о Telegram-боте
[**GetAgentBotTelegramBots**](AIAgentIntegrationsAPI.md#GetAgentBotTelegramBots) | **Get** /api/v1/cloud-ai/bots | Получение списка Telegram-ботов
[**GetAgentBotTelegramBotsByAgent**](AIAgentIntegrationsAPI.md#GetAgentBotTelegramBotsByAgent) | **Get** /api/v1/cloud-ai/agents/{agent_id}/bots | Получение списка Telegram-ботов агента
[**GetAgentIntegrationCrisp**](AIAgentIntegrationsAPI.md#GetAgentIntegrationCrisp) | **Get** /api/v1/cloud-ai/agent/{agent_id}/crisp | Получение интеграции агента с Crisp
[**GetAgentIntegrationJivo**](AIAgentIntegrationsAPI.md#GetAgentIntegrationJivo) | **Get** /api/v1/cloud-ai/agent/{agent_id}/jivo | Получение интеграции агента с Jivo
[**UpdateAgentBotMaxBot**](AIAgentIntegrationsAPI.md#UpdateAgentBotMaxBot) | **Patch** /api/v1/cloud-ai/max-bots/{bot_id} | Изменение бота MAX
[**UpdateAgentBotTelegramBot**](AIAgentIntegrationsAPI.md#UpdateAgentBotTelegramBot) | **Patch** /api/v1/cloud-ai/bots/{bot_id} | Изменение Telegram-бота



## CreateAgentBotMaxBot

> CreateAgentBotTelegramBot201Response CreateAgentBotMaxBot(ctx).CreateMaxBot(createMaxBot).Execute()

Подключение бота MAX к агенту



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
    createMaxBot := *openapiclient.NewCreateMaxBot(int32(1), "f9LHodD0cOKYOelWNfqrDE7kdy9yVnEAZoOjUkbvZh3PjE1M_8SAk33h") // CreateMaxBot | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.CreateAgentBotMaxBot(context.Background()).CreateMaxBot(createMaxBot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.CreateAgentBotMaxBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentBotMaxBot`: CreateAgentBotTelegramBot201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.CreateAgentBotMaxBot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentBotMaxBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMaxBot** | [**CreateMaxBot**](CreateMaxBot.md) |  | 

### Return type

[**CreateAgentBotTelegramBot201Response**](CreateAgentBotTelegramBot201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgentBotTelegramBot

> CreateAgentBotTelegramBot201Response CreateAgentBotTelegramBot(ctx).CreateBot(createBot).Execute()

Подключение Telegram-бота к агенту



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
    createBot := *openapiclient.NewCreateBot(int32(1), "1234567890:ABCdefGHIjklMNOpqrsTUVwxyz") // CreateBot | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.CreateAgentBotTelegramBot(context.Background()).CreateBot(createBot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.CreateAgentBotTelegramBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentBotTelegramBot`: CreateAgentBotTelegramBot201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.CreateAgentBotTelegramBot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentBotTelegramBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBot** | [**CreateBot**](CreateBot.md) |  | 

### Return type

[**CreateAgentBotTelegramBot201Response**](CreateAgentBotTelegramBot201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgentIntegrationCrisp

> GetAgentIntegrationCrisp200Response CreateAgentIntegrationCrisp(ctx, agentId).CreateCrispIntegration(createCrispIntegration).Execute()

Подключение агента к Crisp



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
    agentId := int32(1) // int32 | ID агента
    createCrispIntegration := *openapiclient.NewCreateCrispIntegration("8c842203-7ed8-4e29-a608-7cf78a7d2fcc", "7c3ef21c-1e04-41ce-8c06-5605c1cd7c2f", "cc29e1a5d3c5e3f4c9a9d4a2e1b7f8a3c6d5e4f3a2b1c0d9e8f7a6b5c4d3e2f1") // CreateCrispIntegration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.CreateAgentIntegrationCrisp(context.Background(), agentId).CreateCrispIntegration(createCrispIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.CreateAgentIntegrationCrisp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentIntegrationCrisp`: GetAgentIntegrationCrisp200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.CreateAgentIntegrationCrisp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentIntegrationCrispRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCrispIntegration** | [**CreateCrispIntegration**](CreateCrispIntegration.md) |  | 

### Return type

[**GetAgentIntegrationCrisp200Response**](GetAgentIntegrationCrisp200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgentIntegrationJivo

> GetAgentIntegrationJivo200Response CreateAgentIntegrationJivo(ctx, agentId).CreateJivoIntegration(createJivoIntegration).Execute()

Подключение агента к Jivo



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
    agentId := int32(1) // int32 | ID агента
    createJivoIntegration := *openapiclient.NewCreateJivoIntegration("Ut3O9rsU8DwAMFE") // CreateJivoIntegration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.CreateAgentIntegrationJivo(context.Background(), agentId).CreateJivoIntegration(createJivoIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.CreateAgentIntegrationJivo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentIntegrationJivo`: GetAgentIntegrationJivo200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.CreateAgentIntegrationJivo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentIntegrationJivoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createJivoIntegration** | [**CreateJivoIntegration**](CreateJivoIntegration.md) |  | 

### Return type

[**GetAgentIntegrationJivo200Response**](GetAgentIntegrationJivo200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentBotMaxBot

> DeleteAgentBotMaxBot(ctx, botId).Execute()

Отключение бота MAX



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
    botId := "550e8400-e29b-41d4-a716-446655440000" // string | ID бота (поле `id` из ответа)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentIntegrationsAPI.DeleteAgentBotMaxBot(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.DeleteAgentBotMaxBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | ID бота (поле &#x60;id&#x60; из ответа) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentBotMaxBotRequest struct via the builder pattern


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


## DeleteAgentBotTelegramBot

> DeleteAgentBotTelegramBot(ctx, botId).Execute()

Отключение Telegram-бота



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
    botId := "550e8400-e29b-41d4-a716-446655440000" // string | ID бота (поле `id` из ответа)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentIntegrationsAPI.DeleteAgentBotTelegramBot(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.DeleteAgentBotTelegramBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | ID бота (поле &#x60;id&#x60; из ответа) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentBotTelegramBotRequest struct via the builder pattern


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


## DeleteAgentIntegrationCrisp

> DeleteAgentIntegrationCrisp(ctx, agentId).Execute()

Отключение агента от Crisp



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
    agentId := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentIntegrationsAPI.DeleteAgentIntegrationCrisp(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.DeleteAgentIntegrationCrisp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentIntegrationCrispRequest struct via the builder pattern


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


## DeleteAgentIntegrationJivo

> DeleteAgentIntegrationJivo(ctx, agentId).Execute()

Отключение агента от Jivo



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
    agentId := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentIntegrationsAPI.DeleteAgentIntegrationJivo(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.DeleteAgentIntegrationJivo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentIntegrationJivoRequest struct via the builder pattern


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


## GetAgentBotMaxBot

> CreateAgentBotTelegramBot201Response GetAgentBotMaxBot(ctx, botId).Execute()

Получение информации о боте MAX



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
    botId := "550e8400-e29b-41d4-a716-446655440000" // string | ID бота (поле `id` из ответа)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.GetAgentBotMaxBot(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.GetAgentBotMaxBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentBotMaxBot`: CreateAgentBotTelegramBot201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.GetAgentBotMaxBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | ID бота (поле &#x60;id&#x60; из ответа) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBotMaxBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAgentBotTelegramBot201Response**](CreateAgentBotTelegramBot201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBotMaxBots

> GetAgentBotMaxBots200Response GetAgentBotMaxBots(ctx).Execute()

Получение списка ботов MAX



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
    resp, r, err := apiClient.AIAgentIntegrationsAPI.GetAgentBotMaxBots(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.GetAgentBotMaxBots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentBotMaxBots`: GetAgentBotMaxBots200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.GetAgentBotMaxBots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBotMaxBotsRequest struct via the builder pattern


### Return type

[**GetAgentBotMaxBots200Response**](GetAgentBotMaxBots200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBotMaxBotsByAgent

> GetAgentBotMaxBotsByAgent200Response GetAgentBotMaxBotsByAgent(ctx, agentId).Execute()

Получение списка ботов MAX агента



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
    agentId := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.GetAgentBotMaxBotsByAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.GetAgentBotMaxBotsByAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentBotMaxBotsByAgent`: GetAgentBotMaxBotsByAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.GetAgentBotMaxBotsByAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBotMaxBotsByAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentBotMaxBotsByAgent200Response**](GetAgentBotMaxBotsByAgent200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBotTelegramBot

> CreateAgentBotTelegramBot201Response GetAgentBotTelegramBot(ctx, botId).Execute()

Получение информации о Telegram-боте



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
    botId := "550e8400-e29b-41d4-a716-446655440000" // string | ID бота (поле `id` из ответа)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.GetAgentBotTelegramBot(context.Background(), botId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.GetAgentBotTelegramBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentBotTelegramBot`: CreateAgentBotTelegramBot201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.GetAgentBotTelegramBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | ID бота (поле &#x60;id&#x60; из ответа) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBotTelegramBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAgentBotTelegramBot201Response**](CreateAgentBotTelegramBot201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBotTelegramBots

> GetAgentBotTelegramBots200Response GetAgentBotTelegramBots(ctx).Execute()

Получение списка Telegram-ботов



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
    resp, r, err := apiClient.AIAgentIntegrationsAPI.GetAgentBotTelegramBots(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.GetAgentBotTelegramBots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentBotTelegramBots`: GetAgentBotTelegramBots200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.GetAgentBotTelegramBots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBotTelegramBotsRequest struct via the builder pattern


### Return type

[**GetAgentBotTelegramBots200Response**](GetAgentBotTelegramBots200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBotTelegramBotsByAgent

> GetAgentBotTelegramBotsByAgent200Response GetAgentBotTelegramBotsByAgent(ctx, agentId).Execute()

Получение списка Telegram-ботов агента



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
    agentId := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.GetAgentBotTelegramBotsByAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.GetAgentBotTelegramBotsByAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentBotTelegramBotsByAgent`: GetAgentBotTelegramBotsByAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.GetAgentBotTelegramBotsByAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBotTelegramBotsByAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentBotTelegramBotsByAgent200Response**](GetAgentBotTelegramBotsByAgent200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentIntegrationCrisp

> GetAgentIntegrationCrisp200Response GetAgentIntegrationCrisp(ctx, agentId).Execute()

Получение интеграции агента с Crisp



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
    agentId := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.GetAgentIntegrationCrisp(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.GetAgentIntegrationCrisp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentIntegrationCrisp`: GetAgentIntegrationCrisp200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.GetAgentIntegrationCrisp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentIntegrationCrispRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentIntegrationCrisp200Response**](GetAgentIntegrationCrisp200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentIntegrationJivo

> GetAgentIntegrationJivo200Response GetAgentIntegrationJivo(ctx, agentId).Execute()

Получение интеграции агента с Jivo



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
    agentId := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.GetAgentIntegrationJivo(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.GetAgentIntegrationJivo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentIntegrationJivo`: GetAgentIntegrationJivo200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.GetAgentIntegrationJivo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentIntegrationJivoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentIntegrationJivo200Response**](GetAgentIntegrationJivo200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentBotMaxBot

> CreateAgentBotTelegramBot201Response UpdateAgentBotMaxBot(ctx, botId).UpdateMaxBot(updateMaxBot).Execute()

Изменение бота MAX



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
    botId := "550e8400-e29b-41d4-a716-446655440000" // string | ID бота (поле `id` из ответа)
    updateMaxBot := *openapiclient.NewUpdateMaxBot() // UpdateMaxBot | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.UpdateAgentBotMaxBot(context.Background(), botId).UpdateMaxBot(updateMaxBot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.UpdateAgentBotMaxBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentBotMaxBot`: CreateAgentBotTelegramBot201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.UpdateAgentBotMaxBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | ID бота (поле &#x60;id&#x60; из ответа) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentBotMaxBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMaxBot** | [**UpdateMaxBot**](UpdateMaxBot.md) |  | 

### Return type

[**CreateAgentBotTelegramBot201Response**](CreateAgentBotTelegramBot201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentBotTelegramBot

> CreateAgentBotTelegramBot201Response UpdateAgentBotTelegramBot(ctx, botId).UpdateBot(updateBot).Execute()

Изменение Telegram-бота



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
    botId := "550e8400-e29b-41d4-a716-446655440000" // string | ID бота (поле `id` из ответа)
    updateBot := *openapiclient.NewUpdateBot() // UpdateBot | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentIntegrationsAPI.UpdateAgentBotTelegramBot(context.Background(), botId).UpdateBot(updateBot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentIntegrationsAPI.UpdateAgentBotTelegramBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentBotTelegramBot`: CreateAgentBotTelegramBot201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentIntegrationsAPI.UpdateAgentBotTelegramBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**botId** | **string** | ID бота (поле &#x60;id&#x60; из ответа) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentBotTelegramBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBot** | [**UpdateBot**](UpdateBot.md) |  | 

### Return type

[**CreateAgentBotTelegramBot201Response**](CreateAgentBotTelegramBot201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

