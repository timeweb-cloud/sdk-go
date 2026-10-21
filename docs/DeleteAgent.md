# DeleteAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Причина удаления | [optional] 
**Comment** | Pointer to **string** | Комментарий к причине удаления | [optional] 

## Methods

### NewDeleteAgent

`func NewDeleteAgent() *DeleteAgent`

NewDeleteAgent instantiates a new DeleteAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAgentWithDefaults

`func NewDeleteAgentWithDefaults() *DeleteAgent`

NewDeleteAgentWithDefaults instantiates a new DeleteAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *DeleteAgent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DeleteAgent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DeleteAgent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DeleteAgent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *DeleteAgent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DeleteAgent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DeleteAgent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DeleteAgent) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


