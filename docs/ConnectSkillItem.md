# ConnectSkillItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillId** | **int32** | ID навыка | 
**Version** | Pointer to **int32** | Версия из истории навыка. Если не передано — текущая | [optional] 

## Methods

### NewConnectSkillItem

`func NewConnectSkillItem(skillId int32, ) *ConnectSkillItem`

NewConnectSkillItem instantiates a new ConnectSkillItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectSkillItemWithDefaults

`func NewConnectSkillItemWithDefaults() *ConnectSkillItem`

NewConnectSkillItemWithDefaults instantiates a new ConnectSkillItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillId

`func (o *ConnectSkillItem) GetSkillId() int32`

GetSkillId returns the SkillId field if non-nil, zero value otherwise.

### GetSkillIdOk

`func (o *ConnectSkillItem) GetSkillIdOk() (*int32, bool)`

GetSkillIdOk returns a tuple with the SkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillId

`func (o *ConnectSkillItem) SetSkillId(v int32)`

SetSkillId sets SkillId field to given value.


### GetVersion

`func (o *ConnectSkillItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectSkillItem) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectSkillItem) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectSkillItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


