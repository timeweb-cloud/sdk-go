# \KnowledgeBasesAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdditionalTokenPackageToKnowledgebase**](KnowledgeBasesAPI.md#AddAdditionalTokenPackageToKnowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/add-additional-token-package | Покупка дополнительных пакетов токенов
[**CreateKnowledgebase**](KnowledgeBasesAPI.md#CreateKnowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases | Создание базы знаний
[**CreateKnowledgebaseDocument**](KnowledgeBasesAPI.md#CreateKnowledgebaseDocument) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/documents | Добавление документа по ссылке
[**DeleteDocument**](KnowledgeBasesAPI.md#DeleteDocument) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id} | Удаление документа из базы знаний
[**DeleteKnowledgebase**](KnowledgeBasesAPI.md#DeleteKnowledgebase) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id} | Удаление базы знаний
[**DownloadDocument**](KnowledgeBasesAPI.md#DownloadDocument) | **Get** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id}/download | Скачивание файла документа
[**GetKnowledgebase**](KnowledgeBasesAPI.md#GetKnowledgebase) | **Get** /api/v1/cloud-ai/knowledge-bases/{id} | Получение базы знаний со списком документов
[**GetKnowledgebaseDocumentsV2**](KnowledgeBasesAPI.md#GetKnowledgebaseDocumentsV2) | **Get** /api/v2/cloud-ai/knowledge-bases/{id}/documents | Получение списка документов базы знаний
[**GetKnowledgebaseV3**](KnowledgeBasesAPI.md#GetKnowledgebaseV3) | **Get** /api/v3/cloud-ai/knowledge-bases/{id} | Получение базы знаний
[**GetKnowledgebases**](KnowledgeBasesAPI.md#GetKnowledgebases) | **Get** /api/v1/cloud-ai/knowledge-bases | Получение списка баз знаний (v1)
[**GetKnowledgebasesV2**](KnowledgeBasesAPI.md#GetKnowledgebasesV2) | **Get** /api/v2/cloud-ai/knowledge-bases | Получение списка баз знаний (v2)
[**GetKnowledgebasesV3**](KnowledgeBasesAPI.md#GetKnowledgebasesV3) | **Get** /api/v3/cloud-ai/knowledge-bases | Получение списка баз знаний
[**LinkKnowledgebaseToAgent**](KnowledgeBasesAPI.md#LinkKnowledgebaseToAgent) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/link/{agent_id} | Привязка базы знаний к агенту
[**ReindexDocument**](KnowledgeBasesAPI.md#ReindexDocument) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id}/reindex | Переиндексация документа
[**UnlinkKnowledgebaseFromAgent**](KnowledgeBasesAPI.md#UnlinkKnowledgebaseFromAgent) | **Delete** /api/v1/cloud-ai/knowledge-bases/{id}/link/{agent_id} | Отвязка базы знаний от агента
[**UpdateKnowledgebase**](KnowledgeBasesAPI.md#UpdateKnowledgebase) | **Patch** /api/v1/cloud-ai/knowledge-bases/{id} | Изменение базы знаний
[**UpdateKnowledgebaseAgentLink**](KnowledgeBasesAPI.md#UpdateKnowledgebaseAgentLink) | **Patch** /api/v1/cloud-ai/knowledge-bases/{id}/link/{agent_id} | Изменение режима поиска по базе знаний для агента
[**UpdateKnowledgebaseDocument**](KnowledgeBasesAPI.md#UpdateKnowledgebaseDocument) | **Patch** /api/v1/cloud-ai/knowledge-bases/{id}/documents/{document_id} | Изменение документа
[**UploadFilesToKnowledgebase**](KnowledgeBasesAPI.md#UploadFilesToKnowledgebase) | **Post** /api/v1/cloud-ai/knowledge-bases/{id}/upload | Загрузка файлов в базу знаний



## AddAdditionalTokenPackageToKnowledgebase

> AddAdditionalTokenPackageToKnowledgebase(ctx, id).KnowledgebaseAddTokenPackage(knowledgebaseAddTokenPackage).Execute()

Покупка дополнительных пакетов токенов



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
    knowledgebaseAddTokenPackage := *openapiclient.NewKnowledgebaseAddTokenPackage() // KnowledgebaseAddTokenPackage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.AddAdditionalTokenPackageToKnowledgebase(context.Background(), id).KnowledgebaseAddTokenPackage(knowledgebaseAddTokenPackage).Execute()
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

 **knowledgebaseAddTokenPackage** | [**KnowledgebaseAddTokenPackage**](KnowledgebaseAddTokenPackage.md) |  | 

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
    createKnowledgebase := *openapiclient.NewCreateKnowledgebase("Документация продукта", int32(2449), "network-1234567890abcdef", int32(111)) // CreateKnowledgebase | 

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


## CreateKnowledgebaseDocument

> CreateKnowledgebaseDocument201Response CreateKnowledgebaseDocument(ctx, id).CreateDocument(createDocument).Execute()

Добавление документа по ссылке



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
    createDocument := *openapiclient.NewCreateDocument("FAQ", "link", "https://example.com/docs/faq.md") // CreateDocument | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.CreateKnowledgebaseDocument(context.Background(), id).CreateDocument(createDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.CreateKnowledgebaseDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKnowledgebaseDocument`: CreateKnowledgebaseDocument201Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.CreateKnowledgebaseDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKnowledgebaseDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDocument** | [**CreateDocument**](CreateDocument.md) |  | 

### Return type

[**CreateKnowledgebaseDocument201Response**](CreateKnowledgebaseDocument201Response.md)

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
    documentId := int32(286071) // int32 | ID документа

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

Скачивание файла документа



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
    documentId := int32(286071) // int32 | ID документа

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

Получение базы знаний со списком документов



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

> GetKnowledgebaseDocumentsV2200Response GetKnowledgebaseDocumentsV2(ctx, id).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

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
    limit := int32(20) // int32 | Количество документов на странице. Значения больше 100 уменьшаются до 100 (optional) (default to 10)
    offset := int32(0) // int32 | Количество документов, которые нужно пропустить (optional) (default to 0)
    sortBy := "name" // string | Поле сортировки. `indexing_timestamp` — время последнего изменения документа (optional) (default to "indexing_timestamp")
    sortOrder := "ASC" // string | Порядок сортировки (optional) (default to "DESC")
    search := "faq" // string | Фильтр по подстроке в названии документа (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.GetKnowledgebaseDocumentsV2(context.Background(), id).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
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

 **limit** | **int32** | Количество документов на странице. Значения больше 100 уменьшаются до 100 | [default to 10]
 **offset** | **int32** | Количество документов, которые нужно пропустить | [default to 0]
 **sortBy** | **string** | Поле сортировки. &#x60;indexing_timestamp&#x60; — время последнего изменения документа | [default to &quot;indexing_timestamp&quot;]
 **sortOrder** | **string** | Порядок сортировки | [default to &quot;DESC&quot;]
 **search** | **string** | Фильтр по подстроке в названии документа | 

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


## GetKnowledgebaseV3

> GetKnowledgebaseV3200Response GetKnowledgebaseV3(ctx, id).Execute()

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
    resp, r, err := apiClient.KnowledgeBasesAPI.GetKnowledgebaseV3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.GetKnowledgebaseV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKnowledgebaseV3`: GetKnowledgebaseV3200Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.GetKnowledgebaseV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebaseV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetKnowledgebaseV3200Response**](GetKnowledgebaseV3200Response.md)

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

Получение списка баз знаний (v1)



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


## GetKnowledgebasesV3

> GetKnowledgebasesV3200Response GetKnowledgebasesV3(ctx).Execute()

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
    resp, r, err := apiClient.KnowledgeBasesAPI.GetKnowledgebasesV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.GetKnowledgebasesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKnowledgebasesV3`: GetKnowledgebasesV3200Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.GetKnowledgebasesV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgebasesV3Request struct via the builder pattern


### Return type

[**GetKnowledgebasesV3200Response**](GetKnowledgebasesV3200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkKnowledgebaseToAgent

> LinkKnowledgebaseToAgent(ctx, id, agentId).LinkKnowledgebase(linkKnowledgebase).Execute()

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
    agentId := int32(101) // int32 | ID AI агента
    linkKnowledgebase := *openapiclient.NewLinkKnowledgebase() // LinkKnowledgebase |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.LinkKnowledgebaseToAgent(context.Background(), id, agentId).LinkKnowledgebase(linkKnowledgebase).Execute()
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
**agentId** | **int32** | ID AI агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkKnowledgebaseToAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkKnowledgebase** | [**LinkKnowledgebase**](LinkKnowledgebase.md) |  | 

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
    documentId := int32(286071) // int32 | ID документа

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
    agentId := int32(101) // int32 | ID AI агента

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
**agentId** | **int32** | ID AI агента | 

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

Изменение базы знаний



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


## UpdateKnowledgebaseAgentLink

> UpdateKnowledgebaseAgentLink(ctx, id, agentId).LinkKnowledgebase(linkKnowledgebase).Execute()

Изменение режима поиска по базе знаний для агента



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
    agentId := int32(101) // int32 | ID AI агента
    linkKnowledgebase := *openapiclient.NewLinkKnowledgebase() // LinkKnowledgebase |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KnowledgeBasesAPI.UpdateKnowledgebaseAgentLink(context.Background(), id, agentId).LinkKnowledgebase(linkKnowledgebase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.UpdateKnowledgebaseAgentLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 
**agentId** | **int32** | ID AI агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKnowledgebaseAgentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkKnowledgebase** | [**LinkKnowledgebase**](LinkKnowledgebase.md) |  | 

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


## UpdateKnowledgebaseDocument

> CreateKnowledgebaseDocument201Response UpdateKnowledgebaseDocument(ctx, id, documentId).UpdateDocument(updateDocument).Execute()

Изменение документа



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
    documentId := int32(286071) // int32 | ID документа
    updateDocument := *openapiclient.NewUpdateDocument() // UpdateDocument | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.UpdateKnowledgebaseDocument(context.Background(), id, documentId).UpdateDocument(updateDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.UpdateKnowledgebaseDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKnowledgebaseDocument`: CreateKnowledgebaseDocument201Response
    fmt.Fprintf(os.Stdout, "Response from `KnowledgeBasesAPI.UpdateKnowledgebaseDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID базы знаний | 
**documentId** | **int32** | ID документа | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKnowledgebaseDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDocument** | [**UpdateDocument**](UpdateDocument.md) |  | 

### Return type

[**CreateKnowledgebaseDocument201Response**](CreateKnowledgebaseDocument201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFilesToKnowledgebase

> UploadResult UploadFilesToKnowledgebase(ctx, id).Files(files).Execute()

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
    files := []*os.File{"TODO"} // []*os.File | Файлы для загрузки

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KnowledgeBasesAPI.UploadFilesToKnowledgebase(context.Background(), id).Files(files).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBasesAPI.UploadFilesToKnowledgebase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFilesToKnowledgebase`: UploadResult
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

 **files** | **[]*os.File** | Файлы для загрузки | 

### Return type

[**UploadResult**](UploadResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

