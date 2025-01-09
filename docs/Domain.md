# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedBuyPeriods** | [**[]DomainAllowedBuyPeriodsInner**](DomainAllowedBuyPeriodsInner.md) | Допустимые периоды продления домена. | 
**DaysLeft** | **float32** | Количество дней, оставшихся до конца срока регистрации домена. | 
**DomainStatus** | **string** | Статус домена. | 
**Expiration** | **string** | Дата окончания срока регистрации домена, для доменов без срока окончания регистрации будет приходить 0000-00-00. | 
**Fqdn** | **string** | Полное имя домена. | 
**Id** | **float32** | ID домена. | 
**IsAutoprolongEnabled** | **NullableBool** | Это логическое значение, которое показывает, включено ли автопродление домена. | 
**IsPremium** | **bool** | Это логическое значение, которое показывает, является ли домен премиальным. | 
**IsProlongAllowed** | **bool** | Это логическое значение, которое показывает, можно ли сейчас продлить домен. | 
**IsTechnical** | **bool** | Это логическое значение, которое показывает, является ли домен техническим. | 
**IsWhoisPrivacyEnabled** | **NullableBool** | Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Если приходит null, значит для данной зоны эта услуга не доступна. | 
**LinkedIp** | **NullableString** | Привязанный к домену IP-адрес. | 
**PaidTill** | **NullableString** | До какого числа оплачен домен. | 
**PersonId** | **NullableFloat32** | ID администратора, на которого зарегистрирован домен. | 
**PremiumProlongCost** | **NullableFloat32** | Стоимость премиального домена. | 
**Provider** | **NullableString** | ID регистратора домена. | 
**RequestStatus** | **NullableString** | Статус заявки на продление/регистрацию/трансфер домена. | 
**Subdomains** | [**[]Subdomain**](Subdomain.md) | Список поддоменов. | 
**TldId** | **NullableFloat32** | ID доменной зоны. | 

## Methods

### NewDomain

`func NewDomain(allowedBuyPeriods []DomainAllowedBuyPeriodsInner, daysLeft float32, domainStatus string, expiration string, fqdn string, id float32, isAutoprolongEnabled NullableBool, isPremium bool, isProlongAllowed bool, isTechnical bool, isWhoisPrivacyEnabled NullableBool, linkedIp NullableString, paidTill NullableString, personId NullableFloat32, premiumProlongCost NullableFloat32, provider NullableString, requestStatus NullableString, subdomains []Subdomain, tldId NullableFloat32, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedBuyPeriods

`func (o *Domain) GetAllowedBuyPeriods() []DomainAllowedBuyPeriodsInner`

GetAllowedBuyPeriods returns the AllowedBuyPeriods field if non-nil, zero value otherwise.

### GetAllowedBuyPeriodsOk

`func (o *Domain) GetAllowedBuyPeriodsOk() (*[]DomainAllowedBuyPeriodsInner, bool)`

GetAllowedBuyPeriodsOk returns a tuple with the AllowedBuyPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBuyPeriods

`func (o *Domain) SetAllowedBuyPeriods(v []DomainAllowedBuyPeriodsInner)`

SetAllowedBuyPeriods sets AllowedBuyPeriods field to given value.


### GetDaysLeft

`func (o *Domain) GetDaysLeft() float32`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *Domain) GetDaysLeftOk() (*float32, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *Domain) SetDaysLeft(v float32)`

SetDaysLeft sets DaysLeft field to given value.


### GetDomainStatus

`func (o *Domain) GetDomainStatus() string`

GetDomainStatus returns the DomainStatus field if non-nil, zero value otherwise.

### GetDomainStatusOk

`func (o *Domain) GetDomainStatusOk() (*string, bool)`

GetDomainStatusOk returns a tuple with the DomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStatus

`func (o *Domain) SetDomainStatus(v string)`

SetDomainStatus sets DomainStatus field to given value.


