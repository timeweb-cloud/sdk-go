# SpamFilterIsDisabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включен ли спам-фильтр | 
**Action** | Pointer to **string** | Что делать с письмами, которые попадают в спам. \\  Параметры: \\  &#x60;move_to_directory&#x60; - переместить в паку спам; \\  &#x60;forward&#x60; - переслать письмо на другой адрес; \\  &#x60;delete&#x60; - удалить письмо; \\  &#x60;tag&#x60; - пометить письмо; | [optional] 
**ForwardTo** | Pointer to **string** | Адрес для пересылки при выбранном действии &#x60;forward&#x60; из параметра &#x60;action&#x60;. Не может быть пустым, если &#x60;action&#x60; выбран &#x60;forward&#x60; | [optional] 
**WhiteList** | Pointer to **[]string** | Белый список адресов от которых письма не будут попадать в спам | [optional] 

## Methods

### NewSpamFilterIsDisabled

`func NewSpamFilterIsDisabled(isEnabled bool, ) *SpamFilterIsDisabled`

NewSpamFilterIsDisabled instantiates a new SpamFilterIsDisabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpamFilterIsDisabledWithDefaults

`func NewSpamFilterIsDisabledWithDefaults() *SpamFilterIsDisabled`

NewSpamFilterIsDisabledWithDefaults instantiates a new SpamFilterIsDisabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *SpamFilterIsDisabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SpamFilterIsDisabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SpamFilterIsDisabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetAction

`func (o *SpamFilterIsDisabled) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SpamFilterIsDisabled) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SpamFilterIsDisabled) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SpamFilterIsDisabled) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetForwardTo

`func (o *SpamFilterIsDisabled) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *SpamFilterIsDisabled) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *SpamFilterIsDisabled) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *SpamFilterIsDisabled) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetWhiteList

`func (o *SpamFilterIsDisabled) GetWhiteList() []string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *SpamFilterIsDisabled) GetWhiteListOk() (*[]string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *SpamFilterIsDisabled) SetWhiteList(v []string)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *SpamFilterIsDisabled) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


