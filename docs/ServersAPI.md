# \ServersAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServerIP**](ServersAPI.md#AddServerIP) | **Post** /api/v1/servers/{server_id}/ips | Добавление IP-адреса сервера
[**CloneServer**](ServersAPI.md#CloneServer) | **Post** /api/v1/servers/{server_id}/clone | Клонирование сервера
[**CreateServer**](ServersAPI.md#CreateServer) | **Post** /api/v1/servers | Создание сервера
[**CreateServerDisk**](ServersAPI.md#CreateServerDisk) | **Post** /api/v1/servers/{server_id}/disks | Создание диска сервера
[**CreateServerDiskBackup**](ServersAPI.md#CreateServerDiskBackup) | **Post** /api/v1/servers/{server_id}/disks/{disk_id}/backups | Создание бэкапа диска сервера
[**DeleteServer**](ServersAPI.md#DeleteServer) | **Delete** /api/v1/servers/{server_id} | Удаление сервера
[**DeleteServerDisk**](ServersAPI.md#DeleteServerDisk) | **Delete** /api/v1/servers/{server_id}/disks/{disk_id} | Удаление диска сервера
[**DeleteServerDiskBackup**](ServersAPI.md#DeleteServerDiskBackup) | **Delete** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Удаление бэкапа диска сервера
[**DeleteServerIP**](ServersAPI.md#DeleteServerIP) | **Delete** /api/v1/servers/{server_id}/ips | Удаление IP-адреса сервера
[**GetConfigurators**](ServersAPI.md#GetConfigurators) | **Get** /api/v1/configurator/servers | Получение списка конфигураторов серверов
[**GetOsList**](ServersAPI.md#GetOsList) | **Get** /api/v1/os/servers | Получение списка операционных систем
[**GetServer**](ServersAPI.md#GetServer) | **Get** /api/v1/servers/{server_id} | Получение сервера
[**GetServerDisk**](ServersAPI.md#GetServerDisk) | **Get** /api/v1/servers/{server_id}/disks/{disk_id} | Получение диска сервера
[**GetServerDiskAutoBackupSettings**](ServersAPI.md#GetServerDiskAutoBackupSettings) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/auto-backups | Получить настройки автобэкапов диска сервера
[**GetServerDiskBackup**](ServersAPI.md#GetServerDiskBackup) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Получение бэкапа диска сервера
[**GetServerDiskBackups**](ServersAPI.md#GetServerDiskBackups) | **Get** /api/v1/servers/{server_id}/disks/{disk_id}/backups | Получение списка бэкапов диска сервера
[**GetServerDisks**](ServersAPI.md#GetServerDisks) | **Get** /api/v1/servers/{server_id}/disks | Получение списка дисков сервера
[**GetServerIPs**](ServersAPI.md#GetServerIPs) | **Get** /api/v1/servers/{server_id}/ips | Получение списка IP-адресов сервера
[**GetServerLogs**](ServersAPI.md#GetServerLogs) | **Get** /api/v1/servers/{server_id}/logs | Получение списка логов сервера
[**GetServerStatistics**](ServersAPI.md#GetServerStatistics) | **Get** /api/v1/servers/{server_id}/statistics | Получение статистики сервера
[**GetServers**](ServersAPI.md#GetServers) | **Get** /api/v1/servers | Получение списка серверов
[**GetServersPresets**](ServersAPI.md#GetServersPresets) | **Get** /api/v1/presets/servers | Получение списка тарифов серверов
[**GetSoftware**](ServersAPI.md#GetSoftware) | **Get** /api/v1/software/servers | Получение списка ПО из маркетплейса
[**HardShutdownServer**](ServersAPI.md#HardShutdownServer) | **Post** /api/v1/servers/{server_id}/hard-shutdown | Принудительное выключение сервера
[**ImageUnmountAndServerReload**](ServersAPI.md#ImageUnmountAndServerReload) | **Post** /api/v1/servers/{server_id}/image-unmount | Отмонтирование ISO образа и перезагрузка сервера
[**PerformActionOnBackup**](ServersAPI.md#PerformActionOnBackup) | **Post** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id}/action | Выполнение действия над бэкапом диска сервера
[**PerformActionOnServer**](ServersAPI.md#PerformActionOnServer) | **Post** /api/v1/servers/{server_id}/action | Выполнение действия над сервером
[**RebootServer**](ServersAPI.md#RebootServer) | **Post** /api/v1/servers/{server_id}/reboot | Перезагрузка сервера
[**ResetServerPassword**](ServersAPI.md#ResetServerPassword) | **Post** /api/v1/servers/{server_id}/reset-password | Сброс пароля сервера
[**ShutdownServer**](ServersAPI.md#ShutdownServer) | **Post** /api/v1/servers/{server_id}/shutdown | Выключение сервера
[**StartServer**](ServersAPI.md#StartServer) | **Post** /api/v1/servers/{server_id}/start | Запуск сервера
[**UpdateServer**](ServersAPI.md#UpdateServer) | **Patch** /api/v1/servers/{server_id} | Изменение сервера
[**UpdateServerDisk**](ServersAPI.md#UpdateServerDisk) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id} | Изменение параметров диска сервера
[**UpdateServerDiskAutoBackupSettings**](ServersAPI.md#UpdateServerDiskAutoBackupSettings) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id}/auto-backups | Изменение настроек автобэкапов диска сервера
[**UpdateServerDiskBackup**](ServersAPI.md#UpdateServerDiskBackup) | **Patch** /api/v1/servers/{server_id}/disks/{disk_id}/backups/{backup_id} | Изменение бэкапа диска сервера
[**UpdateServerIP**](ServersAPI.md#UpdateServerIP) | **Patch** /api/v1/servers/{server_id}/ips | Изменение IP-адреса сервера
[**UpdateServerNAT**](ServersAPI.md#UpdateServerNAT) | **Patch** /api/v1/servers/{server_id}/local-networks/nat-mode | Изменение правил маршрутизации трафика сервера (NAT)
[**UpdateServerOSBootMode**](ServersAPI.md#UpdateServerOSBootMode) | **Post** /api/v1/servers/{server_id}/boot-mode | Выбор типа загрузки операционной системы сервера



## AddServerIP

> AddServerIP201Response AddServerIP(ctx, serverId).AddServerIPRequest(addServerIPRequest).Execute()

Добавление IP-адреса сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    addServerIPRequest := *openapiclient.NewAddServerIPRequest("ipv6") // AddServerIPRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.AddServerIP(context.Background(), serverId).AddServerIPRequest(addServerIPRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.AddServerIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServerIP`: AddServerIP201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.AddServerIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServerIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addServerIPRequest** | [**AddServerIPRequest**](AddServerIPRequest.md) |  | 

### Return type

[**AddServerIP201Response**](AddServerIP201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneServer

> CreateServer201Response CloneServer(ctx, serverId).Execute()

Клонирование сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.CloneServer(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.CloneServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneServer`: CreateServer201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.CloneServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateServer201Response**](CreateServer201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServer

> CreateServer201Response CreateServer(ctx).CreateServer(createServer).Execute()

Создание сервера



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
    createServer := *openapiclient.NewCreateServer("name") // CreateServer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.CreateServer(context.Background()).CreateServer(createServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.CreateServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServer`: CreateServer201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.CreateServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServer** | [**CreateServer**](CreateServer.md) |  | 

### Return type

[**CreateServer201Response**](CreateServer201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerDisk

> CreateServerDisk201Response CreateServerDisk(ctx, serverId).CreateServerDiskRequest(createServerDiskRequest).Execute()

Создание диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    createServerDiskRequest := *openapiclient.NewCreateServerDiskRequest(float32(10240)) // CreateServerDiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.CreateServerDisk(context.Background(), serverId).CreateServerDiskRequest(createServerDiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.CreateServerDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServerDisk`: CreateServerDisk201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.CreateServerDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServerDiskRequest** | [**CreateServerDiskRequest**](CreateServerDiskRequest.md) |  | 

### Return type

[**CreateServerDisk201Response**](CreateServerDisk201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServerDiskBackup

> CreateServerDiskBackup201Response CreateServerDiskBackup(ctx, serverId, diskId).CreateServerDiskBackupRequest(createServerDiskBackupRequest).Execute()

Создание бэкапа диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.
    createServerDiskBackupRequest := *openapiclient.NewCreateServerDiskBackupRequest() // CreateServerDiskBackupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.CreateServerDiskBackup(context.Background(), serverId, diskId).CreateServerDiskBackupRequest(createServerDiskBackupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.CreateServerDiskBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServerDiskBackup`: CreateServerDiskBackup201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.CreateServerDiskBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerDiskBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createServerDiskBackupRequest** | [**CreateServerDiskBackupRequest**](CreateServerDiskBackupRequest.md) |  | 

### Return type

[**CreateServerDiskBackup201Response**](CreateServerDiskBackup201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServer

> DeleteServer200Response DeleteServer(ctx, serverId).Hash(hash).Code(code).Execute()

Удаление сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    hash := "15095f25-aac3-4d60-a788-96cb5136f186" // string | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. (optional)
    code := "0000" // string | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена `is_able_to_delete` установлен в значение `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.DeleteServer(context.Background(), serverId).Hash(hash).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.DeleteServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServer`: DeleteServer200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.DeleteServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hash** | **string** | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. | 
 **code** | **string** | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена &#x60;is_able_to_delete&#x60; установлен в значение &#x60;true&#x60; | 

### Return type

[**DeleteServer200Response**](DeleteServer200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerDisk

> DeleteServerDisk(ctx, serverId, diskId).Execute()

Удаление диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.DeleteServerDisk(context.Background(), serverId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.DeleteServerDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerDiskRequest struct via the builder pattern


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


## DeleteServerDiskBackup

> DeleteServerDiskBackup(ctx, serverId, diskId, backupId).Execute()

Удаление бэкапа диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.
    backupId := int32(1051) // int32 | ID бэкапа сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.DeleteServerDiskBackup(context.Background(), serverId, diskId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.DeleteServerDiskBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 
**backupId** | **int32** | ID бэкапа сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerDiskBackupRequest struct via the builder pattern


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


## DeleteServerIP

> DeleteServerIP(ctx, serverId).DeleteServerIPRequest(deleteServerIPRequest).Execute()

Удаление IP-адреса сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    deleteServerIPRequest := *openapiclient.NewDeleteServerIPRequest("1.1.1.1") // DeleteServerIPRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.DeleteServerIP(context.Background(), serverId).DeleteServerIPRequest(deleteServerIPRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.DeleteServerIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteServerIPRequest** | [**DeleteServerIPRequest**](DeleteServerIPRequest.md) |  | 

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


## GetConfigurators

> GetConfigurators200Response GetConfigurators(ctx).Execute()

Получение списка конфигураторов серверов



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
    resp, r, err := apiClient.ServersAPI.GetConfigurators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetConfigurators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurators`: GetConfigurators200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetConfigurators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguratorsRequest struct via the builder pattern


### Return type

[**GetConfigurators200Response**](GetConfigurators200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsList

> GetOsList200Response GetOsList(ctx).Execute()

Получение списка операционных систем



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
    resp, r, err := apiClient.ServersAPI.GetOsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetOsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOsList`: GetOsList200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetOsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsListRequest struct via the builder pattern


### Return type

[**GetOsList200Response**](GetOsList200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServer

> CreateServer201Response GetServer(ctx, serverId).Execute()

Получение сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServer(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServer`: CreateServer201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateServer201Response**](CreateServer201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDisk

> CreateServerDisk201Response GetServerDisk(ctx, serverId, diskId).Execute()

Получение диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerDisk(context.Background(), serverId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerDisk`: CreateServerDisk201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateServerDisk201Response**](CreateServerDisk201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDiskAutoBackupSettings

> GetServerDiskAutoBackupSettings200Response GetServerDiskAutoBackupSettings(ctx, serverId, diskId).Execute()

Получить настройки автобэкапов диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerDiskAutoBackupSettings(context.Background(), serverId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerDiskAutoBackupSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerDiskAutoBackupSettings`: GetServerDiskAutoBackupSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerDiskAutoBackupSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDiskAutoBackupSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetServerDiskAutoBackupSettings200Response**](GetServerDiskAutoBackupSettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDiskBackup

> GetServerDiskBackup200Response GetServerDiskBackup(ctx, serverId, diskId, backupId).Execute()

Получение бэкапа диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.
    backupId := int32(1051) // int32 | ID бэкапа сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerDiskBackup(context.Background(), serverId, diskId, backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerDiskBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerDiskBackup`: GetServerDiskBackup200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerDiskBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 
**backupId** | **int32** | ID бэкапа сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDiskBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetServerDiskBackup200Response**](GetServerDiskBackup200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDiskBackups

> GetServerDiskBackups200Response GetServerDiskBackups(ctx, serverId, diskId).Execute()

Получение списка бэкапов диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerDiskBackups(context.Background(), serverId, diskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerDiskBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerDiskBackups`: GetServerDiskBackups200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerDiskBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDiskBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetServerDiskBackups200Response**](GetServerDiskBackups200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDisks

> GetServerDisks200Response GetServerDisks(ctx, serverId).Execute()

Получение списка дисков сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerDisks(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerDisks`: GetServerDisks200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerDisks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServerDisks200Response**](GetServerDisks200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerIPs

> GetServerIPs200Response GetServerIPs(ctx, serverId).Execute()

Получение списка IP-адресов сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerIPs(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerIPs`: GetServerIPs200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerIPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServerIPs200Response**](GetServerIPs200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerLogs

> GetServerLogs200Response GetServerLogs(ctx, serverId).Limit(limit).Offset(offset).Order(order).Execute()

Получение списка логов сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)
    order := "order_example" // string | Сортировка элементов по дате (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerLogs(context.Background(), serverId).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerLogs`: GetServerLogs200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]
 **order** | **string** | Сортировка элементов по дате | [default to &quot;asc&quot;]

### Return type

[**GetServerLogs200Response**](GetServerLogs200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerStatistics

> GetServerStatistics200Response GetServerStatistics(ctx, serverId).DateFrom(dateFrom).DateTo(dateTo).Execute()

Получение статистики сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    dateFrom := "dateFrom_example" // string | Дата начала сбора статистики. Строка в формате ISO 8061, закодированная в ASCII, пример: `2023-05-25%202023-05-25T14%3A35%3A38`
    dateTo := "dateTo_example" // string | Дата окончания сбора статистики. Строка в формате ISO 8061, закодированная в ASCII, пример: `2023-05-26%202023-05-25T14%3A35%3A38`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerStatistics(context.Background(), serverId).DateFrom(dateFrom).DateTo(dateTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerStatistics`: GetServerStatistics200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerStatisticsRequest struct via the builder pattern


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


## GetServers

> GetServers200Response GetServers(ctx).Limit(limit).Offset(offset).Execute()

Получение списка серверов



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
    resp, r, err := apiClient.ServersAPI.GetServers(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServers`: GetServers200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**GetServers200Response**](GetServers200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServersPresets

> GetServersPresets200Response GetServersPresets(ctx).Execute()

Получение списка тарифов серверов



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
    resp, r, err := apiClient.ServersAPI.GetServersPresets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServersPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServersPresets`: GetServersPresets200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServersPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServersPresetsRequest struct via the builder pattern


### Return type

[**GetServersPresets200Response**](GetServersPresets200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftware

> GetSoftware200Response GetSoftware(ctx).Execute()

Получение списка ПО из маркетплейса



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
    resp, r, err := apiClient.ServersAPI.GetSoftware(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetSoftware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftware`: GetSoftware200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetSoftware`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareRequest struct via the builder pattern


### Return type

[**GetSoftware200Response**](GetSoftware200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HardShutdownServer

> HardShutdownServer(ctx, serverId).Execute()

Принудительное выключение сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.HardShutdownServer(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.HardShutdownServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHardShutdownServerRequest struct via the builder pattern


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


## ImageUnmountAndServerReload

> ImageUnmountAndServerReload(ctx, serverId).Execute()

Отмонтирование ISO образа и перезагрузка сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.ImageUnmountAndServerReload(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ImageUnmountAndServerReload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageUnmountAndServerReloadRequest struct via the builder pattern


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


## PerformActionOnBackup

> PerformActionOnBackup(ctx, serverId, diskId, backupId).PerformActionOnBackupRequest(performActionOnBackupRequest).Execute()

Выполнение действия над бэкапом диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.
    backupId := int32(1051) // int32 | ID бэкапа сервера.
    performActionOnBackupRequest := *openapiclient.NewPerformActionOnBackupRequest("mount") // PerformActionOnBackupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.PerformActionOnBackup(context.Background(), serverId, diskId, backupId).PerformActionOnBackupRequest(performActionOnBackupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.PerformActionOnBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 
**backupId** | **int32** | ID бэкапа сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionOnBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **performActionOnBackupRequest** | [**PerformActionOnBackupRequest**](PerformActionOnBackupRequest.md) |  | 

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


## PerformActionOnServer

> PerformActionOnServer(ctx, serverId).PerformActionOnServerRequest(performActionOnServerRequest).Execute()

Выполнение действия над сервером



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    performActionOnServerRequest := *openapiclient.NewPerformActionOnServerRequest("Action_example") // PerformActionOnServerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.PerformActionOnServer(context.Background(), serverId).PerformActionOnServerRequest(performActionOnServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.PerformActionOnServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionOnServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **performActionOnServerRequest** | [**PerformActionOnServerRequest**](PerformActionOnServerRequest.md) |  | 

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


## RebootServer

> RebootServer(ctx, serverId).Execute()

Перезагрузка сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.RebootServer(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.RebootServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootServerRequest struct via the builder pattern


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


## ResetServerPassword

> ResetServerPassword(ctx, serverId).Execute()

Сброс пароля сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.ResetServerPassword(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ResetServerPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetServerPasswordRequest struct via the builder pattern


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


## ShutdownServer

> ShutdownServer(ctx, serverId).Execute()

Выключение сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.ShutdownServer(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ShutdownServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownServerRequest struct via the builder pattern


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


## StartServer

> StartServer(ctx, serverId).Execute()

Запуск сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.StartServer(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.StartServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartServerRequest struct via the builder pattern


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


## UpdateServer

> CreateServer201Response UpdateServer(ctx, serverId).UpdateServer(updateServer).Execute()

Изменение сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    updateServer := *openapiclient.NewUpdateServer() // UpdateServer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.UpdateServer(context.Background(), serverId).UpdateServer(updateServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.UpdateServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServer`: CreateServer201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.UpdateServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServer** | [**UpdateServer**](UpdateServer.md) |  | 

### Return type

[**CreateServer201Response**](CreateServer201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerDisk

> CreateServerDisk201Response UpdateServerDisk(ctx, serverId, diskId).UpdateServerDiskRequest(updateServerDiskRequest).Execute()

Изменение параметров диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.
    updateServerDiskRequest := *openapiclient.NewUpdateServerDiskRequest(float32(10240)) // UpdateServerDiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.UpdateServerDisk(context.Background(), serverId, diskId).UpdateServerDiskRequest(updateServerDiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.UpdateServerDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerDisk`: CreateServerDisk201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.UpdateServerDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateServerDiskRequest** | [**UpdateServerDiskRequest**](UpdateServerDiskRequest.md) |  | 

### Return type

[**CreateServerDisk201Response**](CreateServerDisk201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerDiskAutoBackupSettings

> GetServerDiskAutoBackupSettings200Response UpdateServerDiskAutoBackupSettings(ctx, serverId, diskId).AutoBackup(autoBackup).Execute()

Изменение настроек автобэкапов диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.
    autoBackup := *openapiclient.NewAutoBackup(true) // AutoBackup | При значении `is_enabled`: `true`, поля `copy_count`, `creation_start_at`, `interval` являются обязательными (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.UpdateServerDiskAutoBackupSettings(context.Background(), serverId, diskId).AutoBackup(autoBackup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.UpdateServerDiskAutoBackupSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerDiskAutoBackupSettings`: GetServerDiskAutoBackupSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.UpdateServerDiskAutoBackupSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerDiskAutoBackupSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoBackup** | [**AutoBackup**](AutoBackup.md) | При значении &#x60;is_enabled&#x60;: &#x60;true&#x60;, поля &#x60;copy_count&#x60;, &#x60;creation_start_at&#x60;, &#x60;interval&#x60; являются обязательными | 

### Return type

[**GetServerDiskAutoBackupSettings200Response**](GetServerDiskAutoBackupSettings200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerDiskBackup

> GetServerDiskBackup200Response UpdateServerDiskBackup(ctx, serverId, diskId, backupId).UpdateServerDiskBackupRequest(updateServerDiskBackupRequest).Execute()

Изменение бэкапа диска сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    diskId := int32(1051) // int32 | ID диска сервера.
    backupId := int32(1051) // int32 | ID бэкапа сервера.
    updateServerDiskBackupRequest := *openapiclient.NewUpdateServerDiskBackupRequest("comment") // UpdateServerDiskBackupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.UpdateServerDiskBackup(context.Background(), serverId, diskId, backupId).UpdateServerDiskBackupRequest(updateServerDiskBackupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.UpdateServerDiskBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerDiskBackup`: GetServerDiskBackup200Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.UpdateServerDiskBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 
**diskId** | **int32** | ID диска сервера. | 
**backupId** | **int32** | ID бэкапа сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerDiskBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateServerDiskBackupRequest** | [**UpdateServerDiskBackupRequest**](UpdateServerDiskBackupRequest.md) |  | 

### Return type

[**GetServerDiskBackup200Response**](GetServerDiskBackup200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerIP

> AddServerIP201Response UpdateServerIP(ctx, serverId).UpdateServerIPRequest(updateServerIPRequest).Execute()

Изменение IP-адреса сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    updateServerIPRequest := *openapiclient.NewUpdateServerIPRequest("1.1.1.1", "1197521-cl1233.tw1.ru") // UpdateServerIPRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.UpdateServerIP(context.Background(), serverId).UpdateServerIPRequest(updateServerIPRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.UpdateServerIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerIP`: AddServerIP201Response
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.UpdateServerIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerIPRequest** | [**UpdateServerIPRequest**](UpdateServerIPRequest.md) |  | 

### Return type

[**AddServerIP201Response**](AddServerIP201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerNAT

> UpdateServerNAT(ctx, serverId).UpdateServerNATRequest(updateServerNATRequest).Execute()

Изменение правил маршрутизации трафика сервера (NAT)



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    updateServerNATRequest := *openapiclient.NewUpdateServerNATRequest("no_nat") // UpdateServerNATRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.UpdateServerNAT(context.Background(), serverId).UpdateServerNATRequest(updateServerNATRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.UpdateServerNAT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerNATRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerNATRequest** | [**UpdateServerNATRequest**](UpdateServerNATRequest.md) |  | 

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


## UpdateServerOSBootMode

> UpdateServerOSBootMode(ctx, serverId).UpdateServerOSBootModeRequest(updateServerOSBootModeRequest).Execute()

Выбор типа загрузки операционной системы сервера



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
    serverId := int32(1051) // int32 | ID облачного сервера.
    updateServerOSBootModeRequest := *openapiclient.NewUpdateServerOSBootModeRequest("default") // UpdateServerOSBootModeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServersAPI.UpdateServerOSBootMode(context.Background(), serverId).UpdateServerOSBootModeRequest(updateServerOSBootModeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.UpdateServerOSBootMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **int32** | ID облачного сервера. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerOSBootModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerOSBootModeRequest** | [**UpdateServerOSBootModeRequest**](UpdateServerOSBootModeRequest.md) |  | 

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

