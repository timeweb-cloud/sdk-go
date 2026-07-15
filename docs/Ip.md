# Ip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип администратора. | 
**Name** | **string** | ФИО индивидуального предпринимателя. | 
**IsResident** | **bool** | Это логическое значение, которое показывает, является ли администратор резидентом РФ. | 
**Birthdate** | **string** | Дата рождения. | 
**PassportDate** | **string** | Дата выдачи паспорта. | 
**PassportNumber** | **string** | Номер паспорта. | 
**PassportPlace** | **string** | Кем выдан паспорт. | 
**PassportSeries** | **string** | Серия паспорта. | 
**Inn** | **string** | ИНН индивидуального предпринимателя. | 
**Postcode** | **string** | Почтовый индекс. | 
**MailingAddress** | **string** | Почтовый адрес. | 
**Phone** | **string** | Контактный телефон. | 
**Email** | **string** | Адрес электронной почты. | 
**CountryCode** | Pointer to **NullableString** | Код страны. Только для нерезидентов РФ (&#x60;is_resident: false&#x60;); для резидентов поле передавать не нужно. | [optional] 

## Methods

### NewIp

`func NewIp(type_ string, name string, isResident bool, birthdate string, passportDate string, passportNumber string, passportPlace string, passportSeries string, inn string, postcode string, mailingAddress string, phone string, email string, ) *Ip`

NewIp instantiates a new Ip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpWithDefaults

`func NewIpWithDefaults() *Ip`

NewIpWithDefaults instantiates a new Ip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Ip) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ip) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ip) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Ip) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ip) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ip) SetName(v string)`

SetName sets Name field to given value.


### GetIsResident

`func (o *Ip) GetIsResident() bool`

GetIsResident returns the IsResident field if non-nil, zero value otherwise.

### GetIsResidentOk

`func (o *Ip) GetIsResidentOk() (*bool, bool)`

GetIsResidentOk returns a tuple with the IsResident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResident

`func (o *Ip) SetIsResident(v bool)`

SetIsResident sets IsResident field to given value.


### GetBirthdate

`func (o *Ip) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *Ip) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *Ip) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.


### GetPassportDate

`func (o *Ip) GetPassportDate() string`

GetPassportDate returns the PassportDate field if non-nil, zero value otherwise.

### GetPassportDateOk

`func (o *Ip) GetPassportDateOk() (*string, bool)`

GetPassportDateOk returns a tuple with the PassportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportDate

`func (o *Ip) SetPassportDate(v string)`

SetPassportDate sets PassportDate field to given value.


### GetPassportNumber

`func (o *Ip) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *Ip) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *Ip) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.


### GetPassportPlace

`func (o *Ip) GetPassportPlace() string`

GetPassportPlace returns the PassportPlace field if non-nil, zero value otherwise.

### GetPassportPlaceOk

`func (o *Ip) GetPassportPlaceOk() (*string, bool)`

GetPassportPlaceOk returns a tuple with the PassportPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportPlace

`func (o *Ip) SetPassportPlace(v string)`

SetPassportPlace sets PassportPlace field to given value.


### GetPassportSeries

`func (o *Ip) GetPassportSeries() string`

GetPassportSeries returns the PassportSeries field if non-nil, zero value otherwise.

### GetPassportSeriesOk

`func (o *Ip) GetPassportSeriesOk() (*string, bool)`

GetPassportSeriesOk returns a tuple with the PassportSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportSeries

`func (o *Ip) SetPassportSeries(v string)`

SetPassportSeries sets PassportSeries field to given value.


### GetInn

`func (o *Ip) GetInn() string`

GetInn returns the Inn field if non-nil, zero value otherwise.

### GetInnOk

`func (o *Ip) GetInnOk() (*string, bool)`

GetInnOk returns a tuple with the Inn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInn

`func (o *Ip) SetInn(v string)`

SetInn sets Inn field to given value.


### GetPostcode

`func (o *Ip) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *Ip) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *Ip) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetMailingAddress

`func (o *Ip) GetMailingAddress() string`

GetMailingAddress returns the MailingAddress field if non-nil, zero value otherwise.

### GetMailingAddressOk

`func (o *Ip) GetMailingAddressOk() (*string, bool)`

GetMailingAddressOk returns a tuple with the MailingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingAddress

`func (o *Ip) SetMailingAddress(v string)`

SetMailingAddress sets MailingAddress field to given value.


### GetPhone

`func (o *Ip) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Ip) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Ip) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *Ip) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Ip) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Ip) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCountryCode

`func (o *Ip) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Ip) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Ip) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Ip) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *Ip) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *Ip) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


