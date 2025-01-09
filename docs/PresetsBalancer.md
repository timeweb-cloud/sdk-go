# PresetsBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID для каждого экземпляра тарифа базы данных. | 
**Description** | **string** | Описание тарифа. | 
**DescriptionShort** | **string** | Краткое описание тарифа. | 
**Bandwidth** | **float32** | Пропускная способность тарифа. | 
**ReplicaCount** | **float32** | Количество реплик. | 
**RequestPerSecond** | **string** | Запросов в секунду. | 
**Price** | **float32** | Стоимость тарифа базы данных. | 
**Location** | **string** | Географическое расположение тарифа. | 

## Methods

### NewPresetsBalancer

`func NewPresetsBalancer(id float32, description string, descriptionShort string, bandwidth float32, replicaCount float32, requestPerSecond string, price float32, location string, ) *PresetsBalancer`

NewPresetsBalancer instantiates a new PresetsBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresetsBalancerWithDefaults

`func NewPresetsBalancerWithDefaults() *PresetsBalancer`

NewPresetsBalancerWithDefaults instantiates a new PresetsBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PresetsBalancer) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PresetsBalancer) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PresetsBalancer) SetId(v float32)`

SetId sets Id field to given value.


### GetDescription

`func (o *PresetsBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PresetsBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PresetsBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionShort

`func (o *PresetsBalancer) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *PresetsBalancer) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *PresetsBalancer) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.


### GetBandwidth

`func (o *PresetsBalancer) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *PresetsBalancer) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *PresetsBalancer) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.


### GetReplicaCount

`func (o *PresetsBalancer) GetReplicaCount() float32`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *PresetsBalancer) GetReplicaCountOk() (*float32, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *PresetsBalancer) SetReplicaCount(v float32)`

SetReplicaCount sets ReplicaCount field to given value.


### GetRequestPerSecond

`func (o *PresetsBalancer) GetRequestPerSecond() string`

GetRequestPerSecond returns the RequestPerSecond field if non-nil, zero value otherwise.

### GetRequestPerSecondOk

`func (o *PresetsBalancer) GetRequestPerSecondOk() (*string, bool)`

GetRequestPerSecondOk returns a tuple with the RequestPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPerSecond

`func (o *PresetsBalancer) SetRequestPerSecond(v string)`

SetRequestPerSecond sets RequestPerSecond field to given value.


### GetPrice

`func (o *PresetsBalancer) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PresetsBalancer) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PresetsBalancer) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetLocation

`func (o *PresetsBalancer) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PresetsBalancer) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PresetsBalancer) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


