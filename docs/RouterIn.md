# RouterIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Имя роутера | 
**Comment** | Pointer to **string** | Описание роутера | [optional] 
**PresetId** | **int32** | ID тарифа | 
**ProjectId** | Pointer to **int32** | ID проекта | [optional] 
**Networks** | [**[]RouterInNetworksInner**](RouterInNetworksInner.md) | Подключаемые сети | 
**Ips** | Pointer to [**[]RouterInIpsInner**](RouterInIpsInner.md) | IP-адреса | [optional] 
**ParentService** | Pointer to [**RouterInParentService**](RouterInParentService.md) |  | [optional] 

## Methods

### NewRouterIn

`func NewRouterIn(name string, presetId int32, networks []RouterInNetworksInner, ) *RouterIn`

NewRouterIn instantiates a new RouterIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterInWithDefaults

`func NewRouterInWithDefaults() *RouterIn`

NewRouterInWithDefaults instantiates a new RouterIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RouterIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouterIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouterIn) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *RouterIn) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RouterIn) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RouterIn) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RouterIn) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPresetId

`func (o *RouterIn) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *RouterIn) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *RouterIn) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.


### GetProjectId

`func (o *RouterIn) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RouterIn) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RouterIn) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RouterIn) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetNetworks

`func (o *RouterIn) GetNetworks() []RouterInNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *RouterIn) GetNetworksOk() (*[]RouterInNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *RouterIn) SetNetworks(v []RouterInNetworksInner)`

SetNetworks sets Networks field to given value.


### GetIps

`func (o *RouterIn) GetIps() []RouterInIpsInner`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *RouterIn) GetIpsOk() (*[]RouterInIpsInner, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *RouterIn) SetIps(v []RouterInIpsInner)`

SetIps sets Ips field to given value.

### HasIps

`func (o *RouterIn) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetParentService

`func (o *RouterIn) GetParentService() RouterInParentService`

GetParentService returns the ParentService field if non-nil, zero value otherwise.

### GetParentServiceOk

`func (o *RouterIn) GetParentServiceOk() (*RouterInParentService, bool)`

GetParentServiceOk returns a tuple with the ParentService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentService

`func (o *RouterIn) SetParentService(v RouterInParentService)`

SetParentService sets ParentService field to given value.

### HasParentService

`func (o *RouterIn) HasParentService() bool`

HasParentService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


