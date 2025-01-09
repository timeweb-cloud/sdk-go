# \MailAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainMailbox**](MailAPI.md#CreateDomainMailbox) | **Post** /api/v1/mail/domains/{domain} | Создание почтового ящика
[**CreateMultipleDomainMailboxes**](MailAPI.md#CreateMultipleDomainMailboxes) | **Post** /api/v1/mail/domains/{domain}/batch | Множественное создание почтовых ящиков
[**DeleteMailbox**](MailAPI.md#DeleteMailbox) | **Delete** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Удаление почтового ящика
[**GetDomainMailInfo**](MailAPI.md#GetDomainMailInfo) | **Get** /api/v1/mail/domains/{domain}/info | Получение почтовой информации о домене
[**GetDomainMailboxes**](MailAPI.md#GetDomainMailboxes) | **Get** /api/v1/mail/domains/{domain} | Получение списка почтовых ящиков домена
[**GetMailQuota**](MailAPI.md#GetMailQuota) | **Get** /api/v1/mail/quota | Получение квоты почты аккаунта
[**GetMailbox**](MailAPI.md#GetMailbox) | **Get** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Получение почтового ящика
[**GetMailboxes**](MailAPI.md#GetMailboxes) | **Get** /api/v1/mail | Получение списка почтовых ящиков аккаунта
[**UpdateDomainMailInfo**](MailAPI.md#UpdateDomainMailInfo) | **Patch** /api/v1/mail/domains/{domain}/info | Изменение почтовой информации о домене
[**UpdateMailQuota**](MailAPI.md#UpdateMailQuota) | **Patch** /api/v1/mail/quota | Изменение квоты почты аккаунта
[**UpdateMailbox**](MailAPI.md#UpdateMailbox) | **Patch** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Изменение почтового ящика



## CreateDomainMailbox

> CreateDomainMailbox201Response CreateDomainMailbox(ctx, domain).CreateDomainMailboxRequest(createDomainMailboxRequest).Execute()

Создание почтового ящика



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
    domain := "somedomain.ru" // string | Полное имя домена
    createDomainMailboxRequest := *openapiclient.NewCreateDomainMailboxRequest("mailbox", "zfAsl-") // CreateDomainMailboxRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.CreateDomainMailbox(context.Background(), domain).CreateDomainMailboxRequest(createDomainMailboxRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.CreateDomainMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainMailbox`: CreateDomainMailbox201Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.CreateDomainMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDomainMailboxRequest** | [**CreateDomainMailboxRequest**](CreateDomainMailboxRequest.md) |  | 

### Return type

[**CreateDomainMailbox201Response**](CreateDomainMailbox201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMultipleDomainMailboxes

> CreateMultipleDomainMailboxes201Response CreateMultipleDomainMailboxes(ctx, domain).CreateMultipleDomainMailboxesRequest(createMultipleDomainMailboxesRequest).Execute()

Множественное создание почтовых ящиков



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
    domain := "somedomain.ru" // string | Полное имя домена
    createMultipleDomainMailboxesRequest := *openapiclient.NewCreateMultipleDomainMailboxesRequest([]openapiclient.CreateMultipleDomainMailboxesRequestMailboxesInner{*openapiclient.NewCreateMultipleDomainMailboxesRequestMailboxesInner("mailbox", "zfAsl-")}) // CreateMultipleDomainMailboxesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.CreateMultipleDomainMailboxes(context.Background(), domain).CreateMultipleDomainMailboxesRequest(createMultipleDomainMailboxesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.CreateMultipleDomainMailboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultipleDomainMailboxes`: CreateMultipleDomainMailboxes201Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.CreateMultipleDomainMailboxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultipleDomainMailboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMultipleDomainMailboxesRequest** | [**CreateMultipleDomainMailboxesRequest**](CreateMultipleDomainMailboxesRequest.md) |  | 

### Return type

[**CreateMultipleDomainMailboxes201Response**](CreateMultipleDomainMailboxes201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMailbox

> DeleteMailbox(ctx, domain, mailbox).Execute()

Удаление почтового ящика



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
    domain := "somedomain.ru" // string | Полное имя домена
    mailbox := "mailbox" // string | Название почтового ящика

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MailAPI.DeleteMailbox(context.Background(), domain, mailbox).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.DeleteMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 
**mailbox** | **string** | Название почтового ящика | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailboxRequest struct via the builder pattern


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


## GetDomainMailInfo

> GetDomainMailInfo200Response GetDomainMailInfo(ctx, domain).Execute()

Получение почтовой информации о домене



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
    domain := "somedomain.ru" // string | Полное имя домена

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.GetDomainMailInfo(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.GetDomainMailInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainMailInfo`: GetDomainMailInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.GetDomainMailInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainMailInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDomainMailInfo200Response**](GetDomainMailInfo200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainMailboxes

> GetMailboxes200Response GetDomainMailboxes(ctx, domain).Limit(limit).Offset(offset).Search(search).Execute()

Получение списка почтовых ящиков домена



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
    domain := "somedomain.ru" // string | Полное имя домена
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)
    search := "search_example" // string | Поиск почтового ящика по названию (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.GetDomainMailboxes(context.Background(), domain).Limit(limit).Offset(offset).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.GetDomainMailboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainMailboxes`: GetMailboxes200Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.GetDomainMailboxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainMailboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]
 **search** | **string** | Поиск почтового ящика по названию | 

### Return type

[**GetMailboxes200Response**](GetMailboxes200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailQuota

> GetMailQuota200Response GetMailQuota(ctx).Execute()

Получение квоты почты аккаунта



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
    resp, r, err := apiClient.MailAPI.GetMailQuota(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.GetMailQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMailQuota`: GetMailQuota200Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.GetMailQuota`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailQuotaRequest struct via the builder pattern


### Return type

[**GetMailQuota200Response**](GetMailQuota200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailbox

> CreateDomainMailbox201Response GetMailbox(ctx, domain, mailbox).Execute()

Получение почтового ящика



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
    domain := "somedomain.ru" // string | Полное имя домена
    mailbox := "mailbox" // string | Название почтового ящика

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.GetMailbox(context.Background(), domain, mailbox).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.GetMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMailbox`: CreateDomainMailbox201Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.GetMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 
**mailbox** | **string** | Название почтового ящика | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateDomainMailbox201Response**](CreateDomainMailbox201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailboxes

> GetMailboxes200Response GetMailboxes(ctx).Limit(limit).Offset(offset).Search(search).Execute()

Получение списка почтовых ящиков аккаунта



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
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)
    search := "search_example" // string | Поиск почтового ящика по названию (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.GetMailboxes(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.GetMailboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMailboxes`: GetMailboxes200Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.GetMailboxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMailboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]
 **search** | **string** | Поиск почтового ящика по названию | 

### Return type

[**GetMailboxes200Response**](GetMailboxes200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainMailInfo

> GetDomainMailInfo200Response UpdateDomainMailInfo(ctx, domain).UpdateDomainMailInfoRequest(updateDomainMailInfoRequest).Execute()

Изменение почтовой информации о домене



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
    domain := "somedomain.ru" // string | Полное имя домена
    updateDomainMailInfoRequest := *openapiclient.NewUpdateDomainMailInfoRequest("test@timeweb.ru") // UpdateDomainMailInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.UpdateDomainMailInfo(context.Background(), domain).UpdateDomainMailInfoRequest(updateDomainMailInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.UpdateDomainMailInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomainMailInfo`: GetDomainMailInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.UpdateDomainMailInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainMailInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDomainMailInfoRequest** | [**UpdateDomainMailInfoRequest**](UpdateDomainMailInfoRequest.md) |  | 

### Return type

[**GetDomainMailInfo200Response**](GetDomainMailInfo200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMailQuota

> GetMailQuota200Response UpdateMailQuota(ctx).UpdateMailQuotaRequest(updateMailQuotaRequest).Execute()

Изменение квоты почты аккаунта



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
    updateMailQuotaRequest := *openapiclient.NewUpdateMailQuotaRequest(float32(6144)) // UpdateMailQuotaRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.UpdateMailQuota(context.Background()).UpdateMailQuotaRequest(updateMailQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.UpdateMailQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMailQuota`: GetMailQuota200Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.UpdateMailQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMailQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMailQuotaRequest** | [**UpdateMailQuotaRequest**](UpdateMailQuotaRequest.md) |  | 

### Return type

[**GetMailQuota200Response**](GetMailQuota200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMailbox

> CreateDomainMailbox201Response UpdateMailbox(ctx, domain, mailbox).UpdateMailbox(updateMailbox).Execute()

Изменение почтового ящика



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
    domain := "somedomain.ru" // string | Полное имя домена
    mailbox := "mailbox" // string | Название почтового ящика
    updateMailbox := *openapiclient.NewUpdateMailbox() // UpdateMailbox | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.UpdateMailbox(context.Background(), domain, mailbox).UpdateMailbox(updateMailbox).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.UpdateMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMailbox`: CreateDomainMailbox201Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.UpdateMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 
**mailbox** | **string** | Название почтового ящика | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMailbox** | [**UpdateMailbox**](UpdateMailbox.md) |  | 

### Return type

[**CreateDomainMailbox201Response**](CreateDomainMailbox201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

