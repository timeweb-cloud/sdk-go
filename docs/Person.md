# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Уникальный идентификатор администратора домена. | 
**Type** | **string** | Тип администратора: &#x60;person&#x60; — физическое лицо, &#x60;org&#x60; — организация, &#x60;ip&#x60; — индивидуальный предприниматель. | 
**Name** | **string** | Имя администратора (ФИО или название организации). | 
**NameEng** | **string** | Имя администратора в латинской транслитерации. | 
**Email** | **string** | Адрес электронной почты администратора. | 
**Phone** | **string** | Контактный телефон администратора. | 
**Postcode** | **string** | Почтовый индекс. | 
**MailingAddress** | **string** | Почтовый адрес. | 
**CountryCode** | **NullableString** | Код страны администратора. | 
**IsResident** | **bool** | Это логическое значение, которое показывает, является ли администратор резидентом РФ. | 
**IsBlank** | **bool** | Это логическое значение, которое показывает, заполнены ли данные администратора не полностью. | 
**IsClosed** | **bool** | Это логическое значение, которое показывает, закрыт ли администратор. | 
**Birthdate** | Pointer to **string** | Дата рождения. Только для типов &#x60;person&#x60; и &#x60;ip&#x60;. | [optional] 
**PassportDate** | Pointer to **string** | Дата выдачи паспорта. Только для типов &#x60;person&#x60; и &#x60;ip&#x60;. | [optional] 
**PassportNumber** | Pointer to **string** | Номер паспорта. Только для типов &#x60;person&#x60; и &#x60;ip&#x60;. | [optional] 
**PassportPlace** | Pointer to **string** | Кем выдан паспорт. Только для типов &#x60;person&#x60; и &#x60;ip&#x60;. | [optional] 
**PassportSeries** | Pointer to **string** | Серия паспорта. Только для типов &#x60;person&#x60; и &#x60;ip&#x60;. | [optional] 
**Inn** | Pointer to **string** | ИНН. Только для типов &#x60;org&#x60; и &#x60;ip&#x60;. | [optional] 
**Kpp** | Pointer to **string** | КПП организации. Только для типа &#x60;org&#x60;. | [optional] 
**LegalAddress** | Pointer to **string** | Юридический адрес организации. Только для типа &#x60;org&#x60;. | [optional] 
**ContactName** | Pointer to **string** | Контактное лицо организации. Только для типа &#x60;org&#x60;. | [optional] 

## Methods

### NewPerson

`func NewPerson(id int32, type_ string, name string, nameEng string, email string, phone string, postcode string, mailingAddress string, countryCode NullableString, isResident bool, isBlank bool, isClosed bool, ) *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Person) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Person) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Person) SetId(v int32)`

SetId sets Id field to given value.


### GetType

`func (o *Person) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Person) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Person) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Person) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Person) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Person) SetName(v string)`

SetName sets Name field to given value.


### GetNameEng

`func (o *Person) GetNameEng() string`

GetNameEng returns the NameEng field if non-nil, zero value otherwise.

### GetNameEngOk

`func (o *Person) GetNameEngOk() (*string, bool)`

GetNameEngOk returns a tuple with the NameEng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameEng

`func (o *Person) SetNameEng(v string)`

SetNameEng sets NameEng field to given value.


### GetEmail

`func (o *Person) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Person) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Person) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *Person) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Person) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Person) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPostcode

`func (o *Person) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *Person) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *Person) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetMailingAddress

`func (o *Person) GetMailingAddress() string`

GetMailingAddress returns the MailingAddress field if non-nil, zero value otherwise.

### GetMailingAddressOk

`func (o *Person) GetMailingAddressOk() (*string, bool)`

GetMailingAddressOk returns a tuple with the MailingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingAddress

`func (o *Person) SetMailingAddress(v string)`

SetMailingAddress sets MailingAddress field to given value.


### GetCountryCode

