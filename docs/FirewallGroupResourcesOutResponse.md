# FirewallGroupResourcesOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Meta** | [**Meta**](Meta.md) |  | 
**Resources** | [**[]FirewallGroupResource**](FirewallGroupResource.md) |  | 

## Methods

### NewFirewallGroupResourcesOutResponse

`func NewFirewallGroupResourcesOutResponse(meta Meta, resources []FirewallGroupResource, ) *FirewallGroupResourcesOutResponse`

NewFirewallGroupResourcesOutResponse instantiates a new FirewallGroupResourcesOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallGroupResourcesOutResponseWithDefaults

`func NewFirewallGroupResourcesOutResponseWithDefaults() *FirewallGroupResourcesOutResponse`

NewFirewallGroupResourcesOutResponseWithDefaults instantiates a new FirewallGroupResourcesOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *FirewallGroupResourcesOutResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *FirewallGroupResourcesOutResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *FirewallGroupResourcesOutResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *FirewallGroupResourcesOutResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *FirewallGroupResourcesOutResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FirewallGroupResourcesOutResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FirewallGroupResourcesOutResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetResources

`func (o *FirewallGroupResourcesOutResponse) GetResources() []FirewallGroupResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *FirewallGroupResourcesOutResponse) GetResourcesOk() (*[]FirewallGroupResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *FirewallGroupResourcesOutResponse) SetResources(v []FirewallGroupResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


