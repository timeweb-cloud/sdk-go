# UpdateDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoprolongEnabled** | Pointer to **bool** | Это логическое значение, которое показывает, включено ли автопродление домена. | [optional] 
**LinkedIp** | Pointer to **NullableString** | Привязанный к домену IP-адрес. | [optional] 

## Methods

### NewUpdateDomain

`func NewUpdateDomain() *UpdateDomain`

NewUpdateDomain instantiates a new UpdateDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainWithDefaults

`func NewUpdateDomainWithDefaults() *UpdateDomain`

NewUpdateDomainWithDefaults instantiates a new UpdateDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoprolongEnabled

`func (o *UpdateDomain) GetIsAutoprolongEnabled() bool`

GetIsAutoprolongEnabled returns the IsAutoprolongEnabled field if non-nil, zero value otherwise.

### GetIsAutoprolongEnabledOk

`func (o *UpdateDomain) GetIsAutoprolongEnabledOk() (*bool, bool)`

GetIsAutoprolongEnabledOk returns a tuple with the IsAutoprolongEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprolongEnabled

`func (o *UpdateDomain) SetIsAutoprolongEnabled(v bool)`

SetIsAutoprolongEnabled sets IsAutoprolongEnabled field to given value.

### HasIsAutoprolongEnabled

`func (o *UpdateDomain) HasIsAutoprolongEnabled() bool`

HasIsAutoprolongEnabled returns a boolean if a field has been set.

### GetLinkedIp

`func (o *UpdateDomain) GetLinkedIp() string`

GetLinkedIp returns the LinkedIp field if non-nil, zero value otherwise.

### GetLinkedIpOk

`func (o *UpdateDomain) GetLinkedIpOk() (*string, bool)`

GetLinkedIpOk returns a tuple with the LinkedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIp

`func (o *UpdateDomain) SetLinkedIp(v string)`

SetLinkedIp sets LinkedIp field to given value.

### HasLinkedIp

`func (o *UpdateDomain) HasLinkedIp() bool`

HasLinkedIp returns a boolean if a field has been set.

### SetLinkedIpNil

`func (o *UpdateDomain) SetLinkedIpNil(b bool)`

 SetLinkedIpNil sets the value for LinkedIp to be an explicit nil

### UnsetLinkedIp
`func (o *UpdateDomain) UnsetLinkedIp()`

UnsetLinkedIp ensures that no value is present for LinkedIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


