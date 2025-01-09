# DedicatedServerAdditionalService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID дополнительной услуги для выделенного сервера. | 
**Price** | **float32** | Стоимость (в рублях) дополнительной услуги для выделенного сервера. | 
**Period** | **string** | Период оплаты. | 
**Description** | **string** | Описание дополнительной услуги выделенного сервера. | 
**ShortDescription** | **string** | Краткое описание дополнительной услуги выделенного сервера. | 
**Name** | **string** | Уникально имя дополнительной услуги выделенного сервера. | 

## Methods

### NewDedicatedServerAdditionalService

`func NewDedicatedServerAdditionalService(id float32, price float32, period string, description string, shortDescription string, name string, ) *DedicatedServerAdditionalService`

NewDedicatedServerAdditionalService instantiates a new DedicatedServerAdditionalService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedServerAdditionalServiceWithDefaults

`func NewDedicatedServerAdditionalServiceWithDefaults() *DedicatedServerAdditionalService`

NewDedicatedServerAdditionalServiceWithDefaults instantiates a new DedicatedServerAdditionalService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DedicatedServerAdditionalService) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DedicatedServerAdditionalService) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DedicatedServerAdditionalService) SetId(v float32)`

SetId sets Id field to given value.


### GetPrice

`func (o *DedicatedServerAdditionalService) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DedicatedServerAdditionalService) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DedicatedServerAdditionalService) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetPeriod

`func (o *DedicatedServerAdditionalService) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DedicatedServerAdditionalService) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DedicatedServerAdditionalService) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetDescription

`func (o *DedicatedServerAdditionalService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DedicatedServerAdditionalService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DedicatedServerAdditionalService) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetShortDescription

`func (o *DedicatedServerAdditionalService) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *DedicatedServerAdditionalService) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *DedicatedServerAdditionalService) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.


### GetName

`func (o *DedicatedServerAdditionalService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DedicatedServerAdditionalService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DedicatedServerAdditionalService) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


