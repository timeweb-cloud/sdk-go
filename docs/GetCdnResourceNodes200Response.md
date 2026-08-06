# GetCdnResourceNodes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserNodes** | [**[]UserNode**](UserNode.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetCdnResourceNodes200Response

`func NewGetCdnResourceNodes200Response(userNodes []UserNode, meta Meta, ) *GetCdnResourceNodes200Response`

NewGetCdnResourceNodes200Response instantiates a new GetCdnResourceNodes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCdnResourceNodes200ResponseWithDefaults

`func NewGetCdnResourceNodes200ResponseWithDefaults() *GetCdnResourceNodes200Response`

NewGetCdnResourceNodes200ResponseWithDefaults instantiates a new GetCdnResourceNodes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserNodes

`func (o *GetCdnResourceNodes200Response) GetUserNodes() []UserNode`

GetUserNodes returns the UserNodes field if non-nil, zero value otherwise.

### GetUserNodesOk

`func (o *GetCdnResourceNodes200Response) GetUserNodesOk() (*[]UserNode, bool)`

GetUserNodesOk returns a tuple with the UserNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNodes

`func (o *GetCdnResourceNodes200Response) SetUserNodes(v []UserNode)`

SetUserNodes sets UserNodes field to given value.


### GetMeta

`func (o *GetCdnResourceNodes200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCdnResourceNodes200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCdnResourceNodes200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


