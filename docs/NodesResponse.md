# NodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | Идентификатор запроса | [optional] 
**Meta** | [**Meta**](Meta.md) |  | 
**Nodes** | [**[]NodeOut**](NodeOut.md) | Массив объектов Нода | 

## Methods

### NewNodesResponse

`func NewNodesResponse(meta Meta, nodes []NodeOut, ) *NodesResponse`

NewNodesResponse instantiates a new NodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodesResponseWithDefaults

`func NewNodesResponseWithDefaults() *NodesResponse`

NewNodesResponseWithDefaults instantiates a new NodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *NodesResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *NodesResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *NodesResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *NodesResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMeta

`func (o *NodesResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NodesResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NodesResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetNodes

`func (o *NodesResponse) GetNodes() []NodeOut`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NodesResponse) GetNodesOk() (*[]NodeOut, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NodesResponse) SetNodes(v []NodeOut)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


