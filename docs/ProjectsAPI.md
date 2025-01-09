# \ProjectsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBalancerToProject**](ProjectsAPI.md#AddBalancerToProject) | **Post** /api/v1/projects/{project_id}/resources/balancers | Добавление балансировщика в проект
[**AddClusterToProject**](ProjectsAPI.md#AddClusterToProject) | **Post** /api/v1/projects/{project_id}/resources/clusters | Добавление кластера в проект
[**AddDatabaseToProject**](ProjectsAPI.md#AddDatabaseToProject) | **Post** /api/v1/projects/{project_id}/resources/databases | Добавление базы данных в проект
[**AddDedicatedServerToProject**](ProjectsAPI.md#AddDedicatedServerToProject) | **Post** /api/v1/projects/{project_id}/resources/dedicated | Добавление выделенного сервера в проект
[**AddServerToProject**](ProjectsAPI.md#AddServerToProject) | **Post** /api/v1/projects/{project_id}/resources/servers | Добавление сервера в проект
[**AddStorageToProject**](ProjectsAPI.md#AddStorageToProject) | **Post** /api/v1/projects/{project_id}/resources/buckets | Добавление хранилища в проект
[**CreateProject**](ProjectsAPI.md#CreateProject) | **Post** /api/v1/projects | Создание проекта
[**DeleteProject**](ProjectsAPI.md#DeleteProject) | **Delete** /api/v1/projects/{project_id} | Удаление проекта
[**GetAccountBalancers**](ProjectsAPI.md#GetAccountBalancers) | **Get** /api/v1/projects/resources/balancers | Получение списка всех балансировщиков на аккаунте
[**GetAccountClusters**](ProjectsAPI.md#GetAccountClusters) | **Get** /api/v1/projects/resources/clusters | Получение списка всех кластеров на аккаунте
[**GetAccountDatabases**](ProjectsAPI.md#GetAccountDatabases) | **Get** /api/v1/projects/resources/databases | Получение списка всех баз данных на аккаунте
[**GetAccountDedicatedServers**](ProjectsAPI.md#GetAccountDedicatedServers) | **Get** /api/v1/projects/resources/dedicated | Получение списка всех выделенных серверов на аккаунте
[**GetAccountServers**](ProjectsAPI.md#GetAccountServers) | **Get** /api/v1/projects/resources/servers | Получение списка всех серверов на аккаунте
[**GetAccountStorages**](ProjectsAPI.md#GetAccountStorages) | **Get** /api/v1/projects/resources/buckets | Получение списка всех хранилищ на аккаунте
[**GetAllProjectResources**](ProjectsAPI.md#GetAllProjectResources) | **Get** /api/v1/projects/{project_id}/resources | Получение всех ресурсов проекта
[**GetProject**](ProjectsAPI.md#GetProject) | **Get** /api/v1/projects/{project_id} | Получение проекта по ID
[**GetProjectBalancers**](ProjectsAPI.md#GetProjectBalancers) | **Get** /api/v1/projects/{project_id}/resources/balancers | Получение списка балансировщиков проекта
[**GetProjectClusters**](ProjectsAPI.md#GetProjectClusters) | **Get** /api/v1/projects/{project_id}/resources/clusters | Получение списка кластеров проекта
[**GetProjectDatabases**](ProjectsAPI.md#GetProjectDatabases) | **Get** /api/v1/projects/{project_id}/resources/databases | Получение списка баз данных проекта
[**GetProjectDedicatedServers**](ProjectsAPI.md#GetProjectDedicatedServers) | **Get** /api/v1/projects/{project_id}/resources/dedicated | Получение списка выделенных серверов проекта
[**GetProjectServers**](ProjectsAPI.md#GetProjectServers) | **Get** /api/v1/projects/{project_id}/resources/servers | Получение списка серверов проекта
[**GetProjectStorages**](ProjectsAPI.md#GetProjectStorages) | **Get** /api/v1/projects/{project_id}/resources/buckets | Получение списка хранилищ проекта
[**GetProjects**](ProjectsAPI.md#GetProjects) | **Get** /api/v1/projects | Получение списка проектов
[**TransferResourceToAnotherProject**](ProjectsAPI.md#TransferResourceToAnotherProject) | **Put** /api/v1/projects/{project_id}/resources/transfer | Перенести ресурс в другой проект
[**UpdateProject**](ProjectsAPI.md#UpdateProject) | **Put** /api/v1/projects/{project_id} | Изменение проекта



## AddBalancerToProject

> AddBalancerToProject200Response AddBalancerToProject(ctx, projectId).AddBalancerToProjectRequest(addBalancerToProjectRequest).Execute()

Добавление балансировщика в проект



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
    projectId := int32(99) // int32 | ID проекта.
    addBalancerToProjectRequest := *openapiclient.NewAddBalancerToProjectRequest(float32(1)) // AddBalancerToProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.AddBalancerToProject(context.Background(), projectId).AddBalancerToProjectRequest(addBalancerToProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddBalancerToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBalancerToProject`: AddBalancerToProject200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AddBalancerToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBalancerToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addBalancerToProjectRequest** | [**AddBalancerToProjectRequest**](AddBalancerToProjectRequest.md) |  | 

### Return type

[**AddBalancerToProject200Response**](AddBalancerToProject200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddClusterToProject

> AddBalancerToProject200Response AddClusterToProject(ctx, projectId).AddClusterToProjectRequest(addClusterToProjectRequest).Execute()

Добавление кластера в проект



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
    projectId := int32(99) // int32 | ID проекта.
    addClusterToProjectRequest := *openapiclient.NewAddClusterToProjectRequest(float32(1)) // AddClusterToProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.AddClusterToProject(context.Background(), projectId).AddClusterToProjectRequest(addClusterToProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddClusterToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClusterToProject`: AddBalancerToProject200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AddClusterToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addClusterToProjectRequest** | [**AddClusterToProjectRequest**](AddClusterToProjectRequest.md) |  | 

### Return type

[**AddBalancerToProject200Response**](AddBalancerToProject200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDatabaseToProject

> AddBalancerToProject200Response AddDatabaseToProject(ctx, projectId).AddDatabaseToProjectRequest(addDatabaseToProjectRequest).Execute()

Добавление базы данных в проект



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
    projectId := int32(99) // int32 | ID проекта.
    addDatabaseToProjectRequest := *openapiclient.NewAddDatabaseToProjectRequest(float32(1)) // AddDatabaseToProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.AddDatabaseToProject(context.Background(), projectId).AddDatabaseToProjectRequest(addDatabaseToProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddDatabaseToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDatabaseToProject`: AddBalancerToProject200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AddDatabaseToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDatabaseToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addDatabaseToProjectRequest** | [**AddDatabaseToProjectRequest**](AddDatabaseToProjectRequest.md) |  | 

### Return type

[**AddBalancerToProject200Response**](AddBalancerToProject200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDedicatedServerToProject

> AddBalancerToProject200Response AddDedicatedServerToProject(ctx, projectId).AddDedicatedServerToProjectRequest(addDedicatedServerToProjectRequest).Execute()

Добавление выделенного сервера в проект



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
    projectId := int32(99) // int32 | ID проекта.
    addDedicatedServerToProjectRequest := *openapiclient.NewAddDedicatedServerToProjectRequest(float32(1)) // AddDedicatedServerToProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.AddDedicatedServerToProject(context.Background(), projectId).AddDedicatedServerToProjectRequest(addDedicatedServerToProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddDedicatedServerToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDedicatedServerToProject`: AddBalancerToProject200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AddDedicatedServerToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDedicatedServerToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addDedicatedServerToProjectRequest** | [**AddDedicatedServerToProjectRequest**](AddDedicatedServerToProjectRequest.md) |  | 

### Return type

[**AddBalancerToProject200Response**](AddBalancerToProject200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddServerToProject

> AddBalancerToProject200Response AddServerToProject(ctx, projectId).AddServerToProjectRequest(addServerToProjectRequest).Execute()

Добавление сервера в проект



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
    projectId := int32(99) // int32 | ID проекта.
    addServerToProjectRequest := *openapiclient.NewAddServerToProjectRequest(float32(1)) // AddServerToProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.AddServerToProject(context.Background(), projectId).AddServerToProjectRequest(addServerToProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddServerToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServerToProject`: AddBalancerToProject200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AddServerToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServerToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addServerToProjectRequest** | [**AddServerToProjectRequest**](AddServerToProjectRequest.md) |  | 

### Return type

[**AddBalancerToProject200Response**](AddBalancerToProject200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddStorageToProject

> AddBalancerToProject200Response AddStorageToProject(ctx, projectId).AddStorageToProjectRequest(addStorageToProjectRequest).Execute()

Добавление хранилища в проект



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
    projectId := int32(99) // int32 | ID проекта.
    addStorageToProjectRequest := *openapiclient.NewAddStorageToProjectRequest(float32(1)) // AddStorageToProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.AddStorageToProject(context.Background(), projectId).AddStorageToProjectRequest(addStorageToProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddStorageToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddStorageToProject`: AddBalancerToProject200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AddStorageToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddStorageToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addStorageToProjectRequest** | [**AddStorageToProjectRequest**](AddStorageToProjectRequest.md) |  | 

### Return type

[**AddBalancerToProject200Response**](AddBalancerToProject200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> CreateProject201Response CreateProject(ctx).CreateProject(createProject).Execute()

Создание проекта



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
    createProject := *openapiclient.NewCreateProject("Name") // CreateProject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.CreateProject(context.Background()).CreateProject(createProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: CreateProject201Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProject** | [**CreateProject**](CreateProject.md) |  | 

### Return type

[**CreateProject201Response**](CreateProject201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx, projectId).Execute()

Удаление проекта



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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


## GetAccountBalancers

> GetProjectBalancers200Response GetAccountBalancers(ctx).Execute()

Получение списка всех балансировщиков на аккаунте



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
    resp, r, err := apiClient.ProjectsAPI.GetAccountBalancers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAccountBalancers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountBalancers`: GetProjectBalancers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAccountBalancers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBalancersRequest struct via the builder pattern


### Return type

[**GetProjectBalancers200Response**](GetProjectBalancers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountClusters

> GetProjectClusters200Response GetAccountClusters(ctx).Execute()

Получение списка всех кластеров на аккаунте



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
    resp, r, err := apiClient.ProjectsAPI.GetAccountClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAccountClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountClusters`: GetProjectClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAccountClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountClustersRequest struct via the builder pattern


### Return type

[**GetProjectClusters200Response**](GetProjectClusters200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountDatabases

> GetProjectDatabases200Response GetAccountDatabases(ctx).Execute()

Получение списка всех баз данных на аккаунте



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
    resp, r, err := apiClient.ProjectsAPI.GetAccountDatabases(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAccountDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountDatabases`: GetProjectDatabases200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAccountDatabases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDatabasesRequest struct via the builder pattern


### Return type

[**GetProjectDatabases200Response**](GetProjectDatabases200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountDedicatedServers

> GetProjectDedicatedServers200Response GetAccountDedicatedServers(ctx).Execute()

Получение списка всех выделенных серверов на аккаунте



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
    resp, r, err := apiClient.ProjectsAPI.GetAccountDedicatedServers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAccountDedicatedServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountDedicatedServers`: GetProjectDedicatedServers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAccountDedicatedServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDedicatedServersRequest struct via the builder pattern


### Return type

[**GetProjectDedicatedServers200Response**](GetProjectDedicatedServers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountServers

> GetProjectServers200Response GetAccountServers(ctx).Execute()

Получение списка всех серверов на аккаунте



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
    resp, r, err := apiClient.ProjectsAPI.GetAccountServers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAccountServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountServers`: GetProjectServers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAccountServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountServersRequest struct via the builder pattern


### Return type

[**GetProjectServers200Response**](GetProjectServers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountStorages

> GetProjectStorages200Response GetAccountStorages(ctx).Execute()

Получение списка всех хранилищ на аккаунте



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
    resp, r, err := apiClient.ProjectsAPI.GetAccountStorages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAccountStorages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountStorages`: GetProjectStorages200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAccountStorages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountStoragesRequest struct via the builder pattern


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


## GetAllProjectResources

> GetAllProjectResources200Response GetAllProjectResources(ctx, projectId).Execute()

Получение всех ресурсов проекта



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.GetAllProjectResources(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetAllProjectResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProjectResources`: GetAllProjectResources200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetAllProjectResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllProjectResources200Response**](GetAllProjectResources200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> CreateProject201Response GetProject(ctx, projectId).Execute()

Получение проекта по ID



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: CreateProject201Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateProject201Response**](CreateProject201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBalancers

> GetProjectBalancers200Response GetProjectBalancers(ctx, projectId).Execute()

Получение списка балансировщиков проекта



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.GetProjectBalancers(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectBalancers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectBalancers`: GetProjectBalancers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectBalancers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProjectBalancers200Response**](GetProjectBalancers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectClusters

> GetProjectClusters200Response GetProjectClusters(ctx, projectId).Execute()

Получение списка кластеров проекта



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.GetProjectClusters(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectClusters`: GetProjectClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProjectClusters200Response**](GetProjectClusters200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDatabases

> GetProjectDatabases200Response GetProjectDatabases(ctx, projectId).Execute()

Получение списка баз данных проекта



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.GetProjectDatabases(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectDatabases`: GetProjectDatabases200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProjectDatabases200Response**](GetProjectDatabases200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDedicatedServers

> GetProjectDedicatedServers200Response GetProjectDedicatedServers(ctx, projectId).Execute()

Получение списка выделенных серверов проекта



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.GetProjectDedicatedServers(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectDedicatedServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectDedicatedServers`: GetProjectDedicatedServers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectDedicatedServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDedicatedServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProjectDedicatedServers200Response**](GetProjectDedicatedServers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectServers

> GetProjectServers200Response GetProjectServers(ctx, projectId).Execute()

Получение списка серверов проекта



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.GetProjectServers(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectServers`: GetProjectServers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProjectServers200Response**](GetProjectServers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectStorages

> GetProjectStorages200Response GetProjectStorages(ctx, projectId).Execute()

Получение списка хранилищ проекта



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
    projectId := int32(99) // int32 | ID проекта.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.GetProjectStorages(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectStorages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectStorages`: GetProjectStorages200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetProjects

> GetProjects200Response GetProjects(ctx).Execute()

Получение списка проектов



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
    resp, r, err := apiClient.ProjectsAPI.GetProjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjects`: GetProjects200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


### Return type

[**GetProjects200Response**](GetProjects200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferResourceToAnotherProject

> AddBalancerToProject200Response TransferResourceToAnotherProject(ctx, projectId).ResourceTransfer(resourceTransfer).Execute()

Перенести ресурс в другой проект



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
    projectId := int32(99) // int32 | ID проекта.
    resourceTransfer := *openapiclient.NewResourceTransfer(float32(1), float32(1), "server") // ResourceTransfer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.TransferResourceToAnotherProject(context.Background(), projectId).ResourceTransfer(resourceTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.TransferResourceToAnotherProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferResourceToAnotherProject`: AddBalancerToProject200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.TransferResourceToAnotherProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferResourceToAnotherProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceTransfer** | [**ResourceTransfer**](ResourceTransfer.md) |  | 

### Return type

[**AddBalancerToProject200Response**](AddBalancerToProject200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> CreateProject201Response UpdateProject(ctx, projectId).UpdateProject(updateProject).Execute()

Изменение проекта



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
    projectId := int32(99) // int32 | ID проекта.
    updateProject := *openapiclient.NewUpdateProject() // UpdateProject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.UpdateProject(context.Background(), projectId).UpdateProject(updateProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: CreateProject201Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | ID проекта. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProject** | [**UpdateProject**](UpdateProject.md) |  | 

### Return type

[**CreateProject201Response**](CreateProject201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

