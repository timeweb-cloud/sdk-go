# SpamProtectionIsEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включен ли спам-фильтр | 
**FilterAction** | Pointer to **string** | Что делать с письмами, которые попадают в спам. \\  Параметры: \\  &#x60;directory&#x60; - переместить в папку спам; \\  &#x60;label&#x60; - пометить письмо; \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**WhiteList** | Pointer to **[]string** | Белый список адресов от которых письма не будут попадать в спам. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewSpamProtectionIsEnabled

`func NewSpamProtectionIsEnabled(isEnabled bool, ) *SpamProtectionIsEnabled`

NewSpamProtectionIsEnabled instantiates a new SpamProtectionIsEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpamProtectionIsEnabledWithDefaults

`func NewSpamProtectionIsEnabledWithDefaults() *SpamProtectionIsEnabled`

NewSpamProtectionIsEnabledWithDefaults instantiates a new SpamProtectionIsEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *SpamProtectionIsEnabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SpamProtectionIsEnabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SpamProtectionIsEnabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetFilterAction

`func (o *SpamProtectionIsEnabled) GetFilterAction() string`

GetFilterAction returns the FilterAction field if non-nil, zero value otherwise.

### GetFilterActionOk

`func (o *SpamProtectionIsEnabled) GetFilterActionOk() (*string, bool)`

GetFilterActionOk returns a tuple with the FilterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAction

`func (o *SpamProtectionIsEnabled) SetFilterAction(v string)`

SetFilterAction sets FilterAction field to given value.

### HasFilterAction

`func (o *SpamProtectionIsEnabled) HasFilterAction() bool`

HasFilterAction returns a boolean if a field has been set.

### GetWhiteList

`func (o *SpamProtectionIsEnabled) GetWhiteList() []string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *SpamProtectionIsEnabled) GetWhiteListOk() (*[]string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *SpamProtectionIsEnabled) SetWhiteList(v []string)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *SpamProtectionIsEnabled) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


