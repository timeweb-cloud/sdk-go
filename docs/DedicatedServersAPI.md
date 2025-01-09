# \DedicatedServersAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDedicatedServer**](DedicatedServersAPI.md#CreateDedicatedServer) | **Post** /api/v1/dedicated-servers | Создание выделенного сервера
[**DeleteDedicatedServer**](DedicatedServersAPI.md#DeleteDedicatedServer) | **Delete** /api/v1/dedicated-servers/{dedicated_id} | Удаление выделенного сервера
[**GetDedicatedServer**](DedicatedServersAPI.md#GetDedicatedServer) | **Get** /api/v1/dedicated-servers/{dedicated_id} | Получение выделенного сервера
[**GetDedicatedServerPresetAdditionalServices**](DedicatedServersAPI.md#GetDedicatedServerPresetAdditionalServices) | **Get** /api/v1/presets/dedicated-servers/{preset_id}/additional-services | Получение дополнительных услуг для выделенного сервера
[**GetDedicatedServers**](DedicatedServersAPI.md#GetDedicatedServers) | **Get** /api/v1/dedicated-servers | Получение списка выделенных серверов
[**GetDedicatedServersPresets**](DedicatedServersAPI.md#GetDedicatedServersPresets) | **Get** /api/v1/presets/dedicated-servers | Получение списка тарифов для выделенного сервера
[**UpdateDedicatedServer**](DedicatedServersAPI.md#UpdateDedicatedServer) | **Patch** /api/v1/dedicated-servers/{dedicated_id} | Обновление выделенного сервера



## CreateDedicatedServer

> CreateDedicatedServer201Response CreateDedicatedServer(ctx).CreateDedicatedServer(createDedicatedServer).Execute()

Создание выделенного сервера



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
    createDedicatedServer := *openapiclient.NewCreateDedicatedServer(float32(2377), float32(81), "P1M", "name") // CreateDedicatedServer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DedicatedServersAPI.CreateDedicatedServer(context.Background()).CreateDedicatedServer(createDedicatedServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServersAPI.CreateDedicatedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDedicatedServer`: CreateDedicatedServer201Response
    fmt.Fprintf(os.Stdout, "Response from `DedicatedServersAPI.CreateDedicatedServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDedicatedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDedicatedServer** | [**CreateDedicatedServer**](CreateDedicatedServer.md) |  | 

### Return type

[**CreateDedicatedServer201Response**](CreateDedicatedServer201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDedicatedServer

> DeleteDedicatedServer(ctx, dedicatedId).Execute()

Удаление выделенного сервера



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
    dedicatedId := int32(1051) // int32 | ID выделенного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DedicatedServersAPI.DeleteDedicatedServer(context.Background(), dedicatedId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServersAPI.DeleteDedicatedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dedicatedId** | **int32** | ID выделенного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDedicatedServerRequest struct via the builder pattern


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


## GetDedicatedServer

> CreateDedicatedServer201Response GetDedicatedServer(ctx, dedicatedId).Execute()

Получение выделенного сервера



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
    dedicatedId := int32(1051) // int32 | ID выделенного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DedicatedServersAPI.GetDedicatedServer(context.Background(), dedicatedId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServersAPI.GetDedicatedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDedicatedServer`: CreateDedicatedServer201Response
    fmt.Fprintf(os.Stdout, "Response from `DedicatedServersAPI.GetDedicatedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dedicatedId** | **int32** | ID выделенного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDedicatedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDedicatedServer201Response**](CreateDedicatedServer201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDedicatedServerPresetAdditionalServices

> GetDedicatedServerPresetAdditionalServices200Response GetDedicatedServerPresetAdditionalServices(ctx, presetId).Execute()

Получение дополнительных услуг для выделенного сервера



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
    presetId := int32(1051) // int32 | ID тарифа выделенного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DedicatedServersAPI.GetDedicatedServerPresetAdditionalServices(context.Background(), presetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServersAPI.GetDedicatedServerPresetAdditionalServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDedicatedServerPresetAdditionalServices`: GetDedicatedServerPresetAdditionalServices200Response
    fmt.Fprintf(os.Stdout, "Response from `DedicatedServersAPI.GetDedicatedServerPresetAdditionalServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presetId** | **int32** | ID тарифа выделенного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDedicatedServerPresetAdditionalServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDedicatedServerPresetAdditionalServices200Response**](GetDedicatedServerPresetAdditionalServices200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDedicatedServers

> GetDedicatedServers200Response GetDedicatedServers(ctx).Execute()

Получение списка выделенных серверов



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
    resp, r, err := apiClient.DedicatedServersAPI.GetDedicatedServers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServersAPI.GetDedicatedServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDedicatedServers`: GetDedicatedServers200Response
    fmt.Fprintf(os.Stdout, "Response from `DedicatedServersAPI.GetDedicatedServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDedicatedServersRequest struct via the builder pattern


### Return type

[**GetDedicatedServers200Response**](GetDedicatedServers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDedicatedServersPresets

> GetDedicatedServersPresets200Response GetDedicatedServersPresets(ctx).Location(location).Execute()

Получение списка тарифов для выделенного сервера



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
    location := "location_example" // string | Получение тарифов определенной локации. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DedicatedServersAPI.GetDedicatedServersPresets(context.Background()).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServersAPI.GetDedicatedServersPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDedicatedServersPresets`: GetDedicatedServersPresets200Response
    fmt.Fprintf(os.Stdout, "Response from `DedicatedServersAPI.GetDedicatedServersPresets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDedicatedServersPresetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | Получение тарифов определенной локации. | 

### Return type

[**GetDedicatedServersPresets200Response**](GetDedicatedServersPresets200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDedicatedServer

> CreateDedicatedServer201Response UpdateDedicatedServer(ctx, dedicatedId).UpdateDedicatedServerRequest(updateDedicatedServerRequest).Execute()

Обновление выделенного сервера



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
    dedicatedId := int32(1051) // int32 | ID выделенного сервера.
    updateDedicatedServerRequest := *openapiclient.NewUpdateDedicatedServerRequest() // UpdateDedicatedServerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DedicatedServersAPI.UpdateDedicatedServer(context.Background(), dedicatedId).UpdateDedicatedServerRequest(updateDedicatedServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DedicatedServersAPI.UpdateDedicatedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDedicatedServer`: CreateDedicatedServer201Response
    fmt.Fprintf(os.Stdout, "Response from `DedicatedServersAPI.UpdateDedicatedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dedicatedId** | **int32** | ID выделенного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDedicatedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDedicatedServerRequest** | [**UpdateDedicatedServerRequest**](UpdateDedicatedServerRequest.md) |  | 

### Return type

[**CreateDedicatedServer201Response**](CreateDedicatedServer201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

