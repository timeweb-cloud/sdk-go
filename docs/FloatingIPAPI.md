# \FloatingIPAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindFloatingIp**](FloatingIPAPI.md#BindFloatingIp) | **Post** /api/v1/floating-ips/{floating_ip_id}/bind | Привязать IP к сервису
[**CreateFloatingIp**](FloatingIPAPI.md#CreateFloatingIp) | **Post** /api/v1/floating-ips | Создание плавающего IP
[**DeleteFloatingIP**](FloatingIPAPI.md#DeleteFloatingIP) | **Delete** /api/v1/floating-ips/{floating_ip_id} | Удаление плавающего IP по ID
[**GetFloatingIp**](FloatingIPAPI.md#GetFloatingIp) | **Get** /api/v1/floating-ips/{floating_ip_id} | Получение плавающего IP
[**GetFloatingIps**](FloatingIPAPI.md#GetFloatingIps) | **Get** /api/v1/floating-ips | Получение списка плавающих IP
[**UnbindFloatingIp**](FloatingIPAPI.md#UnbindFloatingIp) | **Post** /api/v1/floating-ips/{floating_ip_id}/unbind | Отвязать IP от сервиса
[**UpdateFloatingIP**](FloatingIPAPI.md#UpdateFloatingIP) | **Patch** /api/v1/floating-ips/{floating_ip_id} | Изменение плавающего IP по ID



## BindFloatingIp

> BindFloatingIp(ctx, floatingIpId).BindFloatingIp(bindFloatingIp).Execute()

Привязать IP к сервису



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
    floatingIpId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID плавающего IP
    bindFloatingIp := *openapiclient.NewBindFloatingIp("server", float32(24569)) // BindFloatingIp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FloatingIPAPI.BindFloatingIp(context.Background(), floatingIpId).BindFloatingIp(bindFloatingIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPAPI.BindFloatingIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**floatingIpId** | **string** | ID плавающего IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindFloatingIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindFloatingIp** | [**BindFloatingIp**](BindFloatingIp.md) |  | 

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


## CreateFloatingIp

> CreateFloatingIp201Response CreateFloatingIp(ctx).CreateFloatingIp(createFloatingIp).Execute()

Создание плавающего IP



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
    createFloatingIp := *openapiclient.NewCreateFloatingIp(true, openapiclient.availability-zone("spb-1")) // CreateFloatingIp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPAPI.CreateFloatingIp(context.Background()).CreateFloatingIp(createFloatingIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPAPI.CreateFloatingIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFloatingIp`: CreateFloatingIp201Response
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPAPI.CreateFloatingIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFloatingIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFloatingIp** | [**CreateFloatingIp**](CreateFloatingIp.md) |  | 

### Return type

[**CreateFloatingIp201Response**](CreateFloatingIp201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFloatingIP

> DeleteFloatingIP(ctx, floatingIpId).Execute()

Удаление плавающего IP по ID



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
    floatingIpId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID плавающего IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FloatingIPAPI.DeleteFloatingIP(context.Background(), floatingIpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPAPI.DeleteFloatingIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**floatingIpId** | **string** | ID плавающего IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFloatingIPRequest struct via the builder pattern


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


## GetFloatingIp

> CreateFloatingIp201Response GetFloatingIp(ctx, floatingIpId).Execute()

Получение плавающего IP



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
    floatingIpId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID плавающего IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPAPI.GetFloatingIp(context.Background(), floatingIpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPAPI.GetFloatingIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFloatingIp`: CreateFloatingIp201Response
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPAPI.GetFloatingIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**floatingIpId** | **string** | ID плавающего IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFloatingIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateFloatingIp201Response**](CreateFloatingIp201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFloatingIps

> GetFloatingIps200Response GetFloatingIps(ctx).Execute()

Получение списка плавающих IP



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
    resp, r, err := apiClient.FloatingIPAPI.GetFloatingIps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPAPI.GetFloatingIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFloatingIps`: GetFloatingIps200Response
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPAPI.GetFloatingIps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFloatingIpsRequest struct via the builder pattern


### Return type

[**GetFloatingIps200Response**](GetFloatingIps200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindFloatingIp

> UnbindFloatingIp(ctx, floatingIpId).Execute()

Отвязать IP от сервиса



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
    floatingIpId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID плавающего IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FloatingIPAPI.UnbindFloatingIp(context.Background(), floatingIpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPAPI.UnbindFloatingIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**floatingIpId** | **string** | ID плавающего IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindFloatingIpRequest struct via the builder pattern


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


## UpdateFloatingIP

> CreateFloatingIp201Response UpdateFloatingIP(ctx, floatingIpId).UpdateFloatingIp(updateFloatingIp).Execute()

Изменение плавающего IP по ID



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
    floatingIpId := "87fa289f-1513-4c4d-8d49-5707f411f14b" // string | ID плавающего IP
    updateFloatingIp := *openapiclient.NewUpdateFloatingIp() // UpdateFloatingIp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPAPI.UpdateFloatingIP(context.Background(), floatingIpId).UpdateFloatingIp(updateFloatingIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPAPI.UpdateFloatingIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFloatingIP`: CreateFloatingIp201Response
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPAPI.UpdateFloatingIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**floatingIpId** | **string** | ID плавающего IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFloatingIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFloatingIp** | [**UpdateFloatingIp**](UpdateFloatingIp.md) |  | 

### Return type

[**CreateFloatingIp201Response**](CreateFloatingIp201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

