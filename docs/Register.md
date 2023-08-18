# Register

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Тип создаваемой заявки. | 
**Fqdn** | **string** | Полное имя домена. | 
**IsAutoprolongEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли автопродление домена. | [optional] 
**IsWhoisPrivacyEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Опция недоступна для доменов в зонах .ru и .рф. | [optional] 
**Ns** | Pointer to [**[]RegisterNsInner**](RegisterNsInner.md) | Name-серверы для регистрации домена. Если не передавать этот параметр, будут использованы наши стандартные name-серверы. Нужно указать как минимум 2 name-сервера. | [optional] 
**Period** | Pointer to [**DomainPaymentPeriod**](DomainPaymentPeriod.md) |  | [optional] 
**PersonId** | **float32** | Идентификатор администратора, на которого регистрируется домен. | 

## Methods

### NewRegister

`func NewRegister(action string, fqdn string, personId float32, ) *Register`

NewRegister instantiates a new Register object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterWithDefaults

`func NewRegisterWithDefaults() *Register`

NewRegisterWithDefaults instantiates a new Register object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Register) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Register) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Register) SetAction(v string)`

SetAction sets Action field to given value.


### GetFqdn

`func (o *Register) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Register) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Register) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetIsAutoprolongEnabled

`func (o *Register) GetIsAutoprolongEnabled() bool`

GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field if non-nil, zero value otherwise.

### GetIsAutoprolongEnabledOk

`func (o *Register) GetIsAutoprolongEnabledOk() (*bool, bool)`

GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprolongEnabled

`func (o *Register) SetIsAutoprolongEnabled(v bool)`

SetIsAutoprolongEnabled sets IsAutoprolongEnabled field to given value.

### HasIsAutoprolongEnabled

`func (o *Register) HasIsAutoprolongEnabled() bool`

HasIsAutoprolongEnabled returns a boolean if a field has been set.

### GetIsWhoisPrivacyEnabled

`func (o *Register) GetIsWhoisPrivacyEnabled() bool`

GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field if non-nil, zero value otherwise.

### GetIsWhoisPrivacyEnabledOk

`func (o *Register) GetIsWhoisPrivacyEnabledOk() (*bool, bool)`

GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhoisPrivacyEnabled

`func (o *Register) SetIsWhoisPrivacyEnabled(v bool)`

SetIsWhoisPrivacyEnabled sets IsWhoisPrivacyEnabled field to given value.

### HasIsWhoisPrivacyEnabled

`func (o *Register) HasIsWhoisPrivacyEnabled() bool`

HasIsWhoisPrivacyEnabled returns a boolean if a field has been set.

### GetNs

`func (o *Register) GetNs() []RegisterNsInner`

GetNs returns the Ns field if non-nil, zero value otherwise.

### GetNsOk

`func (o *Register) GetNsOk() (*[]RegisterNsInner, bool)`

GetNsOk returns a tuple with the Ns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs

`func (o *Register) SetNs(v []RegisterNsInner)`

SetNs sets Ns field to given value.

### HasNs

`func (o *Register) HasNs() bool`

HasNs returns a boolean if a field has been set.

### GetPeriod

`func (o *Register) GetPeriod() DomainPaymentPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Register) GetPeriodOk() (*DomainPaymentPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Register) SetPeriod(v DomainPaymentPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Register) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPersonId

`func (o *Register) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Register) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Register) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


