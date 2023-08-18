# DomainAllowedBuyPeriodsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**DomainPaymentPeriod**](DomainPaymentPeriod.md) |  | 
**Price** | **float32** | Стоимость домена за указанный период. | 

## Methods

### NewDomainAllowedBuyPeriodsInner

`func NewDomainAllowedBuyPeriodsInner(period DomainPaymentPeriod, price float32, ) *DomainAllowedBuyPeriodsInner`

NewDomainAllowedBuyPeriodsInner instantiates a new DomainAllowedBuyPeriodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAllowedBuyPeriodsInnerWithDefaults

`func NewDomainAllowedBuyPeriodsInnerWithDefaults() *DomainAllowedBuyPeriodsInner`

NewDomainAllowedBuyPeriodsInnerWithDefaults instantiates a new DomainAllowedBuyPeriodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *DomainAllowedBuyPeriodsInner) GetPeriod() DomainPaymentPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainAllowedBuyPeriodsInner) GetPeriodOk() (*DomainPaymentPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainAllowedBuyPeriodsInner) SetPeriod(v DomainPaymentPeriod)`

SetPeriod sets Period field to given value.


### GetPrice

`func (o *DomainAllowedBuyPeriodsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DomainAllowedBuyPeriodsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DomainAllowedBuyPeriodsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


