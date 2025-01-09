# PresetsStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра тарифа хранилища. | 
**Description** | **string** | Описание тарифа. | 
**DescriptionShort** | **string** | Краткое описание тарифа. | 
**Disk** | **float32** | Описание диска хранилища. | 
**Price** | **float32** | Стоимость тарифа хранилища. | 
**Location** | **string** | Географическое расположение тарифа. | 

## Methods

### NewPresetsStorage

`func NewPresetsStorage(id float32, description string, descriptionShort string, disk float32, price float32, location string, ) *PresetsStorage`

NewPresetsStorage instantiates a new PresetsStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresetsStorageWithDefaults

`func NewPresetsStorageWithDefaults() *PresetsStorage`

NewPresetsStorageWithDefaults instantiates a new PresetsStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PresetsStorage) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PresetsStorage) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PresetsStorage) SetId(v float32)`

SetId sets Id field to given value.


### GetDescription

`func (o *PresetsStorage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PresetsStorage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PresetsStorage) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionShort

`func (o *PresetsStorage) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *PresetsStorage) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *PresetsStorage) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetDisk

`func (o *PresetsStorage) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *PresetsStorage) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *PresetsStorage) SetDisk(v float32)`

SetDisk sets Disk field to given value.


### GetPrice

`func (o *PresetsStorage) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PresetsStorage) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PresetsStorage) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetLocation

`func (o *PresetsStorage) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PresetsStorage) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PresetsStorage) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


