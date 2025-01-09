# \VPCAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVPC**](VPCAPI.md#CreateVPC) | **Post** /api/v2/vpcs | Создание VPC
[**DeleteVPC**](VPCAPI.md#DeleteVPC) | **Delete** /api/v1/vpcs/{vpc_id} | Удаление VPC по ID сети
[**GetVPC**](VPCAPI.md#GetVPC) | **Get** /api/v2/vpcs/{vpc_id} | Получение VPC
[**GetVPCPorts**](VPCAPI.md#GetVPCPorts) | **Get** /api/v1/vpcs/{vpc_id}/ports | Получение списка портов для VPC
[**GetVPCServices**](VPCAPI.md#GetVPCServices) | **Get** /api/v2/vpcs/{vpc_id}/services | Получение списка сервисов в VPC
[**GetVPCs**](VPCAPI.md#GetVPCs) | **Get** /api/v2/vpcs | Получение списка VPCs
[**UpdateVPCs**](VPCAPI.md#UpdateVPCs) | **Patch** /api/v2/vpcs/{vpc_id} | Изменение VPC по ID сети



## CreateVPC

> CreateVPC201Response CreateVPC(ctx).CreateVpc(createVpc).Execute()

Создание VPC



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
    createVpc := *openapiclient.NewCreateVpc("Общая сеть", "192.168.0.0/24", "ru-1") // CreateVpc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VPCAPI.CreateVPC(context.Background()).CreateVpc(createVpc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.CreateVPC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVPC`: CreateVPC201Response
    fmt.Fprintf(os.Stdout, "Response from `VPCAPI.CreateVPC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVPCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVpc** | [**CreateVpc**](CreateVpc.md) |  | 

### Return type

[**CreateVPC201Response**](CreateVPC201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVPC

> DeleteVPC(ctx, vpcId).Execute()

Удаление VPC по ID сети



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
    vpcId := "network-1234567890" // string | ID сети

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VPCAPI.DeleteVPC(context.Background(), vpcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.DeleteVPC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | ID сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVPCRequest struct via the builder pattern


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


## GetVPC

> CreateVPC201Response GetVPC(ctx, vpcId).Execute()

Получение VPC



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
    vpcId := "network-1234567890" // string | ID сети

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VPCAPI.GetVPC(context.Background(), vpcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.GetVPC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVPC`: CreateVPC201Response
    fmt.Fprintf(os.Stdout, "Response from `VPCAPI.GetVPC`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | ID сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVPCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateVPC201Response**](CreateVPC201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVPCPorts

> GetVPCPorts200Response GetVPCPorts(ctx, vpcId).Execute()

Получение списка портов для VPC



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
    vpcId := "network-1234567890" // string | ID сети

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VPCAPI.GetVPCPorts(context.Background(), vpcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.GetVPCPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVPCPorts`: GetVPCPorts200Response
    fmt.Fprintf(os.Stdout, "Response from `VPCAPI.GetVPCPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | ID сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVPCPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVPCPorts200Response**](GetVPCPorts200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVPCServices

> GetVPCServices200Response GetVPCServices(ctx, vpcId).Execute()

Получение списка сервисов в VPC



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
    vpcId := "network-1234567890" // string | ID сети

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VPCAPI.GetVPCServices(context.Background(), vpcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.GetVPCServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVPCServices`: GetVPCServices200Response
    fmt.Fprintf(os.Stdout, "Response from `VPCAPI.GetVPCServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | ID сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVPCServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVPCServices200Response**](GetVPCServices200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVPCs

> GetVPCs200Response GetVPCs(ctx).Execute()

Получение списка VPCs



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
    resp, r, err := apiClient.VPCAPI.GetVPCs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.GetVPCs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVPCs`: GetVPCs200Response
    fmt.Fprintf(os.Stdout, "Response from `VPCAPI.GetVPCs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVPCsRequest struct via the builder pattern


### Return type

[**GetVPCs200Response**](GetVPCs200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVPCs

> CreateVPC201Response UpdateVPCs(ctx, vpcId).UpdateVpc(updateVpc).Execute()

Изменение VPC по ID сети



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
    vpcId := "network-1234567890" // string | ID сети
    updateVpc := *openapiclient.NewUpdateVpc() // UpdateVpc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VPCAPI.UpdateVPCs(context.Background(), vpcId).UpdateVpc(updateVpc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.UpdateVPCs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVPCs`: CreateVPC201Response
    fmt.Fprintf(os.Stdout, "Response from `VPCAPI.UpdateVPCs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | ID сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVpc** | [**UpdateVpc**](UpdateVpc.md) |  | 

### Return type

[**CreateVPC201Response**](CreateVPC201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

