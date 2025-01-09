# \DatabasesAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](DatabasesAPI.md#CreateDatabase) | **Post** /api/v1/dbs | Создание базы данных
[**CreateDatabaseBackup**](DatabasesAPI.md#CreateDatabaseBackup) | **Post** /api/v1/dbs/{db_id}/backups | Создание бэкапа базы данных
[**CreateDatabaseCluster**](DatabasesAPI.md#CreateDatabaseCluster) | **Post** /api/v1/databases | Создание кластера базы данных
[**CreateDatabaseInstance**](DatabasesAPI.md#CreateDatabaseInstance) | **Post** /api/v1/databases/{db_cluster_id}/instances | Создание инстанса базы данных
[**CreateDatabaseUser**](DatabasesAPI.md#CreateDatabaseUser) | **Post** /api/v1/databases/{db_cluster_id}/admins | Создание пользователя базы данных
[**DeleteDatabase**](DatabasesAPI.md#DeleteDatabase) | **Delete** /api/v1/dbs/{db_id} | Удаление базы данных
[**DeleteDatabaseBackup**](DatabasesAPI.md#DeleteDatabaseBackup) | **Delete** /api/v1/dbs/{db_id}/backups/{backup_id} | Удаление бэкапа базы данных
[**DeleteDatabaseCluster**](DatabasesAPI.md#DeleteDatabaseCluster) | **Delete** /api/v1/databases/{db_cluster_id} | Удаление кластера базы данных
[**DeleteDatabaseInstance**](DatabasesAPI.md#DeleteDatabaseInstance) | **Delete** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Удаление инстанса базы данных
[**DeleteDatabaseUser**](DatabasesAPI.md#DeleteDatabaseUser) | **Delete** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Удаление пользователя базы данных
[**GetDatabase**](DatabasesAPI.md#GetDatabase) | **Get** /api/v1/dbs/{db_id} | Получение базы данных
[**GetDatabaseAutoBackupsSettings**](DatabasesAPI.md#GetDatabaseAutoBackupsSettings) | **Get** /api/v1/dbs/{db_id}/auto-backups | Получение настроек автобэкапов базы данных
[**GetDatabaseBackup**](DatabasesAPI.md#GetDatabaseBackup) | **Get** /api/v1/dbs/{db_id}/backups/{backup_id} | Получение бэкапа базы данных
[**GetDatabaseBackups**](DatabasesAPI.md#GetDatabaseBackups) | **Get** /api/v1/dbs/{db_id}/backups | Список бэкапов базы данных
[**GetDatabaseCluster**](DatabasesAPI.md#GetDatabaseCluster) | **Get** /api/v1/databases/{db_cluster_id} | Получение кластера базы данных
[**GetDatabaseClusterTypes**](DatabasesAPI.md#GetDatabaseClusterTypes) | **Get** /api/v1/database-types | Получение списка типов кластеров баз данных
[**GetDatabaseClusters**](DatabasesAPI.md#GetDatabaseClusters) | **Get** /api/v1/databases | Получение списка кластеров баз данных
[**GetDatabaseInstance**](DatabasesAPI.md#GetDatabaseInstance) | **Get** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Получение инстанса базы данных
[**GetDatabaseInstances**](DatabasesAPI.md#GetDatabaseInstances) | **Get** /api/v1/databases/{db_cluster_id}/instances | Получение списка инстансов баз данных
[**GetDatabaseParameters**](DatabasesAPI.md#GetDatabaseParameters) | **Get** /api/v1/dbs/parameters | Получение списка параметров баз данных
[**GetDatabaseUser**](DatabasesAPI.md#GetDatabaseUser) | **Get** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Получение пользователя базы данных
[**GetDatabaseUsers**](DatabasesAPI.md#GetDatabaseUsers) | **Get** /api/v1/databases/{db_cluster_id}/admins | Получение списка пользователей базы данных
[**GetDatabases**](DatabasesAPI.md#GetDatabases) | **Get** /api/v1/dbs | Получение списка всех баз данных
[**GetDatabasesPresets**](DatabasesAPI.md#GetDatabasesPresets) | **Get** /api/v1/presets/dbs | Получение списка тарифов для баз данных
[**RestoreDatabaseFromBackup**](DatabasesAPI.md#RestoreDatabaseFromBackup) | **Put** /api/v1/dbs/{db_id}/backups/{backup_id} | Восстановление базы данных из бэкапа
[**UpdateDatabase**](DatabasesAPI.md#UpdateDatabase) | **Patch** /api/v1/dbs/{db_id} | Обновление базы данных
[**UpdateDatabaseAutoBackupsSettings**](DatabasesAPI.md#UpdateDatabaseAutoBackupsSettings) | **Patch** /api/v1/dbs/{db_id}/auto-backups | Изменение настроек автобэкапов базы данных
[**UpdateDatabaseCluster**](DatabasesAPI.md#UpdateDatabaseCluster) | **Patch** /api/v1/databases/{db_cluster_id} | Изменение кластера базы данных
[**UpdateDatabaseInstance**](DatabasesAPI.md#UpdateDatabaseInstance) | **Patch** /api/v1/databases/{db_cluster_id}/instances/{instance_id} | Изменение инстанса базы данных
[**UpdateDatabaseUser**](DatabasesAPI.md#UpdateDatabaseUser) | **Patch** /api/v1/databases/{db_cluster_id}/admins/{admin_id} | Изменение пользователя базы данных



## CreateDatabase

> CreateDatabase201Response CreateDatabase(ctx).CreateDb(createDb).Execute()

Создание базы данных



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
    createDb := *openapiclient.NewCreateDb("password", "default_db", openapiclient.db-type("mysql"), int32(5)) // CreateDb | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.CreateDatabase(context.Background()).CreateDb(createDb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabase`: CreateDatabase201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDb** | [**CreateDb**](CreateDb.md) |  | 

### Return type

[**CreateDatabase201Response**](CreateDatabase201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseBackup

> CreateDatabaseBackup201Response CreateDatabaseBackup(ctx, dbId).Execute()

Создание бэкапа базы данных



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
    dbId := int32(56) // int32 | ID базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.CreateDatabaseBackup(context.Background(), dbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabaseBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseBackup`: CreateDatabaseBackup201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabaseBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDatabaseBackup201Response**](CreateDatabaseBackup201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseCluster

> CreateDatabaseCluster201Response CreateDatabaseCluster(ctx).CreateCluster(createCluster).Execute()

Создание кластера базы данных



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
    createCluster := *openapiclient.NewCreateCluster("default_db", openapiclient.db-type("mysql"), int32(5)) // CreateCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.CreateDatabaseCluster(context.Background()).CreateCluster(createCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabaseCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseCluster`: CreateDatabaseCluster201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabaseCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCluster** | [**CreateCluster**](CreateCluster.md) |  | 

### Return type

[**CreateDatabaseCluster201Response**](CreateDatabaseCluster201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseInstance

> CreateDatabaseInstance201Response CreateDatabaseInstance(ctx, dbClusterId).CreateInstance(createInstance).Execute()

Создание инстанса базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    createInstance := *openapiclient.NewCreateInstance("default_name") // CreateInstance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.CreateDatabaseInstance(context.Background(), dbClusterId).CreateInstance(createInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabaseInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseInstance`: CreateDatabaseInstance201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabaseInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInstance** | [**CreateInstance**](CreateInstance.md) |  | 

### Return type

[**CreateDatabaseInstance201Response**](CreateDatabaseInstance201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabaseUser

> CreateDatabaseUser201Response CreateDatabaseUser(ctx, dbClusterId).CreateAdmin(createAdmin).Execute()

Создание пользователя базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    createAdmin := *openapiclient.NewCreateAdmin("default_login", "bs.:L2f$Tm:SC~", []string{"Privileges_example"}) // CreateAdmin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.CreateDatabaseUser(context.Background(), dbClusterId).CreateAdmin(createAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.CreateDatabaseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseUser`: CreateDatabaseUser201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.CreateDatabaseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAdmin** | [**CreateAdmin**](CreateAdmin.md) |  | 

### Return type

[**CreateDatabaseUser201Response**](CreateDatabaseUser201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> DeleteDatabase200Response DeleteDatabase(ctx, dbId).Hash(hash).Code(code).Execute()

Удаление базы данных



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
    dbId := int32(56) // int32 | ID базы данных
    hash := "15095f25-aac3-4d60-a788-96cb5136f186" // string | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. (optional)
    code := "0000" // string | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена `is_able_to_delete` установлен в значение `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.DeleteDatabase(context.Background(), dbId).Hash(hash).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabase`: DeleteDatabase200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hash** | **string** | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. | 
 **code** | **string** | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена &#x60;is_able_to_delete&#x60; установлен в значение &#x60;true&#x60; | 

### Return type

[**DeleteDatabase200Response**](DeleteDatabase200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseBackup

> DeleteDatabaseBackup(ctx, dbId, backupId).Execute()

Удаление бэкапа базы данных



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
    dbId := int32(56) // int32 | ID базы данных
    backupId := int32(56) // int32 | ID резевной копии

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabasesAPI.DeleteDatabaseBackup(context.Background(), dbId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabaseBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 
**backupId** | **int32** | ID резевной копии | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseBackupRequest struct via the builder pattern


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


## DeleteDatabaseCluster

> DeleteDatabaseCluster200Response DeleteDatabaseCluster(ctx, dbClusterId).Hash(hash).Code(code).Execute()

Удаление кластера базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    hash := "15095f25-aac3-4d60-a788-96cb5136f186" // string | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. (optional)
    code := "0000" // string | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена `is_able_to_delete` установлен в значение `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.DeleteDatabaseCluster(context.Background(), dbClusterId).Hash(hash).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabaseCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabaseCluster`: DeleteDatabaseCluster200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabaseCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hash** | **string** | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. | 
 **code** | **string** | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена &#x60;is_able_to_delete&#x60; установлен в значение &#x60;true&#x60; | 

### Return type

[**DeleteDatabaseCluster200Response**](DeleteDatabaseCluster200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabaseInstance

> DeleteDatabaseInstance(ctx, dbClusterId, instanceId).Execute()

Удаление инстанса базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    instanceId := int32(56) // int32 | ID инстанса базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabasesAPI.DeleteDatabaseInstance(context.Background(), dbClusterId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabaseInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 
**instanceId** | **int32** | ID инстанса базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseInstanceRequest struct via the builder pattern


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


## DeleteDatabaseUser

> DeleteDatabaseUser(ctx, dbClusterId, adminId).Execute()

Удаление пользователя базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    adminId := int32(56) // int32 | ID пользователя базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabasesAPI.DeleteDatabaseUser(context.Background(), dbClusterId, adminId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabaseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 
**adminId** | **int32** | ID пользователя базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseUserRequest struct via the builder pattern


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


## GetDatabase

> CreateDatabase201Response GetDatabase(ctx, dbId).Execute()

Получение базы данных



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
    dbId := int32(56) // int32 | ID базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabase(context.Background(), dbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabase`: CreateDatabase201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDatabase201Response**](CreateDatabase201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseAutoBackupsSettings

> GetDatabaseAutoBackupsSettings200Response GetDatabaseAutoBackupsSettings(ctx, dbId).Execute()

Получение настроек автобэкапов базы данных



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
    dbId := int32(56) // int32 | ID базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseAutoBackupsSettings(context.Background(), dbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseAutoBackupsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseAutoBackupsSettings`: GetDatabaseAutoBackupsSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseAutoBackupsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseAutoBackupsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDatabaseAutoBackupsSettings200Response**](GetDatabaseAutoBackupsSettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseBackup

> CreateDatabaseBackup201Response GetDatabaseBackup(ctx, dbId, backupId).Execute()

Получение бэкапа базы данных



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
    dbId := int32(56) // int32 | ID базы данных
    backupId := int32(56) // int32 | ID резевной копии

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseBackup(context.Background(), dbId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseBackup`: CreateDatabaseBackup201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 
**backupId** | **int32** | ID резевной копии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateDatabaseBackup201Response**](CreateDatabaseBackup201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseBackups

> GetDatabaseBackups200Response GetDatabaseBackups(ctx, dbId).Limit(limit).Offset(offset).Execute()

Список бэкапов базы данных



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
    dbId := int32(56) // int32 | ID базы данных
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseBackups(context.Background(), dbId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseBackups`: GetDatabaseBackups200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**GetDatabaseBackups200Response**](GetDatabaseBackups200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseCluster

> CreateDatabaseCluster201Response GetDatabaseCluster(ctx, dbClusterId).Execute()

Получение кластера базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseCluster(context.Background(), dbClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseCluster`: CreateDatabaseCluster201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDatabaseCluster201Response**](CreateDatabaseCluster201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseClusterTypes

> GetDatabaseClusterTypes200Response GetDatabaseClusterTypes(ctx).Execute()

Получение списка типов кластеров баз данных



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
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseClusterTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseClusterTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseClusterTypes`: GetDatabaseClusterTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseClusterTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseClusterTypesRequest struct via the builder pattern


### Return type

[**GetDatabaseClusterTypes200Response**](GetDatabaseClusterTypes200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseClusters

> GetDatabaseClusters200Response GetDatabaseClusters(ctx).Limit(limit).Offset(offset).Execute()

Получение списка кластеров баз данных



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
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseClusters(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseClusters`: GetDatabaseClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**GetDatabaseClusters200Response**](GetDatabaseClusters200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseInstance

> CreateDatabaseInstance201Response GetDatabaseInstance(ctx, dbClusterId, instanceId).Execute()

Получение инстанса базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    instanceId := int32(56) // int32 | ID инстанса базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseInstance(context.Background(), dbClusterId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseInstance`: CreateDatabaseInstance201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 
**instanceId** | **int32** | ID инстанса базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateDatabaseInstance201Response**](CreateDatabaseInstance201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseInstances

> GetDatabaseInstances200Response GetDatabaseInstances(ctx, dbClusterId).Execute()

Получение списка инстансов баз данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseInstances(context.Background(), dbClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseInstances`: GetDatabaseInstances200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDatabaseInstances200Response**](GetDatabaseInstances200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseParameters

> map[string][]string GetDatabaseParameters(ctx).Execute()

Получение списка параметров баз данных



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
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseParameters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseParameters`: map[string][]string
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseParameters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseParametersRequest struct via the builder pattern


### Return type

[**map[string][]string**](array.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseUser

> CreateDatabaseUser201Response GetDatabaseUser(ctx, dbClusterId, adminId).Execute()

Получение пользователя базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    adminId := int32(56) // int32 | ID пользователя базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseUser(context.Background(), dbClusterId, adminId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseUser`: CreateDatabaseUser201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 
**adminId** | **int32** | ID пользователя базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateDatabaseUser201Response**](CreateDatabaseUser201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabaseUsers

> GetDatabaseUsers200Response GetDatabaseUsers(ctx, dbClusterId).Execute()

Получение списка пользователей базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.GetDatabaseUsers(context.Background(), dbClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabaseUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseUsers`: GetDatabaseUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDatabaseUsers200Response**](GetDatabaseUsers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabases

> GetDatabases200Response GetDatabases(ctx).Limit(limit).Offset(offset).Execute()

Получение списка всех баз данных



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
    resp, r, err := apiClient.DatabasesAPI.GetDatabases(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabases`: GetDatabases200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**GetDatabases200Response**](GetDatabases200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesPresets

> GetDatabasesPresets200Response GetDatabasesPresets(ctx).Execute()

Получение списка тарифов для баз данных



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
    resp, r, err := apiClient.DatabasesAPI.GetDatabasesPresets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabasesPresets`: GetDatabasesPresets200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesPresetsRequest struct via the builder pattern


### Return type

[**GetDatabasesPresets200Response**](GetDatabasesPresets200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDatabaseFromBackup

> RestoreDatabaseFromBackup(ctx, dbId, backupId).Execute()

Восстановление базы данных из бэкапа



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
    dbId := int32(56) // int32 | ID базы данных
    backupId := int32(56) // int32 | ID резевной копии

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DatabasesAPI.RestoreDatabaseFromBackup(context.Background(), dbId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.RestoreDatabaseFromBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 
**backupId** | **int32** | ID резевной копии | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDatabaseFromBackupRequest struct via the builder pattern


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


## UpdateDatabase

> CreateDatabase201Response UpdateDatabase(ctx, dbId).UpdateDb(updateDb).Execute()

Обновление базы данных



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
    dbId := int32(56) // int32 | ID базы данных
    updateDb := *openapiclient.NewUpdateDb() // UpdateDb | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.UpdateDatabase(context.Background(), dbId).UpdateDb(updateDb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.UpdateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDatabase`: CreateDatabase201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.UpdateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDb** | [**UpdateDb**](UpdateDb.md) |  | 

### Return type

[**CreateDatabase201Response**](CreateDatabase201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabaseAutoBackupsSettings

> GetDatabaseAutoBackupsSettings200Response UpdateDatabaseAutoBackupsSettings(ctx, dbId).AutoBackup(autoBackup).Execute()

Изменение настроек автобэкапов базы данных



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
    dbId := int32(56) // int32 | ID базы данных
    autoBackup := *openapiclient.NewAutoBackup(true) // AutoBackup | При значении `is_enabled`: `true`, поля `copy_count`, `creation_start_at`, `interval` являются обязательными (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.UpdateDatabaseAutoBackupsSettings(context.Background(), dbId).AutoBackup(autoBackup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.UpdateDatabaseAutoBackupsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDatabaseAutoBackupsSettings`: GetDatabaseAutoBackupsSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.UpdateDatabaseAutoBackupsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbId** | **int32** | ID базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseAutoBackupsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoBackup** | [**AutoBackup**](AutoBackup.md) | При значении &#x60;is_enabled&#x60;: &#x60;true&#x60;, поля &#x60;copy_count&#x60;, &#x60;creation_start_at&#x60;, &#x60;interval&#x60; являются обязательными | 

### Return type

[**GetDatabaseAutoBackupsSettings200Response**](GetDatabaseAutoBackupsSettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabaseCluster

> CreateDatabaseCluster201Response UpdateDatabaseCluster(ctx, dbClusterId).UpdateCluster(updateCluster).Execute()

Изменение кластера базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    updateCluster := *openapiclient.NewUpdateCluster() // UpdateCluster | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.UpdateDatabaseCluster(context.Background(), dbClusterId).UpdateCluster(updateCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.UpdateDatabaseCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDatabaseCluster`: CreateDatabaseCluster201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.UpdateDatabaseCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCluster** | [**UpdateCluster**](UpdateCluster.md) |  | 

### Return type

[**CreateDatabaseCluster201Response**](CreateDatabaseCluster201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabaseInstance

> CreateDatabaseInstance201Response UpdateDatabaseInstance(ctx, dbClusterId).UpdateInstance(updateInstance).Execute()

Изменение инстанса базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    updateInstance := *openapiclient.NewUpdateInstance() // UpdateInstance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.UpdateDatabaseInstance(context.Background(), dbClusterId).UpdateInstance(updateInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.UpdateDatabaseInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDatabaseInstance`: CreateDatabaseInstance201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.UpdateDatabaseInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInstance** | [**UpdateInstance**](UpdateInstance.md) |  | 

### Return type

[**CreateDatabaseInstance201Response**](CreateDatabaseInstance201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabaseUser

> CreateDatabaseUser201Response UpdateDatabaseUser(ctx, dbClusterId, adminId).UpdateAdmin(updateAdmin).Execute()

Изменение пользователя базы данных



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
    dbClusterId := int32(56) // int32 | ID кластера базы данных
    adminId := int32(56) // int32 | ID пользователя базы данных
    updateAdmin := *openapiclient.NewUpdateAdmin() // UpdateAdmin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesAPI.UpdateDatabaseUser(context.Background(), dbClusterId, adminId).UpdateAdmin(updateAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.UpdateDatabaseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDatabaseUser`: CreateDatabaseUser201Response
    fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.UpdateDatabaseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dbClusterId** | **int32** | ID кластера базы данных | 
**adminId** | **int32** | ID пользователя базы данных | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAdmin** | [**UpdateAdmin**](UpdateAdmin.md) |  | 

### Return type

[**CreateDatabaseUser201Response**](CreateDatabaseUser201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

