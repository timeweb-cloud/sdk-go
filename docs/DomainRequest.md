# DomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | ID пользователя | 
**AuthCode** | **NullableString** | Код авторизации для переноса домена. | 
**Date** | **time.Time** | Дата создания заявки. | 
**DomainBundleId** | **NullableString** | Идентификационный номер бандла, в который входит данная заявка (null - если заявка не входит ни в один бандл). | 
**ErrorCodeTransfer** | **NullableString** | Код ошибки трансфера домена. | 
**Fqdn** | **string** | Полное имя домена. | 
**GroupId** | **float32** | ID группы доменных зон. | 
**Id** | **float32** | ID заявки. | 
**IsAntispamEnabled** | **bool** | Это логическое значение, которое показывает включена ли услуга \&quot;Антиспам\&quot; для домена | 
**IsAutoprolongEnabled** | **bool** | Это логическое значение, которое показывает, включено ли автопродление домена. | 
**IsWhoisPrivacyEnabled** | **bool** | Это логическое значение, которое показывает, включено ли скрытие данных администратора домена для whois. Опция недоступна для доменов в зонах .ru и .рф. | 
**Message** | **NullableString** | Информационное сообщение о заявке. | 
**MoneySource** | **NullableString** | Источник (способ) оплаты заявки. | 
**Period** | [**DomainPaymentPeriod**](DomainPaymentPeriod.md) |  | 
**PersonId** | **float32** | Идентификационный номер персоны для заявки на регистрацию. | 
**Prime** | [**DomainPrimeType**](DomainPrimeType.md) |  | 
**SoonExpire** | **float32** | Количество дней до конца регистрации домена, за которые мы уведомим о необходимости продления. | 
**SortOrder** | **float32** | Это значение используется для сортировки доменных зон в панели управления. | 
**Type** | **string** | Тип заявки. | 

## Methods

### NewDomainRequest

`func NewDomainRequest(accountId string, authCode NullableString, date time.Time, domainBundleId NullableString, errorCodeTransfer NullableString, fqdn string, groupId float32, id float32, isAntispamEnabled bool, isAutoprolongEnabled bool, isWhoisPrivacyEnabled bool, message NullableString, moneySource NullableString, period DomainPaymentPeriod, personId float32, prime DomainPrimeType, soonExpire float32, sortOrder float32, type_ string, ) *DomainRequest`

NewDomainRequest instantiates a new DomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRequestWithDefaults

`func NewDomainRequestWithDefaults() *DomainRequest`

NewDomainRequestWithDefaults instantiates a new DomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DomainRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DomainRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DomainRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAuthCode

`func (o *DomainRequest) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *DomainRequest) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *DomainRequest) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### SetAuthCodeNil

`func (o *DomainRequest) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *DomainRequest) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetDate

`func (o *DomainRequest) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DomainRequest) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DomainRequest) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetDomainBundleId

`func (o *DomainRequest) GetDomainBundleId() string`

GetDomainBundleId returns the DomainBundleId field if non-nil, zero value otherwise.

### GetDomainBundleIdOk

`func (o *DomainRequest) GetDomainBundleIdOk() (*string, bool)`

GetDomainBundleIdOk returns a tuple with the DomainBundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainBundleId

`func (o *DomainRequest) SetDomainBundleId(v string)`

SetDomainBundleId sets DomainBundleId field to given value.


### SetDomainBundleIdNil

`func (o *DomainRequest) SetDomainBundleIdNil(b bool)`

 SetDomainBundleIdNil sets the value for DomainBundleId to be an explicit nil

### UnsetDomainBundleId
`func (o *DomainRequest) UnsetDomainBundleId()`

UnsetDomainBundleId ensures that no value is present for DomainBundleId, not even an explicit nil
### GetErrorCodeTransfer

`func (o *DomainRequest) GetErrorCodeTransfer() string`

GetErrorCodeTransfer returns the ErrorCodeTransfer field if non-nil, zero value otherwise.

### GetErrorCodeTransferOk

`func (o *DomainRequest) GetErrorCodeTransferOk() (*string, bool)`

GetErrorCodeTransferOk returns a tuple with the ErrorCodeTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCodeTransfer

`func (o *DomainRequest) SetErrorCodeTransfer(v string)`

SetErrorCodeTransfer sets ErrorCodeTransfer field to given value.


### SetErrorCodeTransferNil

`func (o *DomainRequest) SetErrorCodeTransferNil(b bool)`

 SetErrorCodeTransferNil sets the value for ErrorCodeTransfer to be an explicit nil

### UnsetErrorCodeTransfer
`func (o *DomainRequest) UnsetErrorCodeTransfer()`

UnsetErrorCodeTransfer ensures that no value is present for ErrorCodeTransfer, not even an explicit nil
### GetFqdn

`func (o *DomainRequest) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DomainRequest) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DomainRequest) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetGroupId

