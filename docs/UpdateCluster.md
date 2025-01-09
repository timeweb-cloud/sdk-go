# UpdateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Название кластера базы данных. | [optional] 
**PresetId** | Pointer to **int32** | ID тарифа. | [optional] 
**Description** | Pointer to **string** | Описание кластера базы данных | [optional] 
**IsEnabledPublicNetwork** | Pointer to **bool** | Доступность публичного IP-адреса | [optional] 

## Methods

### NewUpdateCluster

`func NewUpdateCluster() *UpdateCluster`

NewUpdateCluster instantiates a new UpdateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterWithDefaults

`func NewUpdateClusterWithDefaults() *UpdateCluster`

NewUpdateClusterWithDefaults instantiates a new UpdateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresetId

`func (o *UpdateCluster) GetPresetId() int32`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *UpdateCluster) GetPresetIdOk() (*int32, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *UpdateCluster) SetPresetId(v int32)`

SetPresetId sets PresetId field to given value.

### HasPresetId

`func (o *UpdateCluster) HasPresetId() bool`

HasPresetId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabledPublicNetwork

`func (o *UpdateCluster) GetIsEnabledPublicNetwork() bool`

GetIsEnabledPublicNetwork returns the IsEnabledPublicNetwork field if non-nil, zero value otherwise.

### GetIsEnabledPublicNetworkOk

`func (o *UpdateCluster) GetIsEnabledPublicNetworkOk() (*bool, bool)`

GetIsEnabledPublicNetworkOk returns a tuple with the IsEnabledPublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledPublicNetwork

`func (o *UpdateCluster) SetIsEnabledPublicNetwork(v bool)`

SetIsEnabledPublicNetwork sets IsEnabledPublicNetwork field to given value.

### HasIsEnabledPublicNetwork

`func (o *UpdateCluster) HasIsEnabledPublicNetwork() bool`

HasIsEnabledPublicNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


