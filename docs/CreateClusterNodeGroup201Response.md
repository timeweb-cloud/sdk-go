# CreateClusterNodeGroup201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **string** | Идентификатор запроса, который можно указывать при обращении в службу технической поддержки, чтобы помочь определить проблему. | 
**NodeGroup** | [**NodeGroupOut**](NodeGroupOut.md) |  | 

## Methods

### NewCreateClusterNodeGroup201Response

`func NewCreateClusterNodeGroup201Response(responseId string, nodeGroup NodeGroupOut, ) *CreateClusterNodeGroup201Response`

NewCreateClusterNodeGroup201Response instantiates a new CreateClusterNodeGroup201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterNodeGroup201ResponseWithDefaults

`func NewCreateClusterNodeGroup201ResponseWithDefaults() *CreateClusterNodeGroup201Response`

NewCreateClusterNodeGroup201ResponseWithDefaults instantiates a new CreateClusterNodeGroup201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *CreateClusterNodeGroup201Response) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *CreateClusterNodeGroup201Response) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *CreateClusterNodeGroup201Response) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.


### GetNodeGroup

`func (o *CreateClusterNodeGroup201Response) GetNodeGroup() NodeGroupOut`

GetNodeGroup returns the NodeGroup field if non-nil, zero value otherwise.

### GetNodeGroupOk

`func (o *CreateClusterNodeGroup201Response) GetNodeGroupOk() (*NodeGroupOut, bool)`

GetNodeGroupOk returns a tuple with the NodeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroup

`func (o *CreateClusterNodeGroup201Response) SetNodeGroup(v NodeGroupOut)`

SetNodeGroup sets NodeGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


