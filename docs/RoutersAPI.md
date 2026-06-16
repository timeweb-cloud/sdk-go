# \RoutersAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetworks**](RoutersAPI.md#AddNetworks) | **Post** /api/v1/routers/{router_id}/networks | Подключение сетей к роутеру
[**CreateRouter**](RoutersAPI.md#CreateRouter) | **Post** /api/v1/routers | Создание роутера
[**DeleteDnat**](RoutersAPI.md#DeleteDnat) | **Delete** /api/v1/routers/{router_id}/dnat-rules/{dnat_id} | Удаление правила проброса портов
[**DeleteRouter**](RoutersAPI.md#DeleteRouter) | **Delete** /api/v1/routers/{router_id} | Удаление роутера
[**DeleteRouterNat**](RoutersAPI.md#DeleteRouterNat) | **Delete** /api/v1/routers/{router_id}/networks/{network_name}/nat | Выключение NAT для сети
[**DeleteRouterNetwork**](RoutersAPI.md#DeleteRouterNetwork) | **Delete** /api/v1/routers/{router_id}/networks/{network_name} | Удаление сети из роутера
[**DeleteStaticRoute**](RoutersAPI.md#DeleteStaticRoute) | **Delete** /api/v1/routers/{router_id}/static-routes/{static_route_id} | Удаление статического маршрута
[**GetAvailableStaticRoutes**](RoutersAPI.md#GetAvailableStaticRoutes) | **Get** /api/v1/routers/{router_id}/static-routes/available | Получение доступных подсетей для статических маршрутов
[**GetDnat**](RoutersAPI.md#GetDnat) | **Get** /api/v1/routers/{router_id}/dnat-rules | Получение списка правил проброса портов
[**GetDnatRule**](RoutersAPI.md#GetDnatRule) | **Get** /api/v1/routers/{router_id}/dnat-rules/{dnat_id} | Получение правила проброса портов
[**GetNetworks**](RoutersAPI.md#GetNetworks) | **Get** /api/v1/routers/{router_id}/networks | Получение списка сетей роутера
[**GetRouter**](RoutersAPI.md#GetRouter) | **Get** /api/v1/routers/{router_id} | Получение информации о роутере
[**GetRouterAvailableNetworks**](RoutersAPI.md#GetRouterAvailableNetworks) | **Get** /api/v1/routers/networks/available | Получение списка доступных сетей
[**GetRouterPresets**](RoutersAPI.md#GetRouterPresets) | **Get** /api/v1/presets/routers | Получение списка тарифов роутеров
[**GetRouterStatistics**](RoutersAPI.md#GetRouterStatistics) | **Get** /api/v1/routers/{router_id}/statistics/{time_from}/{period}/{keys} | Получение статистики роутера
[**GetRouters**](RoutersAPI.md#GetRouters) | **Get** /api/v1/routers | Получение списка роутеров
[**GetStaticRoutes**](RoutersAPI.md#GetStaticRoutes) | **Get** /api/v1/routers/{router_id}/static-routes | Получение списка статических маршрутов
[**PatchNetwork**](RoutersAPI.md#PatchNetwork) | **Patch** /api/v1/routers/{router_id}/networks/{network_name} | Обновление информации о сети
[**PatchNetworks**](RoutersAPI.md#PatchNetworks) | **Patch** /api/v1/routers/{router_id}/networks | Обновление сетей роутера
[**PostDnat**](RoutersAPI.md#PostDnat) | **Post** /api/v1/routers/{router_id}/dnat-rules | Добавление правила проброса портов
[**PostStaticRoute**](RoutersAPI.md#PostStaticRoute) | **Post** /api/v1/routers/{router_id}/static-routes | Добавление статического маршрута
[**UpdateRouter**](RoutersAPI.md#UpdateRouter) | **Patch** /api/v1/routers/{router_id} | Обновление информации о роутере
[**UpdateRouterNat**](RoutersAPI.md#UpdateRouterNat) | **Patch** /api/v1/routers/{router_id}/networks/{network_name}/nat | Включение NAT для сети



## AddNetworks

> NetworksResponse AddNetworks(ctx, routerId).NetworkIn(networkIn).Execute()

Подключение сетей к роутеру



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
    routerId := "routerId_example" // string | ID роутера
    networkIn := *openapiclient.NewNetworkIn([]openapiclient.RouterInNetworksInner{*openapiclient.NewRouterInNetworksInner("network-8ea2cef4765c422c9b2a22d6ec283bf3")}) // NetworkIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.AddNetworks(context.Background(), routerId).NetworkIn(networkIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.AddNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNetworks`: NetworksResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.AddNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIn** | [**NetworkIn**](NetworkIn.md) |  | 

### Return type

[**NetworksResponse**](NetworksResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRouter

> RouterResponse CreateRouter(ctx).RouterIn(routerIn).Execute()

Создание роутера



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
    routerIn := *openapiclient.NewRouterIn("Name_example", int32(123), []openapiclient.RouterInNetworksInner{*openapiclient.NewRouterInNetworksInner("network-8ea2cef4765c422c9b2a22d6ec283bf3")}) // RouterIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.CreateRouter(context.Background()).RouterIn(routerIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.CreateRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRouter`: RouterResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.CreateRouter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerIn** | [**RouterIn**](RouterIn.md) |  | 

### Return type

[**RouterResponse**](RouterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDnat

> DeleteDnat(ctx, routerId, dnatId).Execute()

Удаление правила проброса портов



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
    routerId := "routerId_example" // string | ID роутера
    dnatId := "dnatId_example" // string | ID правила

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoutersAPI.DeleteDnat(context.Background(), routerId, dnatId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.DeleteDnat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 
**dnatId** | **string** | ID правила | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnatRequest struct via the builder pattern


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


## DeleteRouter

> DeleteRouter(ctx, routerId).Execute()

Удаление роутера



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
    routerId := "routerId_example" // string | ID роутера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoutersAPI.DeleteRouter(context.Background(), routerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.DeleteRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouterRequest struct via the builder pattern


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


## DeleteRouterNat

> RouterResponse DeleteRouterNat(ctx, routerId, networkName).Execute()

Выключение NAT для сети



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
    routerId := "routerId_example" // string | ID роутера
    networkName := "networkName_example" // string | Имя сети

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.DeleteRouterNat(context.Background(), routerId, networkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.DeleteRouterNat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRouterNat`: RouterResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.DeleteRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 
**networkName** | **string** | Имя сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RouterResponse**](RouterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouterNetwork

> DeleteRouterNetwork(ctx, routerId, networkName).Execute()

Удаление сети из роутера



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
    routerId := "routerId_example" // string | ID роутера
    networkName := "networkName_example" // string | Имя сети

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoutersAPI.DeleteRouterNetwork(context.Background(), routerId, networkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.DeleteRouterNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 
**networkName** | **string** | Имя сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouterNetworkRequest struct via the builder pattern


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


## DeleteStaticRoute

> DeleteStaticRoute(ctx, routerId, staticRouteId).Execute()

Удаление статического маршрута



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
    routerId := "routerId_example" // string | ID роутера
    staticRouteId := "staticRouteId_example" // string | ID статического маршрута

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoutersAPI.DeleteStaticRoute(context.Background(), routerId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.DeleteStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 
**staticRouteId** | **string** | ID статического маршрута | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStaticRouteRequest struct via the builder pattern


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


## GetAvailableStaticRoutes

> AvailableStaticRoutesResponse GetAvailableStaticRoutes(ctx, routerId).Execute()

Получение доступных подсетей для статических маршрутов



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
    routerId := "routerId_example" // string | ID роутера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.GetAvailableStaticRoutes(context.Background(), routerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetAvailableStaticRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableStaticRoutes`: AvailableStaticRoutesResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetAvailableStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvailableStaticRoutesResponse**](AvailableStaticRoutesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnat

> DnatRulesResponse GetDnat(ctx, routerId).Execute()

Получение списка правил проброса портов



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
    routerId := "routerId_example" // string | ID роутера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.GetDnat(context.Background(), routerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetDnat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDnat`: DnatRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetDnat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnatRulesResponse**](DnatRulesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnatRule

> DnatRuleResponse GetDnatRule(ctx, routerId, dnatId).Execute()

Получение правила проброса портов



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
    routerId := "routerId_example" // string | ID роутера
    dnatId := "dnatId_example" // string | ID правила

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.GetDnatRule(context.Background(), routerId, dnatId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetDnatRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDnatRule`: DnatRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetDnatRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 
**dnatId** | **string** | ID правила | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnatRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DnatRuleResponse**](DnatRuleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworks

> NetworksResponse GetNetworks(ctx, routerId).Execute()

Получение списка сетей роутера



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
    routerId := "routerId_example" // string | ID роутера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.GetNetworks(context.Background(), routerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworks`: NetworksResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworksResponse**](NetworksResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouter

> RouterResponse GetRouter(ctx, routerId).Execute()

Получение информации о роутере



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
    routerId := "routerId_example" // string | ID роутера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.GetRouter(context.Background(), routerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouter`: RouterResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouterResponse**](RouterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouterAvailableNetworks

> AvailableNetworksResponse GetRouterAvailableNetworks(ctx).Execute()

Получение списка доступных сетей



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
    resp, r, err := apiClient.RoutersAPI.GetRouterAvailableNetworks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetRouterAvailableNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouterAvailableNetworks`: AvailableNetworksResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetRouterAvailableNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouterAvailableNetworksRequest struct via the builder pattern


### Return type

[**AvailableNetworksResponse**](AvailableNetworksResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouterPresets

> RouterPresetsResponse GetRouterPresets(ctx).Execute()

Получение списка тарифов роутеров



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
    resp, r, err := apiClient.RoutersAPI.GetRouterPresets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetRouterPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouterPresets`: RouterPresetsResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetRouterPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouterPresetsRequest struct via the builder pattern


### Return type

[**RouterPresetsResponse**](RouterPresetsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouterStatistics

> RouterStatisticsResponse GetRouterStatistics(ctx, routerId, timeFrom, period, keys).NodeId(nodeId).Execute()

Получение статистики роутера



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
    routerId := "routerId_example" // string | ID роутера
    timeFrom := "timeFrom_example" // string | Начало периода
    period := "period_example" // string | Период агрегации
    keys := "keys_example" // string | Ключи метрик
    nodeId := "nodeId_example" // string | ID ноды (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.GetRouterStatistics(context.Background(), routerId, timeFrom, period, keys).NodeId(nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetRouterStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouterStatistics`: RouterStatisticsResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetRouterStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 
**timeFrom** | **string** | Начало периода | 
**period** | **string** | Период агрегации | 
**keys** | **string** | Ключи метрик | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouterStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **nodeId** | **string** | ID ноды | 

### Return type

[**RouterStatisticsResponse**](RouterStatisticsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouters

> RoutersResponse GetRouters(ctx).Execute()

Получение списка роутеров



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
    resp, r, err := apiClient.RoutersAPI.GetRouters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetRouters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRouters`: RoutersResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetRouters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutersRequest struct via the builder pattern


### Return type

[**RoutersResponse**](RoutersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaticRoutes

> StaticRoutesResponse GetStaticRoutes(ctx, routerId).Execute()

Получение списка статических маршрутов



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
    routerId := "routerId_example" // string | ID роутера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.GetStaticRoutes(context.Background(), routerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.GetStaticRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStaticRoutes`: StaticRoutesResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.GetStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StaticRoutesResponse**](StaticRoutesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNetwork

> NetworkResponse PatchNetwork(ctx, routerId, networkName).NetworkEdit(networkEdit).Execute()

Обновление информации о сети



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
    routerId := "routerId_example" // string | ID роутера
    networkName := "networkName_example" // string | Имя сети
    networkEdit := *openapiclient.NewNetworkEdit(true) // NetworkEdit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.PatchNetwork(context.Background(), routerId, networkName).NetworkEdit(networkEdit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.PatchNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNetwork`: NetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.PatchNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 
**networkName** | **string** | Имя сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkEdit** | [**NetworkEdit**](NetworkEdit.md) |  | 

### Return type

[**NetworkResponse**](NetworkResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNetworks

> NetworksResponse PatchNetworks(ctx, routerId).NetworkIn(networkIn).Execute()

Обновление сетей роутера



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
    routerId := "routerId_example" // string | ID роутера
    networkIn := *openapiclient.NewNetworkIn([]openapiclient.RouterInNetworksInner{*openapiclient.NewRouterInNetworksInner("network-8ea2cef4765c422c9b2a22d6ec283bf3")}) // NetworkIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.PatchNetworks(context.Background(), routerId).NetworkIn(networkIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.PatchNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNetworks`: NetworksResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.PatchNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIn** | [**NetworkIn**](NetworkIn.md) |  | 

### Return type

[**NetworksResponse**](NetworksResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDnat

> DnatRuleResponse PostDnat(ctx, routerId).DnatIn(dnatIn).Execute()

Добавление правила проброса портов



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
    routerId := "routerId_example" // string | ID роутера
    dnatIn := *openapiclient.NewDnatIn("Protocol_example", *openapiclient.NewDnatInLocal("Ip_example"), *openapiclient.NewDnatInPublic("Ip_example")) // DnatIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.PostDnat(context.Background(), routerId).DnatIn(dnatIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.PostDnat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDnat`: DnatRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.PostDnat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDnatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnatIn** | [**DnatIn**](DnatIn.md) |  | 

### Return type

[**DnatRuleResponse**](DnatRuleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStaticRoute

> StaticRouteResponse PostStaticRoute(ctx, routerId).StaticRouteIn(staticRouteIn).Execute()

Добавление статического маршрута



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
    routerId := "routerId_example" // string | ID роутера
    staticRouteIn := *openapiclient.NewStaticRouteIn("10.0.0.0/24", "192.168.1.4") // StaticRouteIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.PostStaticRoute(context.Background(), routerId).StaticRouteIn(staticRouteIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.PostStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostStaticRoute`: StaticRouteResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.PostStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **staticRouteIn** | [**StaticRouteIn**](StaticRouteIn.md) |  | 

### Return type

[**StaticRouteResponse**](StaticRouteResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouter

> RouterResponse UpdateRouter(ctx, routerId).RouterEdit(routerEdit).Execute()

Обновление информации о роутере



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
    routerId := "routerId_example" // string | ID роутера
    routerEdit := *openapiclient.NewRouterEdit() // RouterEdit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.UpdateRouter(context.Background(), routerId).RouterEdit(routerEdit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.UpdateRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRouter`: RouterResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.UpdateRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routerEdit** | [**RouterEdit**](RouterEdit.md) |  | 

### Return type

[**RouterResponse**](RouterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouterNat

> RouterResponse UpdateRouterNat(ctx, routerId, networkName).NatIn(natIn).Execute()

Включение NAT для сети



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
    routerId := "routerId_example" // string | ID роутера
    networkName := "networkName_example" // string | Имя сети
    natIn := *openapiclient.NewNatIn("203.0.113.10") // NatIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutersAPI.UpdateRouterNat(context.Background(), routerId, networkName).NatIn(natIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutersAPI.UpdateRouterNat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRouterNat`: RouterResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutersAPI.UpdateRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **string** | ID роутера | 
**networkName** | **string** | Имя сети | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **natIn** | [**NatIn**](NatIn.md) |  | 

### Return type

[**RouterResponse**](RouterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

