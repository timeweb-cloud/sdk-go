# FirewallGroupsOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Meta** | [**Meta**](Meta.md) |  | 
**Groups** | [**[]FirewallGroup**](FirewallGroup.md) |  | 

## Methods

### NewFirewallGroupsOutResponse

`func NewFirewallGroupsOutResponse(meta Meta, groups []FirewallGroup, ) *FirewallGroupsOutResponse`

NewFirewallGroupsOutResponse instantiates a new FirewallGroupsOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallGroupsOutResponseWithDefaults

`func NewFirewallGroupsOutResponseWithDefaults() *FirewallGroupsOutResponse`

NewFirewallGroupsOutResponseWithDefaults instantiates a new FirewallGroupsOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *FirewallGroupsOutResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *FirewallGroupsOutResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *FirewallGroupsOutResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *FirewallGroupsOutResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *FirewallGroupsOutResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FirewallGroupsOutResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FirewallGroupsOutResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetGroups

`func (o *FirewallGroupsOutResponse) GetGroups() []FirewallGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FirewallGroupsOutResponse) GetGroupsOk() (*[]FirewallGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FirewallGroupsOutResponse) SetGroups(v []FirewallGroup)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


