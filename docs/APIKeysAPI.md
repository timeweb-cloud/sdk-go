# \APIKeysAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](APIKeysAPI.md#CreateToken) | **Post** /api/v1/auth/api-keys | Создание токена
[**DeleteToken**](APIKeysAPI.md#DeleteToken) | **Delete** /api/v1/auth/api-keys/{token_id} | Удалить токен
[**GetTokens**](APIKeysAPI.md#GetTokens) | **Get** /api/v1/auth/api-keys | Получение списка выпущенных токенов
[**ReissueToken**](APIKeysAPI.md#ReissueToken) | **Put** /api/v1/auth/api-keys/{token_id} | Перевыпустить токен
[**UpdateToken**](APIKeysAPI.md#UpdateToken) | **Patch** /api/v1/auth/api-keys/{token_id} | Изменить токен



## CreateToken

> CreateToken201Response CreateToken(ctx).CreateApiKey(createApiKey).Execute()

Создание токена



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
    createApiKey := *openapiclient.NewCreateApiKey("Example") // CreateApiKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysAPI.CreateToken(context.Background()).CreateApiKey(createApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: CreateToken201Response
    fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiKey** | [**CreateApiKey**](CreateApiKey.md) |  | 

### Return type

[**CreateToken201Response**](CreateToken201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> DeleteToken(ctx, tokenId).Execute()

Удалить токен



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
    tokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID токена

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.APIKeysAPI.DeleteToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.DeleteToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | ID токена | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


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


## GetTokens

> GetTokens200Response GetTokens(ctx).Execute()

Получение списка выпущенных токенов



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
    resp, r, err := apiClient.APIKeysAPI.GetTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.GetTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokens`: GetTokens200Response
    fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


### Return type

[**GetTokens200Response**](GetTokens200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReissueToken

> CreateToken201Response ReissueToken(ctx, tokenId).RefreshApiKey(refreshApiKey).Execute()

Перевыпустить токен



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
    tokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID токена
    refreshApiKey := *openapiclient.NewRefreshApiKey() // RefreshApiKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysAPI.ReissueToken(context.Background(), tokenId).RefreshApiKey(refreshApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.ReissueToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReissueToken`: CreateToken201Response
    fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.ReissueToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | ID токена | 

### Other Parameters

Other parameters are passed through a pointer to a apiReissueTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshApiKey** | [**RefreshApiKey**](RefreshApiKey.md) |  | 

### Return type

[**CreateToken201Response**](CreateToken201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToken

> UpdateToken200Response UpdateToken(ctx, tokenId).EditApiKey(editApiKey).Execute()

Изменить токен



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
    tokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID токена
    editApiKey := *openapiclient.NewEditApiKey() // EditApiKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysAPI.UpdateToken(context.Background(), tokenId).EditApiKey(editApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.UpdateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateToken`: UpdateToken200Response
    fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.UpdateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | ID токена | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editApiKey** | [**EditApiKey**](EditApiKey.md) |  | 

### Return type

[**UpdateToken200Response**](UpdateToken200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

