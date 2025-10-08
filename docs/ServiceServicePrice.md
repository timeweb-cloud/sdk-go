# ServiceServicePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | Идентификатор сервиса | [optional] 
**Cost** | Pointer to **float32** | Стоимость сервиса | [optional] 
**Description** | Pointer to **string** | Описание сервиса | [optional] 
**Type** | Pointer to [**ServiceCostType**](ServiceCostType.md) |  | [optional] 
**NodeGroups** | Pointer to [**[]ServiceServicePriceNodeGroupsInner**](ServiceServicePriceNodeGroupsInner.md) | Группы узлов для Kubernetes кластера | [optional] 

## Methods

### NewServiceServicePrice

`func NewServiceServicePrice() *ServiceServicePrice`

NewServiceServicePrice instantiates a new ServiceServicePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceServicePriceWithDefaults

`func NewServiceServicePriceWithDefaults() *ServiceServicePrice`

NewServiceServicePriceWithDefaults instantiates a new ServiceServicePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceServicePrice) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceServicePrice) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceServicePrice) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceServicePrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCost

`func (o *ServiceServicePrice) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ServiceServicePrice) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ServiceServicePrice) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ServiceServicePrice) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceServicePrice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceServicePrice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceServicePrice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceServicePrice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ServiceServicePrice) GetType() ServiceCostType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceServicePrice) GetTypeOk() (*ServiceCostType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceServicePrice) SetType(v ServiceCostType)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceServicePrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNodeGroups

`func (o *ServiceServicePrice) GetNodeGroups() []ServiceServicePriceNodeGroupsInner`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *ServiceServicePrice) GetNodeGroupsOk() (*[]ServiceServicePriceNodeGroupsInner, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *ServiceServicePrice) SetNodeGroups(v []ServiceServicePriceNodeGroupsInner)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *ServiceServicePrice) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


