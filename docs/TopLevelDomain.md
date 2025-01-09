# TopLevelDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedBuyPeriods** | [**[]TopLevelDomainAllowedBuyPeriodsInner**](TopLevelDomainAllowedBuyPeriodsInner.md) | Список доступных периодов для регистрации/продления доменов со стоимостью. | 
**EarlyRenewPeriod** | **NullableFloat32** | Количество дней до истечение срока регистрации, когда можно продлять домен. | 
**GracePeriod** | **float32** | Количество дней, которые действует льготный период когда вы ещё можете продлить домен, после окончания его регистрации | 
**Id** | **float32** | ID доменной зоны. | 
**IsPublished** | **bool** | Это логическое значение, которое показывает, опубликована ли доменная зона. | 
**IsRegistered** | **bool** | Это логическое значение, которое показывает, зарегистрирована ли доменная зона. | 
**IsWhoisPrivacyDefaultEnabled** | **bool** | Это логическое значение, которое показывает, включено ли по умолчанию скрытие данных администратора для доменной зоны. | 
**IsWhoisPrivacyEnabled** | **bool** | Это логическое значение, которое показывает, доступно ли управление скрытием данных администратора для доменной зоны. | 
**Name** | **string** | Имя доменной зоны. | 
**Price** | **float32** | Цена регистрации домена | 
**ProlongPrice** | **float32** | Цена продления домена. | 
**Registrar** | **string** | Регистратор доменной зоны. | 
**Transfer** | **float32** | Цена услуги трансфера домена. | 
**WhoisPrivacyPrice** | **float32** | Цена услуги скрытия данных администратора для доменной зоны. | 

## Methods

### NewTopLevelDomain

`func NewTopLevelDomain(allowedBuyPeriods []TopLevelDomainAllowedBuyPeriodsInner, earlyRenewPeriod NullableFloat32, gracePeriod float32, id float32, isPublished bool, isRegistered bool, isWhoisPrivacyDefaultEnabled bool, isWhoisPrivacyEnabled bool, name string, price float32, prolongPrice float32, registrar string, transfer float32, whoisPrivacyPrice float32, ) *TopLevelDomain`

NewTopLevelDomain instantiates a new TopLevelDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopLevelDomainWithDefaults

`func NewTopLevelDomainWithDefaults() *TopLevelDomain`

NewTopLevelDomainWithDefaults instantiates a new TopLevelDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedBuyPeriods

`func (o *TopLevelDomain) GetAllowedBuyPeriods() []TopLevelDomainAllowedBuyPeriodsInner`

GetAllowedBuyPeriods returns the AllowedBuyPeriods field if non-nil, zero value otherwise.

### GetAllowedBuyPeriodsOk

`func (o *TopLevelDomain) GetAllowedBuyPeriodsOk() (*[]TopLevelDomainAllowedBuyPeriodsInner, bool)`

GetAllowedBuyPeriodsOk returns a tuple with the AllowedBuyPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBuyPeriods

`func (o *TopLevelDomain) SetAllowedBuyPeriods(v []TopLevelDomainAllowedBuyPeriodsInner)`

SetAllowedBuyPeriods sets AllowedBuyPeriods field to given value.


### GetEarlyRenewPeriod

`func (o *TopLevelDomain) GetEarlyRenewPeriod() float32`

GetEarlyRenewPeriod returns the EarlyRenewPeriod field if non-nil, zero value otherwise.

### GetEarlyRenewPeriodOk

`func (o *TopLevelDomain) GetEarlyRenewPeriodOk() (*float32, bool)`

GetEarlyRenewPeriodOk returns a tuple with the EarlyRenewPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyRenewPeriod

`func (o *TopLevelDomain) SetEarlyRenewPeriod(v float32)`

SetEarlyRenewPeriod sets EarlyRenewPeriod field to given value.


### SetEarlyRenewPeriodNil

`func (o *TopLevelDomain) SetEarlyRenewPeriodNil(b bool)`

 SetEarlyRenewPeriodNil sets the value for EarlyRenewPeriod to be an explicit nil

