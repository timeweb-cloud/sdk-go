# \KnowledgeBasesAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdditionalTokenPackageToKnowledgebase**](KnowledgeBasesAPI.md#AddAdditionalTokenPackageToKnowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/add-additional-token-package | Добавление дополнительного пакета токенов
[**CreateKnowledgebase**](KnowledgeBasesAPI.md#CreateKnowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases | Создание базы знаний
[**DeleteDocument**](KnowledgeBasesAPI.md#DeleteDocument) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id} | Удаление документа из базы знаний
[**DeleteKnowledgebase**](KnowledgeBasesAPI.md#DeleteKnowledgebase) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id} | Удаление базы знаний
[**DownloadDocument**](KnowledgeBasesAPI.md#DownloadDocument) | **Get** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id}/download | Скачивание документа из базы знаний
[**GetKnowledgebase**](KnowledgeBasesAPI.md#GetKnowledgebase) | **Get** /api/v1/cloud-ai/knowledge-bases/{id} | Получение базы знаний
[**GetKnowledgebaseDocumentsV2**](KnowledgeBasesAPI.md#GetKnowledgebaseDocumentsV2) | **Get** /api/v2/cloud-ai/knowledge-bases/{id}/documents | Получение списка документов базы знаний
[**GetKnowledgebaseStatistics**](KnowledgeBasesAPI.md#GetKnowledgebaseStatistics) | **Get** /api/v1/cloud-ai/knowledge-bases/{id}/statistic | Получение статистики использования токенов базы знаний
[**GetKnowledgebases**](KnowledgeBasesAPI.md#GetKnowledgebases) | **Get** /api/v1/cloud-ai/knowledge-bases | Получение списка баз знаний
[**GetKnowledgebasesV2**](KnowledgeBasesAPI.md#GetKnowledgebasesV2) | **Get** /api/v2/cloud-ai/knowledge-bases | Получение списка баз знаний (v2)
[**LinkKnowledgebaseToAgent**](KnowledgeBasesAPI.md#LinkKnowledgebaseToAgent) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/link/{agent_id} | Привязка базы знаний к агенту
[**ReindexDocument**](KnowledgeBasesAPI.md#ReindexDocument) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id}/reindex | Переиндексация документа
[**UnlinkKnowledgebaseFromAgent**](KnowledgeBasesAPI.md#UnlinkKnowledgebaseFromAgent) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id}/link/{agent_id} | Отвязка базы знаний от агента
[**UpdateKnowledgebase**](KnowledgeBasesAPI.md#UpdateKnowledgebase) | **Patch** /api/v1/cloud-ai/knowledge-bases/{id} | Обновление базы знаний
[**UploadFilesToKnowledgebase**](KnowledgeBasesAPI.md#UploadFilesToKnowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/upload | Загрузка файлов в базу знаний



## AddAdditionalTokenPackageToKnowledgebase

> AddAdditionalTokenPackageToKnowledgebase(ctx, id).AddTokenPackage(addTokenPackage).Execute()

Добавление дополнительного пакета токенов



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
    id := int32(1) // int32 | ID базы знаний
    addTokenPackage := *openapiclient.NewAddTokenPackage() // AddTokenPackage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.AddAdditionalTokenPackageToKnowledgebase(context.Background(), id).AddTokenPackage(addTokenPackage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.AddAdditionalTokenPackageToKnowledgebase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAdditionalTokenPackageToKnowledgebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addTokenPackage** | [**AddTokenPackage**](AddTokenPackage.md) |  | 

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


## CreateKnowledgebase

> CreateKnowledgebase201Response CreateKnowledgebase(ctx).CreateKnowledgebase(createKnowledgebase).Execute()

Создание базы знаний



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
    createKnowledgebase := *openapiclient.NewCreateKnowledgebase("Моя база знаний", float32(1), "net_12345", float32(1)) // CreateKnowledgebase | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.CreateKnowledgebase(context.Background()).CreateKnowledgebase(createKnowledgebase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.CreateKnowledgebase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKnowledgebase`: CreateKnowledgebase201Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.CreateKnowledgebase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKnowledgebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKnowledgebase** | [**CreateKnowledgebase**](CreateKnowledgebase.md) |  | 

### Return type

[**CreateKnowledgebase201Response**](CreateKnowledgebase201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> DeleteDocument(ctx, id, documentId).Execute()

Удаление документа из базы знаний



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
    id := int32(1) // int32 | ID базы знаний
    documentId := int32(1) // int32 | ID документа

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.DeleteDocument(context.Background(), id, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.DeleteDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 
**documentId** | **int32** | ID документа | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentRequest struct via the builder pattern


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


## DeleteKnowledgebase

> DeleteKnowledgebase(ctx, id).Execute()

Удаление базы знаний



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
    id := int32(1) // int32 | ID базы знаний

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.DeleteKnowledgebase(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.DeleteKnowledgebase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKnowledgebaseRequest struct via the builder pattern


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


## DownloadDocument

> *os.File DownloadDocument(ctx, id, documentId).Execute()

Скачивание документа из базы знаний



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
    id := int32(1) // int32 | ID базы знаний
    documentId := int32(1) // int32 | ID документа

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.DownloadDocument(context.Background(), id, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.DownloadDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadDocument`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.DownloadDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 
**documentId** | **int32** | ID документа | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgebase

> CreateKnowledgebase201Response GetKnowledgebase(ctx, id).Execute()

Получение базы знаний



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
    id := int32(1) // int32 | ID базы знаний

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.GetKnowledgebase(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.GetKnowledgebase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKnowledgebase`: CreateKnowledgebase201Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.GetKnowledgebase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateKnowledgebase201Response**](CreateKnowledgebase201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgebaseDocumentsV2

> GetKnowledgebaseDocumentsV2200Response GetKnowledgebaseDocumentsV2(ctx, id).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()

Получение списка документов базы знаний



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
    id := int32(1) // int32 | ID базы знаний
    limit := int32(20) // int32 | Количество документов на странице (по умолчанию: 10, максимум: 100) (optional) (default to 10)
    offset := int32(0) // int32 | Количество документов для пропуска (по умолчанию: 0) (optional) (default to 0)
    sortBy := "indexing_timestamp" // string | Поле для сортировки (по умолчанию: indexing_timestamp - время последней индексации документа) (optional) (default to "indexing_timestamp")
    sortOrder := "DESC" // string | Порядок сортировки (по умолчанию: DESC) (optional) (default to "DESC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.GetKnowledgebaseDocumentsV2(context.Background(), id).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.GetKnowledgebaseDocumentsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKnowledgebaseDocumentsV2`: GetKnowledgebaseDocumentsV2200Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.GetKnowledgebaseDocumentsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebaseDocumentsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Количество документов на странице (по умолчанию: 10, максимум: 100) | [default to 10]
 **offset** | **int32** | Количество документов для пропуска (по умолчанию: 0) | [default to 0]
 **sortBy** | **string** | Поле для сортировки (по умолчанию: indexing_timestamp - время последней индексации документа) | [default to &quot;indexing_timestamp&quot;]
 **sortOrder** | **string** | Порядок сортировки (по умолчанию: DESC) | [default to &quot;DESC&quot;]

### Return type

[**GetKnowledgebaseDocumentsV2200Response**](GetKnowledgebaseDocumentsV2200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgebaseStatistics

> GetKnowledgebaseStatistics200Response GetKnowledgebaseStatistics(ctx, id).StartTime(startTime).EndTime(endTime).Interval(interval).Execute()

Получение статистики использования токенов базы знаний



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
    id := int32(1) // int32 | ID базы знаний
    startTime := time.Now() // time.Time | Начало временного диапазона (ISO 8601) (optional)
    endTime := time.Now() // time.Time | Конец временного диапазона (ISO 8601) (optional)
    interval := float32(60) // float32 | Интервал в минутах (по умолчанию 60) (optional) (default to 60)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.GetKnowledgebaseStatistics(context.Background(), id).StartTime(startTime).EndTime(endTime).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.GetKnowledgebaseStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKnowledgebaseStatistics`: GetKnowledgebaseStatistics200Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.GetKnowledgebaseStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebaseStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **time.Time** | Начало временного диапазона (ISO 8601) | 
 **endTime** | **time.Time** | Конец временного диапазона (ISO 8601) | 
 **interval** | **float32** | Интервал в минутах (по умолчанию 60) | [default to 60]

### Return type

[**GetKnowledgebaseStatistics200Response**](GetKnowledgebaseStatistics200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgebases

> GetKnowledgebases200Response GetKnowledgebases(ctx).Execute()

Получение списка баз знаний



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
    resp, r, err := apiClient.KnowledgeBasesAPI.GetKnowledgebases(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.GetKnowledgebases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKnowledgebases`: GetKnowledgebases200Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.GetKnowledgebases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebasesRequest struct via the builder pattern


### Return type

[**GetKnowledgebases200Response**](GetKnowledgebases200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgebasesV2

> GetKnowledgebasesV2200Response GetKnowledgebasesV2(ctx).Execute()

Получение списка баз знаний (v2)



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
    resp, r, err := apiClient.KnowledgeBasesAPI.GetKnowledgebasesV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.GetKnowledgebasesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKnowledgebasesV2`: GetKnowledgebasesV2200Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.GetKnowledgebasesV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebasesV2Request struct via the builder pattern


### Return type

[**GetKnowledgebasesV2200Response**](GetKnowledgebasesV2200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkKnowledgebaseToAgent

> LinkKnowledgebaseToAgent(ctx, id, agentId).Execute()

Привязка базы знаний к агенту



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
    id := int32(1) // int32 | ID базы знаний
    agentId := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.LinkKnowledgebaseToAgent(context.Background(), id, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.LinkKnowledgebaseToAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkKnowledgebaseToAgentRequest struct via the builder pattern


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


## ReindexDocument

> ReindexDocument(ctx, id, documentId).Execute()

Переиндексация документа



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
    id := int32(1) // int32 | ID базы знаний
    documentId := int32(1) // int32 | ID документа

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.ReindexDocument(context.Background(), id, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.ReindexDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 
**documentId** | **int32** | ID документа | 

### Other Parameters

Other parameters are passed through a pointer to a apiReindexDocumentRequest struct via the builder pattern


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


## UnlinkKnowledgebaseFromAgent

> UnlinkKnowledgebaseFromAgent(ctx, id, agentId).Execute()

Отвязка базы знаний от агента



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
    id := int32(1) // int32 | ID базы знаний
    agentId := int32(1) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.UnlinkKnowledgebaseFromAgent(context.Background(), id, agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.UnlinkKnowledgebaseFromAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkKnowledgebaseFromAgentRequest struct via the builder pattern


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


## UpdateKnowledgebase

> CreateKnowledgebase201Response UpdateKnowledgebase(ctx, id).UpdateKnowledgebase(updateKnowledgebase).Execute()

Обновление базы знаний



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
    id := int32(1) // int32 | ID базы знаний
    updateKnowledgebase := *openapiclient.NewUpdateKnowledgebase() // UpdateKnowledgebase | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.UpdateKnowledgebase(context.Background(), id).UpdateKnowledgebase(updateKnowledgebase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.UpdateKnowledgebase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKnowledgebase`: CreateKnowledgebase201Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.UpdateKnowledgebase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKnowledgebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKnowledgebase** | [**UpdateKnowledgebase**](UpdateKnowledgebase.md) |  | 

### Return type

[**CreateKnowledgebase201Response**](CreateKnowledgebase201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFilesToKnowledgebase

> UploadFilesToKnowledgebase200Response UploadFilesToKnowledgebase(ctx, id).Files(files).Execute()

Загрузка файлов в базу знаний



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
    id := int32(1) // int32 | ID базы знаний
    files := []*os.File{"TODO"} // []*os.File | Файлы для загрузки (максимум 10 файлов, 10 МБ каждый)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.UploadFilesToKnowledgebase(context.Background(), id).Files(files).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.UploadFilesToKnowledgebase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFilesToKnowledgebase`: UploadFilesToKnowledgebase200Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.UploadFilesToKnowledgebase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFilesToKnowledgebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **files** | **[]*os.File** | Файлы для загрузки (максимум 10 файлов, 10 МБ каждый) | 

### Return type

[**UploadFilesToKnowledgebase200Response**](UploadFilesToKnowledgebase200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

