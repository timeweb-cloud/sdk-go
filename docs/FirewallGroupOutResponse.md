# FirewallGroupOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Group** | [**FirewallGroup**](FirewallGroup.md) |  | 

## Methods

### NewFirewallGroupOutResponse

`func NewFirewallGroupOutResponse(group FirewallGroup, ) *FirewallGroupOutResponse`

NewFirewallGroupOutResponse instantiates a new FirewallGroupOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallGroupOutResponseWithDefaults

`func NewFirewallGroupOutResponseWithDefaults() *FirewallGroupOutResponse`

NewFirewallGroupOutResponseWithDefaults instantiates a new FirewallGroupOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *FirewallGroupOutResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *FirewallGroupOutResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *FirewallGroupOutResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *FirewallGroupOutResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetGroup

`func (o *FirewallGroupOutResponse) GetGroup() FirewallGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FirewallGroupOutResponse) GetGroupOk() (*FirewallGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FirewallGroupOutResponse) SetGroup(v FirewallGroup)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


