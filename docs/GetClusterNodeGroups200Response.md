# GetClusterNodeGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**Meta** | [**Meta**](Meta.md) |  | 
**NodeGroups** | [**[]NodeGroupOut**](NodeGroupOut.md) | Массив объектов Группа нод | 

## Methods

### NewGetClusterNodeGroups200Response

`func NewGetClusterNodeGroups200Response(responseId string, meta Meta, nodeGroups []NodeGroupOut, ) *GetClusterNodeGroups200Response`

NewGetClusterNodeGroups200Response instantiates a new GetClusterNodeGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterNodeGroups200ResponseWithDefaults

`func NewGetClusterNodeGroups200ResponseWithDefaults() *GetClusterNodeGroups200Response`

NewGetClusterNodeGroups200ResponseWithDefaults instantiates a new GetClusterNodeGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *GetClusterNodeGroups200Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *GetClusterNodeGroups200Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *GetClusterNodeGroups200Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetMeta

`func (o *GetClusterNodeGroups200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetClusterNodeGroups200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetClusterNodeGroups200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetNodeGroups

`func (o *GetClusterNodeGroups200Response) GetNodeGroups() []NodeGroupOut`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *GetClusterNodeGroups200Response) GetNodeGroupsOk() (*[]NodeGroupOut, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *GetClusterNodeGroups200Response) SetNodeGroups(v []NodeGroupOut)`

SetNodeGroups sets NodeGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


