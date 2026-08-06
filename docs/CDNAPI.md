# \CDNAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCdnCertificate**](CDNAPI.md#AddCdnCertificate) | **Post** /api/v1/cdn/certificates | Загрузка собственного сертификата CDN
[**ArchiveCdnCertificateTask**](CDNAPI.md#ArchiveCdnCertificateTask) | **Post** /api/v1/cdn/certificates/tasks/{task_id}/archive | Архивация задачи на выпуск сертификата
[**ClearCdnResourceCache**](CDNAPI.md#ClearCdnResourceCache) | **Post** /api/v1/cdn/http-resources/{resource_id}/clear-cache | Очистка кэша CDN-ресурса
[**CreateCdnResource**](CDNAPI.md#CreateCdnResource) | **Post** /api/v1/cdn/http-resources | Создание CDN-ресурса
[**DeleteCdnCertificate**](CDNAPI.md#DeleteCdnCertificate) | **Delete** /api/v1/cdn/certificates/{certificate_id} | Удаление сертификата CDN
[**DeleteCdnResource**](CDNAPI.md#DeleteCdnResource) | **Delete** /api/v1/cdn/http-resources/{resource_id} | Удаление CDN-ресурса
[**GetCdnCertificateTasks**](CDNAPI.md#GetCdnCertificateTasks) | **Get** /api/v1/cdn/certificates/tasks | Получение списка задач на выпуск сертификатов
[**GetCdnCertificates**](CDNAPI.md#GetCdnCertificates) | **Get** /api/v1/cdn/certificates | Получение списка сертификатов CDN
[**GetCdnOriginNodes**](CDNAPI.md#GetCdnOriginNodes) | **Get** /api/v1/cdn/nodes/origin | Получение списка подсетей узлов CDN
[**GetCdnPresets**](CDNAPI.md#GetCdnPresets) | **Get** /api/v1/cdn/presets | Получение списка тарифов CDN
[**GetCdnResource**](CDNAPI.md#GetCdnResource) | **Get** /api/v1/cdn/http-resources/{resource_id} | Получение CDN-ресурса
[**GetCdnResourceConfiguration**](CDNAPI.md#GetCdnResourceConfiguration) | **Get** /api/v1/cdn/http-resources/{resource_id}/configuration | Получение конфигурации CDN-ресурса
[**GetCdnResourceNodes**](CDNAPI.md#GetCdnResourceNodes) | **Get** /api/v1/cdn/nodes/http-resources/{resource_id} | Получение списка раздающих узлов CDN-ресурса
[**GetCdnResourceStatistics**](CDNAPI.md#GetCdnResourceStatistics) | **Get** /api/v1/cdn/http-resources/{resource_id}/statistics | Получение статистики CDN-ресурса
[**GetCdnResources**](CDNAPI.md#GetCdnResources) | **Get** /api/v1/cdn/http-resources | Получение списка CDN-ресурсов
[**IssueCdnCertificate**](CDNAPI.md#IssueCdnCertificate) | **Post** /api/v1/cdn/certificates/issue | Выпуск сертификата Let&#39;s Encrypt для CDN-ресурса
[**PreloadCdnResourceCache**](CDNAPI.md#PreloadCdnResourceCache) | **Post** /api/v1/cdn/http-resources/{resource_id}/preload-cache | Предварительная загрузка кэша CDN-ресурса
[**ResumeCdnResource**](CDNAPI.md#ResumeCdnResource) | **Post** /api/v1/cdn/http-resources/{resource_id}/resume | Возобновление раздачи CDN-ресурса
[**SuspendCdnResource**](CDNAPI.md#SuspendCdnResource) | **Post** /api/v1/cdn/http-resources/{resource_id}/suspend | Приостановка раздачи CDN-ресурса
[**UpdateCdnResource**](CDNAPI.md#UpdateCdnResource) | **Patch** /api/v1/cdn/http-resources/{resource_id} | Изменение CDN-ресурса



## AddCdnCertificate

> AddCdnCertificate(ctx).AddCertificate(addCertificate).Execute()

Загрузка собственного сертификата CDN



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
    addCertificate := *openapiclient.NewAddCertificate("-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQ...
-----END PRIVATE KEY-----", "-----BEGIN CERTIFICATE-----
MIIFazCCA1OgAwIBAgIRAI...
-----END CERTIFICATE-----") // AddCertificate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CDNAPI.AddCdnCertificate(context.Background()).AddCertificate(addCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.AddCdnCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCdnCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCertificate** | [**AddCertificate**](AddCertificate.md) |  | 

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


## ArchiveCdnCertificateTask

> ArchiveCdnCertificateTask(ctx, taskId).Execute()

Архивация задачи на выпуск сертификата



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
    taskId := int32(42) // int32 | ID задачи на выпуск сертификата

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CDNAPI.ArchiveCdnCertificateTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.ArchiveCdnCertificateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** | ID задачи на выпуск сертификата | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveCdnCertificateTaskRequest struct via the builder pattern


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


## ClearCdnResourceCache

> ClearCdnResourceCache(ctx, resourceId).ClearCache(clearCache).Execute()

Очистка кэша CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса
    clearCache := *openapiclient.NewClearCache("partial") // ClearCache | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CDNAPI.ClearCdnResourceCache(context.Background(), resourceId).ClearCache(clearCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.ClearCdnResourceCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearCdnResourceCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clearCache** | [**ClearCache**](ClearCache.md) |  | 

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


## CreateCdnResource

> CreateCdnResource201Response CreateCdnResource(ctx).CreateHttpResource(createHttpResource).Execute()

Создание CDN-ресурса



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
    createHttpResource := *openapiclient.NewCreateHttpResource("Статика интернет-магазина", int32(4001)) // CreateHttpResource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.CreateCdnResource(context.Background()).CreateHttpResource(createHttpResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.CreateCdnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCdnResource`: CreateCdnResource201Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.CreateCdnResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createHttpResource** | [**CreateHttpResource**](CreateHttpResource.md) |  | 

### Return type

[**CreateCdnResource201Response**](CreateCdnResource201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCdnCertificate

> DeleteCdnCertificate(ctx, certificateId).Execute()

Удаление сертификата CDN



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
    certificateId := int32(5678) // int32 | ID сертификата

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CDNAPI.DeleteCdnCertificate(context.Background(), certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.DeleteCdnCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **int32** | ID сертификата | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdnCertificateRequest struct via the builder pattern


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


## DeleteCdnResource

> DeleteCdnResource(ctx, resourceId).Execute()

Удаление CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CDNAPI.DeleteCdnResource(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.DeleteCdnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdnResourceRequest struct via the builder pattern


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


## GetCdnCertificateTasks

> GetCdnCertificateTasks200Response GetCdnCertificateTasks(ctx).ResourceId(resourceId).Execute()

Получение списка задач на выпуск сертификатов



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
    resourceId := int32(1234) // int32 | Оставить в выдаче только задачи указанного CDN-ресурса. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.GetCdnCertificateTasks(context.Background()).ResourceId(resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnCertificateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnCertificateTasks`: GetCdnCertificateTasks200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnCertificateTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnCertificateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | **int32** | Оставить в выдаче только задачи указанного CDN-ресурса. | 

### Return type

[**GetCdnCertificateTasks200Response**](GetCdnCertificateTasks200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdnCertificates

> GetCdnCertificates200Response GetCdnCertificates(ctx).ResourceId(resourceId).Execute()

Получение списка сертификатов CDN



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
    resourceId := int32(1234) // int32 | Оставить в выдаче только сертификаты, подходящие для доменов указанного CDN-ресурса. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.GetCdnCertificates(context.Background()).ResourceId(resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnCertificates`: GetCdnCertificates200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | **int32** | Оставить в выдаче только сертификаты, подходящие для доменов указанного CDN-ресурса. | 

### Return type

[**GetCdnCertificates200Response**](GetCdnCertificates200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdnOriginNodes

> GetCdnOriginNodes200Response GetCdnOriginNodes(ctx).WithExtraZones(withExtraZones).Execute()

Получение списка подсетей узлов CDN



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
    withExtraZones := true // bool | Добавить в выдачу узлы дополнительных зон раздачи. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.GetCdnOriginNodes(context.Background()).WithExtraZones(withExtraZones).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnOriginNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnOriginNodes`: GetCdnOriginNodes200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnOriginNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnOriginNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withExtraZones** | **bool** | Добавить в выдачу узлы дополнительных зон раздачи. | [default to false]

### Return type

[**GetCdnOriginNodes200Response**](GetCdnOriginNodes200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdnPresets

> GetCdnPresets200Response GetCdnPresets(ctx).Execute()

Получение списка тарифов CDN



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
    resp, r, err := apiClient.CDNAPI.GetCdnPresets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnPresets`: GetCdnPresets200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnPresetsRequest struct via the builder pattern


### Return type

[**GetCdnPresets200Response**](GetCdnPresets200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdnResource

> CreateCdnResource201Response GetCdnResource(ctx, resourceId).Execute()

Получение CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.GetCdnResource(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnResource`: CreateCdnResource201Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCdnResource201Response**](CreateCdnResource201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdnResourceConfiguration

> GetCdnResourceConfiguration200Response GetCdnResourceConfiguration(ctx, resourceId).Execute()

Получение конфигурации CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.GetCdnResourceConfiguration(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnResourceConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnResourceConfiguration`: GetCdnResourceConfiguration200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnResourceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnResourceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCdnResourceConfiguration200Response**](GetCdnResourceConfiguration200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdnResourceNodes

> GetCdnResourceNodes200Response GetCdnResourceNodes(ctx, resourceId).WithExtraZones(withExtraZones).Country(country).Execute()

Получение списка раздающих узлов CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса
    withExtraZones := true // bool | Добавить в выдачу узлы дополнительных зон раздачи. (optional) (default to false)
    country := []string{"Inner_example"} // []string | Оставить в выдаче только основные зоны раздачи в указанных странах. Коды стран в формате ISO 3166-1 alpha-2. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.GetCdnResourceNodes(context.Background(), resourceId).WithExtraZones(withExtraZones).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnResourceNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnResourceNodes`: GetCdnResourceNodes200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnResourceNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnResourceNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withExtraZones** | **bool** | Добавить в выдачу узлы дополнительных зон раздачи. | [default to false]
 **country** | **[]string** | Оставить в выдаче только основные зоны раздачи в указанных странах. Коды стран в формате ISO 3166-1 alpha-2. | 

### Return type

[**GetCdnResourceNodes200Response**](GetCdnResourceNodes200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdnResourceStatistics

> GetCdnResourceStatistics200Response GetCdnResourceStatistics(ctx, resourceId).From(from).To(to).Execute()

Получение статистики CDN-ресурса



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    resourceId := int32(1234) // int32 | ID CDN-ресурса
    from := time.Now() // time.Time | Начало периода в формате ISO 8601. По умолчанию — 6 часов назад. (optional)
    to := time.Now() // time.Time | Конец периода в формате ISO 8601. По умолчанию — текущий момент. Должен быть не раньше `from`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.GetCdnResourceStatistics(context.Background(), resourceId).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnResourceStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnResourceStatistics`: GetCdnResourceStatistics200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnResourceStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnResourceStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Начало периода в формате ISO 8601. По умолчанию — 6 часов назад. | 
 **to** | **time.Time** | Конец периода в формате ISO 8601. По умолчанию — текущий момент. Должен быть не раньше &#x60;from&#x60;. | 

### Return type

[**GetCdnResourceStatistics200Response**](GetCdnResourceStatistics200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCdnResources

> GetCdnResources200Response GetCdnResources(ctx).BucketId(bucketId).Execute()

Получение списка CDN-ресурсов



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
    bucketId := int32(4210) // int32 | Оставить в выдаче только ресурсы, источником контента которых является указанное S3-хранилище. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.GetCdnResources(context.Background()).BucketId(bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.GetCdnResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdnResources`: GetCdnResources200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.GetCdnResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCdnResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketId** | **int32** | Оставить в выдаче только ресурсы, источником контента которых является указанное S3-хранилище. | 

### Return type

[**GetCdnResources200Response**](GetCdnResources200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCdnCertificate

> IssueCdnCertificate(ctx).IssueCertificate(issueCertificate).Execute()

Выпуск сертификата Let's Encrypt для CDN-ресурса



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
    issueCertificate := *openapiclient.NewIssueCertificate(int32(1234)) // IssueCertificate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CDNAPI.IssueCdnCertificate(context.Background()).IssueCertificate(issueCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.IssueCdnCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueCdnCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueCertificate** | [**IssueCertificate**](IssueCertificate.md) |  | 

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


## PreloadCdnResourceCache

> PreloadCdnResourceCache(ctx, resourceId).PreloadCache(preloadCache).Execute()

Предварительная загрузка кэша CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса
    preloadCache := *openapiclient.NewPreloadCache([]string{"Paths_example"}) // PreloadCache | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CDNAPI.PreloadCdnResourceCache(context.Background(), resourceId).PreloadCache(preloadCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.PreloadCdnResourceCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreloadCdnResourceCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preloadCache** | [**PreloadCache**](PreloadCache.md) |  | 

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


## ResumeCdnResource

> CreateCdnResource201Response ResumeCdnResource(ctx, resourceId).Execute()

Возобновление раздачи CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.ResumeCdnResource(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.ResumeCdnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeCdnResource`: CreateCdnResource201Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.ResumeCdnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeCdnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCdnResource201Response**](CreateCdnResource201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendCdnResource

> CreateCdnResource201Response SuspendCdnResource(ctx, resourceId).Execute()

Приостановка раздачи CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.SuspendCdnResource(context.Background(), resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.SuspendCdnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendCdnResource`: CreateCdnResource201Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.SuspendCdnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendCdnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCdnResource201Response**](CreateCdnResource201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCdnResource

> CreateCdnResource201Response UpdateCdnResource(ctx, resourceId).UpdateHttpResource(updateHttpResource).Execute()

Изменение CDN-ресурса



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
    resourceId := int32(1234) // int32 | ID CDN-ресурса
    updateHttpResource := *openapiclient.NewUpdateHttpResource() // UpdateHttpResource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAPI.UpdateCdnResource(context.Background(), resourceId).UpdateHttpResource(updateHttpResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAPI.UpdateCdnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCdnResource`: CreateCdnResource201Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAPI.UpdateCdnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **int32** | ID CDN-ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCdnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHttpResource** | [**UpdateHttpResource**](UpdateHttpResource.md) |  | 

### Return type

[**CreateCdnResource201Response**](CreateCdnResource201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