`func (o *Person) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Person) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Person) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *Person) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *Person) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetIsResident

`func (o *Person) GetIsResident() bool`

GetIsResident returns the IsResident field if non-nil, zero value otherwise.

### GetIsResidentOk

`func (o *Person) GetIsResidentOk() (*bool, bool)`

GetIsResidentOk returns a tuple with the IsResident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResident

`func (o *Person) SetIsResident(v bool)`

SetIsResident sets IsResident field to given value.


### GetIsBlank

`func (o *Person) GetIsBlank() bool`

GetIsBlank returns the IsBlank field if non-nil, zero value otherwise.

### GetIsBlankOk

`func (o *Person) GetIsBlankOk() (*bool, bool)`

GetIsBlankOk returns a tuple with the IsBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlank

`func (o *Person) SetIsBlank(v bool)`

SetIsBlank sets IsBlank field to given value.


### GetIsClosed

`func (o *Person) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *Person) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *Person) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.


### GetBirthdate

`func (o *Person) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *Person) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *Person) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *Person) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### GetPassportDate

`func (o *Person) GetPassportDate() string`

GetPassportDate returns the PassportDate field if non-nil, zero value otherwise.

### GetPassportDateOk

`func (o *Person) GetPassportDateOk() (*string, bool)`

GetPassportDateOk returns a tuple with the PassportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportDate

`func (o *Person) SetPassportDate(v string)`

SetPassportDate sets PassportDate field to given value.

### HasPassportDate

`func (o *Person) HasPassportDate() bool`

HasPassportDate returns a boolean if a field has been set.

### GetPassportNumber

`func (o *Person) GetPassportNumber() string`

GetPassportNumber returns the PassportNumber field if non-nil, zero value otherwise.

### GetPassportNumberOk

`func (o *Person) GetPassportNumberOk() (*string, bool)`

GetPassportNumberOk returns a tuple with the PassportNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportNumber

`func (o *Person) SetPassportNumber(v string)`

SetPassportNumber sets PassportNumber field to given value.

### HasPassportNumber

`func (o *Person) HasPassportNumber() bool`

HasPassportNumber returns a boolean if a field has been set.

### GetPassportPlace

`func (o *Person) GetPassportPlace() string`

GetPassportPlace returns the PassportPlace field if non-nil, zero value otherwise.

### GetPassportPlaceOk

`func (o *Person) GetPassportPlaceOk() (*string, bool)`

GetPassportPlaceOk returns a tuple with the PassportPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportPlace

`func (o *Person) SetPassportPlace(v string)`

SetPassportPlace sets PassportPlace field to given value.

### HasPassportPlace

`func (o *Person) HasPassportPlace() bool`

HasPassportPlace returns a boolean if a field has been set.

### GetPassportSeries

`func (o *Person) GetPassportSeries() string`

GetPassportSeries returns the PassportSeries field if non-nil, zero value otherwise.

### GetPassportSeriesOk

`func (o *Person) GetPassportSeriesOk() (*string, bool)`

GetPassportSeriesOk returns a tuple with the PassportSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassportSeries

`func (o *Person) SetPassportSeries(v string)`

SetPassportSeries sets PassportSeries field to given value.

### HasPassportSeries

`func (o *Person) HasPassportSeries() bool`

HasPassportSeries returns a boolean if a field has been set.

### GetInn

`func (o *Person) GetInn() string`

GetInn returns the Inn field if non-nil, zero value otherwise.

### GetInnOk

`func (o *Person) GetInnOk() (*string, bool)`

GetInnOk returns a tuple with the Inn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInn

`func (o *Person) SetInn(v string)`

SetInn sets Inn field to given value.

### HasInn

`func (o *Person) HasInn() bool`

HasInn returns a boolean if a field has been set.

### GetKpp

`func (o *Person) GetKpp() string`

GetKpp returns the Kpp field if non-nil, zero value otherwise.

### GetKppOk

`func (o *Person) GetKppOk() (*string, bool)`

GetKppOk returns a tuple with the Kpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpp

`func (o *Person) SetKpp(v string)`

SetKpp sets Kpp field to given value.

### HasKpp

`func (o *Person) HasKpp() bool`

HasKpp returns a boolean if a field has been set.

### GetLegalAddress

`func (o *Person) GetLegalAddress() string`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *Person) GetLegalAddressOk() (*string, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *Person) SetLegalAddress(v string)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *Person) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetContactName

`func (o *Person) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *Person) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *Person) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *Person) HasContactName() bool`

HasContactName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


