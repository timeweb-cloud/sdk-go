# GetClusterNodesFromGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Nodes** | [**[]NodeOut**](NodeOut.md) | Массив объектов Нода | 

## Methods

### NewGetClusterNodesFromGroup200Response

`func NewGetClusterNodesFromGroup200Response(responseId string, meta Meta, nodes []NodeOut, ) *GetClusterNodesFromGroup200Response`

NewGetClusterNodesFromGroup200Response instantiates a new GetClusterNodesFromGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterNodesFromGroup200ResponseWithDefaults

`func NewGetClusterNodesFromGroup200ResponseWithDefaults() *GetClusterNodesFromGroup200Response`

NewGetClusterNodesFromGroup200ResponseWithDefaults instantiates a new GetClusterNodesFromGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetClusterNodesFromGroup200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetClusterNodesFromGroup200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetClusterNodesFromGroup200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetClusterNodesFromGroup200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetClusterNodesFromGroup200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetClusterNodesFromGroup200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetNodes

`func (o *GetClusterNodesFromGroup200Response) GetNodes() []NodeOut`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetClusterNodesFromGroup200Response) GetNodesOk() (*[]NodeOut, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetClusterNodesFromGroup200Response) SetNodes(v []NodeOut)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


