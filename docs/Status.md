# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBlocked** | **bool** | Это логическое значение, которое показывает, заблокирован ли аккаунт. | 
**IsPermanentBlocked** | **bool** | Это логическое значение, которое показывает, заблокирован ли аккаунт навсегда. | 
**IsSendBillLetters** | **bool** | Это логическое значение, которое показывает, требуется ли отправлять счета на почту. | 
**CompanyInfo** | [**StatusCompanyInfo**](StatusCompanyInfo.md) |  | 
**LastPasswordChangedAt** | **string** | Значение времени, указанное в комбинированном формате даты и времени ISO8601, которое представляет, когда последний раз изменялся пароль. | 
**YmClientId** | **NullableString** | ID аккаунта для яндекс метрики. | 

## Methods

### NewStatus

`func NewStatus(isBlocked bool, isPermanentBlocked bool, isSendBillLetters bool, companyInfo StatusCompanyInfo, lastPasswordChangedAt string, ymClientId NullableString, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBlocked

`func (o *Status) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *Status) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *Status) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.


### GetIsPermanentBlocked

`func (o *Status) GetIsPermanentBlocked() bool`

GetIsPermanentBlocked returns the IsPermanentBlocked field if non-nil, zero value otherwise.

### GetIsPermanentBlockedOk

`func (o *Status) GetIsPermanentBlockedOk() (*bool, bool)`

GetIsPermanentBlockedOk returns a tuple with the IsPermanentBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanentBlocked

`func (o *Status) SetIsPermanentBlocked(v bool)`

SetIsPermanentBlocked sets IsPermanentBlocked field to given value.


### GetIsSendBillLetters

`func (o *Status) GetIsSendBillLetters() bool`

GetIsSendBillLetters returns the IsSendBillLetters field if non-nil, zero value otherwise.

### GetIsSendBillLettersOk

`func (o *Status) GetIsSendBillLettersOk() (*bool, bool)`

GetIsSendBillLettersOk returns a tuple with the IsSendBillLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSendBillLetters

`func (o *Status) SetIsSendBillLetters(v bool)`

SetIsSendBillLetters sets IsSendBillLetters field to given value.


### GetCompanyInfo

`func (o *Status) GetCompanyInfo() StatusCompanyInfo`

GetCompanyInfo returns the CompanyInfo field if non-nil, zero value otherwise.

### GetCompanyInfoOk

`func (o *Status) GetCompanyInfoOk() (*StatusCompanyInfo, bool)`

GetCompanyInfoOk returns a tuple with the CompanyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfo

`func (o *Status) SetCompanyInfo(v StatusCompanyInfo)`

SetCompanyInfo sets CompanyInfo field to given value.


### GetLastPasswordChangedAt

`func (o *Status) GetLastPasswordChangedAt() string`

GetLastPasswordChangedAt returns the LastPasswordChangedAt field if non-nil, zero value otherwise.

### GetLastPasswordChangedAtOk

`func (o *Status) GetLastPasswordChangedAtOk() (*string, bool)`

GetLastPasswordChangedAtOk returns a tuple with the LastPasswordChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChangedAt

`func (o *Status) SetLastPasswordChangedAt(v string)`

SetLastPasswordChangedAt sets LastPasswordChangedAt field to given value.


### GetYmClientId

`func (o *Status) GetYmClientId() string`

GetYmClientId returns the YmClientId field if non-nil, zero value otherwise.

### GetYmClientIdOk

`func (o *Status) GetYmClientIdOk() (*string, bool)`

GetYmClientIdOk returns a tuple with the YmClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYmClientId

`func (o *Status) SetYmClientId(v string)`

SetYmClientId sets YmClientId field to given value.


### SetYmClientIdNil

`func (o *Status) SetYmClientIdNil(b bool)`

 SetYmClientIdNil sets the value for YmClientId to be an explicit nil

### UnsetYmClientId
`func (o *Status) UnsetYmClientId()`

UnsetYmClientId ensures that no value is present for YmClientId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


