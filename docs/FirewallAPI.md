# \FirewallAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourceToGroup**](FirewallAPI.md#AddResourceToGroup) | **Post** /api/v1/firewall/groups/{group_id}/resources/{resource_id} | Линковка ресурса в firewall group
[**CreateGroup**](FirewallAPI.md#CreateGroup) | **Post** /api/v1/firewall/groups | Создание группы правил
[**CreateGroupRule**](FirewallAPI.md#CreateGroupRule) | **Post** /api/v1/firewall/groups/{group_id}/rules | Создание firewall правила
[**DeleteGroup**](FirewallAPI.md#DeleteGroup) | **Delete** /api/v1/firewall/groups/{group_id} | Удаление группы правил
[**DeleteGroupRule**](FirewallAPI.md#DeleteGroupRule) | **Delete** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Удаление firewall правила
[**DeleteResourceFromGroup**](FirewallAPI.md#DeleteResourceFromGroup) | **Delete** /api/v1/firewall/groups/{group_id}/resources/{resource_id} | Отлинковка ресурса из firewall group
[**GetGroup**](FirewallAPI.md#GetGroup) | **Get** /api/v1/firewall/groups/{group_id} | Получение информации о группе правил
[**GetGroupResources**](FirewallAPI.md#GetGroupResources) | **Get** /api/v1/firewall/groups/{group_id}/resources | Получение слинкованных ресурсов
[**GetGroupRule**](FirewallAPI.md#GetGroupRule) | **Get** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Получение информации о правиле
[**GetGroupRules**](FirewallAPI.md#GetGroupRules) | **Get** /api/v1/firewall/groups/{group_id}/rules | Получение списка правил
[**GetGroups**](FirewallAPI.md#GetGroups) | **Get** /api/v1/firewall/groups | Получение групп правил
[**GetRulesForResource**](FirewallAPI.md#GetRulesForResource) | **Get** /api/v1/firewall/service/{resource_type}/{resource_id} | Получение групп правил для ресурса
[**UpdateGroup**](FirewallAPI.md#UpdateGroup) | **Patch** /api/v1/firewall/groups/{group_id} | Обновление группы правил
[**UpdateGroupRule**](FirewallAPI.md#UpdateGroupRule) | **Patch** /api/v1/firewall/groups/{group_id}/rules/{rule_id} | Обновление firewall правила



## AddResourceToGroup

> FirewallGroupResourceOutResponse AddResourceToGroup(ctx, groupId, resourceId).ResourceType(resourceType).Execute()

Линковка ресурса в firewall group



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
    groupId := "groupId_example" // string | ID группы правил.
    resourceId := "resourceId_example" // string | ID ресурса
    resourceType := openapiclient.ResourceType("server") // ResourceType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.AddResourceToGroup(context.Background(), groupId, resourceId).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.AddResourceToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddResourceToGroup`: FirewallGroupResourceOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.AddResourceToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 
**resourceId** | **string** | ID ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceType** | [**ResourceType**](ResourceType.md) |  | 

### Return type

[**FirewallGroupResourceOutResponse**](FirewallGroupResourceOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> FirewallGroupOutResponse CreateGroup(ctx).FirewallGroupInAPI(firewallGroupInAPI).Policy(policy).Execute()

Создание группы правил



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
    firewallGroupInAPI := *openapiclient.NewFirewallGroupInAPI("Name_example") // FirewallGroupInAPI | 
    policy := "policy_example" // string | Тип группы правил. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.CreateGroup(context.Background()).FirewallGroupInAPI(firewallGroupInAPI).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: FirewallGroupOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallGroupInAPI** | [**FirewallGroupInAPI**](FirewallGroupInAPI.md) |  | 
 **policy** | **string** | Тип группы правил. | 

### Return type

[**FirewallGroupOutResponse**](FirewallGroupOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupRule

> FirewallRuleOutResponse CreateGroupRule(ctx, groupId).FirewallRuleInAPI(firewallRuleInAPI).Execute()

Создание firewall правила



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
    groupId := "groupId_example" // string | ID группы правил.
    firewallRuleInAPI := *openapiclient.NewFirewallRuleInAPI(openapiclient.FirewallRuleDirection("ingress"), openapiclient.FirewallRuleProtocol("tcp")) // FirewallRuleInAPI | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.CreateGroupRule(context.Background(), groupId).FirewallRuleInAPI(firewallRuleInAPI).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.CreateGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupRule`: FirewallRuleOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.CreateGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRuleInAPI** | [**FirewallRuleInAPI**](FirewallRuleInAPI.md) |  | 

### Return type

[**FirewallRuleOutResponse**](FirewallRuleOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()

Удаление группы правил



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
    groupId := "groupId_example" // string | ID группы правил.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallAPI.DeleteGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## DeleteGroupRule

> DeleteGroupRule(ctx, groupId, ruleId).Execute()

Удаление firewall правила



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
    groupId := "groupId_example" // string | ID группы правил.
    ruleId := "ruleId_example" // string | ID правила

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallAPI.DeleteGroupRule(context.Background(), groupId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.DeleteGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 
**ruleId** | **string** | ID правила | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRuleRequest struct via the builder pattern


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


## DeleteResourceFromGroup

> DeleteResourceFromGroup(ctx, groupId, resourceId).ResourceType(resourceType).Execute()

Отлинковка ресурса из firewall group



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
    groupId := "groupId_example" // string | ID группы правил.
    resourceId := "resourceId_example" // string | ID ресурса
    resourceType := openapiclient.ResourceType("server") // ResourceType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallAPI.DeleteResourceFromGroup(context.Background(), groupId, resourceId).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.DeleteResourceFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 
**resourceId** | **string** | ID ресурса | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceType** | [**ResourceType**](ResourceType.md) |  | 

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


## GetGroup

> FirewallGroupOutResponse GetGroup(ctx, groupId).Execute()

Получение информации о группе правил



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
    groupId := "groupId_example" // string | ID группы правил.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.GetGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: FirewallGroupOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallGroupOutResponse**](FirewallGroupOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupResources

> FirewallGroupResourcesOutResponse GetGroupResources(ctx, groupId).Limit(limit).Offset(offset).Execute()

Получение слинкованных ресурсов



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
    groupId := "groupId_example" // string | ID группы правил.
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.GetGroupResources(context.Background(), groupId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetGroupResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupResources`: FirewallGroupResourcesOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetGroupResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**FirewallGroupResourcesOutResponse**](FirewallGroupResourcesOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRule

> FirewallRuleOutResponse GetGroupRule(ctx, ruleId, groupId).Execute()

Получение информации о правиле



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
    ruleId := "ruleId_example" // string | ID правила.
    groupId := "groupId_example" // string | ID группы правил.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.GetGroupRule(context.Background(), ruleId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupRule`: FirewallRuleOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | ID правила. | 
**groupId** | **string** | ID группы правил. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FirewallRuleOutResponse**](FirewallRuleOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRules

> FirewallRulesOutResponse GetGroupRules(ctx, groupId).Limit(limit).Offset(offset).Execute()

Получение списка правил



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
    groupId := "groupId_example" // string | ID группы правил.
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.GetGroupRules(context.Background(), groupId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetGroupRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupRules`: FirewallRulesOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**FirewallRulesOutResponse**](FirewallRulesOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> FirewallGroupsOutResponse GetGroups(ctx).Limit(limit).Offset(offset).Execute()

Получение групп правил



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
    resp, r, err := apiClient.FirewallAPI.GetGroups(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: FirewallGroupsOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**FirewallGroupsOutResponse**](FirewallGroupsOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesForResource

> FirewallGroupsOutResponse GetRulesForResource(ctx, resourceId, resourceType).Limit(limit).Offset(offset).Execute()

Получение групп правил для ресурса



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
    resourceId := "resourceId_example" // string | ID ресурса
    resourceType := openapiclient.ResourceType("server") // ResourceType | 
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение относительно начала списка. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.GetRulesForResource(context.Background(), resourceId, resourceType).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetRulesForResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRulesForResource`: FirewallGroupsOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetRulesForResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | ID ресурса | 
**resourceType** | [**ResourceType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesForResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**FirewallGroupsOutResponse**](FirewallGroupsOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> FirewallGroupOutResponse UpdateGroup(ctx, groupId).FirewallGroupInAPI(firewallGroupInAPI).Execute()

Обновление группы правил



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
    groupId := "groupId_example" // string | ID группы правил.
    firewallGroupInAPI := *openapiclient.NewFirewallGroupInAPI("Name_example") // FirewallGroupInAPI | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.UpdateGroup(context.Background(), groupId).FirewallGroupInAPI(firewallGroupInAPI).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: FirewallGroupOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallGroupInAPI** | [**FirewallGroupInAPI**](FirewallGroupInAPI.md) |  | 

### Return type

[**FirewallGroupOutResponse**](FirewallGroupOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupRule

> FirewallRuleOutResponse UpdateGroupRule(ctx, groupId, ruleId).FirewallRuleInAPI(firewallRuleInAPI).Execute()

Обновление firewall правила



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
    groupId := "groupId_example" // string | ID группы правил.
    ruleId := "ruleId_example" // string | ID правила
    firewallRuleInAPI := *openapiclient.NewFirewallRuleInAPI(openapiclient.FirewallRuleDirection("ingress"), openapiclient.FirewallRuleProtocol("tcp")) // FirewallRuleInAPI | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.UpdateGroupRule(context.Background(), groupId, ruleId).FirewallRuleInAPI(firewallRuleInAPI).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.UpdateGroupRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupRule`: FirewallRuleOutResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.UpdateGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID группы правил. | 
**ruleId** | **string** | ID правила | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firewallRuleInAPI** | [**FirewallRuleInAPI**](FirewallRuleInAPI.md) |  | 

### Return type

[**FirewallRuleOutResponse**](FirewallRuleOutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