### GetExpiration

`func (o *Domain) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Domain) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Domain) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetFqdn

`func (o *Domain) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Domain) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Domain) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetId

`func (o *Domain) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v float32)`

SetId sets Id field to given value.


### GetIsAutoprolongEnabled

`func (o *Domain) GetIsAutoprolongEnabled() bool`

GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field if non-nil, zero value otherwise.

### GetIsAutoprolongEnabledOk

`func (o *Domain) GetIsAutoprolongEnabledOk() (*bool, bool)`

GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprolongEnabled

`func (o *Domain) SetIsAutoprolongEnabled(v bool)`

SetIsAutoprolongEnabled sets IsAutoprolongEnabled field to given value.


### SetIsAutoprolongEnabledNil

`func (o *Domain) SetIsAutoprolongEnabledNil(b bool)`

 SetIsAutoprolongEnabledNil sets the value for IsAutoprolongEnabled to be an explicit nil

### UnsetIsAutoprolongEnabled
`func (o *Domain) UnsetIsAutoprolongEnabled()`

UnsetIsAutoprolongEnabled ensures that no value is present for IsAutoprolongEnabled, not even an explicit nil
### GetIsPremium

`func (o *Domain) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *Domain) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *Domain) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.


### GetIsProlongAllowed

`func (o *Domain) GetIsProlongAllowed() bool`

GetIsProlongAllowed returns the IsProlongAllowed field if non-nil, zero value otherwise.

### GetIsProlongAllowedOk

`func (o *Domain) GetIsProlongAllowedOk() (*bool, bool)`

GetIsProlongAllowedOk returns a tuple with the IsProlongAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProlongAllowed

`func (o *Domain) SetIsProlongAllowed(v bool)`

SetIsProlongAllowed sets IsProlongAllowed field to given value.


### GetIsTechnical

`func (o *Domain) GetIsTechnical() bool`

GetIsTechnical returns the IsTechnical field if non-nil, zero value otherwise.

### GetIsTechnicalOk

`func (o *Domain) GetIsTechnicalOk() (*bool, bool)`

GetIsTechnicalOk returns a tuple with the IsTechnical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTechnical

`func (o *Domain) SetIsTechnical(v bool)`

SetIsTechnical sets IsTechnical field to given value.


### GetIsWhoisPrivacyEnabled

`func (o *Domain) GetIsWhoisPrivacyEnabled() bool`

GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field if non-nil, zero value otherwise.

### GetIsWhoisPrivacyEnabledOk

`func (o *Domain) GetIsWhoisPrivacyEnabledOk() (*bool, bool)`

GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhoisPrivacyEnabled

`func (o *Domain) SetIsWhoisPrivacyEnabled(v bool)`

SetIsWhoisPrivacyEnabled sets IsWhoisPrivacyEnabled field to given value.


### SetIsWhoisPrivacyEnabledNil

`func (o *Domain) SetIsWhoisPrivacyEnabledNil(b bool)`

 SetIsWhoisPrivacyEnabledNil sets the value for IsWhoisPrivacyEnabled to be an explicit nil

### UnsetIsWhoisPrivacyEnabled
`func (o *Domain) UnsetIsWhoisPrivacyEnabled()`

UnsetIsWhoisPrivacyEnabled ensures that no value is present for IsWhoisPrivacyEnabled, not even an explicit nil
### GetLinkedIp

`func (o *Domain) GetLinkedIp() string`

GetLinkedIp returns the LinkedIp field if non-nil, zero value otherwise.

### GetLinkedIpOk

`func (o *Domain) GetLinkedIpOk() (*string, bool)`

GetLinkedIpOk returns a tuple with the LinkedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIp

`func (o *Domain) SetLinkedIp(v string)`

SetLinkedIp sets LinkedIp field to given value.


### SetLinkedIpNil

`func (o *Domain) SetLinkedIpNil(b bool)`

 SetLinkedIpNil sets the value for LinkedIp to be an explicit nil

### UnsetLinkedIp
`func (o *Domain) UnsetLinkedIp()`

UnsetLinkedIp ensures that no value is present for LinkedIp, not even an explicit nil
### GetPaidTill

`func (o *Domain) GetPaidTill() string`

GetPaidTill returns the PaidTill field if non-nil, zero value otherwise.

### GetPaidTillOk

`func (o *Domain) GetPaidTillOk() (*string, bool)`

GetPaidTillOk returns a tuple with the PaidTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTill

`func (o *Domain) SetPaidTill(v string)`

SetPaidTill sets PaidTill field to given value.


### SetPaidTillNil

`func (o *Domain) SetPaidTillNil(b bool)`

 SetPaidTillNil sets the value for PaidTill to be an explicit nil

### UnsetPaidTill
`func (o *Domain) UnsetPaidTill()`

UnsetPaidTill ensures that no value is present for PaidTill, not even an explicit nil
### GetPersonId

`func (o *Domain) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Domain) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Domain) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.


