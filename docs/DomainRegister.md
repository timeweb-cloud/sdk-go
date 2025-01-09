# DomainRegister

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Тип создаваемой заявки. | 
**Fqdn** | **string** | Полное имя домена. | 
**IsAutoprolongEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли автопродление домена. | [optional] 
**IsWhoisPrivacyEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Опция недоступна для доменов в зонах .ru и .рф. | [optional] 
**Period** | Pointer to [**DomainPaymentPeriod**](DomainPaymentPeriod.md) |  | [optional] 
**PersonId** | **float32** | ID администратора, на которого регистрируется домен. | 

## Methods

### NewDomainRegister

`func NewDomainRegister(action string, fqdn string, personId float32, ) *DomainRegister`

NewDomainRegister instantiates a new DomainRegister object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRegisterWithDefaults

`func NewDomainRegisterWithDefaults() *DomainRegister`

NewDomainRegisterWithDefaults instantiates a new DomainRegister object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DomainRegister) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DomainRegister) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DomainRegister) SetAction(v string)`

SetAction sets Action field to given value.


### GetFqdn

`func (o *DomainRegister) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DomainRegister) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DomainRegister) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetIsAutoprolongEnabled

`func (o *DomainRegister) GetIsAutoprolongEnabled() bool`

GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field if non-nil, zero value otherwise.

### GetIsAutoprolongEnabledOk

`func (o *DomainRegister) GetIsAutoprolongEnabledOk() (*bool, bool)`

GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprolongEnabled

`func (o *DomainRegister) SetIsAutoprolongEnabled(v bool)`

SetIsAutoprolongEnabled sets IsAutoprolongEnabled field to given value.

### HasIsAutoprolongEnabled

`func (o *DomainRegister) HasIsAutoprolongEnabled() bool`

HasIsAutoprolongEnabled returns a boolean if a field has been set.

### GetIsWhoisPrivacyEnabled

`func (o *DomainRegister) GetIsWhoisPrivacyEnabled() bool`

GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field if non-nil, zero value otherwise.

### GetIsWhoisPrivacyEnabledOk

`func (o *DomainRegister) GetIsWhoisPrivacyEnabledOk() (*bool, bool)`

GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhoisPrivacyEnabled

`func (o *DomainRegister) SetIsWhoisPrivacyEnabled(v bool)`

SetIsWhoisPrivacyEnabled sets IsWhoisPrivacyEnabled field to given value.

### HasIsWhoisPrivacyEnabled

`func (o *DomainRegister) HasIsWhoisPrivacyEnabled() bool`

HasIsWhoisPrivacyEnabled returns a boolean if a field has been set.

### GetPeriod

`func (o *DomainRegister) GetPeriod() DomainPaymentPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainRegister) GetPeriodOk() (*DomainPaymentPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainRegister) SetPeriod(v DomainPaymentPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DomainRegister) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPersonId

`func (o *DomainRegister) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *DomainRegister) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *DomainRegister) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


