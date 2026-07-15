# Person2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип администратора. | 
**Name** | **string** | ФИО администратора. | 
**IsResident** | **bool** | Это логическое значение, которое показывает, является ли администратор резидентом РФ. | 
**Birthdate** | **string** | Дата рождения. | 
**PassportDate** | **string** | Дата выдачи паспорта. | 
**PassportNumber** | **string** | Номер паспорта. | 
**PassportPlace** | **string** | Кем выдан паспорт. | 
**PassportSeries** | **string** | Серия паспорта. | 
**Postcode** | **string** | Почтовый индекс. | 
**MailingAddress** | **string** | Почтовый адрес. | 
**Phone** | **string** | Контактный телефон. | 
**Email** | **string** | Адрес электронной почты. | 
**CountryCode** | Pointer to **NullableString** | Код страны. Только для нерезидентов РФ (&#x60;is_resident: false&#x60;); для резидентов поле передавать не нужно. | [optional] 

## Methods

### NewPerson2

`func NewPerson2(type_ string, name string, isResident bool, birthdate string, passportDate string, passportNumber string, passportPlace string, passportSeries string, postcode string, mailingAddress string, phone string, email string, ) *Person2`

NewPerson2 instantiates a new Person2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerson2WithDefaults

`func NewPerson2WithDefaults() *Person2`

NewPerson2WithDefaults instantiates a new Person2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Person2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Person2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Person2) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Person2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Person2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Person2) SetName(v string)`

SetName sets Name field to given value.


### GetIsResident

`func (o *Person2) GetIsResident() bool`

GetIsResident returns the IsResident field if non-nil, zero value otherwise.

### GetIsResidentOk

`func (o *Person2) GetIsResidentOk() (*bool, bool)`

GetIsResidentOk returns a tuple with the IsResident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResident

`func (o *Person2) SetIsResident(v bool)`

SetIsResident sets IsResident field to given value.


### GetBirthdate

`func (o *Person2) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *Person2) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *Person2) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.


### GetPassportDate

`func (o *Person2) GetPassportDate() string`

GetPassportDate returns the PassportDate field if non-nil, zero value otherwise.

### GetPassportDateOk

`func (o *Person2) GetPassportDateOk() (*string, bool)`

GetPassportDateOk returns a tuple with the PassportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportDate

`func (o *Person2) SetPassportDate(v string)`

SetPassportDate sets PassportDate field to given value.


### GetPassportNumber

`func (o *Person2) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *Person2) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *Person2) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.


### GetPassportPlace

`func (o *Person2) GetPassportPlace() string`

GetPassportPlace returns the PassportPlace field if non-nil, zero value otherwise.

### GetPassportPlaceOk

`func (o *Person2) GetPassportPlaceOk() (*string, bool)`

GetPassportPlaceOk returns a tuple with the PassportPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportPlace

`func (o *Person2) SetPassportPlace(v string)`

SetPassportPlace sets PassportPlace field to given value.


### GetPassportSeries

`func (o *Person2) GetPassportSeries() string`

GetPassportSeries returns the PassportSeries field if non-nil, zero value otherwise.

### GetPassportSeriesOk

`func (o *Person2) GetPassportSeriesOk() (*string, bool)`

GetPassportSeriesOk returns a tuple with the PassportSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportSeries

`func (o *Person2) SetPassportSeries(v string)`

SetPassportSeries sets PassportSeries field to given value.


### GetPostcode

`func (o *Person2) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *Person2) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *Person2) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetMailingAddress

`func (o *Person2) GetMailingAddress() string`

GetMailingAddress returns the MailingAddress field if non-nil, zero value otherwise.

### GetMailingAddressOk

`func (o *Person2) GetMailingAddressOk() (*string, bool)`

GetMailingAddressOk returns a tuple with the MailingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingAddress

`func (o *Person2) SetMailingAddress(v string)`

SetMailingAddress sets MailingAddress field to given value.


### GetPhone

`func (o *Person2) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Person2) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Person2) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *Person2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Person2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Person2) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCountryCode

`func (o *Person2) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Person2) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Person2) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Person2) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *Person2) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *Person2) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


