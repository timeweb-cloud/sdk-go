# CreateDomainRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Тип создаваемой заявки. | 
**Fqdn** | **string** | Полное имя домена. | 
**IsAutoprolongEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли автопродление домена. | [optional] 
**IsWhoisPrivacyEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Опция недоступна для доменов в зонах .ru и .рф. | [optional] 
**Ns** | Pointer to [**[]RegisterNsInner**](RegisterNsInner.md) | Name-серверы для регистрации домена. Если не передавать этот параметр, будут использованы наши стандартные name-серверы. Нужно указать как минимум 2 name-сервера. | [optional] 
**Period** | Pointer to [**DomainPaymentPeriod**](DomainPaymentPeriod.md) |  | [optional] 
**PersonId** | **float32** | Идентификатор администратора, на которого зарегистрирован домен. | 
**IsAntispamEnabled** | Pointer to **bool** | Это логическое значение, которое показывает включена ли услуга \&quot;Антиспам\&quot; для домена | [optional] 
**Prime** | Pointer to [**DomainPrimeType**](DomainPrimeType.md) |  | [optional] 
**AuthCode** | **string** | Код авторизации для переноса домена. | 

## Methods

### NewCreateDomainRequestRequest

`func NewCreateDomainRequestRequest(action string, fqdn string, personId float32, authCode string, ) *CreateDomainRequestRequest`

NewCreateDomainRequestRequest instantiates a new CreateDomainRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainRequestRequestWithDefaults

`func NewCreateDomainRequestRequestWithDefaults() *CreateDomainRequestRequest`

NewCreateDomainRequestRequestWithDefaults instantiates a new CreateDomainRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CreateDomainRequestRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateDomainRequestRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateDomainRequestRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetFqdn

`func (o *CreateDomainRequestRequest) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateDomainRequestRequest) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateDomainRequestRequest) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetIsAutoprolongEnabled

`func (o *CreateDomainRequestRequest) GetIsAutoprolongEnabled() bool`

GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field if non-nil, zero value otherwise.

### GetIsAutoprolongEnabledOk

`func (o *CreateDomainRequestRequest) GetIsAutoprolongEnabledOk() (*bool, bool)`

GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprolongEnabled

`func (o *CreateDomainRequestRequest) SetIsAutoprolongEnabled(v bool)`

SetIsAutoprolongEnabled sets IsAutoprolongEnabled field to given value.

### HasIsAutoprolongEnabled

`func (o *CreateDomainRequestRequest) HasIsAutoprolongEnabled() bool`

HasIsAutoprolongEnabled returns a boolean if a field has been set.

### GetIsWhoisPrivacyEnabled

`func (o *CreateDomainRequestRequest) GetIsWhoisPrivacyEnabled() bool`

GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field if non-nil, zero value otherwise.

### GetIsWhoisPrivacyEnabledOk

`func (o *CreateDomainRequestRequest) GetIsWhoisPrivacyEnabledOk() (*bool, bool)`

GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhoisPrivacyEnabled

`func (o *CreateDomainRequestRequest) SetIsWhoisPrivacyEnabled(v bool)`

SetIsWhoisPrivacyEnabled sets IsWhoisPrivacyEnabled field to given value.

### HasIsWhoisPrivacyEnabled

`func (o *CreateDomainRequestRequest) HasIsWhoisPrivacyEnabled() bool`

HasIsWhoisPrivacyEnabled returns a boolean if a field has been set.

### GetNs

`func (o *CreateDomainRequestRequest) GetNs() []RegisterNsInner`

GetNs returns the Ns field if non-nil, zero value otherwise.

### GetNsOk

`func (o *CreateDomainRequestRequest) GetNsOk() (*[]RegisterNsInner, bool)`

GetNsOk returns a tuple with the Ns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNs

`func (o *CreateDomainRequestRequest) SetNs(v []RegisterNsInner)`

SetNs sets Ns field to given value.

### HasNs

`func (o *CreateDomainRequestRequest) HasNs() bool`

HasNs returns a boolean if a field has been set.

### GetPeriod

`func (o *CreateDomainRequestRequest) GetPeriod() DomainPaymentPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateDomainRequestRequest) GetPeriodOk() (*DomainPaymentPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateDomainRequestRequest) SetPeriod(v DomainPaymentPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CreateDomainRequestRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPersonId

`func (o *CreateDomainRequestRequest) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *CreateDomainRequestRequest) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *CreateDomainRequestRequest) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.


### GetIsAntispamEnabled

`func (o *CreateDomainRequestRequest) GetIsAntispamEnabled() bool`

GetIsAntispamEnabled returns the IsAntispamEnabled field if non-nil, zero value otherwise.

### GetIsAntispamEnabledOk

`func (o *CreateDomainRequestRequest) GetIsAntispamEnabledOk() (*bool, bool)`

GetIsAntispamEnabledOk returns a tuple with the IsAntispamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAntispamEnabled

`func (o *CreateDomainRequestRequest) SetIsAntispamEnabled(v bool)`

SetIsAntispamEnabled sets IsAntispamEnabled field to given value.

### HasIsAntispamEnabled

`func (o *CreateDomainRequestRequest) HasIsAntispamEnabled() bool`

HasIsAntispamEnabled returns a boolean if a field has been set.

### GetPrime

`func (o *CreateDomainRequestRequest) GetPrime() DomainPrimeType`

GetPrime returns the Prime field if non-nil, zero value otherwise.

### GetPrimeOk

`func (o *CreateDomainRequestRequest) GetPrimeOk() (*DomainPrimeType, bool)`

GetPrimeOk returns a tuple with the Prime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrime

`func (o *CreateDomainRequestRequest) SetPrime(v DomainPrimeType)`

SetPrime sets Prime field to given value.

### HasPrime

`func (o *CreateDomainRequestRequest) HasPrime() bool`

HasPrime returns a boolean if a field has been set.

### GetAuthCode

`func (o *CreateDomainRequestRequest) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *CreateDomainRequestRequest) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *CreateDomainRequestRequest) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