`func (o *DomainRequest) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DomainRequest) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DomainRequest) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetId

`func (o *DomainRequest) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainRequest) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainRequest) SetId(v float32)`

SetId sets Id field to given value.


### GetIsAntispamEnabled

`func (o *DomainRequest) GetIsAntispamEnabled() bool`

GetIsAntispamEnabled returns the IsAntispamEnabled field if non-nil, zero value otherwise.

### GetIsAntispamEnabledOk

`func (o *DomainRequest) GetIsAntispamEnabledOk() (*bool, bool)`

GetIsAntispamEnabledOk returns a tuple with the IsAntispamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAntispamEnabled

`func (o *DomainRequest) SetIsAntispamEnabled(v bool)`

SetIsAntispamEnabled sets IsAntispamEnabled field to given value.


### GetIsAutoprolongEnabled

`func (o *DomainRequest) GetIsAutoprolongEnabled() bool`

GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field if non-nil, zero value otherwise.

### GetIsAutoprolongEnabledOk

`func (o *DomainRequest) GetIsAutoprolongEnabledOk() (*bool, bool)`

GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprolongEnabled

`func (o *DomainRequest) SetIsAutoprolongEnabled(v bool)`

SetIsAutoprolongEnabled sets IsAutoprolongEnabled field to given value.


### GetIsWhoisPrivacyEnabled

`func (o *DomainRequest) GetIsWhoisPrivacyEnabled() bool`

GetIsWhoisPrivacyEnabled returns the IsWhoisPrivacyEnabled field if non-nil, zero value otherwise.

### GetIsWhoisPrivacyEnabledOk

`func (o *DomainRequest) GetIsWhoisPrivacyEnabledOk() (*bool, bool)`

GetIsWhoisPrivacyEnabledOk returns a tuple with the IsWhoisPrivacyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhoisPrivacyEnabled

`func (o *DomainRequest) SetIsWhoisPrivacyEnabled(v bool)`

SetIsWhoisPrivacyEnabled sets IsWhoisPrivacyEnabled field to given value.


### GetMessage

`func (o *DomainRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DomainRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DomainRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *DomainRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DomainRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMoneySource

`func (o *DomainRequest) GetMoneySource() string`

GetMoneySource returns the MoneySource field if non-nil, zero value otherwise.

### GetMoneySourceOk

`func (o *DomainRequest) GetMoneySourceOk() (*string, bool)`

GetMoneySourceOk returns a tuple with the MoneySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneySource

`func (o *DomainRequest) SetMoneySource(v string)`

SetMoneySource sets MoneySource field to given value.


### SetMoneySourceNil

`func (o *DomainRequest) SetMoneySourceNil(b bool)`

 SetMoneySourceNil sets the value for MoneySource to be an explicit nil

### UnsetMoneySource
`func (o *DomainRequest) UnsetMoneySource()`

UnsetMoneySource ensures that no value is present for MoneySource, not even an explicit nil
### GetPeriod

`func (o *DomainRequest) GetPeriod() DomainPaymentPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainRequest) GetPeriodOk() (*DomainPaymentPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainRequest) SetPeriod(v DomainPaymentPeriod)`

SetPeriod sets Period field to given value.


### GetPersonId

`func (o *DomainRequest) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *DomainRequest) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *DomainRequest) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.


### GetPrime

`func (o *DomainRequest) GetPrime() DomainPrimeType`

GetPrime returns the Prime field if non-nil, zero value otherwise.

### GetPrimeOk

`func (o *DomainRequest) GetPrimeOk() (*DomainPrimeType, bool)`

GetPrimeOk returns a tuple with the Prime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrime

`func (o *DomainRequest) SetPrime(v DomainPrimeType)`

SetPrime sets Prime field to given value.


### GetSoonExpire

`func (o *DomainRequest) GetSoonExpire() float32`

GetSoonExpire returns the SoonExpire field if non-nil, zero value otherwise.

### GetSoonExpireOk

`func (o *DomainRequest) GetSoonExpireOk() (*float32, bool)`

GetSoonExpireOk returns a tuple with the SoonExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoonExpire

`func (o *DomainRequest) SetSoonExpire(v float32)`

SetSoonExpire sets SoonExpire field to given value.


### GetSortOrder

`func (o *DomainRequest) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *DomainRequest) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *DomainRequest) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.


### GetType

`func (o *DomainRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


