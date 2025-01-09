# \SSHAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKeyToServer**](SSHAPI.md#AddKeyToServer) | **Post** /api/v1/servers/{server_id}/ssh-keys | Добавление SSH-ключей на сервер
[**CreateKey**](SSHAPI.md#CreateKey) | **Post** /api/v1/ssh-keys | Создание SSH-ключа
[**DeleteKey**](SSHAPI.md#DeleteKey) | **Delete** /api/v1/ssh-keys/{ssh_key_id} | Удаление SSH-ключа по ID
[**DeleteKeyFromServer**](SSHAPI.md#DeleteKeyFromServer) | **Delete** /api/v1/servers/{server_id}/ssh-keys/{ssh_key_id} | Удаление SSH-ключей с сервера
[**GetKey**](SSHAPI.md#GetKey) | **Get** /api/v1/ssh-keys/{ssh_key_id} | Получение SSH-ключа по ID
[**GetKeys**](SSHAPI.md#GetKeys) | **Get** /api/v1/ssh-keys | Получение списка SSH-ключей
[**UpdateKey**](SSHAPI.md#UpdateKey) | **Patch** /api/v1/ssh-keys/{ssh_key_id} | Изменение SSH-ключа по ID



## AddKeyToServer

> AddKeyToServer(ctx, serverId).AddKeyToServerRequest(addKeyToServerRequest).Execute()

Добавление SSH-ключей на сервер



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    addKeyToServerRequest := *openapiclient.NewAddKeyToServerRequest([]float32{float32(1)}) // AddKeyToServerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSHAPI.AddKeyToServer(context.Background(), serverId).AddKeyToServerRequest(addKeyToServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHAPI.AddKeyToServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddKeyToServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addKeyToServerRequest** | [**AddKeyToServerRequest**](AddKeyToServerRequest.md) |  | 

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


## CreateKey

> CreateKey201Response CreateKey(ctx).CreateKeyRequest(createKeyRequest).Execute()

Создание SSH-ключа



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
    createKeyRequest := *openapiclient.NewCreateKeyRequest("ssh-dss AAAAB3NzaC1kc3MAAACBALixbc+uF2kuv83LiB3gg200vHJ7hH0QCC77fqlu+5raX6Y1PJZZORwzWbvU6k8gfrmjMRwQFDNncu3aXUiM9TleOBxFsQ2Hg5hDblSr2bnnyVP0jjPzsYye+0Ebaa+IjEqZz8y0qu8zejCpHqhLP/PlRDS9vtTsadrVvNwFFCyhAAAAFQD7m0AZOd/lHraIgBjXJvfN+qGiAQAAAIAdXSV0n47akGC/g0ejkNzHDSdRFf2fR0FsOos/7f5j+QfstijVBMwzA4Ti+dNJv3BQqQv3Bk5X3o5yUTUL+cM6fa4KSyIhGi14RL0kqjJrNjqVAhwjPH8K9aHDouqQEz56UL72+iGhMxwBHBMWq7cZ2k31mmFgEoYgwNcYuNmtlAAAAIBxIebJMItpq3ez5Hg5GyhpnSU7h0kshh+D5UAAlPmJo8OUQH6avQGf6gpxs3CnQ/im/aCKcb/bYm4V5Y9a7Zq+LrZGQZ79GLWsu/oK0HfQiP+oWwzCGmsdIGblsTRzrKU0FFo7zGmyTsBU48OeD76nKsc5EJ9F9wQC3S2KK38Ynw== test@MBP-test.office.timeweb.ru", false, "test") // CreateKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHAPI.CreateKey(context.Background()).CreateKeyRequest(createKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHAPI.CreateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKey`: CreateKey201Response
    fmt.Fprintf(os.Stdout, "Response from `SSHAPI.CreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKeyRequest** | [**CreateKeyRequest**](CreateKeyRequest.md) |  | 

### Return type

[**CreateKey201Response**](CreateKey201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> DeleteKey(ctx, sshKeyId).Execute()

Удаление SSH-ключа по ID



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
    sshKeyId := int32(1051) // int32 | ID SSH-ключа.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSHAPI.DeleteKey(context.Background(), sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHAPI.DeleteKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | ID SSH-ключа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


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


## DeleteKeyFromServer

> DeleteKeyFromServer(ctx, serverId, sshKeyId).Execute()

Удаление SSH-ключей с сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    sshKeyId := int32(1051) // int32 | ID SSH-ключа.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSHAPI.DeleteKeyFromServer(context.Background(), serverId, sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHAPI.DeleteKeyFromServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**sshKeyId** | **int32** | ID SSH-ключа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyFromServerRequest struct via the builder pattern


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


## GetKey

> GetKey200Response GetKey(ctx, sshKeyId).Execute()

Получение SSH-ключа по ID



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
    sshKeyId := int32(1051) // int32 | ID SSH-ключа.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHAPI.GetKey(context.Background(), sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHAPI.GetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKey`: GetKey200Response
    fmt.Fprintf(os.Stdout, "Response from `SSHAPI.GetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | ID SSH-ключа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetKey200Response**](GetKey200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeys

> GetKeys200Response GetKeys(ctx).Execute()

Получение списка SSH-ключей



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
    resp, r, err := apiClient.SSHAPI.GetKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHAPI.GetKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeys`: GetKeys200Response
    fmt.Fprintf(os.Stdout, "Response from `SSHAPI.GetKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeysRequest struct via the builder pattern


### Return type

[**GetKeys200Response**](GetKeys200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> GetKey200Response UpdateKey(ctx, sshKeyId).UpdateKeyRequest(updateKeyRequest).Execute()

Изменение SSH-ключа по ID



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
    sshKeyId := int32(1051) // int32 | ID SSH-ключа.
    updateKeyRequest := *openapiclient.NewUpdateKeyRequest() // UpdateKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHAPI.UpdateKey(context.Background(), sshKeyId).UpdateKeyRequest(updateKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHAPI.UpdateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKey`: GetKey200Response
    fmt.Fprintf(os.Stdout, "Response from `SSHAPI.UpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sshKeyId** | **int32** | ID SSH-ключа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKeyRequest** | [**UpdateKeyRequest**](UpdateKeyRequest.md) |  | 

### Return type

[**GetKey200Response**](GetKey200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

