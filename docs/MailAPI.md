# \MailAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomainMailbox**](MailAPI.md#CreateDomainMailbox) | **Post** /api/v1/mail/domains/{domain} | Создание почтового ящика
[**CreateDomainMailboxV2**](MailAPI.md#CreateDomainMailboxV2) | **Post** /api/v2/mail/domains/{domain} | Создание почтового ящика
[**CreateMultipleDomainMailboxes**](MailAPI.md#CreateMultipleDomainMailboxes) | **Post** /api/v1/mail/domains/{domain}/batch | Множественное создание почтовых ящиков
[**CreateMultipleDomainMailboxesV2**](MailAPI.md#CreateMultipleDomainMailboxesV2) | **Post** /api/v2/mail/domains/{domain}/batch | Множественное создание почтовых ящиков
[**DeleteMailbox**](MailAPI.md#DeleteMailbox) | **Delete** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Удаление почтового ящика
[**GetAllMailboxesV2**](MailAPI.md#GetAllMailboxesV2) | **Get** /api/v2/mail | Получение списка всех почтовых ящиков аккаунта
[**GetDomainMailInfo**](MailAPI.md#GetDomainMailInfo) | **Get** /api/v1/mail/domains/{domain}/info | Получение почтовой информации о домене
[**GetDomainMailboxes**](MailAPI.md#GetDomainMailboxes) | **Get** /api/v1/mail/domains/{domain} | Получение списка почтовых ящиков домена
[**GetMailbox**](MailAPI.md#GetMailbox) | **Get** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Получение почтового ящика
[**GetMailboxV2**](MailAPI.md#GetMailboxV2) | **Get** /api/v2/mail/domains/{domain}/mailboxes/{mailbox} | Получение почтового ящика
[**GetMailboxes**](MailAPI.md#GetMailboxes) | **Get** /api/v1/mail | Получение списка почтовых ящиков аккаунта
[**UpdateDomainMailInfo**](MailAPI.md#UpdateDomainMailInfo) | **Patch** /api/v1/mail/domains/{domain}/info | Изменение почтовой информации о домене
[**UpdateMailbox**](MailAPI.md#UpdateMailbox) | **Patch** /api/v1/mail/domains/{domain}/mailboxes/{mailbox} | Изменение почтового ящика
[**UpdateMailboxV2**](MailAPI.md#UpdateMailboxV2) | **Patch** /api/v2/mail/domains/{domain}/mailboxes/{mailbox} | Изменение почтового ящика



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


## CreateDomainMailboxV2

> CreateDomainMailboxV2201Response CreateDomainMailboxV2(ctx, domain).CreateDomainMailboxV2Request(createDomainMailboxV2Request).Execute()

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
    createDomainMailboxV2Request := *openapiclient.NewCreateDomainMailboxV2Request("mailbox", "zfAsl-") // CreateDomainMailboxV2Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.CreateDomainMailboxV2(context.Background(), domain).CreateDomainMailboxV2Request(createDomainMailboxV2Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.CreateDomainMailboxV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainMailboxV2`: CreateDomainMailboxV2201Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.CreateDomainMailboxV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainMailboxV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDomainMailboxV2Request** | [**CreateDomainMailboxV2Request**](CreateDomainMailboxV2Request.md) |  | 

### Return type

[**CreateDomainMailboxV2201Response**](CreateDomainMailboxV2201Response.md)

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


## CreateMultipleDomainMailboxesV2

> CreateMultipleDomainMailboxesV2201Response CreateMultipleDomainMailboxesV2(ctx, domain).CreateMultipleDomainMailboxesV2RequestInner(createMultipleDomainMailboxesV2RequestInner).Execute()

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
    createMultipleDomainMailboxesV2RequestInner := []openapiclient.CreateMultipleDomainMailboxesV2RequestInner{*openapiclient.NewCreateMultipleDomainMailboxesV2RequestInner("t1", "@,9FT.Pv5}Y,hl")} // []CreateMultipleDomainMailboxesV2RequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.CreateMultipleDomainMailboxesV2(context.Background(), domain).CreateMultipleDomainMailboxesV2RequestInner(createMultipleDomainMailboxesV2RequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.CreateMultipleDomainMailboxesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultipleDomainMailboxesV2`: CreateMultipleDomainMailboxesV2201Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.CreateMultipleDomainMailboxesV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultipleDomainMailboxesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMultipleDomainMailboxesV2RequestInner** | [**[]CreateMultipleDomainMailboxesV2RequestInner**](CreateMultipleDomainMailboxesV2RequestInner.md) |  | 

### Return type

[**CreateMultipleDomainMailboxesV2201Response**](CreateMultipleDomainMailboxesV2201Response.md)

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


## GetAllMailboxesV2

> GetAllMailboxesV2200Response GetAllMailboxesV2(ctx).Limit(limit).Offset(offset).Search(search).Execute()

Получение списка всех почтовых ящиков аккаунта



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
    resp, r, err := apiClient.MailAPI.GetAllMailboxesV2(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.GetAllMailboxesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllMailboxesV2`: GetAllMailboxesV2200Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.GetAllMailboxesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMailboxesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]
 **search** | **string** | Поиск почтового ящика по названию | 

### Return type

[**GetAllMailboxesV2200Response**](GetAllMailboxesV2200Response.md)

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


## GetMailboxV2

> CreateDomainMailboxV2201Response GetMailboxV2(ctx, domain, mailbox).Execute()

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
    resp, r, err := apiClient.MailAPI.GetMailboxV2(context.Background(), domain, mailbox).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.GetMailboxV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMailboxV2`: CreateDomainMailboxV2201Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.GetMailboxV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 
**mailbox** | **string** | Название почтового ящика | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailboxV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateDomainMailboxV2201Response**](CreateDomainMailboxV2201Response.md)

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


## UpdateMailboxV2

> UpdateMailboxV2200Response UpdateMailboxV2(ctx, domain, mailbox).UpdateMailboxV2(updateMailboxV2).Execute()

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
    updateMailboxV2 := *openapiclient.NewUpdateMailboxV2() // UpdateMailboxV2 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailAPI.UpdateMailboxV2(context.Background(), domain, mailbox).UpdateMailboxV2(updateMailboxV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.UpdateMailboxV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMailboxV2`: UpdateMailboxV2200Response
    fmt.Fprintf(os.Stdout, "Response from `MailAPI.UpdateMailboxV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Полное имя домена | 
**mailbox** | **string** | Название почтового ящика | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMailboxV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMailboxV2** | [**UpdateMailboxV2**](UpdateMailboxV2.md) |  | 

### Return type

[**UpdateMailboxV2200Response**](UpdateMailboxV2200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

