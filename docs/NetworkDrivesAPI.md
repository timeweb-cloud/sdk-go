# \NetworkDrivesAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDrive**](NetworkDrivesAPI.md#CreateNetworkDrive) | **Post** /api/v1/network-drives | Создание сетевого диска
[**DeleteNetworkDrive**](NetworkDrivesAPI.md#DeleteNetworkDrive) | **Delete** /api/v1/network-drives/{network_drive_id} | Удаление сетевого диска по идентификатору
[**GetNetworkDrive**](NetworkDrivesAPI.md#GetNetworkDrive) | **Get** /api/v1/network-drives/{network_drive_id} | Получение сетевого диска
[**GetNetworkDrives**](NetworkDrivesAPI.md#GetNetworkDrives) | **Get** /api/v1/network-drives | Получение списка cетевых дисков
[**GetNetworkDrivesAvailableResources**](NetworkDrivesAPI.md#GetNetworkDrivesAvailableResources) | **Get** /api/v1/network-drives/available-resources | Получение списка сервисов доступных для подключения диска
[**GetNetworkDrivesPresets**](NetworkDrivesAPI.md#GetNetworkDrivesPresets) | **Get** /api/v1/presets/network-drives | Получение списка доступных тарифов для сетевого диска
[**MountNetworkDrive**](NetworkDrivesAPI.md#MountNetworkDrive) | **Post** /api/v1/network-drives/{network_drive_id}/mount | Подключить сетевой диск к сервису
[**UnmountNetworkDrive**](NetworkDrivesAPI.md#UnmountNetworkDrive) | **Post** /api/v1/network-drives/{network_drive_id}/unmount | Отключить сетевой диск от сервиса
[**UpdateNetworkDrive**](NetworkDrivesAPI.md#UpdateNetworkDrive) | **Patch** /api/v1/network-drives/{network_drive_id} | Изменение сетевого диска по ID



## CreateNetworkDrive

> CreateNetworkDrive201Response CreateNetworkDrive(ctx).CreateNetworkDrive(createNetworkDrive).Execute()

Создание сетевого диска



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
    createNetworkDrive := *openapiclient.NewCreateNetworkDrive("Cute Quail", float32(10), int32(5)) // CreateNetworkDrive | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDrivesAPI.CreateNetworkDrive(context.Background()).CreateNetworkDrive(createNetworkDrive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.CreateNetworkDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkDrive`: CreateNetworkDrive201Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkDrivesAPI.CreateNetworkDrive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkDrive** | [**CreateNetworkDrive**](CreateNetworkDrive.md) |  | 

### Return type

[**CreateNetworkDrive201Response**](CreateNetworkDrive201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDrive

> DeleteNetworkDrive(ctx, networkDriveId).Execute()

Удаление сетевого диска по идентификатору



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
    networkDriveId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID сетевого диска

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkDrivesAPI.DeleteNetworkDrive(context.Background(), networkDriveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.DeleteNetworkDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDriveId** | **string** | ID сетевого диска | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDriveRequest struct via the builder pattern


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


## GetNetworkDrive

> CreateNetworkDrive201Response GetNetworkDrive(ctx, networkDriveId).Execute()

Получение сетевого диска



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
    networkDriveId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID сетевого диска

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDrivesAPI.GetNetworkDrive(context.Background(), networkDriveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.GetNetworkDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDrive`: CreateNetworkDrive201Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkDrivesAPI.GetNetworkDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDriveId** | **string** | ID сетевого диска | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateNetworkDrive201Response**](CreateNetworkDrive201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDrives

> GetNetworkDrives200Response GetNetworkDrives(ctx).Execute()

Получение списка cетевых дисков



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
    resp, r, err := apiClient.NetworkDrivesAPI.GetNetworkDrives(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.GetNetworkDrives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDrives`: GetNetworkDrives200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkDrivesAPI.GetNetworkDrives`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDrivesRequest struct via the builder pattern


### Return type

[**GetNetworkDrives200Response**](GetNetworkDrives200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDrivesAvailableResources

> GetNetworkDrivesAvailableResources200Response GetNetworkDrivesAvailableResources(ctx).Execute()

Получение списка сервисов доступных для подключения диска



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
    resp, r, err := apiClient.NetworkDrivesAPI.GetNetworkDrivesAvailableResources(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.GetNetworkDrivesAvailableResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDrivesAvailableResources`: GetNetworkDrivesAvailableResources200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkDrivesAPI.GetNetworkDrivesAvailableResources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDrivesAvailableResourcesRequest struct via the builder pattern


### Return type

[**GetNetworkDrivesAvailableResources200Response**](GetNetworkDrivesAvailableResources200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDrivesPresets

> GetNetworkDrivesPresets200Response GetNetworkDrivesPresets(ctx).Execute()

Получение списка доступных тарифов для сетевого диска



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
    resp, r, err := apiClient.NetworkDrivesAPI.GetNetworkDrivesPresets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.GetNetworkDrivesPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDrivesPresets`: GetNetworkDrivesPresets200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkDrivesAPI.GetNetworkDrivesPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDrivesPresetsRequest struct via the builder pattern


### Return type

[**GetNetworkDrivesPresets200Response**](GetNetworkDrivesPresets200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MountNetworkDrive

> MountNetworkDrive(ctx, networkDriveId).MountNetworkDrive(mountNetworkDrive).Execute()

Подключить сетевой диск к сервису



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
    networkDriveId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID сетевого диска
    mountNetworkDrive := *openapiclient.NewMountNetworkDrive("server", float32(24569)) // MountNetworkDrive | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkDrivesAPI.MountNetworkDrive(context.Background(), networkDriveId).MountNetworkDrive(mountNetworkDrive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.MountNetworkDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDriveId** | **string** | ID сетевого диска | 

### Other Parameters

Other parameters are passed through a pointer to a apiMountNetworkDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mountNetworkDrive** | [**MountNetworkDrive**](MountNetworkDrive.md) |  | 

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


## UnmountNetworkDrive

> UnmountNetworkDrive(ctx, networkDriveId).Execute()

Отключить сетевой диск от сервиса



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
    networkDriveId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID сетевого диска

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkDrivesAPI.UnmountNetworkDrive(context.Background(), networkDriveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.UnmountNetworkDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDriveId** | **string** | ID сетевого диска | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmountNetworkDriveRequest struct via the builder pattern


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


## UpdateNetworkDrive

> CreateNetworkDrive201Response UpdateNetworkDrive(ctx, networkDriveId).UpdateNetworkDrive(updateNetworkDrive).Execute()

Изменение сетевого диска по ID



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
    networkDriveId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID сетевого диска
    updateNetworkDrive := *openapiclient.NewUpdateNetworkDrive() // UpdateNetworkDrive | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDrivesAPI.UpdateNetworkDrive(context.Background(), networkDriveId).UpdateNetworkDrive(updateNetworkDrive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDrivesAPI.UpdateNetworkDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkDrive`: CreateNetworkDrive201Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkDrivesAPI.UpdateNetworkDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkDriveId** | **string** | ID сетевого диска | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDrive** | [**UpdateNetworkDrive**](UpdateNetworkDrive.md) |  | 

### Return type

[**CreateNetworkDrive201Response**](CreateNetworkDrive201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

