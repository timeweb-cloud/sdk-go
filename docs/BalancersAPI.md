# \BalancersAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIPsToBalancer**](BalancersAPI.md#AddIPsToBalancer) | **Post** /api/v1/balancers/{balancer_id}/ips | Добавление IP-адресов к балансировщику
[**CreateBalancer**](BalancersAPI.md#CreateBalancer) | **Post** /api/v1/balancers | Создание бaлансировщика
[**CreateBalancerRule**](BalancersAPI.md#CreateBalancerRule) | **Post** /api/v1/balancers/{balancer_id}/rules | Создание правила для балансировщика
[**DeleteBalancer**](BalancersAPI.md#DeleteBalancer) | **Delete** /api/v1/balancers/{balancer_id} | Удаление балансировщика
[**DeleteBalancerRule**](BalancersAPI.md#DeleteBalancerRule) | **Delete** /api/v1/balancers/{balancer_id}/rules/{rule_id} | Удаление правила для балансировщика
[**DeleteIPsFromBalancer**](BalancersAPI.md#DeleteIPsFromBalancer) | **Delete** /api/v1/balancers/{balancer_id}/ips | Удаление IP-адресов из балансировщика
[**GetBalancer**](BalancersAPI.md#GetBalancer) | **Get** /api/v1/balancers/{balancer_id} | Получение бaлансировщика
[**GetBalancerIPs**](BalancersAPI.md#GetBalancerIPs) | **Get** /api/v1/balancers/{balancer_id}/ips | Получение списка IP-адресов балансировщика
[**GetBalancerRules**](BalancersAPI.md#GetBalancerRules) | **Get** /api/v1/balancers/{balancer_id}/rules | Получение правил балансировщика
[**GetBalancers**](BalancersAPI.md#GetBalancers) | **Get** /api/v1/balancers | Получение списка всех бaлансировщиков
[**GetBalancersPresets**](BalancersAPI.md#GetBalancersPresets) | **Get** /api/v1/presets/balancers | Получение списка тарифов для балансировщика
[**UpdateBalancer**](BalancersAPI.md#UpdateBalancer) | **Patch** /api/v1/balancers/{balancer_id} | Обновление балансировщика
[**UpdateBalancerRule**](BalancersAPI.md#UpdateBalancerRule) | **Patch** /api/v1/balancers/{balancer_id}/rules/{rule_id} | Обновление правила для балансировщика



## AddIPsToBalancer

> AddIPsToBalancer(ctx, balancerId).AddIPsToBalancerRequest(addIPsToBalancerRequest).Execute()

Добавление IP-адресов к балансировщику



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
    balancerId := int32(56) // int32 | ID балансировщика
    addIPsToBalancerRequest := *openapiclient.NewAddIPsToBalancerRequest([]string{"192.168.0.0"}) // AddIPsToBalancerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BalancersAPI.AddIPsToBalancer(context.Background(), balancerId).AddIPsToBalancerRequest(addIPsToBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.AddIPsToBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddIPsToBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addIPsToBalancerRequest** | [**AddIPsToBalancerRequest**](AddIPsToBalancerRequest.md) |  | 

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


## CreateBalancer

> CreateBalancer200Response CreateBalancer(ctx).CreateBalancer(createBalancer).Execute()

Создание бaлансировщика



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
    createBalancer := *openapiclient.NewCreateBalancer("Cute Quail", "roundrobin", true, true, true, true, "https", float32(80), "/", float32(10), float32(5), float32(3), float32(2), float32(5)) // CreateBalancer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.CreateBalancer(context.Background()).CreateBalancer(createBalancer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.CreateBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBalancer`: CreateBalancer200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.CreateBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBalancer** | [**CreateBalancer**](CreateBalancer.md) |  | 

### Return type

[**CreateBalancer200Response**](CreateBalancer200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBalancerRule

> CreateBalancerRule200Response CreateBalancerRule(ctx, balancerId).CreateRule(createRule).Execute()

Создание правила для балансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика
    createRule := *openapiclient.NewCreateRule("https", float32(80), "https", float32(80)) // CreateRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.CreateBalancerRule(context.Background(), balancerId).CreateRule(createRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.CreateBalancerRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBalancerRule`: CreateBalancerRule200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.CreateBalancerRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBalancerRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRule** | [**CreateRule**](CreateRule.md) |  | 

### Return type

[**CreateBalancerRule200Response**](CreateBalancerRule200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBalancer

> DeleteBalancer200Response DeleteBalancer(ctx, balancerId).Hash(hash).Code(code).Execute()

Удаление балансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика
    hash := "15095f25-aac3-4d60-a788-96cb5136f186" // string | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. (optional)
    code := "0000" // string | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена `is_able_to_delete` установлен в значение `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.DeleteBalancer(context.Background(), balancerId).Hash(hash).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.DeleteBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBalancer`: DeleteBalancer200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.DeleteBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hash** | **string** | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. | 
 **code** | **string** | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена &#x60;is_able_to_delete&#x60; установлен в значение &#x60;true&#x60; | 

### Return type

[**DeleteBalancer200Response**](DeleteBalancer200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBalancerRule

> DeleteBalancerRule(ctx, balancerId, ruleId).Execute()

Удаление правила для балансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика
    ruleId := int32(56) // int32 | ID правила для балансировщика

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BalancersAPI.DeleteBalancerRule(context.Background(), balancerId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.DeleteBalancerRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 
**ruleId** | **int32** | ID правила для балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBalancerRuleRequest struct via the builder pattern


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


## DeleteIPsFromBalancer

> DeleteIPsFromBalancer(ctx, balancerId).AddIPsToBalancerRequest(addIPsToBalancerRequest).Execute()

Удаление IP-адресов из балансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика
    addIPsToBalancerRequest := *openapiclient.NewAddIPsToBalancerRequest([]string{"192.168.0.0"}) // AddIPsToBalancerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BalancersAPI.DeleteIPsFromBalancer(context.Background(), balancerId).AddIPsToBalancerRequest(addIPsToBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.DeleteIPsFromBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPsFromBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addIPsToBalancerRequest** | [**AddIPsToBalancerRequest**](AddIPsToBalancerRequest.md) |  | 

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


## GetBalancer

> CreateBalancer200Response GetBalancer(ctx, balancerId).Execute()

Получение бaлансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.GetBalancer(context.Background(), balancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.GetBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalancer`: CreateBalancer200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.GetBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateBalancer200Response**](CreateBalancer200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalancerIPs

> GetBalancerIPs200Response GetBalancerIPs(ctx, balancerId).Execute()

Получение списка IP-адресов балансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.GetBalancerIPs(context.Background(), balancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.GetBalancerIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalancerIPs`: GetBalancerIPs200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.GetBalancerIPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancerIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBalancerIPs200Response**](GetBalancerIPs200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalancerRules

> GetBalancerRules200Response GetBalancerRules(ctx, balancerId).Execute()

Получение правил балансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.GetBalancerRules(context.Background(), balancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.GetBalancerRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalancerRules`: GetBalancerRules200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.GetBalancerRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancerRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBalancerRules200Response**](GetBalancerRules200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalancers

> GetBalancers200Response GetBalancers(ctx).Limit(limit).Offset(offset).Execute()

Получение списка всех бaлансировщиков



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.GetBalancers(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.GetBalancers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalancers`: GetBalancers200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.GetBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**GetBalancers200Response**](GetBalancers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalancersPresets

> GetBalancersPresets200Response GetBalancersPresets(ctx).Execute()

Получение списка тарифов для балансировщика



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
    resp, r, err := apiClient.BalancersAPI.GetBalancersPresets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.GetBalancersPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalancersPresets`: GetBalancersPresets200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.GetBalancersPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancersPresetsRequest struct via the builder pattern


### Return type

[**GetBalancersPresets200Response**](GetBalancersPresets200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBalancer

> CreateBalancer200Response UpdateBalancer(ctx, balancerId).UpdateBalancer(updateBalancer).Execute()

Обновление балансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика
    updateBalancer := *openapiclient.NewUpdateBalancer() // UpdateBalancer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.UpdateBalancer(context.Background(), balancerId).UpdateBalancer(updateBalancer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.UpdateBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBalancer`: CreateBalancer200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.UpdateBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBalancer** | [**UpdateBalancer**](UpdateBalancer.md) |  | 

### Return type

[**CreateBalancer200Response**](CreateBalancer200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBalancerRule

> CreateBalancerRule200Response UpdateBalancerRule(ctx, balancerId, ruleId).UpdateRule(updateRule).Execute()

Обновление правила для балансировщика



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
    balancerId := int32(56) // int32 | ID балансировщика
    ruleId := int32(56) // int32 | ID правила для балансировщика
    updateRule := *openapiclient.NewUpdateRule("https", float32(80), "https", float32(80)) // UpdateRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancersAPI.UpdateBalancerRule(context.Background(), balancerId, ruleId).UpdateRule(updateRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancersAPI.UpdateBalancerRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBalancerRule`: CreateBalancerRule200Response
    fmt.Fprintf(os.Stdout, "Response from `BalancersAPI.UpdateBalancerRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**balancerId** | **int32** | ID балансировщика | 
**ruleId** | **int32** | ID правила для балансировщика | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBalancerRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRule** | [**UpdateRule**](UpdateRule.md) |  | 

### Return type

[**CreateBalancerRule200Response**](CreateBalancerRule200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

