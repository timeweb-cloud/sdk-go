# Org

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип администратора. | 
**Name** | **string** | Название организации. | 
**IsResident** | **bool** | Это логическое значение, которое показывает, является ли администратор резидентом РФ. | 
**ContactName** | **string** | Контактное лицо организации. | 
**Inn** | **string** | ИНН организации. | 
**Kpp** | Pointer to **string** | КПП организации. | [optional] 
**LegalAddress** | **string** | Юридический адрес организации. | 
**Postcode** | **string** | Почтовый индекс. | 
**MailingAddress** | **string** | Почтовый адрес. | 
**Phone** | **string** | Контактный телефон. | 
**Email** | **string** | Адрес электронной почты. | 
**CountryCode** | Pointer to **NullableString** | Код страны. Только для нерезидентов РФ (&#x60;is_resident: false&#x60;); для резидентов поле передавать не нужно. | [optional] 

## Methods

### NewOrg

`func NewOrg(type_ string, name string, isResident bool, contactName string, inn string, legalAddress string, postcode string, mailingAddress string, phone string, email string, ) *Org`

NewOrg instantiates a new Org object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWithDefaults

`func NewOrgWithDefaults() *Org`

NewOrgWithDefaults instantiates a new Org object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Org) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Org) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Org) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Org) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Org) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Org) SetName(v string)`

SetName sets Name field to given value.


### GetIsResident

`func (o *Org) GetIsResident() bool`

GetIsResident returns the IsResident field if non-nil, zero value otherwise.

### GetIsResidentOk

`func (o *Org) GetIsResidentOk() (*bool, bool)`

GetIsResidentOk returns a tuple with the IsResident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResident

`func (o *Org) SetIsResident(v bool)`

SetIsResident sets IsResident field to given value.


### GetContactName

`func (o *Org) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *Org) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *Org) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### GetInn

`func (o *Org) GetInn() string`

GetInn returns the Inn field if non-nil, zero value otherwise.

### GetInnOk

`func (o *Org) GetInnOk() (*string, bool)`

GetInnOk returns a tuple with the Inn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInn

`func (o *Org) SetInn(v string)`

SetInn sets Inn field to given value.


### GetKpp

`func (o *Org) GetKpp() string`

GetKpp returns the Kpp field if non-nil, zero value otherwise.

### GetKppOk

`func (o *Org) GetKppOk() (*string, bool)`

GetKppOk returns a tuple with the Kpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpp

`func (o *Org) SetKpp(v string)`

SetKpp sets Kpp field to given value.

### HasKpp

`func (o *Org) HasKpp() bool`

HasKpp returns a boolean if a field has been set.

### GetLegalAddress

`func (o *Org) GetLegalAddress() string`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *Org) GetLegalAddressOk() (*string, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *Org) SetLegalAddress(v string)`

SetLegalAddress sets LegalAddress field to given value.


### GetPostcode

`func (o *Org) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *Org) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *Org) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetMailingAddress

`func (o *Org) GetMailingAddress() string`

GetMailingAddress returns the MailingAddress field if non-nil, zero value otherwise.

### GetMailingAddressOk

`func (o *Org) GetMailingAddressOk() (*string, bool)`

GetMailingAddressOk returns a tuple with the MailingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingAddress

`func (o *Org) SetMailingAddress(v string)`

SetMailingAddress sets MailingAddress field to given value.


### GetPhone

`func (o *Org) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Org) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Org) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *Org) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Org) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Org) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCountryCode

`func (o *Org) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Org) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Org) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Org) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *Org) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *Org) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


