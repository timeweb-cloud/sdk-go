# DomainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Адрес для сбора почты с ошибочных ящиков | 
**Used** | **float32** | Использованное место на почте (в Мб). | 

## Methods

### NewDomainInfo

`func NewDomainInfo(email string, used float32, ) *DomainInfo`

NewDomainInfo instantiates a new DomainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInfoWithDefaults

`func NewDomainInfoWithDefaults() *DomainInfo`

NewDomainInfoWithDefaults instantiates a new DomainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *DomainInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DomainInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DomainInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUsed

`func (o *DomainInfo) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *DomainInfo) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *DomainInfo) SetUsed(v float32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


