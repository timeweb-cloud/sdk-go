# NodeGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | Pointer to **string** | ID запроса | [optional] 
**NodeGroup** | [**NodeGroupOut**](NodeGroupOut.md) |  | 

## Methods

### NewNodeGroupResponse

`func NewNodeGroupResponse(nodeGroup NodeGroupOut, ) *NodeGroupResponse`

NewNodeGroupResponse instantiates a new NodeGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeGroupResponseWithDefaults

`func NewNodeGroupResponseWithDefaults() *NodeGroupResponse`

NewNodeGroupResponseWithDefaults instantiates a new NodeGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *NodeGroupResponse) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *NodeGroupResponse) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *NodeGroupResponse) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *NodeGroupResponse) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetNodeGroup

`func (o *NodeGroupResponse) GetNodeGroup() NodeGroupOut`

GetNodeGroup returns the NodeGroup field if non-nil, zero value otherwise.

### GetNodeGroupOk

`func (o *NodeGroupResponse) GetNodeGroupOk() (*NodeGroupOut, bool)`

GetNodeGroupOk returns a tuple with the NodeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroup

`func (o *NodeGroupResponse) SetNodeGroup(v NodeGroupOut)`

SetNodeGroup sets NodeGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


