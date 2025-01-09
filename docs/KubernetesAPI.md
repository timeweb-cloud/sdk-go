# \KubernetesAPI

All URIs are relative to *https://api.timeweb.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](KubernetesAPI.md#CreateCluster) | **Post** /api/v1/k8s/clusters | Создание кластера
[**CreateClusterNodeGroup**](KubernetesAPI.md#CreateClusterNodeGroup) | **Post** /api/v1/k8s/clusters/{cluster_id}/groups | Создание группы нод
[**DeleteCluster**](KubernetesAPI.md#DeleteCluster) | **Delete** /api/v1/k8s/clusters/{cluster_id} | Удаление кластера
[**DeleteClusterNode**](KubernetesAPI.md#DeleteClusterNode) | **Delete** /api/v1/k8s/clusters/{cluster_id}/nodes/{node_id} | Удаление ноды
[**DeleteClusterNodeGroup**](KubernetesAPI.md#DeleteClusterNodeGroup) | **Delete** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id} | Удаление группы нод
[**GetCluster**](KubernetesAPI.md#GetCluster) | **Get** /api/v1/k8s/clusters/{cluster_id} | Получение информации о кластере
[**GetClusterKubeconfig**](KubernetesAPI.md#GetClusterKubeconfig) | **Get** /api/v1/k8s/clusters/{cluster_id}/kubeconfig | Получение файла kubeconfig
[**GetClusterNodeGroup**](KubernetesAPI.md#GetClusterNodeGroup) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id} | Получение информации о группе нод
[**GetClusterNodeGroups**](KubernetesAPI.md#GetClusterNodeGroups) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups | Получение групп нод кластера
[**GetClusterNodes**](KubernetesAPI.md#GetClusterNodes) | **Get** /api/v1/k8s/clusters/{cluster_id}/nodes | Получение списка нод
[**GetClusterNodesFromGroup**](KubernetesAPI.md#GetClusterNodesFromGroup) | **Get** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Получение списка нод, принадлежащих группе
[**GetClusterResources**](KubernetesAPI.md#GetClusterResources) | **Get** /api/v1/k8s/clusters/{cluster_id}/resources | Получение ресурсов кластера
[**GetClusters**](KubernetesAPI.md#GetClusters) | **Get** /api/v1/k8s/clusters | Получение списка кластеров
[**GetK8SNetworkDrivers**](KubernetesAPI.md#GetK8SNetworkDrivers) | **Get** /api/v1/k8s/network-drivers | Получение списка сетевых драйверов k8s
[**GetK8SVersions**](KubernetesAPI.md#GetK8SVersions) | **Get** /api/v1/k8s/k8s-versions | Получение списка версий k8s
[**GetKubernetesPresets**](KubernetesAPI.md#GetKubernetesPresets) | **Get** /api/v1/presets/k8s | Получение списка тарифов
[**IncreaseCountOfNodesInGroup**](KubernetesAPI.md#IncreaseCountOfNodesInGroup) | **Post** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Увеличение количества нод в группе на указанное количество
[**ReduceCountOfNodesInGroup**](KubernetesAPI.md#ReduceCountOfNodesInGroup) | **Delete** /api/v1/k8s/clusters/{cluster_id}/groups/{group_id}/nodes | Уменьшение количества нод в группе на указанное количество
[**UpdateCluster**](KubernetesAPI.md#UpdateCluster) | **Patch** /api/v1/k8s/clusters/{cluster_id} | Обновление информации о кластере



## CreateCluster

> ClusterResponse CreateCluster(ctx).ClusterIn(clusterIn).Execute()

Создание кластера



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
    clusterIn := *openapiclient.NewClusterIn("Name_example", "K8sVersion_example", "NetworkDriver_example", int32(123)) // ClusterIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.CreateCluster(context.Background()).ClusterIn(clusterIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterIn** | [**ClusterIn**](ClusterIn.md) |  | 

### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClusterNodeGroup

> NodeGroupResponse CreateClusterNodeGroup(ctx, clusterId).NodeGroupIn(nodeGroupIn).Execute()

Создание группы нод



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
    clusterId := int32(56) // int32 | ID кластера
    nodeGroupIn := *openapiclient.NewNodeGroupIn("Name_example", int32(123)) // NodeGroupIn | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.CreateClusterNodeGroup(context.Background(), clusterId).NodeGroupIn(nodeGroupIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.CreateClusterNodeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClusterNodeGroup`: NodeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.CreateClusterNodeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterNodeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeGroupIn** | [**NodeGroupIn**](NodeGroupIn.md) |  | 

### Return type

[**NodeGroupResponse**](NodeGroupResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> DeleteCluster200Response DeleteCluster(ctx, clusterId).Hash(hash).Code(code).Execute()

Удаление кластера



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
    clusterId := int32(56) // int32 | ID кластера
    hash := "15095f25-aac3-4d60-a788-96cb5136f186" // string | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. (optional)
    code := "0000" // string | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена `is_able_to_delete` установлен в значение `true` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.DeleteCluster(context.Background(), clusterId).Hash(hash).Code(code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCluster`: DeleteCluster200Response
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hash** | **string** | Хеш, который совместно с кодом авторизации надо отправить для удаления, если включено подтверждение удаления сервисов через Телеграм. | 
 **code** | **string** | Код подтверждения, который придет к вам в Телеграм, после запроса удаления, если включено подтверждение удаления сервисов.  При помощи API токена сервисы можно удалять без подтверждения, если параметр токена &#x60;is_able_to_delete&#x60; установлен в значение &#x60;true&#x60; | 

### Return type

[**DeleteCluster200Response**](DeleteCluster200Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterNode

> DeleteClusterNode(ctx, clusterId, nodeId).Execute()

Удаление ноды



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
    clusterId := int32(56) // int32 | ID кластера
    nodeId := int32(56) // int32 | ID группы нод

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.DeleteClusterNode(context.Background(), clusterId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.DeleteClusterNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 
**nodeId** | **int32** | ID группы нод | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterNodeRequest struct via the builder pattern


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


## DeleteClusterNodeGroup

> DeleteClusterNodeGroup(ctx, clusterId, groupId).Execute()

Удаление группы нод



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
    clusterId := int32(56) // int32 | ID кластера
    groupId := int32(56) // int32 | ID группы

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.DeleteClusterNodeGroup(context.Background(), clusterId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.DeleteClusterNodeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 
**groupId** | **int32** | ID группы | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterNodeGroupRequest struct via the builder pattern


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


## GetCluster

> ClusterResponse GetCluster(ctx, clusterId).Execute()

Получение информации о кластере



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
    clusterId := int32(56) // int32 | ID кластера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.GetCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterKubeconfig

> string GetClusterKubeconfig(ctx, clusterId).Execute()

Получение файла kubeconfig



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
    clusterId := int32(56) // int32 | ID кластера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.GetClusterKubeconfig(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetClusterKubeconfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterKubeconfig`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetClusterKubeconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNodeGroup

> NodeGroupResponse GetClusterNodeGroup(ctx, clusterId, groupId).Execute()

Получение информации о группе нод



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
    clusterId := int32(56) // int32 | ID кластера
    groupId := int32(56) // int32 | ID группы

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.GetClusterNodeGroup(context.Background(), clusterId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetClusterNodeGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNodeGroup`: NodeGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetClusterNodeGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 
**groupId** | **int32** | ID группы | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNodeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeGroupResponse**](NodeGroupResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNodeGroups

> NodeGroupsResponse GetClusterNodeGroups(ctx, clusterId).Execute()

Получение групп нод кластера



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
    clusterId := int32(56) // int32 | ID кластера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.GetClusterNodeGroups(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetClusterNodeGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNodeGroups`: NodeGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetClusterNodeGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNodeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodeGroupsResponse**](NodeGroupsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNodes

> NodesResponse GetClusterNodes(ctx, clusterId).Execute()

Получение списка нод



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
    clusterId := int32(56) // int32 | ID кластера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.GetClusterNodes(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNodes`: NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetClusterNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodesResponse**](NodesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterNodesFromGroup

> NodesResponse GetClusterNodesFromGroup(ctx, clusterId, groupId).Limit(limit).Offset(offset).Execute()

Получение списка нод, принадлежащих группе



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
    clusterId := int32(56) // int32 | ID кластера
    groupId := int32(56) // int32 | ID группы
    limit := int32(56) // int32 | Обозначает количество записей, которое необходимо вернуть. (optional) (default to 100)
    offset := int32(56) // int32 | Указывает на смещение, относительно начала списка. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.GetClusterNodesFromGroup(context.Background(), clusterId, groupId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetClusterNodesFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterNodesFromGroup`: NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetClusterNodesFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 
**groupId** | **int32** | ID группы | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterNodesFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение, относительно начала списка. | [default to 0]

### Return type

[**NodesResponse**](NodesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterResources

> ResourcesResponse GetClusterResources(ctx, clusterId).Execute()

Получение ресурсов кластера



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
    clusterId := int32(56) // int32 | ID кластера

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.GetClusterResources(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetClusterResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterResources`: ResourcesResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetClusterResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourcesResponse**](ResourcesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> ClustersResponse GetClusters(ctx).Limit(limit).Offset(offset).Execute()

Получение списка кластеров



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
    resp, r, err := apiClient.KubernetesAPI.GetClusters(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusters`: ClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Обозначает количество записей, которое необходимо вернуть. | [default to 100]
 **offset** | **int32** | Указывает на смещение относительно начала списка. | [default to 0]

### Return type

[**ClustersResponse**](ClustersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetK8SNetworkDrivers

> NetworkDriversResponse GetK8SNetworkDrivers(ctx).Execute()

Получение списка сетевых драйверов k8s



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
    resp, r, err := apiClient.KubernetesAPI.GetK8SNetworkDrivers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetK8SNetworkDrivers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetK8SNetworkDrivers`: NetworkDriversResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetK8SNetworkDrivers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetK8SNetworkDriversRequest struct via the builder pattern


### Return type

[**NetworkDriversResponse**](NetworkDriversResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetK8SVersions

> K8SVersionsResponse GetK8SVersions(ctx).Execute()

Получение списка версий k8s



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
    resp, r, err := apiClient.KubernetesAPI.GetK8SVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetK8SVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetK8SVersions`: K8SVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetK8SVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetK8SVersionsRequest struct via the builder pattern


### Return type

[**K8SVersionsResponse**](K8SVersionsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesPresets

> PresetsResponse GetKubernetesPresets(ctx).Execute()

Получение списка тарифов



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
    resp, r, err := apiClient.KubernetesAPI.GetKubernetesPresets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.GetKubernetesPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesPresets`: PresetsResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.GetKubernetesPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesPresetsRequest struct via the builder pattern


### Return type

[**PresetsResponse**](PresetsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncreaseCountOfNodesInGroup

> NodesResponse IncreaseCountOfNodesInGroup(ctx, clusterId, groupId).NodeCount(nodeCount).Execute()

Увеличение количества нод в группе на указанное количество



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
    clusterId := int32(56) // int32 | ID кластера
    groupId := int32(56) // int32 | ID группы
    nodeCount := *openapiclient.NewNodeCount(int32(2)) // NodeCount | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.IncreaseCountOfNodesInGroup(context.Background(), clusterId, groupId).NodeCount(nodeCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.IncreaseCountOfNodesInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncreaseCountOfNodesInGroup`: NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.IncreaseCountOfNodesInGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 
**groupId** | **int32** | ID группы | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncreaseCountOfNodesInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeCount** | [**NodeCount**](NodeCount.md) |  | 

### Return type

[**NodesResponse**](NodesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReduceCountOfNodesInGroup

> ReduceCountOfNodesInGroup(ctx, clusterId, groupId).NodeCount(nodeCount).Execute()

Уменьшение количества нод в группе на указанное количество



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
    clusterId := int32(56) // int32 | ID кластера
    groupId := int32(56) // int32 | ID группы
    nodeCount := *openapiclient.NewNodeCount(int32(2)) // NodeCount | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesAPI.ReduceCountOfNodesInGroup(context.Background(), clusterId, groupId).NodeCount(nodeCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.ReduceCountOfNodesInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 
**groupId** | **int32** | ID группы | 

### Other Parameters

Other parameters are passed through a pointer to a apiReduceCountOfNodesInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeCount** | [**NodeCount**](NodeCount.md) |  | 

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


## UpdateCluster

> ClusterResponse UpdateCluster(ctx, clusterId).ClusterEdit(clusterEdit).Execute()

Обновление информации о кластере



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
    clusterId := int32(56) // int32 | ID кластера
    clusterEdit := *openapiclient.NewClusterEdit() // ClusterEdit | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesAPI.UpdateCluster(context.Background(), clusterId).ClusterEdit(clusterEdit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int32** | ID кластера | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterEdit** | [**ClusterEdit**](ClusterEdit.md) |  | 

### Return type

[**ClusterResponse**](ClusterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

