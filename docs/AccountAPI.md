# \AccountAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCountriesToAllowedList**](AccountAPI.md#AddCountriesToAllowedList) | **Post** /api/v1/auth/access/countries | Добавление стран в список разрешенных
[**AddIPsToAllowedList**](AccountAPI.md#AddIPsToAllowedList) | **Post** /api/v1/auth/access/ips | Добавление IP-адресов в список разрешенных
[**DeleteCountriesFromAllowedList**](AccountAPI.md#DeleteCountriesFromAllowedList) | **Delete** /api/v1/auth/access/countries | Удаление стран из списка разрешенных
[**DeleteIPsFromAllowedList**](AccountAPI.md#DeleteIPsFromAllowedList) | **Delete** /api/v1/auth/access/ips | Удаление IP-адресов из списка разрешенных
[**GetAccountStatus**](AccountAPI.md#GetAccountStatus) | **Get** /api/v1/account/status | Получение статуса аккаунта
[**GetAuthAccessSettings**](AccountAPI.md#GetAuthAccessSettings) | **Get** /api/v1/auth/access | Получить информацию о ограничениях авторизации пользователя
[**GetCountries**](AccountAPI.md#GetCountries) | **Get** /api/v1/auth/access/countries | Получение списка стран
[**GetFinances**](AccountAPI.md#GetFinances) | **Get** /api/v1/account/finances | Получение платежной информации
[**GetNotificationSettings**](AccountAPI.md#GetNotificationSettings) | **Get** /api/v1/account/notification-settings | Получение настроек уведомлений аккаунта
[**UpdateAuthRestrictionsByCountries**](AccountAPI.md#UpdateAuthRestrictionsByCountries) | **Post** /api/v1/auth/access/countries/enabled | Включение/отключение ограничений по стране
[**UpdateAuthRestrictionsByIP**](AccountAPI.md#UpdateAuthRestrictionsByIP) | **Post** /api/v1/auth/access/ips/enabled | Включение/отключение ограничений по IP-адресу
[**UpdateNotificationSettings**](AccountAPI.md#UpdateNotificationSettings) | **Patch** /api/v1/account/notification-settings | Изменение настроек уведомлений аккаунта



## AddCountriesToAllowedList

> AddCountriesToAllowedList201Response AddCountriesToAllowedList(ctx).AddCountriesToAllowedListRequest(addCountriesToAllowedListRequest).Execute()

Добавление стран в список разрешенных



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
    addCountriesToAllowedListRequest := *openapiclient.NewAddCountriesToAllowedListRequest([]string{"RU"}) // AddCountriesToAllowedListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AddCountriesToAllowedList(context.Background()).AddCountriesToAllowedListRequest(addCountriesToAllowedListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AddCountriesToAllowedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCountriesToAllowedList`: AddCountriesToAllowedList201Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AddCountriesToAllowedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCountriesToAllowedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCountriesToAllowedListRequest** | [**AddCountriesToAllowedListRequest**](AddCountriesToAllowedListRequest.md) |  | 

### Return type

[**AddCountriesToAllowedList201Response**](AddCountriesToAllowedList201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddIPsToAllowedList

> AddIPsToAllowedList201Response AddIPsToAllowedList(ctx).AddIPsToAllowedListRequest(addIPsToAllowedListRequest).Execute()

Добавление IP-адресов в список разрешенных



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
    addIPsToAllowedListRequest := *openapiclient.NewAddIPsToAllowedListRequest([]string{"123.231.125.128"}) // AddIPsToAllowedListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AddIPsToAllowedList(context.Background()).AddIPsToAllowedListRequest(addIPsToAllowedListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AddIPsToAllowedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIPsToAllowedList`: AddIPsToAllowedList201Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AddIPsToAllowedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIPsToAllowedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addIPsToAllowedListRequest** | [**AddIPsToAllowedListRequest**](AddIPsToAllowedListRequest.md) |  | 

### Return type

[**AddIPsToAllowedList201Response**](AddIPsToAllowedList201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCountriesFromAllowedList

> DeleteCountriesFromAllowedList200Response DeleteCountriesFromAllowedList(ctx).DeleteCountriesFromAllowedListRequest(deleteCountriesFromAllowedListRequest).Execute()

Удаление стран из списка разрешенных



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
    deleteCountriesFromAllowedListRequest := *openapiclient.NewDeleteCountriesFromAllowedListRequest([]string{"RU"}) // DeleteCountriesFromAllowedListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.DeleteCountriesFromAllowedList(context.Background()).DeleteCountriesFromAllowedListRequest(deleteCountriesFromAllowedListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteCountriesFromAllowedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCountriesFromAllowedList`: DeleteCountriesFromAllowedList200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteCountriesFromAllowedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCountriesFromAllowedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteCountriesFromAllowedListRequest** | [**DeleteCountriesFromAllowedListRequest**](DeleteCountriesFromAllowedListRequest.md) |  | 

### Return type

[**DeleteCountriesFromAllowedList200Response**](DeleteCountriesFromAllowedList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIPsFromAllowedList

> DeleteIPsFromAllowedList200Response DeleteIPsFromAllowedList(ctx).DeleteIPsFromAllowedListRequest(deleteIPsFromAllowedListRequest).Execute()

Удаление IP-адресов из списка разрешенных



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
    deleteIPsFromAllowedListRequest := *openapiclient.NewDeleteIPsFromAllowedListRequest([]string{"123.231.125.128"}) // DeleteIPsFromAllowedListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.DeleteIPsFromAllowedList(context.Background()).DeleteIPsFromAllowedListRequest(deleteIPsFromAllowedListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteIPsFromAllowedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIPsFromAllowedList`: DeleteIPsFromAllowedList200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteIPsFromAllowedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPsFromAllowedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteIPsFromAllowedListRequest** | [**DeleteIPsFromAllowedListRequest**](DeleteIPsFromAllowedListRequest.md) |  | 

### Return type

[**DeleteIPsFromAllowedList200Response**](DeleteIPsFromAllowedList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountStatus

> GetAccountStatus200Response GetAccountStatus(ctx).Execute()

Получение статуса аккаунта



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
    resp, r, err := apiClient.AccountAPI.GetAccountStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountStatus`: GetAccountStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountStatusRequest struct via the builder pattern


### Return type

[**GetAccountStatus200Response**](GetAccountStatus200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthAccessSettings

> GetAuthAccessSettings200Response GetAuthAccessSettings(ctx).Execute()

Получить информацию о ограничениях авторизации пользователя



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
    resp, r, err := apiClient.AccountAPI.GetAuthAccessSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAuthAccessSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthAccessSettings`: GetAuthAccessSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAuthAccessSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthAccessSettingsRequest struct via the builder pattern


### Return type

[**GetAuthAccessSettings200Response**](GetAuthAccessSettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> GetCountries200Response GetCountries(ctx).Execute()

Получение списка стран



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
    resp, r, err := apiClient.AccountAPI.GetCountries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountries`: GetCountries200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


### Return type

[**GetCountries200Response**](GetCountries200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinances

> GetFinances200Response GetFinances(ctx).Execute()

Получение платежной информации



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
    resp, r, err := apiClient.AccountAPI.GetFinances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetFinances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinances`: GetFinances200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetFinances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancesRequest struct via the builder pattern


### Return type

[**GetFinances200Response**](GetFinances200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSettings

> GetNotificationSettings200Response GetNotificationSettings(ctx).Execute()

Получение настроек уведомлений аккаунта



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
    resp, r, err := apiClient.AccountAPI.GetNotificationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetNotificationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationSettings`: GetNotificationSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetNotificationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSettingsRequest struct via the builder pattern


### Return type

[**GetNotificationSettings200Response**](GetNotificationSettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthRestrictionsByCountries

> UpdateAuthRestrictionsByCountries(ctx).UpdateAuthRestrictionsByCountriesRequest(updateAuthRestrictionsByCountriesRequest).Execute()

Включение/отключение ограничений по стране



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
    updateAuthRestrictionsByCountriesRequest := *openapiclient.NewUpdateAuthRestrictionsByCountriesRequest(true) // UpdateAuthRestrictionsByCountriesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountAPI.UpdateAuthRestrictionsByCountries(context.Background()).UpdateAuthRestrictionsByCountriesRequest(updateAuthRestrictionsByCountriesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAuthRestrictionsByCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthRestrictionsByCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAuthRestrictionsByCountriesRequest** | [**UpdateAuthRestrictionsByCountriesRequest**](UpdateAuthRestrictionsByCountriesRequest.md) |  | 

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


## UpdateAuthRestrictionsByIP

> UpdateAuthRestrictionsByIP(ctx).UpdateAuthRestrictionsByCountriesRequest(updateAuthRestrictionsByCountriesRequest).Execute()

Включение/отключение ограничений по IP-адресу



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
    updateAuthRestrictionsByCountriesRequest := *openapiclient.NewUpdateAuthRestrictionsByCountriesRequest(true) // UpdateAuthRestrictionsByCountriesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountAPI.UpdateAuthRestrictionsByIP(context.Background()).UpdateAuthRestrictionsByCountriesRequest(updateAuthRestrictionsByCountriesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAuthRestrictionsByIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthRestrictionsByIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAuthRestrictionsByCountriesRequest** | [**UpdateAuthRestrictionsByCountriesRequest**](UpdateAuthRestrictionsByCountriesRequest.md) |  | 

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


## UpdateNotificationSettings

> GetNotificationSettings200Response UpdateNotificationSettings(ctx).UpdateNotificationSettingsRequest(updateNotificationSettingsRequest).Execute()

Изменение настроек уведомлений аккаунта



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
    updateNotificationSettingsRequest := *openapiclient.NewUpdateNotificationSettingsRequest() // UpdateNotificationSettingsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.UpdateNotificationSettings(context.Background()).UpdateNotificationSettingsRequest(updateNotificationSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateNotificationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationSettings`: GetNotificationSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateNotificationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateNotificationSettingsRequest** | [**UpdateNotificationSettingsRequest**](UpdateNotificationSettingsRequest.md) |  | 

### Return type

[**GetNotificationSettings200Response**](GetNotificationSettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