### SetPersonIdNil

`func (o *Domain) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *Domain) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetPremiumProlongCost

`func (o *Domain) GetPremiumProlongCost() float32`

GetPremiumProlongCost returns the PremiumProlongCost field if non-nil, zero value otherwise.

### GetPremiumProlongCostOk

`func (o *Domain) GetPremiumProlongCostOk() (*float32, bool)`

GetPremiumProlongCostOk returns a tuple with the PremiumProlongCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumProlongCost

`func (o *Domain) SetPremiumProlongCost(v float32)`

SetPremiumProlongCost sets PremiumProlongCost field to given value.


### SetPremiumProlongCostNil

`func (o *Domain) SetPremiumProlongCostNil(b bool)`

 SetPremiumProlongCostNil sets the value for PremiumProlongCost to be an explicit nil

### UnsetPremiumProlongCost
`func (o *Domain) UnsetPremiumProlongCost()`

UnsetPremiumProlongCost ensures that no value is present for PremiumProlongCost, not even an explicit nil
### GetProvider

`func (o *Domain) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Domain) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Domain) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *Domain) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *Domain) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetRequestStatus

`func (o *Domain) GetRequestStatus() string`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *Domain) GetRequestStatusOk() (*string, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *Domain) SetRequestStatus(v string)`

SetRequestStatus sets RequestStatus field to given value.


### SetRequestStatusNil

`func (o *Domain) SetRequestStatusNil(b bool)`

 SetRequestStatusNil sets the value for RequestStatus to be an explicit nil

### UnsetRequestStatus
`func (o *Domain) UnsetRequestStatus()`

UnsetRequestStatus ensures that no value is present for RequestStatus, not even an explicit nil
### GetSubdomains

`func (o *Domain) GetSubdomains() []Subdomain`

GetSubdomains returns the Subdomains field if non-nil, zero value otherwise.

### GetSubdomainsOk

`func (o *Domain) GetSubdomainsOk() (*[]Subdomain, bool)`

GetSubdomainsOk returns a tuple with the Subdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomains

`func (o *Domain) SetSubdomains(v []Subdomain)`

SetSubdomains sets Subdomains field to given value.


### GetTldId

`func (o *Domain) GetTldId() float32`

GetTldId returns the TldId field if non-nil, zero value otherwise.

### GetTldIdOk

`func (o *Domain) GetTldIdOk() (*float32, bool)`

GetTldIdOk returns a tuple with the TldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTldId

`func (o *Domain) SetTldId(v float32)`

SetTldId sets TldId field to given value.


### SetTldIdNil

`func (o *Domain) SetTldIdNil(b bool)`

 SetTldIdNil sets the value for TldId to be an explicit nil

### UnsetTldId
`func (o *Domain) UnsetTldId()`

UnsetTldId ensures that no value is present for TldId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