### UnsetEarlyRenewPeriod
`func (o *TopLevelDomain) UnsetEarlyRenewPeriod()`

UnsetEarlyRenewPeriod ensures that no value is present for EarlyRenewPeriod, not even an explicit nil
### GetGracePeriod

`func (o *TopLevelDomain) GetGracePeriod() float32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *TopLevelDomain) GetGracePeriodOk() (*float32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *TopLevelDomain) SetGracePeriod(v float32)`

SetGracePeriod sets GracePeriod field to given value.


### GetId

`func (o *TopLevelDomain) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopLevelDomain) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopLevelDomain) SetId(v float32)`

SetId sets Id field to given value.


### GetIsPublished

`func (o *TopLevelDomain) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *TopLevelDomain) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *TopLevelDomain) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetIsRegistered

`func (o *TopLevelDomain) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *TopLevelDomain) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *TopLevelDomain) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.


### GetIsWhoisPrivacyDefaultEnabled

`func (o *TopLevelDomain) GetIsWhoisPrivacyDefaultEnabled() bool`

GetIsWhoisPrivacyDefaultEnabled returns the IsWhoisPrivacyDefaultEnabled field if non-nil, zero value otherwise.

### GetIsWhoisPrivacyDefaultEnabledOk

`func (o *TopLevelDomain) GetIsWhoisPrivacyDefaultEnabledOk() (*bool, bool)`

GetIsWhoisPrivacyDefaultEnabledOk returns a tuple with the IsWhoisPrivacyDefaultEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhoisPrivacyDefaultEnabled

`func (o *TopLevelDomain) SetIsWhoisPrivacyDefaultEnabled(v bool)`

SetIsWhoisPrivacyDefaultEnabled sets IsWhoisPrivacyDefaultEnabled field to given value.


### GetIsWhoisPrivacyEnabled

`func (o *TopLevelDomain) GetIsWhoisPrivacyEnabled() bool`

GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field if non-nil, zero value otherwise.

### GetIsWhoisPrivacyEnabledOk

`func (o *TopLevelDomain) GetIsWhoisPrivacyEnabledOk() (*bool, bool)`

GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhoisPrivacyEnabled

`func (o *TopLevelDomain) SetIsWhoisPrivacyEnabled(v bool)`

SetIsWhoisPrivacyEnabled sets IsWhoisPrivacyEnabled field to given value.


### GetName

`func (o *TopLevelDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopLevelDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopLevelDomain) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *TopLevelDomain) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TopLevelDomain) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TopLevelDomain) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetProlongPrice

`func (o *TopLevelDomain) GetProlongPrice() float32`

GetProlongPrice returns the ProlongPrice field if non-nil, zero value otherwise.

### GetProlongPriceOk

`func (o *TopLevelDomain) GetProlongPriceOk() (*float32, bool)`

GetProlongPriceOk returns a tuple with the ProlongPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongPrice

`func (o *TopLevelDomain) SetProlongPrice(v float32)`

SetProlongPrice sets ProlongPrice field to given value.


### GetRegistrar

`func (o *TopLevelDomain) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *TopLevelDomain) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *TopLevelDomain) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.


### GetTransfer

`func (o *TopLevelDomain) GetTransfer() float32`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *TopLevelDomain) GetTransferOk() (*float32, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *TopLevelDomain) SetTransfer(v float32)`

SetTransfer sets Transfer field to given value.


### GetWhoisPrivacyPrice

`func (o *TopLevelDomain) GetWhoisPrivacyPrice() float32`

GetWhoisPrivacyPrice returns the WhoisPrivacyPrice field if non-nil, zero value otherwise.

### GetWhoisPrivacyPriceOk

`func (o *TopLevelDomain) GetWhoisPrivacyPriceOk() (*float32, bool)`

GetWhoisPrivacyPriceOk returns a tuple with the WhoisPrivacyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisPrivacyPrice

`func (o *TopLevelDomain) SetWhoisPrivacyPrice(v float32)`

SetWhoisPrivacyPrice sets WhoisPrivacyPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


