# Preset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID тарифа. Указывается в поле &#x60;preset_id&#x60; при создании и изменении ресурса. | 
**ServiceName** | **string** | Название тарифа. | 
**Description** | **string** | Описание тарифа. | 
**DescriptionShort** | **string** | Краткое описание тарифа. | 
**Cost** | **float32** | Абонентская плата за тариф в рублях. | 
**TrafficRateCost** | **float32** | Стоимость единицы исходящего трафика по тарифу аккаунта, в рублях. | 
**RequestRateCost** | **float32** | Стоимость единицы запросов по тарифу аккаунта, в рублях. | 

## Methods

### NewPreset

`func NewPreset(id int32, serviceName string, description string, descriptionShort string, cost float32, trafficRateCost float32, requestRateCost float32, ) *Preset`

NewPreset instantiates a new Preset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresetWithDefaults

`func NewPresetWithDefaults() *Preset`

NewPresetWithDefaults instantiates a new Preset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Preset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Preset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Preset) SetId(v int32)`

SetId sets Id field to given value.


### GetServiceName

`func (o *Preset) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Preset) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Preset) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetDescription

`func (o *Preset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Preset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Preset) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionShort

`func (o *Preset) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *Preset) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *Preset) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetCost

`func (o *Preset) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Preset) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Preset) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetTrafficRateCost

`func (o *Preset) GetTrafficRateCost() float32`

GetTrafficRateCost returns the TrafficRateCost field if non-nil, zero value otherwise.

### GetTrafficRateCostOk

`func (o *Preset) GetTrafficRateCostOk() (*float32, bool)`

GetTrafficRateCostOk returns a tuple with the TrafficRateCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficRateCost

`func (o *Preset) SetTrafficRateCost(v float32)`

SetTrafficRateCost sets TrafficRateCost field to given value.


### GetRequestRateCost

`func (o *Preset) GetRequestRateCost() float32`

GetRequestRateCost returns the RequestRateCost field if non-nil, zero value otherwise.

### GetRequestRateCostOk

`func (o *Preset) GetRequestRateCostOk() (*float32, bool)`

GetRequestRateCostOk returns a tuple with the RequestRateCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRateCost

`func (o *Preset) SetRequestRateCost(v float32)`

SetRequestRateCost sets RequestRateCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


