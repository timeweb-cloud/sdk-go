# \AppsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProvider**](AppsAPI.md#AddProvider) | **Post** /api/v1/vcs-provider | Привязка vcs провайдера
[**CreateApp**](AppsAPI.md#CreateApp) | **Post** /api/v1/apps | Создание приложения
[**CreateDeploy**](AppsAPI.md#CreateDeploy) | **Post** /api/v1/apps/{app_id}/deploy | Запуск деплоя приложения
[**DeleteApp**](AppsAPI.md#DeleteApp) | **Delete** /api/v1/apps/{app_id} | Удаление приложения
[**DeleteProvider**](AppsAPI.md#DeleteProvider) | **Delete** /api/v1/vcs-provider/{provider_id} | Отвязка vcs провайдера от аккаунта
[**DeployAction**](AppsAPI.md#DeployAction) | **Post** /api/v1/apps/{app_id}/deploy/{deploy_id}/stop | Остановка деплоя приложения
[**GetApp**](AppsAPI.md#GetApp) | **Get** /api/v1/apps/{app_id} | Получение приложения по id
[**GetAppDeploys**](AppsAPI.md#GetAppDeploys) | **Get** /api/v1/apps/{app_id}/deploys | Получение списка деплоев приложения
[**GetAppLogs**](AppsAPI.md#GetAppLogs) | **Get** /api/v1/apps/{app_id}/logs | Получение логов приложения
[**GetAppStatistics**](AppsAPI.md#GetAppStatistics) | **Get** /api/v1/apps/{app_id}/statistics | Получение статистики приложения
[**GetApps**](AppsAPI.md#GetApps) | **Get** /api/v1/apps | Получение списка приложений
[**GetAppsPresets**](AppsAPI.md#GetAppsPresets) | **Get** /api/v1/presets/apps | Получение списка доступных тарифов для приложения
[**GetBranches**](AppsAPI.md#GetBranches) | **Get** /api/v1/vcs-provider/{provider_id}/repository/{repository_id} | Получение списка веток репозитория
[**GetCommits**](AppsAPI.md#GetCommits) | **Get** /api/v1/vcs-provider/{provider_id}/repository/{repository_id}/branch | Получение списка коммитов ветки репозитория
[**GetDeployLogs**](AppsAPI.md#GetDeployLogs) | **Get** /api/v1/apps/{app_id}/deploy/{deploy_id}/logs | Получение логов деплоя приложения
[**GetDeploySettings**](AppsAPI.md#GetDeploySettings) | **Get** /api/v1/deploy-settings/apps | Получение списка дефолтных настроек деплоя для приложения
[**GetFrameworks**](AppsAPI.md#GetFrameworks) | **Get** /api/v1/frameworks/apps | Получение списка доступных фреймворков для приложения
[**GetProviders**](AppsAPI.md#GetProviders) | **Get** /api/v1/vcs-provider | Получение списка vcs провайдеров
[**GetRepositories**](AppsAPI.md#GetRepositories) | **Get** /api/v1/vcs-provider/{provider_id} | Получение списка репозиториев vcs провайдера
[**UpdateAppSettings**](AppsAPI.md#UpdateAppSettings) | **Patch** /api/v1/apps/{app_id} | Изменение настроек приложения
[**UpdateAppState**](AppsAPI.md#UpdateAppState) | **Patch** /api/v1/apps/{app_id}/action/{action} | Изменение состояния приложения



## AddProvider

> AddProvider201Response AddProvider(ctx).AddGithub(addGithub).Execute()

Привязка vcs провайдера



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
    addGithub := *openapiclient.NewAddGithub("GitHub", "ProviderToken_example") // AddGithub | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.AddProvider(context.Background()).AddGithub(addGithub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AddProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProvider`: AddProvider201Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addGithub** | [**AddGithub**](AddGithub.md) |  | 

### Return type

[**AddProvider201Response**](AddProvider201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApp

> CreateApp201Response CreateApp(ctx).CreateApp(createApp).Execute()

Создание приложения



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
    createApp := *openapiclient.NewCreateApp("b392b624-ee80-496e-8b61-47276acac2da", "backend", "eb4e1796-78ce-4232-a824-3c3f4dfaf98f", "npm run build", "master", true, "7d3a6dcca79d2c29e3fe7456ca1598e70d484d7c", "app_example", "example comment", float32(42), openapiclient.frameworks("django")) // CreateApp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.CreateApp(context.Background()).CreateApp(createApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CreateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApp`: CreateApp201Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CreateApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApp** | [**CreateApp**](CreateApp.md) |  | 

### Return type

[**CreateApp201Response**](CreateApp201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeploy

> CreateDeploy201Response CreateDeploy(ctx, appId).CreateDeployRequest(createDeployRequest).Execute()

Запуск деплоя приложения



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
    appId := "appId_example" // string | 
    createDeployRequest := *openapiclient.NewCreateDeployRequest() // CreateDeployRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.CreateDeploy(context.Background(), appId).CreateDeployRequest(createDeployRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CreateDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeploy`: CreateDeploy201Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CreateDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeployRequest** | [**CreateDeployRequest**](CreateDeployRequest.md) |  | 

### Return type

[**CreateDeploy201Response**](CreateDeploy201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> DeleteApp(ctx, appId).Execute()

Удаление приложения



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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppsAPI.DeleteApp(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


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


## DeleteProvider

> DeleteProvider(ctx, providerId).Execute()

Отвязка vcs провайдера от аккаунта



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppsAPI.DeleteProvider(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderRequest struct via the builder pattern


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


## DeployAction

> CreateDeploy201Response DeployAction(ctx, appId, deployId).Execute()

Остановка деплоя приложения



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
    appId := "appId_example" // string | 
    deployId := "deployId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.DeployAction(context.Background(), appId, deployId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeployAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployAction`: CreateDeploy201Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.DeployAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**deployId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateDeploy201Response**](CreateDeploy201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApp

> CreateApp201Response GetApp(ctx, appId).Execute()

Получение приложения по id



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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetApp(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApp`: CreateApp201Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateApp201Response**](CreateApp201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppDeploys

> GetAppDeploys200Response GetAppDeploys(ctx, appId).Limit(limit).Offset(offset).Execute()

Получение списка деплоев приложения



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
    appId := "appId_example" // string | 
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 5)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetAppDeploys(context.Background(), appId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppDeploys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppDeploys`: GetAppDeploys200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppDeploys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppDeploysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 5]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**GetAppDeploys200Response**](GetAppDeploys200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppLogs

> GetAppLogs200Response GetAppLogs(ctx, appId).Execute()

Получение логов приложения



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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetAppLogs(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppLogs`: GetAppLogs200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAppLogs200Response**](GetAppLogs200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppStatistics

> GetServerStatistics200Response GetAppStatistics(ctx, appId).DateFrom(dateFrom).DateTo(dateTo).Execute()

Получение статистики приложения



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
    appId := "appId_example" // string | 
    dateFrom := "dateFrom_example" // string | Дата начала сбора статистики. Строка в формате ISO 8061, закодированная в ASCII, пример: `2023-05-25%202023-05-25T14%3A35%3A38`
    dateTo := "dateTo_example" // string | Дата окончания сбора статистики. Строка в формате ISO 8061, закодированная в ASCII, пример: `2023-05-26%202023-05-25T14%3A35%3A38`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetAppStatistics(context.Background(), appId).DateFrom(dateFrom).DateTo(dateTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppStatistics`: GetServerStatistics200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateFrom** | **string** | Дата начала сбора статистики. Строка в формате ISO 8061, закодированная в ASCII, пример: &#x60;2023-05-25%202023-05-25T14%3A35%3A38&#x60; | 
 **dateTo** | **string** | Дата окончания сбора статистики. Строка в формате ISO 8061, закодированная в ASCII, пример: &#x60;2023-05-26%202023-05-25T14%3A35%3A38&#x60; | 

### Return type

[**GetServerStatistics200Response**](GetServerStatistics200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApps

> GetApps200Response GetApps(ctx).Execute()

Получение списка приложений



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
    resp, r, err := apiClient.AppsAPI.GetApps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApps`: GetApps200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


### Return type

[**GetApps200Response**](GetApps200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppsPresets

> AppsPresets GetAppsPresets(ctx, appId).Execute()

Получение списка доступных тарифов для приложения



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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetAppsPresets(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppsPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppsPresets`: AppsPresets
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppsPresets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsPresetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppsPresets**](AppsPresets.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranches

> GetBranches200Response GetBranches(ctx, providerId, repositoryId).Execute()

Получение списка веток репозитория



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
    providerId := "providerId_example" // string | 
    repositoryId := "repositoryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetBranches(context.Background(), providerId, repositoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBranches`: GetBranches200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**repositoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBranches200Response**](GetBranches200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommits

> GetCommits200Response GetCommits(ctx, accountId, providerId, repositoryId).Name(name).Execute()

Получение списка коммитов ветки репозитория



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
    accountId := "accountId_example" // string | 
    providerId := "providerId_example" // string | 
    repositoryId := "repositoryId_example" // string | 
    name := "name_example" // string | Название ветки

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetCommits(context.Background(), accountId, providerId, repositoryId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommits`: GetCommits200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**providerId** | **string** |  | 
**repositoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **string** | Название ветки | 

### Return type

[**GetCommits200Response**](GetCommits200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployLogs

> GetDeployLogs200Response GetDeployLogs(ctx, appId, deployId).Debug(debug).Execute()

Получение логов деплоя приложения



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
    appId := "appId_example" // string | 
    deployId := "deployId_example" // string | 
    debug := true // bool | Управляет выводом логов деплоя (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetDeployLogs(context.Background(), appId, deployId).Debug(debug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetDeployLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployLogs`: GetDeployLogs200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetDeployLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**deployId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **debug** | **bool** | Управляет выводом логов деплоя | 

### Return type

[**GetDeployLogs200Response**](GetDeployLogs200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploySettings

> GetDeploySettings200Response GetDeploySettings(ctx, appId).Execute()

Получение списка дефолтных настроек деплоя для приложения



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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetDeploySettings(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetDeploySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploySettings`: GetDeploySettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetDeploySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeploySettings200Response**](GetDeploySettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworks

> AvailableFrameworks GetFrameworks(ctx, appId).Execute()

Получение списка доступных фреймворков для приложения



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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetFrameworks(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetFrameworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFrameworks`: AvailableFrameworks
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetFrameworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvailableFrameworks**](AvailableFrameworks.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviders

> GetProviders200Response GetProviders(ctx).Execute()

Получение списка vcs провайдеров



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
    resp, r, err := apiClient.AppsAPI.GetProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProviders`: GetProviders200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersRequest struct via the builder pattern


### Return type

[**GetProviders200Response**](GetProviders200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositories

> GetRepositories200Response GetRepositories(ctx, providerId).Execute()

Получение списка репозиториев vcs провайдера



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.GetRepositories(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositories`: GetRepositories200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRepositories200Response**](GetRepositories200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSettings

> UpdateAppSettings200Response UpdateAppSettings(ctx, appId).UpdeteSettings(updeteSettings).Execute()

Изменение настроек приложения



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
    appId := "appId_example" // string | 
    updeteSettings := *openapiclient.NewUpdeteSettings() // UpdeteSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.UpdateAppSettings(context.Background(), appId).UpdeteSettings(updeteSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppSettings`: UpdateAppSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updeteSettings** | [**UpdeteSettings**](UpdeteSettings.md) |  | 

### Return type

[**UpdateAppSettings200Response**](UpdateAppSettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppState

> UpdateAppState(ctx, appId, action).Execute()

Изменение состояния приложения



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
    appId := "appId_example" // string | 
    action := "action_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppsAPI.UpdateAppState(context.Background(), appId, action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateAppState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**action** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppStateRequest struct via the builder pattern


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

