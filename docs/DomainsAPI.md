# \DomainsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDomain**](DomainsAPI.md#AddDomain) | **Post** /api/v1/add-domain/{fqdn} | Добавление домена на аккаунт
[**AddSubdomain**](DomainsAPI.md#AddSubdomain) | **Post** /api/v1/domains/{fqdn}/subdomains/{subdomain_fqdn} | Добавление поддомена
[**CheckDomain**](DomainsAPI.md#CheckDomain) | **Get** /api/v1/check-domain/{fqdn} | Проверить, доступен ли домен для регистрации
[**CreateDomainDNSRecord**](DomainsAPI.md#CreateDomainDNSRecord) | **Post** /api/v1/domains/{fqdn}/dns-records | Добавить информацию о DNS-записи для домена или поддомена
[**CreateDomainRequest**](DomainsAPI.md#CreateDomainRequest) | **Post** /api/v1/domains-requests | Создание заявки на регистрацию/продление/трансфер домена
[**DeleteDomain**](DomainsAPI.md#DeleteDomain) | **Delete** /api/v1/domains/{fqdn} | Удаление домена
[**DeleteDomainDNSRecord**](DomainsAPI.md#DeleteDomainDNSRecord) | **Delete** /api/v1/domains/{fqdn}/dns-records/{record_id} | Удалить информацию о DNS-записи для домена или поддомена
[**DeleteSubdomain**](DomainsAPI.md#DeleteSubdomain) | **Delete** /api/v1/domains/{fqdn}/subdomains/{subdomain_fqdn} | Удаление поддомена
[**GetDomain**](DomainsAPI.md#GetDomain) | **Get** /api/v1/domains/{fqdn} | Получение информации о домене
[**GetDomainDNSRecords**](DomainsAPI.md#GetDomainDNSRecords) | **Get** /api/v1/domains/{fqdn}/dns-records | Получить информацию обо всех пользовательских DNS-записях домена или поддомена
[**GetDomainDefaultDNSRecords**](DomainsAPI.md#GetDomainDefaultDNSRecords) | **Get** /api/v1/domains/{fqdn}/default-dns-records | Получить информацию обо всех DNS-записях по умолчанию домена или поддомена
[**GetDomainNameServers**](DomainsAPI.md#GetDomainNameServers) | **Get** /api/v1/domains/{fqdn}/name-servers | Получение списка name-серверов домена
[**GetDomainRequest**](DomainsAPI.md#GetDomainRequest) | **Get** /api/v1/domains-requests/{request_id} | Получение заявки на регистрацию/продление/трансфер домена
[**GetDomainRequests**](DomainsAPI.md#GetDomainRequests) | **Get** /api/v1/domains-requests | Получение списка заявок на регистрацию/продление/трансфер домена
[**GetDomains**](DomainsAPI.md#GetDomains) | **Get** /api/v1/domains | Получение списка всех доменов
[**GetTLD**](DomainsAPI.md#GetTLD) | **Get** /api/v1/tlds/{tld_id} | Получить информацию о доменной зоне по ID
[**GetTLDs**](DomainsAPI.md#GetTLDs) | **Get** /api/v1/tlds | Получить информацию о доменных зонах
[**UpdateDomainAutoProlongation**](DomainsAPI.md#UpdateDomainAutoProlongation) | **Patch** /api/v1/domains/{fqdn} | Включение/выключение автопродления домена
[**UpdateDomainDNSRecord**](DomainsAPI.md#UpdateDomainDNSRecord) | **Patch** /api/v1/domains/{fqdn}/dns-records/{record_id} | Обновить информацию о DNS-записи домена или поддомена
[**UpdateDomainNameServers**](DomainsAPI.md#UpdateDomainNameServers) | **Put** /api/v1/domains/{fqdn}/name-servers | Изменение name-серверов домена
[**UpdateDomainRequest**](DomainsAPI.md#UpdateDomainRequest) | **Patch** /api/v1/domains-requests/{request_id} | Оплата/обновление заявки на регистрацию/продление/трансфер домена



## AddDomain

> AddDomain(ctx, fqdn).Execute()

Добавление домена на аккаунт



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
    fqdn := "somedomain.ru" // string | Полное имя домена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.AddDomain(context.Background(), fqdn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.AddDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDomainRequest struct via the builder pattern


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


## AddSubdomain

> AddSubdomain201Response AddSubdomain(ctx, fqdn, subdomainFqdn).Execute()

Добавление поддомена



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
    fqdn := "somedomain.ru" // string | Полное имя домена.
    subdomainFqdn := "sub.somedomain.ru" // string | Полное имя поддомена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.AddSubdomain(context.Background(), fqdn, subdomainFqdn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.AddSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSubdomain`: AddSubdomain201Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.AddSubdomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 
**subdomainFqdn** | **string** | Полное имя поддомена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSubdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddSubdomain201Response**](AddSubdomain201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckDomain

> CheckDomain200Response CheckDomain(ctx, fqdn).Execute()

Проверить, доступен ли домен для регистрации



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
    fqdn := "somedomain.ru" // string | Полное имя домена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.CheckDomain(context.Background(), fqdn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.CheckDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckDomain`: CheckDomain200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.CheckDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckDomain200Response**](CheckDomain200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDomainDNSRecord

> CreateDomainDNSRecord201Response CreateDomainDNSRecord(ctx, fqdn).CreateDns(createDns).Execute()

Добавить информацию о DNS-записи для домена или поддомена



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
    fqdn := "somedomain.ru" // string | Полное имя домена или поддомена.
    createDns := *openapiclient.NewCreateDns("A", "192.168.111.123") // CreateDns | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.CreateDomainDNSRecord(context.Background(), fqdn).CreateDns(createDns).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.CreateDomainDNSRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainDNSRecord`: CreateDomainDNSRecord201Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.CreateDomainDNSRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена или поддомена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainDNSRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDns** | [**CreateDns**](CreateDns.md) |  | 

### Return type

[**CreateDomainDNSRecord201Response**](CreateDomainDNSRecord201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDomainRequest

> CreateDomainRequest201Response CreateDomainRequest(ctx).DomainRegister(domainRegister).Execute()

Создание заявки на регистрацию/продление/трансфер домена



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
    domainRegister := *openapiclient.NewDomainRegister("register", "somedomain.ru", float32(123)) // DomainRegister | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.CreateDomainRequest(context.Background()).DomainRegister(domainRegister).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.CreateDomainRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainRequest`: CreateDomainRequest201Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.CreateDomainRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainRegister** | [**DomainRegister**](DomainRegister.md) |  | 

### Return type

[**CreateDomainRequest201Response**](CreateDomainRequest201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomain

> DeleteDomain(ctx, fqdn).Execute()

Удаление домена



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
    fqdn := "somedomain.ru" // string | Полное имя домена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.DeleteDomain(context.Background(), fqdn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


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


## DeleteDomainDNSRecord

> DeleteDomainDNSRecord(ctx, fqdn, recordId).Execute()

Удалить информацию о DNS-записи для домена или поддомена



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
    fqdn := "somedomain.ru" // string | Полное имя домена или поддомена.
    recordId := int32(123) // int32 | ID DNS-записи домена или поддомена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.DeleteDomainDNSRecord(context.Background(), fqdn, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteDomainDNSRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена или поддомена. | 
**recordId** | **int32** | ID DNS-записи домена или поддомена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainDNSRecordRequest struct via the builder pattern


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


## DeleteSubdomain

> DeleteSubdomain(ctx, fqdn, subdomainFqdn).Execute()

Удаление поддомена



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
    fqdn := "somedomain.ru" // string | Полное имя домена.
    subdomainFqdn := "sub.somedomain.ru" // string | Полное имя поддомена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.DeleteSubdomain(context.Background(), fqdn, subdomainFqdn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 
**subdomainFqdn** | **string** | Полное имя поддомена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubdomainRequest struct via the builder pattern


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


## GetDomain

> GetDomain200Response GetDomain(ctx, fqdn).Execute()

Получение информации о домене



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
    fqdn := "somedomain.ru" // string | Полное имя домена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetDomain(context.Background(), fqdn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomain`: GetDomain200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDomain200Response**](GetDomain200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainDNSRecords

> GetDomainDNSRecords200Response GetDomainDNSRecords(ctx, fqdn).Limit(limit).Offset(offset).Execute()

Получить информацию обо всех пользовательских DNS-записях домена или поддомена



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
    fqdn := "somedomain.ru" // string | Полное имя домена или поддомена.
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetDomainDNSRecords(context.Background(), fqdn).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomainDNSRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainDNSRecords`: GetDomainDNSRecords200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomainDNSRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена или поддомена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainDNSRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**GetDomainDNSRecords200Response**](GetDomainDNSRecords200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainDefaultDNSRecords

> GetDomainDNSRecords200Response GetDomainDefaultDNSRecords(ctx, fqdn).Limit(limit).Offset(offset).Execute()

Получить информацию обо всех DNS-записях по умолчанию домена или поддомена



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
    fqdn := "somedomain.ru" // string | Полное имя домена или поддомена.
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetDomainDefaultDNSRecords(context.Background(), fqdn).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomainDefaultDNSRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainDefaultDNSRecords`: GetDomainDNSRecords200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomainDefaultDNSRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена или поддомена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainDefaultDNSRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**GetDomainDNSRecords200Response**](GetDomainDNSRecords200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainNameServers

> GetDomainNameServers200Response GetDomainNameServers(ctx, fqdn).Execute()

Получение списка name-серверов домена



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
    fqdn := "somedomain.ru" // string | Полное имя домена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetDomainNameServers(context.Background(), fqdn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomainNameServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainNameServers`: GetDomainNameServers200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomainNameServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainNameServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDomainNameServers200Response**](GetDomainNameServers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainRequest

> CreateDomainRequest201Response GetDomainRequest(ctx, requestId).Execute()

Получение заявки на регистрацию/продление/трансфер домена



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
    requestId := int32(123) // int32 | ID заявки на регистрацию/продление/трансфер домена.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetDomainRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomainRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainRequest`: CreateDomainRequest201Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomainRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **int32** | ID заявки на регистрацию/продление/трансфер домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDomainRequest201Response**](CreateDomainRequest201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainRequests

> GetDomainRequests200Response GetDomainRequests(ctx).PersonId(personId).Execute()

Получение списка заявок на регистрацию/продление/трансфер домена



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
    personId := int32(123) // int32 | ID администратора, на которого зарегистрирован домен. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetDomainRequests(context.Background()).PersonId(personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomainRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainRequests`: GetDomainRequests200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomainRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personId** | **int32** | ID администратора, на которого зарегистрирован домен. | 

### Return type

[**GetDomainRequests200Response**](GetDomainRequests200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomains

> GetDomains200Response GetDomains(ctx).Limit(limit).Offset(offset).IdnName(idnName).LinkedIp(linkedIp).Order(order).Sort(sort).Execute()

Получение списка всех доменов



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
    idnName := "xn--e1afmkfd.xn--p1ai" // string | Интернационализированное доменное имя. (optional)
    linkedIp := "192.168.1.1" // string | Привязанный к домену IP-адрес. (optional)
    order := "asc" // string | Порядок доменов. (optional)
    sort := "idn_name" // string | Сортировка доменов. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetDomains(context.Background()).Limit(limit).Offset(offset).IdnName(idnName).LinkedIp(linkedIp).Order(order).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomains`: GetDomains200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]
 **idnName** | **string** | Интернационализированное доменное имя. | 
 **linkedIp** | **string** | Привязанный к домену IP-адрес. | 
 **order** | **string** | Порядок доменов. | 
 **sort** | **string** | Сортировка доменов. | 

### Return type

[**GetDomains200Response**](GetDomains200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTLD

> GetTLD200Response GetTLD(ctx, tldId).Execute()

Получить информацию о доменной зоне по ID



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
    tldId := int32(123) // int32 | ID доменной зоны.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetTLD(context.Background(), tldId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetTLD``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTLD`: GetTLD200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetTLD`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tldId** | **int32** | ID доменной зоны. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTLDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTLD200Response**](GetTLD200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTLDs

> GetTLDs200Response GetTLDs(ctx).IsPublished(isPublished).IsRegistered(isRegistered).Execute()

Получить информацию о доменных зонах



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
    isPublished := true // bool | Это логическое значение, которое показывает, опубликована ли доменная зона. (optional)
    isRegistered := true // bool | Это логическое значение, которое показывает, зарегистрирована ли доменная зона. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.GetTLDs(context.Background()).IsPublished(isPublished).IsRegistered(isRegistered).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetTLDs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTLDs`: GetTLDs200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetTLDs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTLDsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isPublished** | **bool** | Это логическое значение, которое показывает, опубликована ли доменная зона. | 
 **isRegistered** | **bool** | Это логическое значение, которое показывает, зарегистрирована ли доменная зона. | 

### Return type

[**GetTLDs200Response**](GetTLDs200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainAutoProlongation

> UpdateDomainAutoProlongation200Response UpdateDomainAutoProlongation(ctx, fqdn).UpdateDomain(updateDomain).Execute()

Включение/выключение автопродления домена



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
    fqdn := "somedomain.ru" // string | Полное имя домена.
    updateDomain := *openapiclient.NewUpdateDomain() // UpdateDomain | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.UpdateDomainAutoProlongation(context.Background(), fqdn).UpdateDomain(updateDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.UpdateDomainAutoProlongation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomainAutoProlongation`: UpdateDomainAutoProlongation200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.UpdateDomainAutoProlongation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainAutoProlongationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDomain** | [**UpdateDomain**](UpdateDomain.md) |  | 

### Return type

[**UpdateDomainAutoProlongation200Response**](UpdateDomainAutoProlongation200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainDNSRecord

> CreateDomainDNSRecord201Response UpdateDomainDNSRecord(ctx, fqdn, recordId).CreateDns(createDns).Execute()

Обновить информацию о DNS-записи домена или поддомена



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
    fqdn := "somedomain.ru" // string | Полное имя домена или поддомена.
    recordId := int32(123) // int32 | ID DNS-записи домена или поддомена.
    createDns := *openapiclient.NewCreateDns("A", "192.168.111.123") // CreateDns | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.UpdateDomainDNSRecord(context.Background(), fqdn, recordId).CreateDns(createDns).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.UpdateDomainDNSRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomainDNSRecord`: CreateDomainDNSRecord201Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.UpdateDomainDNSRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена или поддомена. | 
**recordId** | **int32** | ID DNS-записи домена или поддомена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainDNSRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDns** | [**CreateDns**](CreateDns.md) |  | 

### Return type

[**CreateDomainDNSRecord201Response**](CreateDomainDNSRecord201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainNameServers

> GetDomainNameServers200Response UpdateDomainNameServers(ctx, fqdn).UpdateDomainNameServers(updateDomainNameServers).Execute()

Изменение name-серверов домена



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
    fqdn := "somedomain.ru" // string | Полное имя домена.
    updateDomainNameServers := *openapiclient.NewUpdateDomainNameServers([]openapiclient.UpdateDomainNameServersNameServersInner{*openapiclient.NewUpdateDomainNameServersNameServersInner("ns1.timeweb.ru")}) // UpdateDomainNameServers | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.UpdateDomainNameServers(context.Background(), fqdn).UpdateDomainNameServers(updateDomainNameServers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.UpdateDomainNameServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomainNameServers`: GetDomainNameServers200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.UpdateDomainNameServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fqdn** | **string** | Полное имя домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainNameServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDomainNameServers** | [**UpdateDomainNameServers**](UpdateDomainNameServers.md) |  | 

### Return type

[**GetDomainNameServers200Response**](GetDomainNameServers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainRequest

> CreateDomainRequest201Response UpdateDomainRequest(ctx, requestId).Use(use).Execute()

Оплата/обновление заявки на регистрацию/продление/трансфер домена



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
    requestId := int32(123) // int32 | ID заявки на регистрацию/продление/трансфер домена.
    use := *openapiclient.NewUse("use") // Use | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.UpdateDomainRequest(context.Background(), requestId).Use(use).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.UpdateDomainRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomainRequest`: CreateDomainRequest201Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.UpdateDomainRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **int32** | ID заявки на регистрацию/продление/трансфер домена. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **use** | [**Use**](Use.md) |  | 

### Return type

[**CreateDomainRequest201Response**](CreateDomainRequest201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

