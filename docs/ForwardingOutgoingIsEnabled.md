# ForwardingOutgoingIsEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Включена ли пересылка исходящих писем. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 
**OutgoingTo** | Pointer to **string** | Адрес для пересылки исходящих писем. \\  Если передан параметр &#x60;is_enabled&#x60;: &#x60;false&#x60;, то значение передавать нельзя | [optional] 

## Methods

### NewForwardingOutgoingIsEnabled

`func NewForwardingOutgoingIsEnabled() *ForwardingOutgoingIsEnabled`

NewForwardingOutgoingIsEnabled instantiates a new ForwardingOutgoingIsEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardingOutgoingIsEnabledWithDefaults

`func NewForwardingOutgoingIsEnabledWithDefaults() *ForwardingOutgoingIsEnabled`

NewForwardingOutgoingIsEnabledWithDefaults instantiates a new ForwardingOutgoingIsEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ForwardingOutgoingIsEnabled) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ForwardingOutgoingIsEnabled) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ForwardingOutgoingIsEnabled) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ForwardingOutgoingIsEnabled) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetOutgoingTo

`func (o *ForwardingOutgoingIsEnabled) GetOutgoingTo() string`

GetOutgoingTo returns the OutgoingTo field if non-nil, zero value otherwise.

### GetOutgoingToOk

`func (o *ForwardingOutgoingIsEnabled) GetOutgoingToOk() (*string, bool)`

GetOutgoingToOk returns a tuple with the OutgoingTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTo

`func (o *ForwardingOutgoingIsEnabled) SetOutgoingTo(v string)`

SetOutgoingTo sets OutgoingTo field to given value.

### HasOutgoingTo

`func (o *ForwardingOutgoingIsEnabled) HasOutgoingTo() bool`

HasOutgoingTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


