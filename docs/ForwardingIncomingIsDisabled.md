# ForwardingIncomingIsDisabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Включена ли пересылка входящик писем | 
**IsDeleteMessages** | Pointer to **bool** | Удалять ли входящие письма | [optional] 
**IncomingList** | Pointer to **[]string** | Список адресов для пересылки. Не должен быть пустым при передачи параметра &#x60;is_enabled&#x60;: &#x60;true&#x60; | [optional] 

## Methods

### NewForwardingIncomingIsDisabled

`func NewForwardingIncomingIsDisabled(isEnabled bool, ) *ForwardingIncomingIsDisabled`

NewForwardingIncomingIsDisabled instantiates a new ForwardingIncomingIsDisabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardingIncomingIsDisabledWithDefaults

`func NewForwardingIncomingIsDisabledWithDefaults() *ForwardingIncomingIsDisabled`

NewForwardingIncomingIsDisabledWithDefaults instantiates a new ForwardingIncomingIsDisabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ForwardingIncomingIsDisabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ForwardingIncomingIsDisabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ForwardingIncomingIsDisabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsDeleteMessages

`func (o *ForwardingIncomingIsDisabled) GetIsDeleteMessages() bool`

GetIsDeleteMessages returns the IsDeleteMessages field if non-nil, zero value otherwise.

### GetIsDeleteMessagesOk

`func (o *ForwardingIncomingIsDisabled) GetIsDeleteMessagesOk() (*bool, bool)`

GetIsDeleteMessagesOk returns a tuple with the IsDeleteMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteMessages

`func (o *ForwardingIncomingIsDisabled) SetIsDeleteMessages(v bool)`

SetIsDeleteMessages sets IsDeleteMessages field to given value.

### HasIsDeleteMessages

`func (o *ForwardingIncomingIsDisabled) HasIsDeleteMessages() bool`

HasIsDeleteMessages returns a boolean if a field has been set.

### GetIncomingList

`func (o *ForwardingIncomingIsDisabled) GetIncomingList() []string`

GetIncomingList returns the IncomingList field if non-nil, zero value otherwise.

### GetIncomingListOk

`func (o *ForwardingIncomingIsDisabled) GetIncomingListOk() (*[]string, bool)`

GetIncomingListOk returns a tuple with the IncomingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingList

`func (o *ForwardingIncomingIsDisabled) SetIncomingList(v []string)`

SetIncomingList sets IncomingList field to given value.

### HasIncomingList

`func (o *ForwardingIncomingIsDisabled) HasIncomingList() bool`

HasIncomingList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


