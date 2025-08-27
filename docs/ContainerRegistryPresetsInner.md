# ContainerRegistryPresetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID тарифа | 
**Description** | **string** | Описание тарифа | 
**DescriptionShort** | **string** | Краткое описание тарифа | 
**Disk** | **int32** | Количество пространства в ГБ | 
**Price** | **float32** | Стоимость | 
**Location** | Pointer to **string** | Локация | [optional] 

## Methods

### NewContainerRegistryPresetsInner

`func NewContainerRegistryPresetsInner(id int32, description string, descriptionShort string, disk int32, price float32, ) *ContainerRegistryPresetsInner`

NewContainerRegistryPresetsInner instantiates a new ContainerRegistryPresetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryPresetsInnerWithDefaults

`func NewContainerRegistryPresetsInnerWithDefaults() *ContainerRegistryPresetsInner`

NewContainerRegistryPresetsInnerWithDefaults instantiates a new ContainerRegistryPresetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerRegistryPresetsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerRegistryPresetsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerRegistryPresetsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetDescription

`func (o *ContainerRegistryPresetsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerRegistryPresetsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerRegistryPresetsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionShort

`func (o *ContainerRegistryPresetsInner) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *ContainerRegistryPresetsInner) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *ContainerRegistryPresetsInner) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetDisk

`func (o *ContainerRegistryPresetsInner) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ContainerRegistryPresetsInner) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ContainerRegistryPresetsInner) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetPrice

`func (o *ContainerRegistryPresetsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ContainerRegistryPresetsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ContainerRegistryPresetsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetLocation

`func (o *ContainerRegistryPresetsInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ContainerRegistryPresetsInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ContainerRegistryPresetsInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ContainerRegistryPresetsInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


