# \MCPAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectAgentToolToAgent**](MCPAPI.md#ConnectAgentToolToAgent) | **Post** /api/v1/cloud-ai/agents/{agent_id}/tools | Подключение MCP-инструмента к агенту
[**CreateAgentTool**](MCPAPI.md#CreateAgentTool) | **Post** /api/v1/cloud-ai/tools | Подключение собственного MCP-сервера
[**DeleteAgentTool**](MCPAPI.md#DeleteAgentTool) | **Delete** /api/v1/cloud-ai/tools/{id} | Удаление MCP-инструмента
[**DeleteMcpSearchSession**](MCPAPI.md#DeleteMcpSearchSession) | **Delete** /api/v1/mcp/search | Завершение сессии сервера MCP с поиском инструментов
[**DeleteMcpSession**](MCPAPI.md#DeleteMcpSession) | **Delete** /api/v1/mcp | Завершение сессии сервера MCP
[**DisconnectAgentToolFromAgent**](MCPAPI.md#DisconnectAgentToolFromAgent) | **Delete** /api/v1/cloud-ai/agents/{agent_id}/tools/{tool_id} | Отключение MCP-инструмента от агента
[**GetAgentTool**](MCPAPI.md#GetAgentTool) | **Get** /api/v1/cloud-ai/tools/{id} | Получение MCP-инструмента
[**GetAgentToolOAuth2AuthorizeUrl**](MCPAPI.md#GetAgentToolOAuth2AuthorizeUrl) | **Get** /api/v1/cloud-ai/tools/oauth2/authorize-url/{tool_id} | Получение ссылки на OAuth2-авторизацию инструмента
[**GetAgentToolTemplatePrices**](MCPAPI.md#GetAgentToolTemplatePrices) | **Get** /api/v1/cloud-ai/tools/templates/prices | Получение цен системных инструментов
[**GetAgentToolTemplates**](MCPAPI.md#GetAgentToolTemplates) | **Get** /api/v1/cloud-ai/tools/templates/marketplace | Получение шаблонов MCP-инструментов из маркетплейса
[**GetAgentTools**](MCPAPI.md#GetAgentTools) | **Get** /api/v1/cloud-ai/tools | Получение списка MCP-инструментов
[**GetAgentToolsByAgent**](MCPAPI.md#GetAgentToolsByAgent) | **Get** /api/v1/cloud-ai/agents/{agent_id}/tools | Получение MCP-инструментов агента
[**InstallAgentToolFromTemplate**](MCPAPI.md#InstallAgentToolFromTemplate) | **Post** /api/v1/cloud-ai/tools/install | Установка MCP-инструмента из маркетплейса
[**OpenMcpSearchStream**](MCPAPI.md#OpenMcpSearchStream) | **Get** /api/v1/mcp/search | Открытие потока событий сервера MCP с поиском инструментов
[**OpenMcpStream**](MCPAPI.md#OpenMcpStream) | **Get** /api/v1/mcp | Открытие потока событий сервера MCP
[**RefreshAgentToolMethods**](MCPAPI.md#RefreshAgentToolMethods) | **Post** /api/v1/cloud-ai/tools/{id}/refresh-methods | Обновление списка методов MCP-инструмента
[**SendMcpMessage**](MCPAPI.md#SendMcpMessage) | **Post** /api/v1/mcp | Отправка JSON-RPC-сообщения серверу MCP
[**SendMcpSearchMessage**](MCPAPI.md#SendMcpSearchMessage) | **Post** /api/v1/mcp/search | Отправка JSON-RPC-сообщения серверу MCP с поиском инструментов
[**UpdateAgentTool**](MCPAPI.md#UpdateAgentTool) | **Put** /api/v1/cloud-ai/tools/{id} | Изменение MCP-инструмента
[**UpdateAgentToolConnection**](MCPAPI.md#UpdateAgentToolConnection) | **Put** /api/v1/cloud-ai/agents/{agent_id}/tools/{tool_id} | Изменение разрешённых методов MCP-инструмента у агента



## ConnectAgentToolToAgent

> ConnectAgentToolToAgent(ctx, agentId).ConnectAgentTool(connectAgentTool).Execute()

Подключение MCP-инструмента к агенту



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
    agentId := int32(113191) // int32 | ID агента
    connectAgentTool := *openapiclient.NewConnectAgentTool(int32(417), []string{"AllowedMethods_example"}) // ConnectAgentTool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MCPAPI.ConnectAgentToolToAgent(context.Background(), agentId).ConnectAgentTool(connectAgentTool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.ConnectAgentToolToAgent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiConnectAgentToolToAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectAgentTool** | [**ConnectAgentTool**](ConnectAgentTool.md) |  | 

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


## CreateAgentTool

> CreateAgentTool201Response CreateAgentTool(ctx).CreateTool(createTool).Execute()

Подключение собственного MCP-сервера



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
    createTool := *openapiclient.NewCreateTool("Мой MCP-сервер", "MCP", "http", "https://mcp.example.com/mcp") // CreateTool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.CreateAgentTool(context.Background()).CreateTool(createTool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.CreateAgentTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentTool`: CreateAgentTool201Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.CreateAgentTool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTool** | [**CreateTool**](CreateTool.md) |  | 

### Return type

[**CreateAgentTool201Response**](CreateAgentTool201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentTool

> DeleteAgentTool(ctx, id).Execute()

Удаление MCP-инструмента



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
    id := int32(417) // int32 | ID инструмента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MCPAPI.DeleteAgentTool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.DeleteAgentTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID инструмента | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentToolRequest struct via the builder pattern


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


## DeleteMcpSearchSession

> DeleteMcpSearchSession(ctx).Execute()

Завершение сессии сервера MCP с поиском инструментов



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
    r, err := apiClient.MCPAPI.DeleteMcpSearchSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.DeleteMcpSearchSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMcpSearchSessionRequest struct via the builder pattern


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


## DeleteMcpSession

> DeleteMcpSession(ctx).Execute()

Завершение сессии сервера MCP



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
    r, err := apiClient.MCPAPI.DeleteMcpSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.DeleteMcpSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMcpSessionRequest struct via the builder pattern


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


## DisconnectAgentToolFromAgent

> DisconnectAgentToolFromAgent(ctx, agentId, toolId).Execute()

Отключение MCP-инструмента от агента



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
    agentId := int32(113191) // int32 | ID агента
    toolId := int32(417) // int32 | ID инструмента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MCPAPI.DisconnectAgentToolFromAgent(context.Background(), agentId, toolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.DisconnectAgentToolFromAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 
**toolId** | **int32** | ID инструмента | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectAgentToolFromAgentRequest struct via the builder pattern


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


## GetAgentTool

> CreateAgentTool201Response GetAgentTool(ctx, id).Execute()

Получение MCP-инструмента



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
    id := int32(417) // int32 | ID инструмента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.GetAgentTool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.GetAgentTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentTool`: CreateAgentTool201Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.GetAgentTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID инструмента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAgentTool201Response**](CreateAgentTool201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentToolOAuth2AuthorizeUrl

> AuthorizeUrl GetAgentToolOAuth2AuthorizeUrl(ctx, toolId).RedirectUrl(redirectUrl).Execute()

Получение ссылки на OAuth2-авторизацию инструмента



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
    toolId := int32(417) // int32 | ID инструмента
    redirectUrl := "https://timeweb.cloud/my/cloud-ai" // string | Куда вернуть пользователя после авторизации. Принимаются только адреса на доменах `timeweb.cloud` и `timeweb.net` и их поддоменах; другие адреса молча игнорируются. После успешной авторизации к адресу добавляются `oauth_success=true&tool_id=<id>`, после ошибки — `oauth_error=<текст>&tool_id=<id>`. Без параметра страница возврата показывает результат в виде JSON. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.GetAgentToolOAuth2AuthorizeUrl(context.Background(), toolId).RedirectUrl(redirectUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.GetAgentToolOAuth2AuthorizeUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentToolOAuth2AuthorizeUrl`: AuthorizeUrl
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.GetAgentToolOAuth2AuthorizeUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **int32** | ID инструмента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentToolOAuth2AuthorizeUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redirectUrl** | **string** | Куда вернуть пользователя после авторизации. Принимаются только адреса на доменах &#x60;timeweb.cloud&#x60; и &#x60;timeweb.net&#x60; и их поддоменах; другие адреса молча игнорируются. После успешной авторизации к адресу добавляются &#x60;oauth_success&#x3D;true&amp;tool_id&#x3D;&lt;id&gt;&#x60;, после ошибки — &#x60;oauth_error&#x3D;&lt;текст&gt;&amp;tool_id&#x3D;&lt;id&gt;&#x60;. Без параметра страница возврата показывает результат в виде JSON. | 

### Return type

[**AuthorizeUrl**](AuthorizeUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentToolTemplatePrices

> GetAgentToolTemplatePrices200Response GetAgentToolTemplatePrices(ctx).Locale(locale).Execute()

Получение цен системных инструментов



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
    locale := "ru" // string | Валюта цены: `ru` — в рублях, `en` — в долларах США. По умолчанию `ru`. (optional) (default to "ru")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.GetAgentToolTemplatePrices(context.Background()).Locale(locale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.GetAgentToolTemplatePrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentToolTemplatePrices`: GetAgentToolTemplatePrices200Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.GetAgentToolTemplatePrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentToolTemplatePricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** | Валюта цены: &#x60;ru&#x60; — в рублях, &#x60;en&#x60; — в долларах США. По умолчанию &#x60;ru&#x60;. | [default to &quot;ru&quot;]

### Return type

[**GetAgentToolTemplatePrices200Response**](GetAgentToolTemplatePrices200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentToolTemplates

> GetAgentToolTemplates200Response GetAgentToolTemplates(ctx).Execute()

Получение шаблонов MCP-инструментов из маркетплейса



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
    resp, r, err := apiClient.MCPAPI.GetAgentToolTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.GetAgentToolTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentToolTemplates`: GetAgentToolTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.GetAgentToolTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentToolTemplatesRequest struct via the builder pattern


### Return type

[**GetAgentToolTemplates200Response**](GetAgentToolTemplates200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentTools

> GetAgentTools200Response GetAgentTools(ctx).Execute()

Получение списка MCP-инструментов



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
    resp, r, err := apiClient.MCPAPI.GetAgentTools(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.GetAgentTools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentTools`: GetAgentTools200Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.GetAgentTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentToolsRequest struct via the builder pattern


### Return type

[**GetAgentTools200Response**](GetAgentTools200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentToolsByAgent

> GetAgentToolsByAgent200Response GetAgentToolsByAgent(ctx, agentId).Execute()

Получение MCP-инструментов агента



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
    agentId := int32(113191) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.GetAgentToolsByAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.GetAgentToolsByAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentToolsByAgent`: GetAgentToolsByAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.GetAgentToolsByAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentToolsByAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentToolsByAgent200Response**](GetAgentToolsByAgent200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallAgentToolFromTemplate

> CreateAgentTool201Response InstallAgentToolFromTemplate(ctx).InstallTool(installTool).Execute()

Установка MCP-инструмента из маркетплейса



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
    installTool := *openapiclient.NewInstallTool(int32(7)) // InstallTool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.InstallAgentToolFromTemplate(context.Background()).InstallTool(installTool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.InstallAgentToolFromTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallAgentToolFromTemplate`: CreateAgentTool201Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.InstallAgentToolFromTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstallAgentToolFromTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installTool** | [**InstallTool**](InstallTool.md) |  | 

### Return type

[**CreateAgentTool201Response**](CreateAgentTool201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenMcpSearchStream

> string OpenMcpSearchStream(ctx).Accept(accept).MCPProtocolVersion(mCPProtocolVersion).Execute()

Открытие потока событий сервера MCP с поиском инструментов



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
    accept := "text/event-stream" // string | Должен содержать `text/event-stream`, иначе сервер ответит 406
    mCPProtocolVersion := "2025-06-18" // string | Версия протокола, согласованная в `initialize`. Неподдерживаемое значение — ответ 400 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.OpenMcpSearchStream(context.Background()).Accept(accept).MCPProtocolVersion(mCPProtocolVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.OpenMcpSearchStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenMcpSearchStream`: string
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.OpenMcpSearchStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenMcpSearchStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Должен содержать &#x60;text/event-stream&#x60;, иначе сервер ответит 406 | 
 **mCPProtocolVersion** | **string** | Версия протокола, согласованная в &#x60;initialize&#x60;. Неподдерживаемое значение — ответ 400 | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenMcpStream

> string OpenMcpStream(ctx).Accept(accept).MCPProtocolVersion(mCPProtocolVersion).Execute()

Открытие потока событий сервера MCP



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
    accept := "text/event-stream" // string | Должен содержать `text/event-stream`, иначе сервер ответит 406
    mCPProtocolVersion := "2025-06-18" // string | Версия протокола, согласованная в `initialize`. Неподдерживаемое значение — ответ 400 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.OpenMcpStream(context.Background()).Accept(accept).MCPProtocolVersion(mCPProtocolVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.OpenMcpStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenMcpStream`: string
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.OpenMcpStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenMcpStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Должен содержать &#x60;text/event-stream&#x60;, иначе сервер ответит 406 | 
 **mCPProtocolVersion** | **string** | Версия протокола, согласованная в &#x60;initialize&#x60;. Неподдерживаемое значение — ответ 400 | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshAgentToolMethods

> CreateAgentTool201Response RefreshAgentToolMethods(ctx, id).Execute()

Обновление списка методов MCP-инструмента



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
    id := int32(417) // int32 | ID инструмента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.RefreshAgentToolMethods(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.RefreshAgentToolMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshAgentToolMethods`: CreateAgentTool201Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.RefreshAgentToolMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID инструмента | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshAgentToolMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAgentTool201Response**](CreateAgentTool201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMcpMessage

> JsonRpcResultResponse SendMcpMessage(ctx).Accept(accept).JsonRpcRequest(jsonRpcRequest).MCPProtocolVersion(mCPProtocolVersion).XMcpConfirmationUi(xMcpConfirmationUi).Execute()

Отправка JSON-RPC-сообщения серверу MCP



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
    accept := "application/json, text/event-stream" // string | Должен содержать и `application/json`, и `text/event-stream`, иначе сервер ответит 406
    jsonRpcRequest := *openapiclient.NewJsonRpcRequest("2.0", "1", "tools/call") // JsonRpcRequest | 
    mCPProtocolVersion := "2025-06-18" // string | Версия протокола, согласованная в `initialize`. Неподдерживаемое значение — ответ 400 (optional)
    xMcpConfirmationUi := "text" // string | Как `request_user_confirmation` спрашивает пользователя. `text` (по умолчанию) — пользователь отвечает в чате, на `/api/v1/mcp` инструмент в этом режиме не выводится. `buttons` — инструмент возвращает `ui_action` с кнопками для клиентов, которые умеют их показывать. Регистр не важен, другие значения игнорируются (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.SendMcpMessage(context.Background()).Accept(accept).JsonRpcRequest(jsonRpcRequest).MCPProtocolVersion(mCPProtocolVersion).XMcpConfirmationUi(xMcpConfirmationUi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.SendMcpMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendMcpMessage`: JsonRpcResultResponse
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.SendMcpMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMcpMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Должен содержать и &#x60;application/json&#x60;, и &#x60;text/event-stream&#x60;, иначе сервер ответит 406 | 
 **jsonRpcRequest** | [**JsonRpcRequest**](JsonRpcRequest.md) |  | 
 **mCPProtocolVersion** | **string** | Версия протокола, согласованная в &#x60;initialize&#x60;. Неподдерживаемое значение — ответ 400 | 
 **xMcpConfirmationUi** | **string** | Как &#x60;request_user_confirmation&#x60; спрашивает пользователя. &#x60;text&#x60; (по умолчанию) — пользователь отвечает в чате, на &#x60;/api/v1/mcp&#x60; инструмент в этом режиме не выводится. &#x60;buttons&#x60; — инструмент возвращает &#x60;ui_action&#x60; с кнопками для клиентов, которые умеют их показывать. Регистр не важен, другие значения игнорируются | 

### Return type

[**JsonRpcResultResponse**](JsonRpcResultResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMcpSearchMessage

> JsonRpcResultResponse SendMcpSearchMessage(ctx).Accept(accept).JsonRpcRequest(jsonRpcRequest).MCPProtocolVersion(mCPProtocolVersion).XToolAccess(xToolAccess).XAllowedTools(xAllowedTools).XMcpConfirmationUi(xMcpConfirmationUi).Execute()

Отправка JSON-RPC-сообщения серверу MCP с поиском инструментов



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
    accept := "application/json, text/event-stream" // string | Должен содержать и `application/json`, и `text/event-stream`, иначе сервер ответит 406
    jsonRpcRequest := *openapiclient.NewJsonRpcRequest("2.0", "1", "tools/call") // JsonRpcRequest | 
    mCPProtocolVersion := "2025-06-18" // string | Версия протокола, согласованная в `initialize`. Неподдерживаемое значение — ответ 400 (optional)
    xToolAccess := "read-only" // string | Только операции чтения: поиск и определения пропускают изменяющие инструменты, а `execute_tool` отказывает с `tool_not_allowed`. Единственное допустимое значение — `read-only` (регистр не важен); любое другое, в том числе пустое, — ответ 400 с ошибкой `-32600` (optional)
    xAllowedTools := "list_servers,get_server,server_action" // string | Имена разрешённых операций через запятую — все остальные недоступны. Неизвестные имена игнорируются, пустое значение не разрешает ничего. Вместе с `x-tool-access` операция должна пройти оба ограничения (optional)
    xMcpConfirmationUi := "text" // string | Как `request_user_confirmation` спрашивает пользователя. `text` (по умолчанию) — пользователь отвечает в чате, на `/api/v1/mcp` инструмент в этом режиме не выводится. `buttons` — инструмент возвращает `ui_action` с кнопками для клиентов, которые умеют их показывать. Регистр не важен, другие значения игнорируются (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.SendMcpSearchMessage(context.Background()).Accept(accept).JsonRpcRequest(jsonRpcRequest).MCPProtocolVersion(mCPProtocolVersion).XToolAccess(xToolAccess).XAllowedTools(xAllowedTools).XMcpConfirmationUi(xMcpConfirmationUi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.SendMcpSearchMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendMcpSearchMessage`: JsonRpcResultResponse
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.SendMcpSearchMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMcpSearchMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Должен содержать и &#x60;application/json&#x60;, и &#x60;text/event-stream&#x60;, иначе сервер ответит 406 | 
 **jsonRpcRequest** | [**JsonRpcRequest**](JsonRpcRequest.md) |  | 
 **mCPProtocolVersion** | **string** | Версия протокола, согласованная в &#x60;initialize&#x60;. Неподдерживаемое значение — ответ 400 | 
 **xToolAccess** | **string** | Только операции чтения: поиск и определения пропускают изменяющие инструменты, а &#x60;execute_tool&#x60; отказывает с &#x60;tool_not_allowed&#x60;. Единственное допустимое значение — &#x60;read-only&#x60; (регистр не важен); любое другое, в том числе пустое, — ответ 400 с ошибкой &#x60;-32600&#x60; | 
 **xAllowedTools** | **string** | Имена разрешённых операций через запятую — все остальные недоступны. Неизвестные имена игнорируются, пустое значение не разрешает ничего. Вместе с &#x60;x-tool-access&#x60; операция должна пройти оба ограничения | 
 **xMcpConfirmationUi** | **string** | Как &#x60;request_user_confirmation&#x60; спрашивает пользователя. &#x60;text&#x60; (по умолчанию) — пользователь отвечает в чате, на &#x60;/api/v1/mcp&#x60; инструмент в этом режиме не выводится. &#x60;buttons&#x60; — инструмент возвращает &#x60;ui_action&#x60; с кнопками для клиентов, которые умеют их показывать. Регистр не важен, другие значения игнорируются | 

### Return type

[**JsonRpcResultResponse**](JsonRpcResultResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentTool

> CreateAgentTool201Response UpdateAgentTool(ctx, id).UpdateTool(updateTool).Execute()

Изменение MCP-инструмента



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
    id := int32(417) // int32 | ID инструмента
    updateTool := *openapiclient.NewUpdateTool() // UpdateTool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MCPAPI.UpdateAgentTool(context.Background(), id).UpdateTool(updateTool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.UpdateAgentTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentTool`: CreateAgentTool201Response
    fmt.Fprintf(os.Stdout, "Response from `MCPAPI.UpdateAgentTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID инструмента | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTool** | [**UpdateTool**](UpdateTool.md) |  | 

### Return type

[**CreateAgentTool201Response**](CreateAgentTool201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentToolConnection

> UpdateAgentToolConnection(ctx, agentId, toolId).UpdateAgentTool(updateAgentTool).Execute()

Изменение разрешённых методов MCP-инструмента у агента



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
    agentId := int32(113191) // int32 | ID агента
    toolId := int32(417) // int32 | ID инструмента
    updateAgentTool := *openapiclient.NewUpdateAgentTool([]string{"AllowedMethods_example"}) // UpdateAgentTool | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MCPAPI.UpdateAgentToolConnection(context.Background(), agentId, toolId).UpdateAgentTool(updateAgentTool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.UpdateAgentToolConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 
**toolId** | **int32** | ID инструмента | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentToolConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAgentTool** | [**UpdateAgentTool**](UpdateAgentTool.md) |  | 

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

