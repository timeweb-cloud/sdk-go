# \SnapshotsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitRestorePoint**](SnapshotsAPI.md#CommitRestorePoint) | **Post** /api/v1/restore-points/{vds_id}/commit | Фиксация снапшота
[**CreateRestorePoint**](SnapshotsAPI.md#CreateRestorePoint) | **Post** /api/v1/restore-points/{vds_id}/create | Создание снапшота
[**GetRestorePoint**](SnapshotsAPI.md#GetRestorePoint) | **Get** /api/v1/restore-points/{vds_id} | Получение снапшота сервера
[**GetRestorePoints**](SnapshotsAPI.md#GetRestorePoints) | **Get** /api/v1/restore-points | Получение списка снапшотов
[**RollbackRestorePoint**](SnapshotsAPI.md#RollbackRestorePoint) | **Post** /api/v1/restore-points/{vds_id}/rollback | Откат к снапшоту



## CommitRestorePoint

> CommitRestorePoint(ctx, vdsId).Execute()

Фиксация снапшота



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
    vdsId := int32(1051) // int32 | ID облачного сервера (VDS).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SnapshotsAPI.CommitRestorePoint(context.Background(), vdsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CommitRestorePoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdsId** | **int32** | ID облачного сервера (VDS). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitRestorePointRequest struct via the builder pattern


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


## CreateRestorePoint

> GetRestorePoint200Response CreateRestorePoint(ctx, vdsId).Execute()

Создание снапшота



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
    vdsId := int32(1051) // int32 | ID облачного сервера (VDS).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.CreateRestorePoint(context.Background(), vdsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateRestorePoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRestorePoint`: GetRestorePoint200Response
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateRestorePoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdsId** | **int32** | ID облачного сервера (VDS). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRestorePointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRestorePoint200Response**](GetRestorePoint200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestorePoint

> GetRestorePoint200Response GetRestorePoint(ctx, vdsId).Execute()

Получение снапшота сервера



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
    vdsId := int32(1051) // int32 | ID облачного сервера (VDS).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.GetRestorePoint(context.Background(), vdsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.GetRestorePoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestorePoint`: GetRestorePoint200Response
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.GetRestorePoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdsId** | **int32** | ID облачного сервера (VDS). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestorePointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRestorePoint200Response**](GetRestorePoint200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestorePoints

> GetRestorePoints200Response GetRestorePoints(ctx).Execute()

Получение списка снапшотов



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
    resp, r, err := apiClient.SnapshotsAPI.GetRestorePoints(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.GetRestorePoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestorePoints`: GetRestorePoints200Response
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.GetRestorePoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestorePointsRequest struct via the builder pattern


### Return type

[**GetRestorePoints200Response**](GetRestorePoints200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackRestorePoint

> RollbackRestorePoint(ctx, vdsId).Execute()

Откат к снапшоту



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
    vdsId := int32(1051) // int32 | ID облачного сервера (VDS).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SnapshotsAPI.RollbackRestorePoint(context.Background(), vdsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RollbackRestorePoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vdsId** | **int32** | ID облачного сервера (VDS). | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackRestorePointRequest struct via the builder pattern


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

