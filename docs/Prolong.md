# Prolong

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Тип создаваемой заявки. | 
**Fqdn** | **string** | Полное имя домена. | 
**IsAntispamEnabled** | Pointer to **bool** | Это логическое значение, которое показывает включена ли услуга \&quot;Антиспам\&quot; для домена | [optional] 
**IsAutoprolongEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли автопродление домена. | [optional] 
**IsWhoisPrivacyEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Опция недоступна для доменов в зонах .ru и .рф. | [optional] 
**Period** | Pointer to [**DomainPaymentPeriod**](DomainPaymentPeriod.md) |  | [optional] 
**PersonId** | Pointer to **float32** | Идентификатор администратора, на которого зарегистрирован домен. | [optional] 
**Prime** | Pointer to [**DomainPrimeType**](DomainPrimeType.md) |  | [optional] 

## Methods

### NewProlong

`func NewProlong(action string, fqdn string, ) *Prolong`

NewProlong instantiates a new Prolong object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProlongWithDefaults

`func NewProlongWithDefaults() *Prolong`

NewProlongWithDefaults instantiates a new Prolong object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Prolong) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Prolong) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Prolong) SetAction(v string)`

SetAction sets Action field to given value.


### GetFqdn

`func (o *Prolong) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Prolong) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Prolong) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetIsAntispamEnabled

`func (o *Prolong) GetIsAntispamEnabled() bool`

GetIsAntispamEnabled returns the IsAntispamEnabled field if non-nil, zero value otherwise.

### GetIsAntispamEnabledOk

`func (o *Prolong) GetIsAntispamEnabledOk() (*bool, bool)`

GetIsAntispamEnabledOk returns a tuple with the IsAntispamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAntispamEnabled

`func (o *Prolong) SetIsAntispamEnabled(v bool)`

SetIsAntispamEnabled sets IsAntispamEnabled field to given value.

### HasIsAntispamEnabled

`func (o *Prolong) HasIsAntispamEnabled() bool`

HasIsAntispamEnabled returns a boolean if a field has been set.

### GetIsAutoprolongEnabled

`func (o *Prolong) GetIsAutoprolongEnabled() bool`

GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field if non-nil, zero value otherwise.

### GetIsAutoprolongEnabledOk

`func (o *Prolong) GetIsAutoprolongEnabledOk() (*bool, bool)`

GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprolongEnabled

`func (o *Prolong) SetIsAutoprolongEnabled(v bool)`

SetIsAutoprolongEnabled sets IsAutoprolongEnabled field to given value.

### HasIsAutoprolongEnabled

`func (o *Prolong) HasIsAutoprolongEnabled() bool`

HasIsAutoprolongEnabled returns a boolean if a field has been set.

### GetIsWhoisPrivacyEnabled

`func (o *Prolong) GetIsWhoisPrivacyEnabled() bool`

GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field if non-nil, zero value otherwise.

### GetIsWhoisPrivacyEnabledOk

`func (o *Prolong) GetIsWhoisPrivacyEnabledOk() (*bool, bool)`

GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhoisPrivacyEnabled

`func (o *Prolong) SetIsWhoisPrivacyEnabled(v bool)`

SetIsWhoisPrivacyEnabled sets IsWhoisPrivacyEnabled field to given value.

### HasIsWhoisPrivacyEnabled

`func (o *Prolong) HasIsWhoisPrivacyEnabled() bool`

HasIsWhoisPrivacyEnabled returns a boolean if a field has been set.

### GetPeriod

`func (o *Prolong) GetPeriod() DomainPaymentPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Prolong) GetPeriodOk() (*DomainPaymentPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Prolong) SetPeriod(v DomainPaymentPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Prolong) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPersonId

`func (o *Prolong) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Prolong) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Prolong) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Prolong) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetPrime

`func (o *Prolong) GetPrime() DomainPrimeType`

GetPrime returns the Prime field if non-nil, zero value otherwise.

### GetPrimeOk

`func (o *Prolong) GetPrimeOk() (*DomainPrimeType, bool)`

GetPrimeOk returns a tuple with the Prime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrime

`func (o *Prolong) SetPrime(v DomainPrimeType)`

SetPrime sets Prime field to given value.

### HasPrime

`func (o *Prolong) HasPrime() bool`

HasPrime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


