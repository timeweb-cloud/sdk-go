# FirewallGroupResourceOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса. | [optional] 
**Resource** | [**FirewallGroupResource**](FirewallGroupResource.md) |  | 

## Methods

### NewFirewallGroupResourceOutResponse

`func NewFirewallGroupResourceOutResponse(resource FirewallGroupResource, ) *FirewallGroupResourceOutResponse`

NewFirewallGroupResourceOutResponse instantiates a new FirewallGroupResourceOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallGroupResourceOutResponseWithDefaults

`func NewFirewallGroupResourceOutResponseWithDefaults() *FirewallGroupResourceOutResponse`

NewFirewallGroupResourceOutResponseWithDefaults instantiates a new FirewallGroupResourceOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *FirewallGroupResourceOutResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *FirewallGroupResourceOutResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *FirewallGroupResourceOutResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *FirewallGroupResourceOutResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetResource

`func (o *FirewallGroupResourceOutResponse) GetResource() FirewallGroupResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *FirewallGroupResourceOutResponse) GetResourceOk() (*FirewallGroupResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *FirewallGroupResourceOutResponse) SetResource(v FirewallGroupResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


