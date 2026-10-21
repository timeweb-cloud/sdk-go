# \MCPAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMcpSearchSession**](MCPAPI.md#DeleteMcpSearchSession) | **Delete** /api/v1/mcp/search | Завершение сессии сервера MCP с поиском инструментов
[**DeleteMcpSession**](MCPAPI.md#DeleteMcpSession) | **Delete** /api/v1/mcp | Завершение сессии сервера MCP
[**OpenMcpSearchStream**](MCPAPI.md#OpenMcpSearchStream) | **Get** /api/v1/mcp/search | Открытие потока событий сервера MCP с поиском инструментов
[**OpenMcpStream**](MCPAPI.md#OpenMcpStream) | **Get** /api/v1/mcp | Открытие потока событий сервера MCP
[**SendMcpMessage**](MCPAPI.md#SendMcpMessage) | **Post** /api/v1/mcp | Отправка JSON-RPC-сообщения серверу MCP
[**SendMcpSearchMessage**](MCPAPI.md#SendMcpSearchMessage) | **Post** /api/v1/mcp/search | Отправка JSON-RPC-сообщения серверу MCP с поиском инструментов



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

