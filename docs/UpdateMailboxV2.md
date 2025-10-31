# UpdateMailboxV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Пароль почтового ящика | [optional] 
**Comment** | Pointer to **string** | Комментарий к почтовому ящику | [optional] 
**OwnerFullName** | Pointer to **string** | ФИО владельца почтового ящика | [optional] 
**SpamProtectionSettings** | Pointer to [**SpamProtectionIsEnabled**](SpamProtectionIsEnabled.md) |  | [optional] 
**ForwardSettings** | Pointer to [**ForwardIsEnabled**](ForwardIsEnabled.md) |  | [optional] 
**AutoreplySettings** | Pointer to [**AutoreplyIsEnabled**](AutoreplyIsEnabled.md) |  | [optional] 
**OutgoingSettings** | Pointer to [**OutgoingIsEnabled**](OutgoingIsEnabled.md) |  | [optional] 

## Methods

### NewUpdateMailboxV2

`func NewUpdateMailboxV2() *UpdateMailboxV2`

NewUpdateMailboxV2 instantiates a new UpdateMailboxV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMailboxV2WithDefaults

`func NewUpdateMailboxV2WithDefaults() *UpdateMailboxV2`

NewUpdateMailboxV2WithDefaults instantiates a new UpdateMailboxV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdateMailboxV2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateMailboxV2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateMailboxV2) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateMailboxV2) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetComment

`func (o *UpdateMailboxV2) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateMailboxV2) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateMailboxV2) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateMailboxV2) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetOwnerFullName

`func (o *UpdateMailboxV2) GetOwnerFullName() string`

GetOwnerFullName returns the OwnerFullName field if non-nil, zero value otherwise.

### GetOwnerFullNameOk

`func (o *UpdateMailboxV2) GetOwnerFullNameOk() (*string, bool)`

GetOwnerFullNameOk returns a tuple with the OwnerFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFullName

`func (o *UpdateMailboxV2) SetOwnerFullName(v string)`

SetOwnerFullName sets OwnerFullName field to given value.

### HasOwnerFullName

`func (o *UpdateMailboxV2) HasOwnerFullName() bool`

HasOwnerFullName returns a boolean if a field has been set.

### GetSpamProtectionSettings

`func (o *UpdateMailboxV2) GetSpamProtectionSettings() SpamProtectionIsEnabled`

GetSpamProtectionSettings returns the SpamProtectionSettings field if non-nil, zero value otherwise.

### GetSpamProtectionSettingsOk

`func (o *UpdateMailboxV2) GetSpamProtectionSettingsOk() (*SpamProtectionIsEnabled, bool)`

GetSpamProtectionSettingsOk returns a tuple with the SpamProtectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamProtectionSettings

`func (o *UpdateMailboxV2) SetSpamProtectionSettings(v SpamProtectionIsEnabled)`

SetSpamProtectionSettings sets SpamProtectionSettings field to given value.

### HasSpamProtectionSettings

`func (o *UpdateMailboxV2) HasSpamProtectionSettings() bool`

HasSpamProtectionSettings returns a boolean if a field has been set.

### GetForwardSettings

`func (o *UpdateMailboxV2) GetForwardSettings() ForwardIsEnabled`

GetForwardSettings returns the ForwardSettings field if non-nil, zero value otherwise.

### GetForwardSettingsOk

`func (o *UpdateMailboxV2) GetForwardSettingsOk() (*ForwardIsEnabled, bool)`

GetForwardSettingsOk returns a tuple with the ForwardSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardSettings

`func (o *UpdateMailboxV2) SetForwardSettings(v ForwardIsEnabled)`

SetForwardSettings sets ForwardSettings field to given value.

### HasForwardSettings

`func (o *UpdateMailboxV2) HasForwardSettings() bool`

HasForwardSettings returns a boolean if a field has been set.

### GetAutoreplySettings

`func (o *UpdateMailboxV2) GetAutoreplySettings() AutoreplyIsEnabled`

GetAutoreplySettings returns the AutoreplySettings field if non-nil, zero value otherwise.

### GetAutoreplySettingsOk

`func (o *UpdateMailboxV2) GetAutoreplySettingsOk() (*AutoreplyIsEnabled, bool)`

GetAutoreplySettingsOk returns a tuple with the AutoreplySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoreplySettings

`func (o *UpdateMailboxV2) SetAutoreplySettings(v AutoreplyIsEnabled)`

SetAutoreplySettings sets AutoreplySettings field to given value.

### HasAutoreplySettings

`func (o *UpdateMailboxV2) HasAutoreplySettings() bool`

HasAutoreplySettings returns a boolean if a field has been set.

### GetOutgoingSettings

`func (o *UpdateMailboxV2) GetOutgoingSettings() OutgoingIsEnabled`

GetOutgoingSettings returns the OutgoingSettings field if non-nil, zero value otherwise.

### GetOutgoingSettingsOk

`func (o *UpdateMailboxV2) GetOutgoingSettingsOk() (*OutgoingIsEnabled, bool)`

GetOutgoingSettingsOk returns a tuple with the OutgoingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingSettings

`func (o *UpdateMailboxV2) SetOutgoingSettings(v OutgoingIsEnabled)`

SetOutgoingSettings sets OutgoingSettings field to given value.

### HasOutgoingSettings

`func (o *UpdateMailboxV2) HasOutgoingSettings() bool`

HasOutgoingSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


