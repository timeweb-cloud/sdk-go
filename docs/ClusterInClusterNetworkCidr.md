# ClusterInClusterNetworkCidr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodsNetwork** | Pointer to **string** | Диапазон адресов подов. Подсеть должна принадлежать диапазонам 10.0.0.0/x, 192.168.0.0/x, 172.16.0.0/x. Максимальный размер для подов CIDR 10.0.0.0/x - /8, 192.168.0.0/x - /16, 172.16.0.0/x - /12, минимальный CIDR в этих диапазонах - /28. | [optional] 
**ServicesNetwork** | Pointer to **string** | Диапазон адресов сервисов. Подсеть должна принадлежать диапазонам 10.0.0.0/x, 192.168.0.0/x, 172.16.0.0/x. Максимальный размер CIDR для сервисов 10.0.0.0/x - /12, 192.168.0.0/x - /16, 172.16.0.0/x - /12, минимальный CIDR в этих диапазонах - /28. | [optional] 

## Methods

### NewClusterInClusterNetworkCidr

`func NewClusterInClusterNetworkCidr() *ClusterInClusterNetworkCidr`

NewClusterInClusterNetworkCidr instantiates a new ClusterInClusterNetworkCidr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInClusterNetworkCidrWithDefaults

`func NewClusterInClusterNetworkCidrWithDefaults() *ClusterInClusterNetworkCidr`

NewClusterInClusterNetworkCidrWithDefaults instantiates a new ClusterInClusterNetworkCidr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodsNetwork

`func (o *ClusterInClusterNetworkCidr) GetPodsNetwork() string`

GetPodsNetwork returns the PodsNetwork field if non-nil, zero value otherwise.

### GetPodsNetworkOk

`func (o *ClusterInClusterNetworkCidr) GetPodsNetworkOk() (*string, bool)`

GetPodsNetworkOk returns a tuple with the PodsNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsNetwork

`func (o *ClusterInClusterNetworkCidr) SetPodsNetwork(v string)`

SetPodsNetwork sets PodsNetwork field to given value.

### HasPodsNetwork

`func (o *ClusterInClusterNetworkCidr) HasPodsNetwork() bool`

HasPodsNetwork returns a boolean if a field has been set.

### GetServicesNetwork

`func (o *ClusterInClusterNetworkCidr) GetServicesNetwork() string`

GetServicesNetwork returns the ServicesNetwork field if non-nil, zero value otherwise.

### GetServicesNetworkOk

`func (o *ClusterInClusterNetworkCidr) GetServicesNetworkOk() (*string, bool)`

GetServicesNetworkOk returns a tuple with the ServicesNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesNetwork

`func (o *ClusterInClusterNetworkCidr) SetServicesNetwork(v string)`

SetServicesNetwork sets ServicesNetwork field to given value.

### HasServicesNetwork

`func (o *ClusterInClusterNetworkCidr) HasServicesNetwork() bool`

HasServicesNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


