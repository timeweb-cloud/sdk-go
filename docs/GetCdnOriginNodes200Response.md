# GetCdnOriginNodes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginNodes** | [**[]OriginNode**](OriginNode.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetCdnOriginNodes200Response

`func NewGetCdnOriginNodes200Response(originNodes []OriginNode, meta Meta, ) *GetCdnOriginNodes200Response`

NewGetCdnOriginNodes200Response instantiates a new GetCdnOriginNodes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCdnOriginNodes200ResponseWithDefaults

`func NewGetCdnOriginNodes200ResponseWithDefaults() *GetCdnOriginNodes200Response`

NewGetCdnOriginNodes200ResponseWithDefaults instantiates a new GetCdnOriginNodes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginNodes

`func (o *GetCdnOriginNodes200Response) GetOriginNodes() []OriginNode`

GetOriginNodes returns the OriginNodes field if non-nil, zero value otherwise.

### GetOriginNodesOk

`func (o *GetCdnOriginNodes200Response) GetOriginNodesOk() (*[]OriginNode, bool)`

GetOriginNodesOk returns a tuple with the OriginNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginNodes

`func (o *GetCdnOriginNodes200Response) SetOriginNodes(v []OriginNode)`

SetOriginNodes sets OriginNodes field to given value.


### GetMeta

`func (o *GetCdnOriginNodes200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCdnOriginNodes200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCdnOriginNodes200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


