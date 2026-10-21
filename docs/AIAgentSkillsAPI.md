# \AIAgentSkillsAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyAgentSkillCatalogUpdate**](AIAgentSkillsAPI.md#ApplyAgentSkillCatalogUpdate) | **Post** /api/v1/cloud-ai/skills/{skill_id}/apply-catalog-update | Применение обновления навыка из каталога
[**ConnectAgentSkillsToAgent**](AIAgentSkillsAPI.md#ConnectAgentSkillsToAgent) | **Post** /api/v1/cloud-ai/agents/{agent_id}/skills/connect | Подключение нескольких навыков к агенту
[**CreateAgentSkill**](AIAgentSkillsAPI.md#CreateAgentSkill) | **Post** /api/v1/cloud-ai/skills | Создание навыка
[**DeleteAgentSkill**](AIAgentSkillsAPI.md#DeleteAgentSkill) | **Delete** /api/v1/cloud-ai/skills/{skill_id} | Удаление навыка
[**DisconnectAgentSkillFromAgent**](AIAgentSkillsAPI.md#DisconnectAgentSkillFromAgent) | **Delete** /api/v1/cloud-ai/agents/{agent_id}/skills/{skill_id} | Отключение навыка от агента
[**DuplicateAgentSkill**](AIAgentSkillsAPI.md#DuplicateAgentSkill) | **Post** /api/v1/cloud-ai/skills/{skill_id}/duplicate | Копирование навыка
[**GetAgentSkill**](AIAgentSkillsAPI.md#GetAgentSkill) | **Get** /api/v1/cloud-ai/skills/{skill_id} | Получение навыка
[**GetAgentSkillCatalog**](AIAgentSkillsAPI.md#GetAgentSkillCatalog) | **Get** /api/v1/cloud-ai/skill-catalog | Получение каталога навыков
[**GetAgentSkillLibrary**](AIAgentSkillsAPI.md#GetAgentSkillLibrary) | **Get** /api/v1/cloud-ai/skills | Получение списка навыков
[**GetAgentSkillVersion**](AIAgentSkillsAPI.md#GetAgentSkillVersion) | **Get** /api/v1/cloud-ai/skills/{skill_id}/versions/{version} | Получение версии навыка
[**GetAgentSkillVersions**](AIAgentSkillsAPI.md#GetAgentSkillVersions) | **Get** /api/v1/cloud-ai/skills/{skill_id}/versions | Получение истории версий навыка
[**GetAgentSkillsForAgent**](AIAgentSkillsAPI.md#GetAgentSkillsForAgent) | **Get** /api/v1/cloud-ai/agents/{agent_id}/skills | Получение навыков агента
[**InstallAgentSkillFromCatalog**](AIAgentSkillsAPI.md#InstallAgentSkillFromCatalog) | **Post** /api/v1/cloud-ai/skill-catalog/{catalog_item_id}/install | Установка навыка из каталога
[**InstallAndConnectAgentSkill**](AIAgentSkillsAPI.md#InstallAndConnectAgentSkill) | **Post** /api/v1/cloud-ai/agents/{agent_id}/skills/install-and-connect | Установка навыка из каталога и включение у агента
[**RollbackAgentSkillVersion**](AIAgentSkillsAPI.md#RollbackAgentSkillVersion) | **Post** /api/v1/cloud-ai/skills/{skill_id}/versions/{version}/rollback | Переключение навыка на версию из истории
[**SetAgentSkillForAgent**](AIAgentSkillsAPI.md#SetAgentSkillForAgent) | **Put** /api/v1/cloud-ai/agents/{agent_id}/skills/{skill_id} | Включение и выключение навыка у агента
[**SyncAgentSkillAgentsVersions**](AIAgentSkillsAPI.md#SyncAgentSkillAgentsVersions) | **Post** /api/v1/cloud-ai/skills/{skill_id}/agents/sync-versions | Перевод всех агентов навыка на текущую версию
[**SyncAgentSkillVersionsForAgent**](AIAgentSkillsAPI.md#SyncAgentSkillVersionsForAgent) | **Post** /api/v1/cloud-ai/agents/{agent_id}/skills/sync-versions | Перевод всех навыков агента на текущие версии
[**UpdateAgentSkill**](AIAgentSkillsAPI.md#UpdateAgentSkill) | **Patch** /api/v1/cloud-ai/skills/{skill_id} | Изменение навыка
[**UpdateAgentSkillVersionForAgent**](AIAgentSkillsAPI.md#UpdateAgentSkillVersionForAgent) | **Post** /api/v1/cloud-ai/agents/{agent_id}/skills/{skill_id}/update-version | Перевод агента на текущую версию навыка



## ApplyAgentSkillCatalogUpdate

> CreateAgentSkill201Response ApplyAgentSkillCatalogUpdate(ctx, skillId).Execute()

Применение обновления навыка из каталога



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
    skillId := int32(12) // int32 | ID навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.ApplyAgentSkillCatalogUpdate(context.Background(), skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.ApplyAgentSkillCatalogUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyAgentSkillCatalogUpdate`: CreateAgentSkill201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.ApplyAgentSkillCatalogUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyAgentSkillCatalogUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAgentSkill201Response**](CreateAgentSkill201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectAgentSkillsToAgent

> GetAgentSkillsForAgent200Response ConnectAgentSkillsToAgent(ctx, agentId).ConnectSkills(connectSkills).Execute()

Подключение нескольких навыков к агенту



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
    agentId := int32(113191) // int32 | ID агента
    connectSkills := *openapiclient.NewConnectSkills() // ConnectSkills | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.ConnectAgentSkillsToAgent(context.Background(), agentId).ConnectSkills(connectSkills).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.ConnectAgentSkillsToAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectAgentSkillsToAgent`: GetAgentSkillsForAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.ConnectAgentSkillsToAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectAgentSkillsToAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectSkills** | [**ConnectSkills**](ConnectSkills.md) |  | 

### Return type

[**GetAgentSkillsForAgent200Response**](GetAgentSkillsForAgent200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgentSkill

> CreateAgentSkill201Response CreateAgentSkill(ctx).CreateSkill(createSkill).Execute()

Создание навыка



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
    createSkill := *openapiclient.NewCreateSkill("Возврат товара", "Клиент спрашивает про возврат, обмен, отказ от заказа или сроки возврата денег", "Товар надлежащего качества принимается в течение 14 дней...") // CreateSkill | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.CreateAgentSkill(context.Background()).CreateSkill(createSkill).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.CreateAgentSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentSkill`: CreateAgentSkill201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.CreateAgentSkill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSkill** | [**CreateSkill**](CreateSkill.md) |  | 

### Return type

[**CreateAgentSkill201Response**](CreateAgentSkill201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentSkill

> DeleteAgentSkill(ctx, skillId).Execute()

Удаление навыка



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
    skillId := int32(12) // int32 | ID навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentSkillsAPI.DeleteAgentSkill(context.Background(), skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.DeleteAgentSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentSkillRequest struct via the builder pattern


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


## DisconnectAgentSkillFromAgent

> DisconnectAgentSkillFromAgent(ctx, agentId, skillId).Execute()

Отключение навыка от агента



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
    agentId := int32(113191) // int32 | ID агента
    skillId := int32(12) // int32 | ID навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AIAgentSkillsAPI.DisconnectAgentSkillFromAgent(context.Background(), agentId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.DisconnectAgentSkillFromAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectAgentSkillFromAgentRequest struct via the builder pattern


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


## DuplicateAgentSkill

> CreateAgentSkill201Response DuplicateAgentSkill(ctx, skillId).Execute()

Копирование навыка



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
    skillId := int32(12) // int32 | ID навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.DuplicateAgentSkill(context.Background(), skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.DuplicateAgentSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuplicateAgentSkill`: CreateAgentSkill201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.DuplicateAgentSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateAgentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAgentSkill201Response**](CreateAgentSkill201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSkill

> GetAgentSkill200Response GetAgentSkill(ctx, skillId).Execute()

Получение навыка



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
    skillId := int32(12) // int32 | ID навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.GetAgentSkill(context.Background(), skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.GetAgentSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentSkill`: GetAgentSkill200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.GetAgentSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentSkill200Response**](GetAgentSkill200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSkillCatalog

> GetAgentSkillCatalog200Response GetAgentSkillCatalog(ctx).Execute()

Получение каталога навыков



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
    resp, r, err := apiClient.AIAgentSkillsAPI.GetAgentSkillCatalog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.GetAgentSkillCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentSkillCatalog`: GetAgentSkillCatalog200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.GetAgentSkillCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSkillCatalogRequest struct via the builder pattern


### Return type

[**GetAgentSkillCatalog200Response**](GetAgentSkillCatalog200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSkillLibrary

> GetAgentSkillLibrary200Response GetAgentSkillLibrary(ctx).Search(search).Execute()

Получение списка навыков



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
    search := "возврат" // string | Поиск по названию навыка и по тексту инструкции (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.GetAgentSkillLibrary(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.GetAgentSkillLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentSkillLibrary`: GetAgentSkillLibrary200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.GetAgentSkillLibrary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSkillLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Поиск по названию навыка и по тексту инструкции | 

### Return type

[**GetAgentSkillLibrary200Response**](GetAgentSkillLibrary200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSkillVersion

> GetAgentSkillVersion200Response GetAgentSkillVersion(ctx, skillId, version).Execute()

Получение версии навыка



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
    skillId := int32(12) // int32 | ID навыка
    version := int32(2) // int32 | Номер версии навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.GetAgentSkillVersion(context.Background(), skillId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.GetAgentSkillVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentSkillVersion`: GetAgentSkillVersion200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.GetAgentSkillVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 
**version** | **int32** | Номер версии навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSkillVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAgentSkillVersion200Response**](GetAgentSkillVersion200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSkillVersions

> GetAgentSkillVersions200Response GetAgentSkillVersions(ctx, skillId).Execute()

Получение истории версий навыка



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
    skillId := int32(12) // int32 | ID навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.GetAgentSkillVersions(context.Background(), skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.GetAgentSkillVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentSkillVersions`: GetAgentSkillVersions200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.GetAgentSkillVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSkillVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentSkillVersions200Response**](GetAgentSkillVersions200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentSkillsForAgent

> GetAgentSkillsForAgent200Response GetAgentSkillsForAgent(ctx, agentId).Execute()

Получение навыков агента



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
    agentId := int32(113191) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.GetAgentSkillsForAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.GetAgentSkillsForAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentSkillsForAgent`: GetAgentSkillsForAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.GetAgentSkillsForAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentSkillsForAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentSkillsForAgent200Response**](GetAgentSkillsForAgent200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallAgentSkillFromCatalog

> InstallAgentSkillFromCatalog200Response InstallAgentSkillFromCatalog(ctx, catalogItemId).Execute()

Установка навыка из каталога



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
    catalogItemId := int32(3) // int32 | ID навыка в каталоге

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.InstallAgentSkillFromCatalog(context.Background(), catalogItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.InstallAgentSkillFromCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallAgentSkillFromCatalog`: InstallAgentSkillFromCatalog200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.InstallAgentSkillFromCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogItemId** | **int32** | ID навыка в каталоге | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallAgentSkillFromCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstallAgentSkillFromCatalog200Response**](InstallAgentSkillFromCatalog200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallAndConnectAgentSkill

> InstallAndConnectAgentSkill200Response InstallAndConnectAgentSkill(ctx, agentId).InstallAndConnectSkill(installAndConnectSkill).Execute()

Установка навыка из каталога и включение у агента



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
    agentId := int32(113191) // int32 | ID агента
    installAndConnectSkill := *openapiclient.NewInstallAndConnectSkill(int32(3)) // InstallAndConnectSkill | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.InstallAndConnectAgentSkill(context.Background(), agentId).InstallAndConnectSkill(installAndConnectSkill).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.InstallAndConnectAgentSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallAndConnectAgentSkill`: InstallAndConnectAgentSkill200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.InstallAndConnectAgentSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallAndConnectAgentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **installAndConnectSkill** | [**InstallAndConnectSkill**](InstallAndConnectSkill.md) |  | 

### Return type

[**InstallAndConnectAgentSkill200Response**](InstallAndConnectAgentSkill200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackAgentSkillVersion

> CreateAgentSkill201Response RollbackAgentSkillVersion(ctx, skillId, version).Execute()

Переключение навыка на версию из истории



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
    skillId := int32(12) // int32 | ID навыка
    version := int32(2) // int32 | Номер версии навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.RollbackAgentSkillVersion(context.Background(), skillId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.RollbackAgentSkillVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackAgentSkillVersion`: CreateAgentSkill201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.RollbackAgentSkillVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 
**version** | **int32** | Номер версии навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackAgentSkillVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateAgentSkill201Response**](CreateAgentSkill201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAgentSkillForAgent

> SetAgentSkillForAgent200Response SetAgentSkillForAgent(ctx, agentId, skillId).SetAgentSkill(setAgentSkill).Execute()

Включение и выключение навыка у агента



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
    agentId := int32(113191) // int32 | ID агента
    skillId := int32(12) // int32 | ID навыка
    setAgentSkill := *openapiclient.NewSetAgentSkill(true) // SetAgentSkill | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.SetAgentSkillForAgent(context.Background(), agentId, skillId).SetAgentSkill(setAgentSkill).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.SetAgentSkillForAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAgentSkillForAgent`: SetAgentSkillForAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.SetAgentSkillForAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAgentSkillForAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setAgentSkill** | [**SetAgentSkill**](SetAgentSkill.md) |  | 

### Return type

[**SetAgentSkillForAgent200Response**](SetAgentSkillForAgent200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAgentSkillAgentsVersions

> SyncAgentSkillAgentsVersions200Response SyncAgentSkillAgentsVersions(ctx, skillId).Execute()

Перевод всех агентов навыка на текущую версию



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
    skillId := int32(12) // int32 | ID навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.SyncAgentSkillAgentsVersions(context.Background(), skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.SyncAgentSkillAgentsVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncAgentSkillAgentsVersions`: SyncAgentSkillAgentsVersions200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.SyncAgentSkillAgentsVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAgentSkillAgentsVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncAgentSkillAgentsVersions200Response**](SyncAgentSkillAgentsVersions200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAgentSkillVersionsForAgent

> SyncAgentSkillVersionsForAgent200Response SyncAgentSkillVersionsForAgent(ctx, agentId).Execute()

Перевод всех навыков агента на текущие версии



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
    agentId := int32(113191) // int32 | ID агента

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.SyncAgentSkillVersionsForAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.SyncAgentSkillVersionsForAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncAgentSkillVersionsForAgent`: SyncAgentSkillVersionsForAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.SyncAgentSkillVersionsForAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncAgentSkillVersionsForAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncAgentSkillVersionsForAgent200Response**](SyncAgentSkillVersionsForAgent200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentSkill

> CreateAgentSkill201Response UpdateAgentSkill(ctx, skillId).UpdateSkill(updateSkill).Execute()

Изменение навыка



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
    skillId := int32(12) // int32 | ID навыка
    updateSkill := *openapiclient.NewUpdateSkill() // UpdateSkill | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.UpdateAgentSkill(context.Background(), skillId).UpdateSkill(updateSkill).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.UpdateAgentSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentSkill`: CreateAgentSkill201Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.UpdateAgentSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSkill** | [**UpdateSkill**](UpdateSkill.md) |  | 

### Return type

[**CreateAgentSkill201Response**](CreateAgentSkill201Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentSkillVersionForAgent

> SetAgentSkillForAgent200Response UpdateAgentSkillVersionForAgent(ctx, agentId, skillId).Execute()

Перевод агента на текущую версию навыка



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
    agentId := int32(113191) // int32 | ID агента
    skillId := int32(12) // int32 | ID навыка

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIAgentSkillsAPI.UpdateAgentSkillVersionForAgent(context.Background(), agentId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIAgentSkillsAPI.UpdateAgentSkillVersionForAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentSkillVersionForAgent`: SetAgentSkillForAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `AIAgentSkillsAPI.UpdateAgentSkillVersionForAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **int32** | ID агента | 
**skillId** | **int32** | ID навыка | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentSkillVersionForAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SetAgentSkillForAgent200Response**](SetAgentSkillForAgent200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

