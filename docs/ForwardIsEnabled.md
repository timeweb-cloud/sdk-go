# ForwardIsEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включена ли пересылка входящих писем | 
**ForwardList** | **[]string** | Список адресов для пересылки. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | 
**IsLeaveMessages** | Pointer to **bool** | Оставлять ли копии входящих писем в почтовом ящике (не удалять). \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewForwardIsEnabled

`func NewForwardIsEnabled(isEnabled bool, forwardList []string, ) *ForwardIsEnabled`

NewForwardIsEnabled instantiates a new ForwardIsEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardIsEnabledWithDefaults

`func NewForwardIsEnabledWithDefaults() *ForwardIsEnabled`

NewForwardIsEnabledWithDefaults instantiates a new ForwardIsEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ForwardIsEnabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ForwardIsEnabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ForwardIsEnabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetForwardList

`func (o *ForwardIsEnabled) GetForwardList() []string`

GetForwardList returns the ForwardList field if non-nil, zero value otherwise.

### GetForwardListOk

`func (o *ForwardIsEnabled) GetForwardListOk() (*[]string, bool)`

GetForwardListOk returns a tuple with the ForwardList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardList

`func (o *ForwardIsEnabled) SetForwardList(v []string)`

SetForwardList sets ForwardList field to given value.


### GetIsLeaveMessages

`func (o *ForwardIsEnabled) GetIsLeaveMessages() bool`

GetIsLeaveMessages returns the IsLeaveMessages field if non-nil, zero value otherwise.

### GetIsLeaveMessagesOk

`func (o *ForwardIsEnabled) GetIsLeaveMessagesOk() (*bool, bool)`

GetIsLeaveMessagesOk returns a tuple with the IsLeaveMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeaveMessages

`func (o *ForwardIsEnabled) SetIsLeaveMessages(v bool)`

SetIsLeaveMessages sets IsLeaveMessages field to given value.

### HasIsLeaveMessages

`func (o *ForwardIsEnabled) HasIsLeaveMessages() bool`

HasIsLeaveMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


