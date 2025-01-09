# \S3API

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStorageSubdomainCertificate**](S3API.md#AddStorageSubdomainCertificate) | **Post** /api/v1/storages/certificates/generate | Добавление сертификата для поддомена хранилища
[**AddStorageSubdomains**](S3API.md#AddStorageSubdomains) | **Post** /api/v1/storages/buckets/{bucket_id}/subdomains | Добавление поддоменов для хранилища
[**CopyStorageFile**](S3API.md#CopyStorageFile) | **Post** /api/v1/storages/buckets/{bucket_id}/object-manager/copy | Копирование файла/директории в хранилище
[**CreateFolderInStorage**](S3API.md#CreateFolderInStorage) | **Post** /api/v1/storages/buckets/{bucket_id}/object-manager/mkdir | Создание директории в хранилище
[**CreateStorage**](S3API.md#CreateStorage) | **Post** /api/v1/storages/buckets | Создание хранилища
[**DeleteStorage**](S3API.md#DeleteStorage) | **Delete** /api/v1/storages/buckets/{bucket_id} | Удаление хранилища на аккаунте
[**DeleteStorageFile**](S3API.md#DeleteStorageFile) | **Delete** /api/v1/storages/buckets/{bucket_id}/object-manager/remove | Удаление файла/директории в хранилище
[**DeleteStorageSubdomains**](S3API.md#DeleteStorageSubdomains) | **Delete** /api/v1/storages/buckets/{bucket_id}/subdomains | Удаление поддоменов хранилища
[**GetStorageFilesList**](S3API.md#GetStorageFilesList) | **Get** /api/v1/storages/buckets/{bucket_id}/object-manager/list | Получение списка файлов в хранилище по префиксу
[**GetStorageSubdomains**](S3API.md#GetStorageSubdomains) | **Get** /api/v1/storages/buckets/{bucket_id}/subdomains | Получение списка поддоменов хранилища
[**GetStorageTransferStatus**](S3API.md#GetStorageTransferStatus) | **Get** /api/v1/storages/buckets/{bucket_id}/transfer-status | Получение статуса переноса хранилища от стороннего S3 в Timeweb Cloud
[**GetStorageUsers**](S3API.md#GetStorageUsers) | **Get** /api/v1/storages/users | Получение списка пользователей хранилищ аккаунта
[**GetStorages**](S3API.md#GetStorages) | **Get** /api/v1/storages/buckets | Получение списка хранилищ аккаунта
[**GetStoragesPresets**](S3API.md#GetStoragesPresets) | **Get** /api/v1/presets/storages | Получение списка тарифов для хранилищ
[**RenameStorageFile**](S3API.md#RenameStorageFile) | **Post** /api/v1/storages/buckets/{bucket_id}/object-manager/rename | Переименование файла/директории в хранилище
[**TransferStorage**](S3API.md#TransferStorage) | **Post** /api/v1/storages/transfer | Перенос хранилища от стороннего провайдера S3 в Timeweb Cloud
[**UpdateStorage**](S3API.md#UpdateStorage) | **Patch** /api/v1/storages/buckets/{bucket_id} | Изменение хранилища на аккаунте
[**UpdateStorageUser**](S3API.md#UpdateStorageUser) | **Patch** /api/v1/storages/users/{user_id} | Изменение пароля пользователя-администратора хранилища
[**UploadFileToStorage**](S3API.md#UploadFileToStorage) | **Post** /api/v1/storages/buckets/{bucket_id}/object-manager/upload | Загрузка файлов в хранилище



## AddStorageSubdomainCertificate

> AddStorageSubdomainCertificate(ctx).AddStorageSubdomainCertificateRequest(addStorageSubdomainCertificateRequest).Execute()

Добавление сертификата для поддомена хранилища



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
    addStorageSubdomainCertificateRequest := *openapiclient.NewAddStorageSubdomainCertificateRequest() // AddStorageSubdomainCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3API.AddStorageSubdomainCertificate(context.Background()).AddStorageSubdomainCertificateRequest(addStorageSubdomainCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.AddStorageSubdomainCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddStorageSubdomainCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addStorageSubdomainCertificateRequest** | [**AddStorageSubdomainCertificateRequest**](AddStorageSubdomainCertificateRequest.md) |  | 

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


## AddStorageSubdomains

> AddStorageSubdomains200Response AddStorageSubdomains(ctx, bucketId).AddStorageSubdomainsRequest(addStorageSubdomainsRequest).Execute()

Добавление поддоменов для хранилища



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
    bucketId := int32(1051) // int32 | ID хранилища.
    addStorageSubdomainsRequest := *openapiclient.NewAddStorageSubdomainsRequest([]string{"test.example.com"}) // AddStorageSubdomainsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.AddStorageSubdomains(context.Background(), bucketId).AddStorageSubdomainsRequest(addStorageSubdomainsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.AddStorageSubdomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddStorageSubdomains`: AddStorageSubdomains200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.AddStorageSubdomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddStorageSubdomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addStorageSubdomainsRequest** | [**AddStorageSubdomainsRequest**](AddStorageSubdomainsRequest.md) |  | 

### Return type

[**AddStorageSubdomains200Response**](AddStorageSubdomains200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyStorageFile

> CopyStorageFile(ctx, bucketId).CopyStorageFileRequest(copyStorageFileRequest).Execute()

Копирование файла/директории в хранилище



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
    bucketId := int32(1051) // int32 | ID хранилища.
    copyStorageFileRequest := *openapiclient.NewCopyStorageFileRequest("new_path", []string{"test1/test2"}) // CopyStorageFileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3API.CopyStorageFile(context.Background(), bucketId).CopyStorageFileRequest(copyStorageFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.CopyStorageFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyStorageFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copyStorageFileRequest** | [**CopyStorageFileRequest**](CopyStorageFileRequest.md) |  | 

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


## CreateFolderInStorage

> CreateFolderInStorage(ctx, bucketId).CreateFolderInStorageRequest(createFolderInStorageRequest).Execute()

Создание директории в хранилище



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
    bucketId := int32(1051) // int32 | ID хранилища.
    createFolderInStorageRequest := *openapiclient.NewCreateFolderInStorageRequest("dir") // CreateFolderInStorageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3API.CreateFolderInStorage(context.Background(), bucketId).CreateFolderInStorageRequest(createFolderInStorageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.CreateFolderInStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderInStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFolderInStorageRequest** | [**CreateFolderInStorageRequest**](CreateFolderInStorageRequest.md) |  | 

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


## CreateStorage

> CreateStorage201Response CreateStorage(ctx).CreateStorageRequest(createStorageRequest).Execute()

Создание хранилища



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
    createStorageRequest := *openapiclient.NewCreateStorageRequest("test", "private", float32(1)) // CreateStorageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.CreateStorage(context.Background()).CreateStorageRequest(createStorageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.CreateStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorage`: CreateStorage201Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.CreateStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStorageRequest** | [**CreateStorageRequest**](CreateStorageRequest.md) |  | 

### Return type

[**CreateStorage201Response**](CreateStorage201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorage

> DeleteStorage200Response DeleteStorage(ctx, bucketId).Hash(hash).Code(code).Execute()

Удаление хранилища на аккаунте



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
    bucketId := int32(1051) // int32 | ID хранилища.
    hash := "15095f25-aac3-4d60-a788-96cb5136f186" // string | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. (optional)
    code := "0000" // string | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена `is_able_to_delete` установлен в значение `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.DeleteStorage(context.Background(), bucketId).Hash(hash).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.DeleteStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorage`: DeleteStorage200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.DeleteStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hash** | **string** | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. | 
 **code** | **string** | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена &#x60;is_able_to_delete&#x60; установлен в значение &#x60;true&#x60; | 

### Return type

[**DeleteStorage200Response**](DeleteStorage200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorageFile

> DeleteStorageFile(ctx, bucketId).DeleteStorageFileRequest(deleteStorageFileRequest).IsMultipart(isMultipart).Execute()

Удаление файла/директории в хранилище



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
    bucketId := int32(1051) // int32 | ID хранилища.
    deleteStorageFileRequest := *openapiclient.NewDeleteStorageFileRequest([]string{"test1/test2"}) // DeleteStorageFileRequest | 
    isMultipart := true // bool | Это логическое значение, которое используется для обозначения multipart-загрузки. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3API.DeleteStorageFile(context.Background(), bucketId).DeleteStorageFileRequest(deleteStorageFileRequest).IsMultipart(isMultipart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.DeleteStorageFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteStorageFileRequest** | [**DeleteStorageFileRequest**](DeleteStorageFileRequest.md) |  | 
 **isMultipart** | **bool** | Это логическое значение, которое используется для обозначения multipart-загрузки. | 

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


## DeleteStorageSubdomains

> AddStorageSubdomains200Response DeleteStorageSubdomains(ctx, bucketId).AddStorageSubdomainsRequest(addStorageSubdomainsRequest).Execute()

Удаление поддоменов хранилища



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
    bucketId := int32(1051) // int32 | ID хранилища.
    addStorageSubdomainsRequest := *openapiclient.NewAddStorageSubdomainsRequest([]string{"test.example.com"}) // AddStorageSubdomainsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.DeleteStorageSubdomains(context.Background(), bucketId).AddStorageSubdomainsRequest(addStorageSubdomainsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.DeleteStorageSubdomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorageSubdomains`: AddStorageSubdomains200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.DeleteStorageSubdomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageSubdomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addStorageSubdomainsRequest** | [**AddStorageSubdomainsRequest**](AddStorageSubdomainsRequest.md) |  | 

### Return type

[**AddStorageSubdomains200Response**](AddStorageSubdomains200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageFilesList

> GetStorageFilesList200Response GetStorageFilesList(ctx, bucketId).Prefix(prefix).IsMultipart(isMultipart).Execute()

Получение списка файлов в хранилище по префиксу



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
    bucketId := int32(1051) // int32 | ID хранилища.
    prefix := "example" // string | Префикс для поиска файла. (optional)
    isMultipart := true // bool | Это логическое значение, которое используется для обозначения multipart-загрузки. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.GetStorageFilesList(context.Background(), bucketId).Prefix(prefix).IsMultipart(isMultipart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetStorageFilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageFilesList`: GetStorageFilesList200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.GetStorageFilesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** | Префикс для поиска файла. | 
 **isMultipart** | **bool** | Это логическое значение, которое используется для обозначения multipart-загрузки. | 

### Return type

[**GetStorageFilesList200Response**](GetStorageFilesList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageSubdomains

> GetStorageSubdomains200Response GetStorageSubdomains(ctx, bucketId).Execute()

Получение списка поддоменов хранилища



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
    bucketId := int32(1051) // int32 | ID хранилища.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.GetStorageSubdomains(context.Background(), bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetStorageSubdomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageSubdomains`: GetStorageSubdomains200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.GetStorageSubdomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageSubdomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStorageSubdomains200Response**](GetStorageSubdomains200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageTransferStatus

> GetStorageTransferStatus200Response GetStorageTransferStatus(ctx, bucketId).Execute()

Получение статуса переноса хранилища от стороннего S3 в Timeweb Cloud



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
    bucketId := int32(1051) // int32 | ID хранилища.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.GetStorageTransferStatus(context.Background(), bucketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetStorageTransferStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageTransferStatus`: GetStorageTransferStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.GetStorageTransferStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageTransferStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStorageTransferStatus200Response**](GetStorageTransferStatus200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageUsers

> GetStorageUsers200Response GetStorageUsers(ctx).Execute()

Получение списка пользователей хранилищ аккаунта



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
    resp, r, err := apiClient.S3API.GetStorageUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetStorageUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageUsers`: GetStorageUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.GetStorageUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageUsersRequest struct via the builder pattern


### Return type

[**GetStorageUsers200Response**](GetStorageUsers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorages

> GetProjectStorages200Response GetStorages(ctx).Execute()

Получение списка хранилищ аккаунта



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
    resp, r, err := apiClient.S3API.GetStorages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetStorages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorages`: GetProjectStorages200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.GetStorages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoragesRequest struct via the builder pattern


### Return type

[**GetProjectStorages200Response**](GetProjectStorages200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoragesPresets

> GetStoragesPresets200Response GetStoragesPresets(ctx).Execute()

Получение списка тарифов для хранилищ



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
    resp, r, err := apiClient.S3API.GetStoragesPresets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.GetStoragesPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoragesPresets`: GetStoragesPresets200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.GetStoragesPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoragesPresetsRequest struct via the builder pattern


### Return type

[**GetStoragesPresets200Response**](GetStoragesPresets200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameStorageFile

> RenameStorageFile(ctx, bucketId).RenameStorageFileRequest(renameStorageFileRequest).Execute()

Переименование файла/директории в хранилище



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
    bucketId := int32(1051) // int32 | ID хранилища.
    renameStorageFileRequest := *openapiclient.NewRenameStorageFileRequest("new_filename", "old_filename") // RenameStorageFileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3API.RenameStorageFile(context.Background(), bucketId).RenameStorageFileRequest(renameStorageFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.RenameStorageFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameStorageFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameStorageFileRequest** | [**RenameStorageFileRequest**](RenameStorageFileRequest.md) |  | 

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


## TransferStorage

> TransferStorage(ctx).TransferStorageRequest(transferStorageRequest).Execute()

Перенос хранилища от стороннего провайдера S3 в Timeweb Cloud



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
    transferStorageRequest := *openapiclient.NewTransferStorageRequest("access_key", "secret_key", "ru-1", true, "https://s3.test.ru", "bucket_name", "new_bucket_name") // TransferStorageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3API.TransferStorage(context.Background()).TransferStorageRequest(transferStorageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.TransferStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferStorageRequest** | [**TransferStorageRequest**](TransferStorageRequest.md) |  | 

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


## UpdateStorage

> CreateStorage201Response UpdateStorage(ctx, bucketId).UpdateStorageRequest(updateStorageRequest).Execute()

Изменение хранилища на аккаунте



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
    bucketId := int32(1051) // int32 | ID хранилища.
    updateStorageRequest := *openapiclient.NewUpdateStorageRequest() // UpdateStorageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.UpdateStorage(context.Background(), bucketId).UpdateStorageRequest(updateStorageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.UpdateStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorage`: CreateStorage201Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.UpdateStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStorageRequest** | [**UpdateStorageRequest**](UpdateStorageRequest.md) |  | 

### Return type

[**CreateStorage201Response**](CreateStorage201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStorageUser

> UpdateStorageUser200Response UpdateStorageUser(ctx, userId).UpdateStorageUserRequest(updateStorageUserRequest).Execute()

Изменение пароля пользователя-администратора хранилища



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
    userId := int32(1051) // int32 | ID пользователя хранилища.
    updateStorageUserRequest := *openapiclient.NewUpdateStorageUserRequest("password") // UpdateStorageUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3API.UpdateStorageUser(context.Background(), userId).UpdateStorageUserRequest(updateStorageUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.UpdateStorageUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorageUser`: UpdateStorageUser200Response
    fmt.Fprintf(os.Stdout, "Response from `S3API.UpdateStorageUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | ID пользователя хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorageUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStorageUserRequest** | [**UpdateStorageUserRequest**](UpdateStorageUserRequest.md) |  | 

### Return type

[**UpdateStorageUser200Response**](UpdateStorageUser200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileToStorage

> UploadFileToStorage(ctx, bucketId).Files(files).Path(path).Execute()

Загрузка файлов в хранилище



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
    bucketId := int32(1051) // int32 | ID хранилища.
    files := []*os.File{"TODO"} // []*os.File | 
    path := "test1/tes2" // string | Путь до директории в хранилище (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.S3API.UploadFileToStorage(context.Background(), bucketId).Files(files).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3API.UploadFileToStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketId** | **int32** | ID хранилища. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileToStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **files** | **[]*os.File** |  | 
 **path** | **string** | Путь до директории в хранилище | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

